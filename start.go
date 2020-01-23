package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"pdok-capabilities-gen/config"
	"pdok-capabilities-gen/wcs201"
	"pdok-capabilities-gen/wfs200"
	"pdok-capabilities-gen/wms130"
	"pdok-capabilities-gen/wmts100"
	"regexp"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"
)

func envString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func writeFile(filename string, buffer []byte) {
	err := ioutil.WriteFile(filename, buffer, 0777)
	if err != nil {
		log.Fatalf("Could not write to file %s : %v ", filename, err)
	}
}

func buildCapabilities(v interface{}, g config.Global) []byte {
	si, _ := xml.MarshalIndent(v, "", " ")
	t := template.Must(template.New("capabilities").Parse(string(si)))
	buf := &bytes.Buffer{}
	err := t.Execute(buf, g)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(buf.String(), "/>"))
}

func buildWFS2_0_0(config config.Config) {
	// retrieve default set
	wfs200base := wfs200.GetBase()
	// merge with specific set skipping featuretypelist, this is a custom operation
	mergo.Merge(&config.Services.WFS200Config.Wfs200, wfs200base, mergo.WithTransformers(wfs200.Wfs201Transfomer{}))

	// can we apply generic base feature template to config ?
	if len(wfs200base.FeatureTypeList.FeatureType) > 0 {
		for index := range config.Services.WFS200Config.Wfs200.FeatureTypeList.FeatureType {
			mergo.Merge(&config.Services.WFS200Config.Wfs200.FeatureTypeList.FeatureType[index], wfs200base.FeatureTypeList.FeatureType[0])
		}
	}

	if &config.Services.WFS200Config.Wfs200.OperationsMetadata.ExtendedCapabilities != nil {
		config.Services.WFS200Config.Wfs200.Namespaces.XmlnsInspireCommon = "http://inspire.ec.europa.eu/schemas/common/1.0"
		config.Services.WFS200Config.Wfs200.Namespaces.XmlnsInspireDls = "http://inspire.ec.europa.eu/schemas/inspire_dls/1.0"
	}

	buf := buildCapabilities(config.Services.WFS200Config.Wfs200, config.Global)
	writeFile(config.Services.WFS200Config.Filename, buf)
}

func buildWMS1_3_0(config config.Config) {
	wms130base := wms130.GetBase()

	// merge with specific set skipping layer, this is a custom operation
	mergo.Merge(&config.Services.WMS130Config.Wms130, wms130base, mergo.WithTransformers(wms130.Wms130Transfomer{}))

	if len(wms130base.Capability.Layer) > 0 {
		for index := range config.Services.WMS130Config.Wms130.Capability.Layer {
			merge(&config.Services.WMS130Config.Wms130.Capability.Layer[index], wms130base.Capability.Layer[0])
		}
	}

	if &config.Services.WMS130Config.Wms130.Capability.ExtendedCapabilities != nil {
		config.Services.WMS130Config.Wms130.Namespaces.XmlnsInspireCommon = "http://inspire.ec.europa.eu/schemas/common/1.0"
		config.Services.WMS130Config.Wms130.Namespaces.XmlnsInspireVs = "http://inspire.ec.europa.eu/schemas/inspire_vs/1.0"
		config.Services.WMS130Config.Wms130.Namespaces.SchemaLocation = wms130base.Namespaces.SchemaLocation + " " + "http://inspire.ec.europa.eu/schemas/inspire_vs/1.0 http://inspire.ec.europa.eu/schemas/inspire_vs/1.0/inspire_vs.xsd"
	}

	buf := buildCapabilities(config.Services.WMS130Config.Wms130, config.Global)
	writeFile(config.Services.WMS130Config.Filename, buf)
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

func buildWMTS1_0_0(config config.Config) {
	wmts100base := wmts100.GetBase()

	mergo.Merge(&config.Services.WMTS100Config.Wmts100, wmts100base, mergo.WithTransformers(wmts100.Wmts100Transfomer{}))

	// Cleanup unused TileMatrixSets
	var tms []wmts100.TileMatrixSet
	for i, t := range config.Services.WMTS100Config.Wmts100.Contents.TileMatrixSet {
		if !config.Services.WMTS100Config.Wmts100.Contents.GetTilematrixsets()[t.Identifier] {
			tms = append(config.Services.WMTS100Config.Wmts100.Contents.TileMatrixSet[:i], wmts100base.Contents.TileMatrixSet[i+1:]...)
		}
	}
	config.Services.WMTS100Config.Wmts100.Contents.TileMatrixSet = tms

	buf := buildCapabilities(config.Services.WMTS100Config.Wmts100, config.Global)
	writeFile(config.Services.WMTS100Config.Filename, buf)
}

func buildWCS2_0_1(config config.Config) {
	wcs201base := wcs201.GetBase()

	mergo.Merge(&config.Services.WCS201Config.Wcs201, wcs201base)

	buf := buildCapabilities(config.Services.WCS201Config.Wcs201, config.Global)
	writeFile(config.Services.WCS201Config.Filename, buf)
}

func main() {
	serviceconfigpath := flag.String("c", envString("SERVICECONFIG", ""), "path of the service configuration")
	flag.Parse()

	serviceconfig, err := ioutil.ReadFile(*serviceconfigpath)
	if err != nil {
		log.Fatalf("error: %v, with file: %v", err, *serviceconfigpath)
	}

	var config config.Config
	if err = yaml.Unmarshal(serviceconfig, &config); err != nil {
		log.Fatalf("error: %v", err)
	}

	if config.Services.WFS200Config.Filename != "" {
		buildWFS2_0_0(config)
	}

	if config.Services.WMS130Config.Filename != "" {
		buildWMS1_3_0(config)
	}

	if config.Services.WMTS100Config.Filename != "" {
		buildWMTS1_0_0(config)
	}

	if config.Services.WCS201Config.Filename != "" {
		buildWCS2_0_1(config)
	}
}
