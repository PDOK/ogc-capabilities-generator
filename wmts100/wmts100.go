package wmts100

import (
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"

	wmts100_response "github.com/pdok/ogc-specifications/pkg/wmts100/response"
)

// GetBase function to get a filled "template" based on the wmts100.yaml config
func GetBase() wmts100_response.GetCapabilities {
	wmts100 := wmts100_response.GetCapabilities{}
	base, err := ioutil.ReadFile("./base/wmts100.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wmts100); err != nil {
		log.Fatalf("error: %v", err)
	}
	return wmts100
}

// Wmts100Transfomer struct
type Wmts100Transfomer struct {
}

// Transformer skip FeatureTypeList when merging Base to config, this is a custom operation
func (t Wmts100Transfomer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(wmts100_response.TileMatrixSet{}) {
		return func(dst, src reflect.Value) error {
			// NOOP
			return nil
		}
	}
	return nil
}
