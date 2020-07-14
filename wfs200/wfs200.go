package wfs200

import (
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"

	wfs200_response "github.com/pdok/ogc-specifications/pkg/wfs200/response"
)

// GetBase function to get a filled "template" based on the wfs200.yaml config
func GetBase() wfs200_response.GetCapabilities {
	wfs200 := wfs200_response.GetCapabilities{}
	base, err := ioutil.ReadFile("./base/wfs200.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wfs200); err != nil {
		log.Fatalf("error: %v", err)
	}
	return wfs200
}

// Wfs201Transfomer struct
type Wfs201Transfomer struct {
}

// Transformer skip FeatureTypeList when merging Base to config, this is a custom operation
func (t Wfs201Transfomer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(wfs200_response.FeatureTypeList{}) {
		return func(dst, src reflect.Value) error {
			// NOOP
			return nil
		}
	}
	return nil
}
