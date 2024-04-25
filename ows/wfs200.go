package ows

import (
	"log"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"

	"github.com/pdok/ogc-specifications/pkg/wfs200"
)

// WFS200Base is the base WFS 2.0.0 GetCapabilities doc
var WFS200Base wfs200.GetCapabilitiesResponse

func init() {
	wfs200Response := wfs200.GetCapabilitiesResponse{}
	base, err := os.ReadFile("./base/wfs200.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wfs200Response); err != nil {
		log.Fatalf("error: %v", err)
	}
	WFS200Base = wfs200Response
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
