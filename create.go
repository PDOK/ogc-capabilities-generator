package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"ogc-capabilities-generator/config"
	"ogc-capabilities-generator/ows"
	"ogc-capabilities-generator/validate"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"

	wms130 "github.com/pdok/ogc-specifications/pkg/wms130"
)

func envString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func writeFile(filename string, buffer []byte) {
	makeDirIfNotExists(filename)
	err := ioutil.WriteFile(filename, buffer, 0777)
	if err != nil {
		log.Fatalf("Could not write to file %s : %v ", filename, err)
	}
}

func buildCapabilities(v interface{}, g config.Global) ([]byte, error) {

	if g.Prefix == `` {
		return nil, errors.New("no dataset prefix defined")
	}

	si, _ := xml.MarshalIndent(v, "", " ")
	t := template.Must(template.New("capabilities").Parse(string(si)))
	buf := &bytes.Buffer{}
	err := t.Execute(buf, g)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(buf.String(), "/>")), nil
}

func buildWFS2_0_0(config config.Config) error {

	// retrieve default set
	wfs200base := ows.WFS200Base
	// merge with specific set skipping featuretypelist, this is a custom operation
	mergo.Merge(&config.Services.WFS200Config.Wfs200, wfs200base, mergo.WithTransformers(ows.WFS200Transfomer{}))

	// can we apply generic base feature template to config ?
	if len(wfs200base.Capabilities.FeatureTypeList.FeatureType) > 0 {
		for index := range config.Services.WFS200Config.Wfs200.FeatureTypeList.FeatureType {
			mergo.Merge(&config.Services.WFS200Config.Wfs200.FeatureTypeList.FeatureType[index], wfs200base.Capabilities.FeatureTypeList.FeatureType[0])
		}
	}

	if config.Services.WFS200Config.Wfs200.Capabilities.OperationsMetadata.ExtendedCapabilities != nil {
		config.Services.WFS200Config.Wfs200.Namespaces.XmlnsInspireCommon = "http://inspire.ec.europa.eu/schemas/common/1.0"
		config.Services.WFS200Config.Wfs200.Namespaces.XmlnsInspireDls = "http://inspire.ec.europa.eu/schemas/inspire_dls/1.0"
	}

	buf, err := buildCapabilities(config.Services.WFS200Config.Wfs200, config.Global)
	if err != nil {
		return err
	}

	validate.ValidateCapabilities(&config, buf, config.Services.WFS200Config.Wfs200.SchemaLocation)

	writeFile(config.Services.WFS200Config.Filename, buf)

	return nil
}

func buildWMS1_3_0(config config.Config) error {
	wms130base := ows.WMS130Base

	// merge with specific set skipping layer, this is a custom operation
	mergo.Merge(&config.Services.WMS130Config.Wms130, wms130base, mergo.WithTransformers(ows.WMS130Transfomer{}))

	if len(wms130base.Capabilities.Layer) > 0 {
		for index := range config.Services.WMS130Config.Wms130.Capabilities.Layer {
			merge(&config.Services.WMS130Config.Wms130.Capabilities.Layer[index], wms130base.Capabilities.Layer[0])
		}
	}

	if config.Services.WMS130Config.Wms130.Capabilities.WMSCapabilities.ExtendedCapabilities != nil {
		config.Services.WMS130Config.Wms130.Namespaces.XmlnsInspireCommon = "http://inspire.ec.europa.eu/schemas/common/1.0"
		config.Services.WMS130Config.Wms130.Namespaces.XmlnsInspireVs = "http://inspire.ec.europa.eu/schemas/inspire_vs/1.0"
		config.Services.WMS130Config.Wms130.Namespaces.SchemaLocation = wms130base.Namespaces.SchemaLocation + " " + "http://inspire.ec.europa.eu/schemas/inspire_vs/1.0 http://inspire.ec.europa.eu/schemas/inspire_vs/1.0/inspire_vs.xsd"
	}

	buf, err := buildCapabilities(config.Services.WMS130Config.Wms130, config.Global)
	if err != nil {
		return err
	}

	validate.ValidateCapabilities(&config, buf, config.Services.WMS130Config.Wms130.SchemaLocation)

	writeFile(config.Services.WMS130Config.Filename, buf)

	return nil
}

// recursive fill
func merge(dst *wms130.Layer, src wms130.Layer) {

	if len(dst.Layer) > 0 {
		for index := range dst.Layer {
			merge(dst.Layer[index], src)
		}
	}
	mergo.Merge(dst, src)
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

func buildWMTS1_0_0(config config.Config) error {
	wmts100base := ows.WMTS100Base

	mergo.Merge(&config.Services.WMTS100Config.Wmts100, wmts100base, mergo.WithTransformers(ows.WMTS100Transfomer{}))

	buf, err := buildCapabilities(config.Services.WMTS100Config.Wmts100, config.Global)
	if err != nil {
		return err
	}

	validate.ValidateCapabilities(&config, buf, config.Services.WMTS100Config.Wmts100.SchemaLocation)

	writeFile(config.Services.WMTS100Config.Filename, buf)

	return nil
}

func buildWCS2_0_1(config config.Config) error {
	wcs201base := ows.WCS201Base

	mergo.Merge(&config.Services.WCS201Config.Wcs201, wcs201base)

	buf, err := buildCapabilities(config.Services.WCS201Config.Wcs201, config.Global)
	if err != nil {
		return err
	}

	validate.ValidateCapabilities(&config, buf, config.Services.WCS201Config.Wcs201.SchemaLocation)

	writeFile(config.Services.WCS201Config.Filename, buf)

	return nil
}

func main() {
	serviceconfigpath := flag.String("c", envString("SERVICECONFIG", ""), "path of the service configuration")
	flag.Parse()

	serviceconfig, err := ioutil.ReadFile(*serviceconfigpath)
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
