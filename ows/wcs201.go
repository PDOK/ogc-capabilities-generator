package ows

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/pdok/ogc-specifications/pkg/wcs201"
)

// WCS201Base is the base WCS 2.0.1 GetCapabilities doc
var WCS201Base wcs201.GetCapabilitiesResponse

func init() {
	wcs201Response := wcs201.GetCapabilitiesResponse{}
	base, err := ioutil.ReadFile("./base/wcs201.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wcs201Response); err != nil {
		log.Fatalf("error: %v", err)
	}
	WCS201Base = wcs201Response
}
