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
	wfs200base := wfs200.GetBase()

	mergo.Merge(&wfs200base.ServiceIdentification, &config.Services.Wfs200.ServiceIdentification)
	mergo.Merge(&wfs200base.FeatureTypeList, &config.Services.Wfs200.FeatureTypeList)

	if &config.Services.Wfs200.ExtendedCapabilities != nil {
		wfs200base.Namespaces.XmlnsInspireCommon = "http://inspire.ec.europa.eu/schemas/common/1.0"
		wfs200base.Namespaces.XmlnsInspireDls = "http://inspire.ec.europa.eu/schemas/inspire_dls/1.0"
		wfs200base.OperationsMetadata.ExtendedCapabilities = &config.Services.Wfs200.ExtendedCapabilities
	}

	buf := buildCapabilities(wfs200base, config.Global)
	writeFile(config.Services.Wfs200.Filename, buf)
}

func buildWMS1_3_0(config config.Config) {
	wms130base := wms130.GetBase()

	mergo.Merge(&wms130base.Service, &config.Services.Wms130.Service)
	mergo.Merge(&wms130base.Capability, &config.Services.Wms130.Capability)

	if &config.Services.Wms130.Capability.ExtendedCapabilities != nil {
		wms130base.Namespaces.XmlnsInspireCommon = "http://inspire.ec.europa.eu/schemas/common/1.0"
		wms130base.Namespaces.XmlnsInspireVs = "http://inspire.ec.europa.eu/schemas/inspire_vs/1.0"
		wms130base.Namespaces.SchemaLocation = wms130base.Namespaces.SchemaLocation + " " + "http://inspire.ec.europa.eu/schemas/inspire_vs/1.0 http://inspire.ec.europa.eu/schemas/inspire_vs/1.0/inspire_vs.xsd"
	}

	buf := buildCapabilities(wms130base, config.Global)
	writeFile(config.Services.Wms130.Filename, buf)
}

func buildWMTS1_0_0(config config.Config) {
	wmts100base := wmts100.GetBase()

	mergo.Merge(&wmts100base.ServiceIdentification, &config.Services.Wmts100.ServiceIdentification)
	mergo.Merge(&wmts100base.Contents, &config.Services.Wmts100.Contents)

	// Cleanup unused TileMatrixSets
	var tms []wmts100.TileMatrixSet
	for i, t := range wmts100base.Contents.TileMatrixSet {
		if !config.Services.Wmts100.Contents.GetTilematrixsets()[t.Identifier] {
			tms = append(wmts100base.Contents.TileMatrixSet[:i], wmts100base.Contents.TileMatrixSet[i+1:]...)
		}
	}
	wmts100base.Contents.TileMatrixSet = tms

	mergo.Merge(&wmts100base.ServiceMetadataURL, &config.Services.Wmts100.ServiceMetadataURL)

	buf := buildCapabilities(wmts100base, config.Global)
	writeFile(config.Services.Wmts100.Filename, buf)
}

func buildWCS2_0_1(config config.Config) {
	wcs201 := wcs201.GetBase()

	mergo.Merge(&wcs201.ServiceIdentification, &config.Services.Wcs201.ServiceIdentification)
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

	if config.Services.Wfs200.Filename != "" {
		buildWFS2_0_0(config)
	}

	if config.Services.Wms130.Filename != "" {
		buildWMS1_3_0(config)
	}

	if config.Services.Wmts100.Filename != "" {
		buildWMTS1_0_0(config)
	}

	if config.Services.Wcs201.Filename != "" {
		buildWCS2_0_1(config)
	}
}
