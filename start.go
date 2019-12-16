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

func main() {

	serviceconfigpath := flag.String("c", envString("SERVICECONFIG", ""), "path of the service configuration")
	flag.Parse()

	serviceconfig, err := ioutil.ReadFile(*serviceconfigpath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	config := builder.Config{}
	if err = yaml.Unmarshal(serviceconfig, &config); err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, w := range config.Services.WFS {

		if w.Version == "2.0.0" {

			basewfs2_0_0 := builder.WFS_2_0_0{}
			base, err := ioutil.ReadFile("./base/wfs_2_0_0.yaml")
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			if err = yaml.Unmarshal(base, &basewfs2_0_0); err != nil {
				log.Fatalf("error: %v", err)
			}

			mergo.Merge(&basewfs2_0_0.ServiceIdentification, &w.ServiceIdentification)
			mergo.Merge(&basewfs2_0_0.FeatureTypeList, &w.FeatureTypeList)

			// si, _ := xml.MarshalIndent(basewfs2_0_0, "", " ")
			// re := regexp.MustCompile(`><.*>`)
			// t := template.Must(template.New("capabilities").Parse(re.ReplaceAllString(xml.Header+string(si), "/>")))
			// buf := new(bytes.Buffer)
			// err = t.Execute(buf, config.Global)

			si, _ := xml.MarshalIndent(basewfs2_0_0, "", " ")
			t := template.Must(template.New("capabilities").Parse(string(si)))
			buf := new(bytes.Buffer)
			err = t.Execute(buf, config.Global)

			re := regexp.MustCompile(`><.*>`)

			//fmt.Println(buf.String())

			writeFile("wfs_2_0_0.xml", []byte(xml.Header+re.ReplaceAllString(buf.String(), "/>")))

		}
	}

	// for _, wms := range service.Service.WMS {
	// 	fmt.Println(wms)
	// }

	// capabilitiesBuilder := createBuilder(service, *onlineResourceURL)
	// if capabilitiesBuilder == nil {
	// 	log.Fatalf("Could not create capabilitiesBuilder")
	// }

	// var writer bytes.Buffer
	// err = capabilitiesBuilder.Build(&writer)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// writeFile(*outputCapabilitiesPtr, writer.Bytes())

}
