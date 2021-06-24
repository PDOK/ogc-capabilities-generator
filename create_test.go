package main

import (
	"bytes"
	"encoding/xml"
	"github.com/ajankovic/xdiff"
	"github.com/ajankovic/xdiff/parser"
	"log"
	"ogc-capabilities-generator/config"
	"strings"
	"testing"
)

type dummy struct {
	XMLName           xml.Name `xml:"dummy"`
	Namespace         string   `xml:"namespace"`
	Prefix            string   `xml:"prefix"`
	Onlineresourceurl string   `xml:"onlineresourceurl"`
	Path              string   `xml:"path"`
	Version           string   `xml:"version"`
	Empty             string   `xml:"empty,omitempty"`
}

var expectedresult = `<?xml version="1.0" encoding="UTF-8"?>
<dummy>
 <namespace>namespace</namespace>
 <prefix>prefix</prefix>
 <onlineresourceurl>onlineresourceurl</onlineresourceurl>
 <path>path</path>
 <version>version</version>
 <empty/>
</dummy>`

func TestBuildCapabilities(t *testing.T) {

	g := config.Global{Namespace: "namespace", Prefix: "prefix", Onlineresourceurl: "onlineresourceurl", Path: "path", Version: "version"}
	d := dummy{Namespace: "{{.Namespace}}", Prefix: "{{.Prefix}}", Onlineresourceurl: "{{.Onlineresourceurl}}", Path: "{{.Path}}", Version: "{{.Version}}", Empty: "{{.Empty}}"}

	buf, _ := buildCapabilities(d, g)

	if expectedresult != string(buf) {
		t.Errorf("Expected %s but was not, got: %s", expectedresult, string(buf))
	}
}

func TestBuildPrefixCapabilities(t *testing.T) {

	type prefixtest struct {
		Prefix string `xml:"xmlns:{{.Prefix}},attr"`
	}

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<prefixtest xmlns:prefix="namespace"/>`

	g := config.Global{Namespace: "namespace", Prefix: "prefix"}
	d := prefixtest{Prefix: "{{.Namespace}}"}

	buf, _ := buildCapabilities(d, g)

	if expected != string(buf) {
		t.Errorf("Expected %s but was not, got: %s", expected, string(buf))
	}
}

func TestBuildPrefixCapabilitiesWeirdCharacter(t *testing.T) {

	type prefixtest struct {
		Prefix string `xml:"xmlns:{{.Prefix}},attr"`
	}

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<prefixtest xmlns:prefix-prefix="namespace"/>`

	g := config.Global{Namespace: "namespace", Prefix: "prefix-prefix"}
	d := prefixtest{Prefix: "{{.Namespace}}"}

	buf, _ := buildCapabilities(d, g)

	if expected != string(buf) {
		t.Errorf("Expected %s but was not, got: %s", expected, string(buf))
	}
}

func TestBuildEmptyPrefixCapabilities(t *testing.T) {

	type prefixtest struct {
		Prefix string `xml:"xmlns:{{.Prefix}},attr"`
	}

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<prefixtest xmlns:prefix="namespace"/>`

	g := config.Global{}
	d := prefixtest{Prefix: "{{.Namespace}}"}

	_, err := buildCapabilities(d, g)

	if err == nil {
		t.Errorf("Expected %s but was not, got: %s", expected, err.Error())
	}

	if err != nil {
		if err.Error() != "No dataset prefix defined" {
			t.Errorf("Expected %s but was not, got: %s", expected, err.Error())
		}
	}
}

func TestBuildCapabilitiesMissingEmpty(t *testing.T) {

	g := config.Global{Namespace: "namespace", Prefix: "prefix", Onlineresourceurl: "onlineresourceurl", Path: "path", Version: "version"}
	d := dummy{Namespace: "{{.Namespace}}", Prefix: "{{.Prefix}}", Onlineresourceurl: "{{.Onlineresourceurl}}", Path: "{{.Path}}", Version: "{{.Version}}"}

	buf, _ := buildCapabilities(d, g)

	result := string(buf)

	if expectedresult == result || strings.Contains(result, "<empty/>") {
		t.Errorf("Expected <empty/> but was not, got nothing")
	}
}

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
