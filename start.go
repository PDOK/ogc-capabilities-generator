package main

import (
	"bytes"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"pdok-capabilities-gen/builder"
	"pdok-capabilities-gen/util"
	"strings"
)

var wfs = util.GetTemplates("templates/wfs/*")
var wms = util.GetTemplates("templates/wms/*")
var wmts = util.GetTemplates("templates/wmts/*")

func main() {

	typePtr := flag.String("service-type", envString("SERVICE_TYPE", ""), "wfs or wms or wmts")
	versionPtr := flag.String("service-version", envString("SERVICE_VERSION", ""), "wfs[1.1.0, 2.0.0] or wms[1.3.0] or wmts[1.0.0]")
	serviceConfigPathPtr := flag.String("service-config", envString("SERVICE_CONFIG_PATH", ""), "location of the service config")
	serviceDefConfigPathPtr := flag.String("service-def-config", envString("SERVICE_DEF_CONFIG_PATH", "config/serviceDef.yaml"), "location of the service definition config")
	outputCapabilitiesPtr := flag.String("service-output-path", envString("SERVICE_CAPABILITIES_PATH", ""), "location of service config")
	flag.Parse()

	if err := validate(*typePtr, *versionPtr); err != nil {
		log.Fatalf("error: %v", err)
	}

	serviceConfig, err := ioutil.ReadFile(*serviceConfigPathPtr)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	service := builder.Service{}
	if err = yaml.Unmarshal(serviceConfig, &service); err != nil {
		log.Fatalf("error: %v", err)
	}

	serviceDef := getServiceDef(*serviceDefConfigPathPtr)
	capabilitiesBuilder := createBuilder(serviceDef, service, *typePtr, *versionPtr)
	if capabilitiesBuilder == nil {
		log.Fatalf("Could not create capabilitiesBuilder with %s[%s]", *typePtr, *versionPtr)
	}

	var writer bytes.Buffer
	err = capabilitiesBuilder.Build(&writer)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	writeFile(*outputCapabilitiesPtr, writer.Bytes())

}

func createBuilder(serviceDef builder.ServiceDef, service builder.Service, serviceType, serviceVersion string) builder.Builder {

	if serviceType == "wfs" {
		return builder.NewWfsCapabilities(wfs, serviceDef, serviceVersion).
			AddServiceIdentification(serviceDef.Identification[service.Identification]).
			AddDataset(serviceDef.Datasets[service.Dataset]).
			AddServiceProvider(serviceDef.Organizations[service.Organization]).
			AddFeatures(serviceDef.Datasets[service.Dataset], service.Features, serviceDef).
			AddOperationsMetadata(serviceDef.Datasets[service.Dataset], service)
	}

	if serviceType == "wms" {
		return builder.NewWmsCapabilities(wms, serviceDef, serviceVersion).
			AddDataset(serviceDef.Datasets[service.Dataset]).
			AddServiceProvider(serviceDef.Organizations[service.Organization], serviceDef.Identification[service.Identification]).
			AddLayers(serviceDef.Datasets[service.Dataset], service.Layers, serviceDef, service).
			AddCapabilityRequest(serviceDef.Datasets[service.Dataset], service).AddInspireCommon(service)
	}

	if serviceType == "wmts" {
		return builder.NewWmtsCapabilities(wmts, serviceDef, serviceVersion).
			AddDataset(serviceDef.Datasets[service.Dataset]).
			AddLayers(serviceDef.Datasets[service.Dataset], service.Layers, serviceDef, service)
	}

	return nil
}

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

func validate(typeStr, versionStr string) error {
	tsl := strings.ToLower(typeStr)
	serviceMap := getServiceMap()

	versions, ok := serviceMap[tsl]
	if !ok {
		return fmt.Errorf("service type unknown : '%s', possible values are : wfs, wms, wmts", typeStr)
	}

	vpl := strings.ToLower(versionStr)
	if !contains(vpl, versions) {
		return fmt.Errorf("version unknown for %s['%s'], possible versions are %s", typeStr, versionStr, arrayToString(serviceMap[typeStr], ", "))
	}
	return nil
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
