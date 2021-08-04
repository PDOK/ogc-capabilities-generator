package ows

import (
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"

	"github.com/pdok/ogc-specifications/pkg/wmts100"
)

// WMTS100Base is the base WMTS 1.0.0 GetCapabilities doc
var WMTS100Base wmts100.GetCapabilitiesResponse

func init() {
	wmts100Response := wmts100.GetCapabilitiesResponse{}
	base, err := ioutil.ReadFile("./base/wmts100.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wmts100Response); err != nil {
		log.Fatalf("error: %v", err)
	}
	WMTS100Base = wmts100Response
}

// WMTS100Transfomer struct
type WMTS100Transfomer struct {
}

// Transformer selectively includes TileMatrixSet and TileMatrix definitions
func (t WMTS100Transfomer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf([]wmts100.TileMatrixSet{}) {
		return func(dst, src reflect.Value) error {
			// determine which TileMatrixSet and TileMatrix items
			// to include by examining 'dst'
			content := make(map[string]map[string]bool)
			for _, tms := range dst.Interface().([]wmts100.TileMatrixSet) {
				if tms.TileMatrix != nil {
					content[tms.Identifier] = make(map[string]bool)
					for _, tm := range tms.TileMatrix {
						content[tms.Identifier][tm.Identifier] = true
					}
				}
			}
			// obtain TileMatrixSet and TileMatrix items from 'src'
			merged := []wmts100.TileMatrixSet{}
			for _, tms := range src.Interface().([]wmts100.TileMatrixSet) {
				tmsCopy := tms
				if _, ok := content[tms.Identifier]; ok {
					tmsCopy.TileMatrix = []wmts100.TileMatrix{}
					for _, tm := range tms.TileMatrix {
						if _, ok := content[tms.Identifier][tm.Identifier]; ok {
							tmsCopy.TileMatrix = append(tmsCopy.TileMatrix, tm)
						}
					}
				}
				merged = append(merged, tmsCopy)
			}
			dst.Set(reflect.ValueOf(merged))
			return nil
		}
	}
	return nil
}
