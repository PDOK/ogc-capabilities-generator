package main

import (
	"encoding/xml"
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

	buf, _ := buildCapabilities(d, g, false)

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

	buf, _ := buildCapabilities(d, g, false)

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

	buf, _ := buildCapabilities(d, g, false)

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

	_, err := buildCapabilities(d, g, true)

	if err == nil {
		t.Errorf("Expected %s but was not, got: %s", expected, err.Error())
	}

	if err != nil {
		if err.Error() != "no dataset prefix defined" {
			t.Errorf("Expected %s but was not, got: %s", expected, err.Error())
		}
	}
}

func TestBuildCapabilitiesMissingEmpty(t *testing.T) {

	g := config.Global{Namespace: "namespace", Prefix: "prefix", Onlineresourceurl: "onlineresourceurl", Path: "path", Version: "version"}
	d := dummy{Namespace: "{{.Namespace}}", Prefix: "{{.Prefix}}", Onlineresourceurl: "{{.Onlineresourceurl}}", Path: "{{.Path}}", Version: "{{.Version}}"}

	buf, _ := buildCapabilities(d, g, false)

	result := string(buf)

	if expectedresult == result || strings.Contains(result, "<empty/>") {
		t.Errorf("Expected <empty/> but was not, got nothing")
	}
}
