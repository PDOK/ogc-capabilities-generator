package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"flag"
	"log"
	"ogc-capabilities-generator/config"
	"ogc-capabilities-generator/ows"
	"ogc-capabilities-generator/validate"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	"dario.cat/mergo"
	"gopkg.in/yaml.v3"

	"github.com/pdok/ogc-specifications/pkg/wms130"
)

func envString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func writeFile(filename string, buffer []byte) error {
	makeDirIfNotExists(filename)
	err := os.WriteFile(filename, buffer, 0777)
	if err != nil {
		return err
	}
	return nil
}

func buildCapabilities(v interface{}, g config.Global) ([]byte, error) {

	if g.Prefix == `` {
		return nil, errors.New("no dataset prefix defined")
	}

	si, _ := xml.MarshalIndent(v, "", "\t")
	t := template.Must(template.New("capabilities").Parse(string(si)))
	buf := &bytes.Buffer{}
	err := t.Execute(buf, g)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`><.*>`)
	originalCapabilities := buf.String()
	regexedCapabilities := xml.Header + re.ReplaceAllString(originalCapabilities, "/>")
	return []byte(regexedCapabilities), nil
}

func buildWFS2_0_0(cfg config.Config) error {

	// retrieve default set
	wfs200base := ows.WFS200Base
	target := cfg.Services.WFS200Config.Wfs200

	// merge with specific set skipping featuretypelist, this is a custom operation
	err := mergo.Merge(&target, wfs200base, mergo.WithTransformers(ows.WFS200Transfomer{}))
	if err != nil {
		return err
	}

	// can we apply generic base feature template to config ?
	if len(wfs200base.Capabilities.FeatureTypeList.FeatureType) > 0 {
		for index := range target.FeatureTypeList.FeatureType {
			err := mergo.Merge(&target.FeatureTypeList.FeatureType[index], wfs200base.Capabilities.FeatureTypeList.FeatureType[0])
			if err != nil {
				return err
			}
		}
	}

	if target.Capabilities.OperationsMetadata.ExtendedCapabilities != nil {
		target.Namespaces.XmlnsInspireCommon = "http://inspire.ec.europa.eu/schemas/common/1.0"
		target.Namespaces.XmlnsInspireDls = "http://inspire.ec.europa.eu/schemas/inspire_dls/1.0"
	}

	buf, err := buildCapabilities(target, cfg.Global)
	if err != nil {
		return err
	}

	err = validate.ValidateCapabilities(&cfg, buf, target.SchemaLocation)
	if err != nil {
		return err
	}

	err = writeFile(cfg.Services.WFS200Config.Filename, buf)
	if err != nil {
		return err
	}

	return nil
}

func buildWMS1_3_0(cfg config.Config) error {
	wms130base := ows.WMS130Base
	configBase := cfg.Services.WMS130Config.Wms130

	// merge with specific set skipping layer, this is a custom operation
	err := mergo.Merge(&configBase, wms130base, mergo.WithTransformers(ows.WMS130Transfomer{}))
	if err != nil {
		return err
	}

	if len(wms130base.Capabilities.Layer) > 0 {
		for index := range configBase.Capabilities.Layer {
			err = merge(&configBase.Capabilities.Layer[index], wms130base.Capabilities.Layer[0])
			if err != nil {
				return err
			}
		}
	}

	if configBase.Capabilities.WMSCapabilities.ExtendedCapabilities != nil {
		configBase.Namespaces.XmlnsInspireCommon = "http://inspire.ec.europa.eu/schemas/common/1.0"
		configBase.Namespaces.XmlnsInspireVs = "http://inspire.ec.europa.eu/schemas/inspire_vs/1.0"
		configBase.Namespaces.SchemaLocation = wms130base.Namespaces.SchemaLocation + " " + "http://inspire.ec.europa.eu/schemas/inspire_vs/1.0 http://inspire.ec.europa.eu/schemas/inspire_vs/1.0/inspire_vs.xsd"
	}

	buf, err := buildCapabilities(configBase, cfg.Global)
	if err != nil {
		return err
	}

	writeFile(cfg.Services.WMS130Config.Filename, buf)

	err = validate.ValidateCapabilities(&cfg, buf, configBase.SchemaLocation)
	if err != nil {
		return err
	}

	return nil
}

// recursive fill
func merge(dst *wms130.Layer, src wms130.Layer) error {
	if len(dst.Layer) > 0 {
		for index := range dst.Layer {
			if err := merge(dst.Layer[index], src); err != nil {
				return err
			}
		}
	}
	return mergo.Merge(dst, src)
}

func makeDirIfNotExists(filename string) {
	dir, _ := filepath.Split(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func buildWMTS1_0_0(cfg config.Config) error {
	wmts100base := ows.WMTS100Base

	err := mergo.Merge(&cfg.Services.WMTS100Config.Wmts100, wmts100base, mergo.WithTransformers(ows.WMTS100Transfomer{}))
	if err != nil {
		return err
	}

	buf, err := buildCapabilities(cfg.Services.WMTS100Config.Wmts100, cfg.Global)
	if err != nil {
		return err
	}

	//validate.ValidateCapabilities(&cfg, buf, cfg.Services.WMTS100Config.Wmts100.SchemaLocation)

	writeFile(cfg.Services.WMTS100Config.Filename, buf)

	return nil
}

func buildWCS2_0_1(cfg config.Config) error {
	wcs201base := ows.WCS201Base

	err := mergo.Merge(&cfg.Services.WCS201Config.Wcs201, wcs201base)
	if err != nil {
		return err
	}

	buf, err := buildCapabilities(cfg.Services.WCS201Config.Wcs201, cfg.Global)
	if err != nil {
		return err
	}

	err = validate.ValidateCapabilities(&cfg, buf, cfg.Services.WCS201Config.Wcs201.SchemaLocation)
	if err != nil {
		return err
	}

	writeFile(cfg.Services.WCS201Config.Filename, buf)

	return nil
}

func main() {
	serviceconfigpath := flag.String("c", envString("SERVICECONFIG", ""), "path of the service configuration")
	flag.Parse()

	serviceconfig, err := os.ReadFile(*serviceconfigpath)
	if err != nil {
		log.Fatalf("error: %v, with file: %v", err, *serviceconfigpath)
	}

	var cfg config.Config
	if err = yaml.Unmarshal(serviceconfig, &cfg); err != nil {
		log.Fatalf("error: %v", err)
	}

	if cfg.Services.WFS200Config.Filename != "" {
		if err := buildWFS2_0_0(cfg); err != nil {
			log.Fatalf("error: %v", err)
		}
	}

	if cfg.Services.WMS130Config.Filename != "" {
		if err := buildWMS1_3_0(cfg); err != nil {
			log.Fatalf("error: %v", err)
		}
	}

	if cfg.Services.WMTS100Config.Filename != "" {
		if err := buildWMTS1_0_0(cfg); err != nil {
			log.Fatalf("error: %v", err)
		}
	}

	if cfg.Services.WCS201Config.Filename != "" {
		if err := buildWCS2_0_1(cfg); err != nil {
			log.Fatalf("error: %v", err)
		}
	}
}
