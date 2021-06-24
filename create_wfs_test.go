package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"ogc-capabilities-generator/config"
	"ogc-capabilities-generator/ows"
	"testing"
)

var wfs200base = ows.WFS200Base

func TestBuildWFSTest(t *testing.T) {
	var serviceconfig, err = ioutil.ReadFile("./test/input/wfs-hh.yaml")
	if err != nil {
		log.Fatalf("error: %v, with file: wfs-hh.yaml", err)
	}

	var config config.Config
	if err := yaml.Unmarshal(serviceconfig, &config); err != nil {
		log.Fatalf("error: %v", err)
	}

	buildWFS2_0_0(config)

	testResultPath := "./test/output/wfs200_hh.xml"
	expectedresultPath := "./test/output/expected/wfs200_hh.xml"

	diff := compareXML(testResultPath, expectedresultPath)
	if len(diff) > 0 {
		t.Errorf("unexpected differences found between actual (%s) and expected (%s) output: %s", testResultPath, expectedresultPath, writeDiff(diff))
	}
}
