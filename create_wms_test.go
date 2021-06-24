package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"ogc-capabilities-generator/config"
	"testing"
)

func TestBuildWMSTest(t *testing.T) {
	var serviceconfig, err = ioutil.ReadFile("./test/input/wms-hh.yaml")
	if err != nil {
		log.Fatalf("error: %v, with file: wfs-hh.yaml", err)
	}

	var config config.Config
	if err := yaml.Unmarshal(serviceconfig, &config); err != nil {
		log.Fatalf("error: %v", err)
	}

	buildWMS1_3_0(config)

	testResultPath := "./test/output/wms130_hh.xml"
	expectedResultPath := "./test/output/expected/wms130_hh.xml"

	diff := compareXML(testResultPath, expectedResultPath)
	if len(diff) > 0 {
		t.Errorf("unexpected differences found between actual (%s) and expected (%s) output: %s", testResultPath, expectedResultPath, writeDiff(diff))
	}
}
