package main

import (
	"encoding/xml"
	capabilities2 "github.com/pdok/ogc-specifications/pkg/wmts100/capabilities"
	wmts100_response "github.com/pdok/ogc-specifications/pkg/wmts100/response"
	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
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

func TestBuildCapabilitiesWmtsValidate(t *testing.T) {

	g := config.Global{Namespace: "namespace", Prefix: "prefix", Onlineresourceurl: "onlineresourceurl", Path: "path", Version: "version"}

	layer := capabilities2.Layer{}
	contents := capabilities2.Contents{Layer: []capabilities2.Layer{layer}}
	namespaces := wmts100_response.Namespaces{"http://www.opengis.net/wmts/1.0", "http://www.opengis.net/ows/1.1", "http://www.w3.org/1999/xlink", "http://www.w3.org/2001/XMLSchema-instance", "http://www.opengis.net/gml", "1.0.0", "http://www.opengis.net/wmts/1.0 http://schemas.opengis.net/wmts/1.0/wmtsGetCapabilities_response.xsd"}
	capabilities := wmts100_response.GetCapabilities{Contents: contents, Namespaces: namespaces}
	//wmts100Config := config.WMTS100Config{"./output/test-wmts100.xml", capabilities}

	buf, _ := buildCapabilities(capabilities, g)

	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()
	xsdhandler, err := xsdvalidate.NewXsdHandlerUrl("http://schemas.opengis.net/wmts/1.0/wmtsGetCapabilities_response.xsd", xsdvalidate.ParsErrDefault)
	if err != nil {
		panic(err)
	}
	defer xsdhandler.Free()

	err = xsdhandler.ValidateMem(buf, xsdvalidate.ValidErrDefault)

	if err != nil {
		t.Errorf(err.Error())
	}

	//if expectedresult != string(buf) {
	//	t.Errorf("Expected %s but was not, got: %s", expectedresult, string(buf))
	//}
}
