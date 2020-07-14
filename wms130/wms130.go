package wms130

import (
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"

	wms130_response "github.com/pdok/ogc-specifications/pkg/wms130/response"
)

// GetBase function to get a filled "template" based on the wms130.yaml config
func GetBase() wms130_response.GetCapabilities {
	wms130 := wms130_response.GetCapabilities{}
	base, err := ioutil.ReadFile("./base/wms130.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wms130); err != nil {
		log.Fatalf("error: %v", err)
	}
	return wms130
}

// Wms130Transfomer struct
type Wms130Transfomer struct {
}

// Transformer skip FeatureTypeList when merging Base to config, this is a custom operation
func (t Wms130Transfomer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(wms130_response.Layer{}) {
		return func(dst, src reflect.Value) error {
			// NOOP
			return nil
		}
	}
	return nil
}
