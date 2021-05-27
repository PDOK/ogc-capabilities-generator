package ows

import (
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"

	wfs200 "github.com/pdok/ogc-specifications/pkg/wfs200"
)

// WFS200Base is the base WFS 2.0.0 GetCapabilities doc
var WFS200Base wfs200.GetCapabilitiesResponse

func init() {
	wfs200 := wfs200.GetCapabilitiesResponse{}
	base, err := ioutil.ReadFile("./base/wfs200.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wfs200); err != nil {
		log.Fatalf("error: %v", err)
	}
	WFS200Base = wfs200
}

// WFS200Transfomer struct
type WFS200Transfomer struct {
}

// Transformer skip FeatureTypeList when merging Base to config, this is a custom operation
func (t WFS200Transfomer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(wfs200.FeatureTypeList{}) {
		return func(dst, src reflect.Value) error {
			// NOOP
			return nil
		}
	}
	return nil
}
