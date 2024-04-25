package ows

import (
	"log"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"

	"github.com/pdok/ogc-specifications/pkg/wmts100"
)

// WMTS100Base is the base WMTS 1.0.0 GetCapabilities doc
var WMTS100Base wmts100.GetCapabilitiesResponse

func init() {
	wmts100Response := wmts100.GetCapabilitiesResponse{}
	base, err := os.ReadFile("./base/wmts100.yaml")
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
			tmsLinks := make(map[string]bool)
			srcContents := src.Interface().(wmts100.Contents)
			contents := dst.Interface().(wmts100.Contents)
			for _, lyr := range contents.Layer {
				for _, tmsl := range lyr.TileMatrixSetLink {
					tmsLinks[tmsl.TileMatrixSet] = true
				}
			}
			contents.TileMatrixSet = mergeTilematrixSets(contents.TileMatrixSet, srcContents.TileMatrixSet, tmsLinks)
			dst.Set(reflect.ValueOf(contents))
			return nil
		}
	}
	return nil
}

// mergeTilematrixSets selectively includes TileMatrixSet and TileMatrix definitions when:
// - TileMatrixSet is in dst, but not in src, dst TileMatrixSet is included
// - dst TileMatrixSet exists, but is empty, src TileMatrixSet is included completely
// - dst TileMatrixSet has a subset of TileMatrices identified, those TileMatrices are included from src
// - TileMatrixSet in src, but not in dst, src TileMatrixSet is _not_ included
func mergeTilematrixSets(dst []wmts100.TileMatrixSet, src []wmts100.TileMatrixSet, tmsLinks map[string]bool) []wmts100.TileMatrixSet {
	content := make(map[string]map[string]bool)
	for _, tms := range dst {
		if tms.TileMatrix != nil {
			content[tms.Identifier] = make(map[string]bool)
			for _, tm := range tms.TileMatrix {
				content[tms.Identifier][tm.Identifier] = true
			}
		} else {
			content[tms.Identifier] = nil
		}
	}

	// obtain TileMatrixSet and TileMatrix items from 'src'
	srcTileMatrices := make(map[string]wmts100.TileMatrixSet)
	for _, tms := range src {
		tmsCopy := tms
		if dstTmMap, ok := content[tms.Identifier]; ok {
			if dstTmMap != nil {
				tmsCopy.TileMatrix = []wmts100.TileMatrix{}
				for _, tm := range tms.TileMatrix {
					if _, ok := dstTmMap[tm.Identifier]; ok {
						tmsCopy.TileMatrix = append(tmsCopy.TileMatrix, tm)
					}
				}
			}
			srcTileMatrices[tms.Identifier] = tmsCopy
		}
	}

	var merged []wmts100.TileMatrixSet
	for _, tms := range dst {
		// keep dst order and filter linked tms
		if _, ok := tmsLinks[tms.Identifier]; ok {
			if srcTms, ok := srcTileMatrices[tms.Identifier]; ok {
				merged = append(merged, srcTms)
			} else {
				merged = append(merged, tms)
			}
		}
	}

	return merged
}
