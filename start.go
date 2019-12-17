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

func writeFile(name string, data []byte) {
	err := ioutil.WriteFile(name, data, 0777)
	if err != nil {
		log.Fatalf("Could not write to file %s : %v ", name, err)
	}
}

func envString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func buildWMS1_3_0(config config.Config) {
	basewms1_3_0 := wms130.Wms130{}
	base, err := ioutil.ReadFile("./wms130/wms130.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &basewms1_3_0); err != nil {
		log.Fatalf("error: %v", err)
	}

	mergo.Merge(&basewms1_3_0.Service, &config.Services.Wms130.Service)
	mergo.Merge(&basewms1_3_0.Capability, &config.Services.Wms130.Capability)

	// if &config.Services.WFS_2_0_0.ExtendedCapabilities != nil {
	// 	basewms1_3_0.WFS_Namespaces.XmlnsInspireCommon = "http://inspire.ec.europa.eu/schemas/common/1.0"
	// 	basewms1_3_0.WFS_Namespaces.XmlnsInspireVs = "http://inspire.ec.europa.eu/schemas/inspire_vs/1.0"
	// 	basewms1_3_0.OperationsMetadata.ExtendedCapabilities = &config.Services.WMS_1_3_0.ExtendedCapabilities
	// }

	si, _ := xml.MarshalIndent(basewms1_3_0, "", " ")
	t := template.Must(template.New("capabilities").Parse(string(si)))
	buf := new(bytes.Buffer)
	err = t.Execute(buf, config.Global)

	re := regexp.MustCompile(`><.*>`)
	writeFile("wms_1_3_0.xml", []byte(xml.Header+re.ReplaceAllString(buf.String(), "/>")))
}

func buildWFS2_0_0(config config.Config) {
	basewfs2_0_0 := wfs200.Wfs200{}
	base, err := ioutil.ReadFile("./wfs200/wfs200.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &basewfs2_0_0); err != nil {
		log.Fatalf("error: %v", err)
	}

	mergo.Merge(&basewfs2_0_0.ServiceIdentification, &config.Services.Wfs200.ServiceIdentification)
	mergo.Merge(&basewfs2_0_0.FeatureTypeList, &config.Services.Wfs200.FeatureTypeList)

	if &config.Services.Wfs200.ExtendedCapabilities != nil {
		basewfs2_0_0.WFS_Namespaces.XmlnsInspireCommon = "http://inspire.ec.europa.eu/schemas/common/1.0"
		basewfs2_0_0.WFS_Namespaces.XmlnsInspireDls = "http://inspire.ec.europa.eu/schemas/inspire_dls/1.0"
		basewfs2_0_0.OperationsMetadata.ExtendedCapabilities = &config.Services.Wfs200.ExtendedCapabilities
	}

	si, _ := xml.MarshalIndent(basewfs2_0_0, "", " ")
	t := template.Must(template.New("capabilities").Parse(string(si)))
	buf := new(bytes.Buffer)
	err = t.Execute(buf, config.Global)

	re := regexp.MustCompile(`><.*>`)
	writeFile("wfs_2_0_0.xml", []byte(xml.Header+re.ReplaceAllString(buf.String(), "/>")))
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
