package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"pdok-capabilities-gen/builder"
	"pdok-capabilities-gen/wfs2_0_0"
	"pdok-capabilities-gen/wms1_3_0"
	"regexp"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"
)

var config builder.Config

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

func buildWMS1_3_0() {
	basewms1_3_0 := wms1_3_0.WMS_1_3_0{}
	base, err := ioutil.ReadFile("./base/wms_1_3_0.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &basewms1_3_0); err != nil {
		log.Fatalf("error: %v", err)
	}

	mergo.Merge(&basewms1_3_0.Service, &config.Services.WMS_1_3_0.Service)
	mergo.Merge(&basewms1_3_0.Capability, &config.Services.WMS_1_3_0.Capability)

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

func buildWFS2_0_0() {
	basewfs2_0_0 := wfs2_0_0.WFS_2_0_0{}
	base, err := ioutil.ReadFile("./base/wfs_2_0_0.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &basewfs2_0_0); err != nil {
		log.Fatalf("error: %v", err)
	}

	mergo.Merge(&basewfs2_0_0.ServiceIdentification, &config.Services.WFS_2_0_0.ServiceIdentification)
	mergo.Merge(&basewfs2_0_0.FeatureTypeList, &config.Services.WFS_2_0_0.FeatureTypeList)

	if &config.Services.WFS_2_0_0.ExtendedCapabilities != nil {
		basewfs2_0_0.WFS_Namespaces.XmlnsInspireCommon = "http://inspire.ec.europa.eu/schemas/common/1.0"
		basewfs2_0_0.WFS_Namespaces.XmlnsInspireDls = "http://inspire.ec.europa.eu/schemas/inspire_dls/1.0"
		basewfs2_0_0.OperationsMetadata.ExtendedCapabilities = &config.Services.WFS_2_0_0.ExtendedCapabilities
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

	if err = yaml.Unmarshal(serviceconfig, &config); err != nil {
		log.Fatalf("error: %v", err)
	}

	if &config.Services.WFS_2_0_0 != nil {
		buildWFS2_0_0()
	}

	if &config.Services.WMS_1_3_0 != nil {
		buildWMS1_3_0()
	}
}
