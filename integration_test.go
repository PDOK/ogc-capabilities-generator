package main

import (
	"bytes"
	"log"
	"ogc-capabilities-generator/config"
	"os"
	"path"
	"testing"

	"github.com/ajankovic/xdiff"
	"github.com/ajankovic/xdiff/parser"
	"gopkg.in/yaml.v3"
)

const configBasePath = "./examples/config"
const expectedBasePath = "./examples/capabilities"

func compareXML(testResult string, expectedResult string) []xdiff.Delta {
	p := parser.New()
	// Parse the XTree.
	left, err := p.ParseFile(testResult)
	if err != nil {
		log.Fatalf("unable to parse testResult %s. error: %v", testResult, err)
	}
	right, err := p.ParseFile(expectedResult)
	if err != nil {
		log.Fatalf("unable to parse expectedResult %s. error: %v", expectedResult, err)
	}
	diff, err := xdiff.Compare(left, right)
	if err != nil {
		log.Fatalf("unable to compare testResult %s with expectedResult %s. error: %v", testResult, expectedResult, err)
	}
	return diff
}

func writeDiff(diff []xdiff.Delta) string {
	buf := new(bytes.Buffer)
	enc := xdiff.NewTextEncoder(buf)
	if err := enc.Encode(diff); err != nil {
		log.Fatalf("unable to write diff. error: %v", err)
	}
	return buf.String()
}

func readConfig(config_path string) (*config.Config, error) {
	configPath := path.Join(configBasePath, config_path)
	var serviceconfig, err = os.ReadFile(configPath)
	if err != nil {
		// log.Fatalf("error: %v, with file: %v", err, configPath)
		return nil, err
	}

	var config *config.Config
	if err := yaml.Unmarshal(serviceconfig, &config); err != nil {
		// log.Fatalf("error: %v", err)
		return nil, err
	}
	return config, nil
}

func TestIntegrationWMS130(t *testing.T) {
	configPath := "wms_1_3_0.yaml"
	config, err := readConfig(configPath)
	if err != nil {
		log.Fatalf("error: %v, with file: %v", err, configPath)
	}
	err = buildWMS1_3_0(*config)
	if err != nil {
		log.Fatalf("error: %v, with file: %v", err, configPath)
	}

	testResultPath := path.Join(config.Services.WMS130Config.Filename)
	expectedresultPath := path.Join(expectedBasePath, "wms_capabilities_1_3_0.xml")
	diff := compareXML(testResultPath, expectedresultPath)
	if len(diff) > 0 {
		t.Errorf("unexpected differences found between actual (%s) and expected (%s) output: %s", testResultPath, expectedresultPath, writeDiff(diff))
	}
}

func TestIntegrationWMS130Inspire(t *testing.T) {
	configPath := "wms_1_3_0_inspire.yaml"
	config, err := readConfig(configPath)
	if err != nil {
		log.Fatalf("error: %v, with file: %v", err, configPath)
	}
	err = buildWMS1_3_0(*config)
	if err != nil {
		log.Fatalf("error: %v, with file: %v", err, configPath)
	}

	testResultPath := path.Join(config.Services.WMS130Config.Filename)
	expectedresultPath := path.Join(expectedBasePath, "wms_capabilities_1_3_0_inspire.xml")
	diff := compareXML(testResultPath, expectedresultPath)
	if len(diff) > 0 {
		t.Errorf("unexpected differences found between actual (%s) and expected (%s) output: %s", testResultPath, expectedresultPath, writeDiff(diff))
	}
}

func TestIntegrationWFS200(t *testing.T) {
	configPath := "wfs_2_0_0.yaml"
	config, err := readConfig(configPath)
	if err != nil {
		log.Fatalf("error: %v, with file: %v", err, configPath)
	}
	err = buildWFS2_0_0(*config)
	if err != nil {
		log.Fatalf("error: %v, with file: %v", err, configPath)
	}
	testResultPath := path.Join(config.Services.WFS200Config.Filename)
	expectedresultPath := path.Join(expectedBasePath, "wfs_capabilities_2_0_0.xml")
	diff := compareXML(testResultPath, expectedresultPath)
	if len(diff) > 0 {
		t.Errorf("unexpected differences found between actual (%s) and expected (%s) output: %s", testResultPath, expectedresultPath, writeDiff(diff))
	}
}

func TestIntegrationWFS200Inspire(t *testing.T) {
	configPath := "wfs_2_0_0_inspire.yaml"
	config, err := readConfig(configPath)
	if err != nil {
		log.Fatalf("error: %v, with file: %v", err, configPath)
	}
	err = buildWFS2_0_0(*config)
	if err != nil {
		log.Fatalf("error: %v, with file: %v", err, configPath)
	}
	testResultPath := path.Join(config.Services.WFS200Config.Filename)
	expectedresultPath := path.Join(expectedBasePath, "wfs_capabilities_2_0_0_inspire.xml")
	diff := compareXML(testResultPath, expectedresultPath)
	if len(diff) > 0 {
		t.Errorf("unexpected differences found between actual (%s) and expected (%s) output: %s", testResultPath, expectedresultPath, writeDiff(diff))
	}
}

func TestIntegrationWMTS100(t *testing.T) {
	configPath := "wmts_1_0_0.yaml"
	config, err := readConfig(configPath)
	if err != nil {
		log.Fatalf("error: %v, with file: %v", err, configPath)
	}
	err = buildWMTS1_0_0(*config)
	if err != nil {
		log.Fatalf("error: %v, with file: %v", err, configPath)
	}
	testResultPath := path.Join(config.Services.WMTS100Config.Filename)
	expectedresultPath := path.Join(expectedBasePath, "wmts_capabilities_1_0_0.xml")
	diff := compareXML(testResultPath, expectedresultPath)
	if len(diff) > 0 {
		t.Errorf("unexpected differences found between actual (%s) and expected (%s) output: %s", testResultPath, expectedresultPath, writeDiff(diff))
	}
}
