package ows

import (
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"

	wms130_response "github.com/pdok/ogc-specifications/pkg/wms130/response"
)

// WMS130Base is the base WMS 1.3.0 GetCapabilities doc
var WMS130Base wms130_response.GetCapabilities

func init() {
	wms130 := wms130_response.GetCapabilities{}
	base, err := ioutil.ReadFile("./base/wms130.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wms130); err != nil {
		log.Fatalf("error: %v", err)
	}
	WMS130Base = wms130
}

// WMS130Transfomer struct
type WMS130Transfomer struct {
}

// Transformer skip FeatureTypeList when merging Base to config, this is a custom operation
func (t WMS130Transfomer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(wms130_response.Layer{}) {
		return func(dst, src reflect.Value) error {
			// NOOP
			return nil
		}
	}
	return nil
}
