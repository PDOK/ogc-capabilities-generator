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
	// Transformer only copies over tilematrixsets from the base (src) if there is a layer present
	// having a tilematrixlink referring to the tilematrixset identifier in the overlayer (dst).
	// When a tilematrixset is present in both base (src) and overlay(dst) with matching identifiers,
	// the tilematrixset from the overlay takes precendent.
	if typ == reflect.TypeOf(wmts100.Contents{}) {
		return func(dst, src reflect.Value) error {
			// collect all tilematrixset references from Layer/TileMatrixSetLink
			tms_links := make(map[string]bool)
			src_contents := src.Interface().(wmts100.Contents)
			contents := dst.Interface().(wmts100.Contents)
			for _, lyr := range contents.Layer {
				for _, tmsl := range lyr.TileMatrixSetLink {
					tms_links[tmsl.TileMatrixSet] = true
				}
			}
			// generate set of all tilematrixset definitions in both base and overlay
			tms_set := make(map[string]wmts100.TileMatrixSet)
			// overwrite tilematrixsets from src by dst; note the order of `append()`
			for _, tms := range append(src_contents.TileMatrixSet, contents.TileMatrixSet...) {
				tms_set[tms.Identifier] = tms
			}
			// iterate over the set of all tilematrixset definitions and check if a layer refers
			// to this tilematrixset definition, if so include this tilematrix set in the output
			merged := []wmts100.TileMatrixSet{}
			for _, tms := range tms_set {
				tmsCopy := tms
				if _, ok := tms_links[tms.Identifier]; ok {
					tmsCopy.TileMatrix = []wmts100.TileMatrix{}
					for _, tm := range tms.TileMatrix {
						tmsCopy.TileMatrix = append(tmsCopy.TileMatrix, tm)
					}
					merged = append(merged, tmsCopy)
				}
			}
			contents.TileMatrixSet = merged
			dst.Set(reflect.ValueOf(contents))
			return nil
		}
	}
	return nil
}
