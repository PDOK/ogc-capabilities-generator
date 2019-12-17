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
	"pdok-capabilities-gen/wfs200"
	"pdok-capabilities-gen/wms130"
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

func writeCapabilitiesFile(filename string, v interface{}, g config.Global) {
	si, _ := xml.MarshalIndent(v, "", " ")
	t := template.Must(template.New("capabilities").Parse(string(si)))
	buf := new(bytes.Buffer)
	err := t.Execute(buf, g)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	re := regexp.MustCompile(`><.*>`)
	err = ioutil.WriteFile(filename, []byte(xml.Header+re.ReplaceAllString(buf.String(), "/>")), 0777)
	if err != nil {
		log.Fatalf("Could not write to file %s : %v ", filename, err)
	}
}

func buildWMS1_3_0(config config.Config) {
	wms130 := wms130.GetBase()

	mergo.Merge(&wms130.Service, &config.Services.Wms130.Service)
	mergo.Merge(&wms130.Capability, &config.Services.Wms130.Capability)

	writeCapabilitiesFile("wms130.xml", wms130, config.Global)
}

func buildWFS2_0_0(config config.Config) {
	wfs200 := wfs200.GetBase()

	mergo.Merge(&wfs200.ServiceIdentification, &config.Services.Wfs200.ServiceIdentification)
	mergo.Merge(&wfs200.FeatureTypeList, &config.Services.Wfs200.FeatureTypeList)

	if &config.Services.Wfs200.ExtendedCapabilities != nil {
		wfs200.WFS_Namespaces.XmlnsInspireCommon = "http://inspire.ec.europa.eu/schemas/common/1.0"
		wfs200.WFS_Namespaces.XmlnsInspireDls = "http://inspire.ec.europa.eu/schemas/inspire_dls/1.0"
		wfs200.OperationsMetadata.ExtendedCapabilities = &config.Services.Wfs200.ExtendedCapabilities
	}

	writeCapabilitiesFile("wfs200.xml", wfs200, config.Global)
}

func main() {
	serviceconfigpath := flag.String("c", envString("SERVICECONFIG", ""), "path of the service configuration")
	flag.Parse()

	serviceconfig, err := ioutil.ReadFile(*serviceconfigpath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var config config.Config
	if err = yaml.Unmarshal(serviceconfig, &config); err != nil {
		log.Fatalf("error: %v", err)
	}

	if &config.Services.Wfs200 != nil {
		buildWFS2_0_0(config)
	}

	if &config.Services.Wms130 != nil {
		buildWMS1_3_0(config)
	}
}
