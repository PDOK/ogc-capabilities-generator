package wcs201

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	wcs201_response "github.com/pdok/ogc-specifications/pkg/wcs201/response"
)

// GetBase function to get a filled "template" based on the wmts100.yaml config
func GetBase() wcs201_response.GetCapabilities {
	wcs201 := wcs201_response.GetCapabilities{}
	base, err := ioutil.ReadFile("./base/wcs201.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wcs201); err != nil {
		log.Fatalf("error: %v", err)
	}
	return wcs201
}
