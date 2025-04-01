package ows

import (
	"log"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"

	"github.com/pdok/ogc-specifications/pkg/wms130"
)

// WMS130Base is the base WMS 1.3.0 GetCapabilities doc
var WMS130Base wms130.GetCapabilitiesResponse

func init() {
	wms130Response := wms130.GetCapabilitiesResponse{}
	base, err := os.ReadFile("../base/wms130.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wms130Response); err != nil {
		log.Fatalf("error: %v", err)
	}
	WMS130Base = wms130Response
}

// WMS130Transfomer struct
type WMS130Transfomer struct {
}

// Transformer skip LayerList when merging Base to config, this is a custom operation
func (t WMS130Transfomer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(wms130.Layer{}) {
		return func(dst, src reflect.Value) error {
			// NOOP
			return nil
		}
	}
	if typ == reflect.TypeOf(wms130.DCPType{}) {
		return func(dst, src reflect.Value) error {
			// NOOP
			return nil
		}
	}
	return nil
}
