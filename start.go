package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"pdok-capabilities-gen/builder"
	"pdok-capabilities-gen/util"

	"gopkg.in/yaml.v2"
)

var wfs = util.GetTemplates("templates/wfs/*")
var wms = util.GetTemplates("templates/wms/*")
var wmts = util.GetTemplates("templates/wmts/*")

// func createBuilder(service builder.Service, onlineResourceURL string) builder.Builder {

// 	if serviceType == "wfs" {
// 		return builder.NewWfsCapabilities(wfs, serviceDef, serviceVersion).
// 			AddServiceIdentification(serviceDef.Identification[service.Identification]).
// 			AddDataset(serviceDef.Datasets[service.Dataset]).
// 			AddServiceProvider(serviceDef.Organizations[service.Organization]).
// 			AddFeatures(serviceDef.Datasets[service.Dataset], service.Features, serviceDef).
// 			AddOperationsMetadata(serviceDef.Datasets[service.Dataset], service)
// 	}

// 	if serviceType == "wms" {
// 		return builder.NewWmsCapabilities(wms, serviceDef, serviceVersion).
// 			AddDataset(serviceDef.Datasets[service.Dataset]).
// 			AddServiceProvider(serviceDef.Organizations[service.Organization], serviceDef.Identification[service.Identification]).
// 			AddLayers(serviceDef.Datasets[service.Dataset], service.Layers, serviceDef, service).
// 			AddCapabilityRequest(serviceDef.Datasets[service.Dataset], service).AddInspireCommon(service)
// 	}

// 	if serviceType == "wmts" {
// 		return builder.NewWmtsCapabilities(wmts, serviceDef, serviceVersion, onlineResourceURL).
// 			AddDataset(serviceDef.Datasets[service.Dataset]).
// 			AddLayers(serviceDef.Datasets[service.Dataset], service.Layers, serviceDef, service)
// 	}

// 	return nil
// }

func writeFile(name string, data []byte) {
	err := ioutil.WriteFile(name, data, 0777)
	if err != nil {
		log.Fatalf("Could not write to file %s : %v ", name, err)
	}
}

func getServiceDef(serviceDefConfigPath string) builder.ServiceDef {
	serviceDefConfig, err := ioutil.ReadFile(serviceDefConfigPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	serviceDef := builder.ServiceDef{}
	if err = yaml.Unmarshal(serviceDefConfig, &serviceDef); err != nil {
		log.Fatalf("error: %v", err)
	}
	return serviceDef
}

func envString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func getServiceMap() map[string][]string {
	serviceMap := map[string][]string{
		"wms":  {"1.3.0"},
		"wmts": {"1.0.0"},
		"wfs":  {"1.1.0", "2.0.0"},
	}
	return serviceMap
}

func contains(value string, list []string) bool {
	for _, version := range list {
		if version == value {
			return true
		}
	}
	return false
}

func arrayToString(aList []string, delim string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(aList); i++ {
		buffer.WriteString(aList[i])
		if i != len(aList)-1 {
			buffer.WriteString(delim)
		}
	}

	return buffer.String()
}

func main() {

	serviceConfigPath := flag.String("c", envString("SERVICE_CONFIG_PATH", ""), "location of the service config")
	// outputCapabilitiesPtr := flag.String("service-output-path", envString("SERVICE_CAPABILITIES_PATH", ""), "location of service config")
	// onlineResourceURL := flag.String("service-online-resource", envString("SERVICE_CAPABILITIES_ONLINERESOURCE", ""), "onlineresource URL used in documents")
	flag.Parse()

	serviceConfig, err := ioutil.ReadFile(*serviceConfigPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//service := builder.Service{}
	service := builder.Config{}
	if err = yaml.Unmarshal(serviceConfig, &service); err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(service)

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
