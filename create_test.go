package main

import (
	"encoding/xml"
	"pdok-capabilities-gen/config"
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

	buf := buildCapabilities(d, g)

	if expectedresult != string(buf) {
		t.Errorf("Expected %s but was not, got: %s", expectedresult, string(buf))
	}
}

func TestBuildCapabilitiesMissingEmpty(t *testing.T) {

	g := config.Global{Namespace: "namespace", Prefix: "prefix", Onlineresourceurl: "onlineresourceurl", Path: "path", Version: "version"}
	d := dummy{Namespace: "{{.Namespace}}", Prefix: "{{.Prefix}}", Onlineresourceurl: "{{.Onlineresourceurl}}", Path: "{{.Path}}", Version: "{{.Version}}"}

	buf := buildCapabilities(d, g)

	result := string(buf)

	if expectedresult == result || strings.Contains(result, "<empty/>") {
		t.Errorf("Expected <empty/> but was not, got nothing")
	}
}
