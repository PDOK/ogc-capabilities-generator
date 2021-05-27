package ows

import (
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"

	wmts100 "github.com/pdok/ogc-specifications/pkg/wmts100"
)

// WMTS100Base is the base WMTS 1.0.0 GetCapabilities doc
var WMTS100Base wmts100.GetCapabilitiesResponse

func init() {
	wmts100 := wmts100.GetCapabilitiesResponse{}
	base, err := ioutil.ReadFile("./base/wmts100.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wmts100); err != nil {
		log.Fatalf("error: %v", err)
	}
	WMTS100Base = wmts100
}

// WMTS100Transfomer struct
type WMTS100Transfomer struct {
}

// Transformer skip FeatureTypeList when merging Base to config, this is a custom operation
func (t WMTS100Transfomer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(wmts100.TileMatrixSet{}) {
		return func(dst, src reflect.Value) error {
			// NOOP
			return nil
		}
	}
	return nil
}
