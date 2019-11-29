package builder

import (
	"bytes"

	"github.com/antchfx/xmlquery"

	"log"
	"pdok-capabilities-gen/util"
	"reflect"
	"strings"
	"testing"
	"text/template"
)

var templatesWmts = util.GetTemplates("../templates/wmts/*")

func TestNewWmtsCapabilities(t *testing.T) {
	c := &WmtsCapabilities{
		wmtsTemplates: &MockTemplateExecutor{[]string{}, make(map[string]interface{})},
	}

	c.AddDataset(Dataset{}).AddLayers(Dataset{}, []Layer{}, ServiceDef{}, Service{
		Dataset:            "",
		Organization:       "",
		Identification:     "",
		Inspire:            false,
		UrlWmts:            "",
		MetadataResourceId: "",
		Operations:         nil,
		Boundingbox:        "",
		Constraints:        nil,
		Features:           nil,
	})
	t.Run("test Build", func(t *testing.T) {

		var buff bytes.Buffer
		err := c.BuildWmts("1.0.0", &buff)
		if err != nil {
			log.Printf("error: %v", err)

		}

		println(string(buff.Bytes()))

	})

}

func TestCapabilities_BuildOrder_wmts(t *testing.T) {

	m := &MockTemplateExecutor{[]string{}, make(map[string]interface{})}
	c := &WmtsCapabilities{
		wmtsTemplates: m,
	}
	c.AddDataset(Dataset{}).AddLayers(Dataset{}, []Layer{}, ServiceDef{}, Service{
		Dataset:            "",
		Organization:       "",
		Identification:     "",
		Inspire:            false,
		UrlWmts:            "",
		MetadataResourceId: "",
		Operations:         nil,
		Boundingbox:        "",
		Constraints:        nil,
		Features:           nil,
	})

	c.Version = "1.0.0"

	orderTemplates := []string{"layer.tpl", "tilematrixset_28992.tpl", "tilematrixset_3857.tpl", "tilematrixset_25831.tpl", "wmts_1_0_0.tpl"}
	orderInterface := []string{"map[string]interface {}", "[]builder.Tilematrixset", "[]builder.Tilematrixset", "[]builder.Tilematrixset", "map[string]interface {}"}
	t.Run("test Build", func(t *testing.T) {
		var buff bytes.Buffer
		err := c.Build(&buff)

		if err != nil {
			t.Errorf("Build yielded an error, not expected")
			return
		}
		d := m.Invoked

		if len(m.InvokeOrder) != len(orderTemplates) {
			t.Errorf("Wrong number of invoked templatesWmts got %v, want %v", len(m.InvokeOrder), len(orderTemplates))
		}

		for index := range orderTemplates {
			if m.InvokeOrder[index] != orderTemplates[index] {
				t.Errorf("Wrong orderTemplates of template invocation() = %v, want %v", m.InvokeOrder[index], orderTemplates[index])
			}

			if m.InvokeOrder[index] == orderTemplates[index] {
				i := d[m.InvokeOrder[index]]

				value := reflect.ValueOf(i).Type().String()

				if value != orderInterface[index] {
					t.Errorf("Wrong template/interface  invocation() = %v, want %v", value, orderInterface[index])
				}

			}
		}

	})

}

func TestCapabilities_BuildWrongVersion_wmts(t *testing.T) {
	c := &WmtsCapabilities{
		wmtsTemplates: &MockTemplateExecutor{[]string{}, make(map[string]interface{})},
	}
	c.AddDataset(Dataset{}).AddLayers(Dataset{}, []Layer{}, ServiceDef{}, Service{
		Dataset:            "",
		Organization:       "",
		Identification:     "",
		Inspire:            false,
		UrlWmts:            "",
		MetadataResourceId: "",
		Operations:         nil,
		Boundingbox:        "",
		Constraints:        nil,
		Features:           nil,
	})
	c.Version = "1.0.2"
	t.Run("test Build", func(t *testing.T) {
		var buff bytes.Buffer
		err := c.Build(&buff)
		if err != nil {
			log.Printf("error: %v", err)
		}
		println(string(buff.Bytes()))

	})

}

func TestCapabilities_AddLayers_wmts(t *testing.T) {

	c := &WmtsCapabilities{
		wmtsTemplates: templatesWmts,
	}
	layer := []Layer{{
		Name:        "krw_oppervlaktewaterlichamen_lijnen_rws_act",
		Title:       "",
		Abstract:    "",
		Keywords:    nil,
		Boundingbox: "WGS84",
		Crs:         nil,
		Srs:         []string{"28992"},
	}}

	srs := make(map[string]Srs)
	srs["28992"] = Srs{
		Name:  28992,
		Value: "28992",
	}

	bbox := make(map[string]Boundingbox)
	bbox["WGS84"] = Boundingbox{
		Coordinates: []string{"1", "2", "3", "4"},
		Srs:         0,
	}
	c.Version = "1.0.0"
	servicedef := ServiceDef{
		Datasets:       nil,
		Organizations:  nil,
		Identification: nil,
		Crs:            nil,
		Format:         nil,
		Boundingbox:    bbox,
		Operations:     nil,
		Constraints:    nil,
	}
	dataset := Dataset{
		Name:               "eendataset",
		MetadataResourceId: "dataset-metadata-resource-id&><",
	}

	tests := []struct {
		name  string
		xpath string
		want  string
	}{

		{"want_layer_name", "//ows:Title/text()", "krw_oppervlaktewaterlichamen_lijnen_rws_act"},
		{"want_lowercorner", "//ows:LowerCorner/text()", "1 2"},
		{"want_uppercorner", "//ows:UpperCorner/text()", "3 4"},
		{"want_identifier", "//ows:Identifier/text()", "krw_oppervlaktewaterlichamen_lijnen_rws_act"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := c.AddLayers(dataset, layer, servicedef, Service{}).GenerateLayers(TplLayers)

			if err != nil {
				log.Printf("error: %v", err)
			}

			doc, err := xmlquery.Parse(strings.NewReader(got))
			if err != nil {
				log.Printf("error: %v", err)
			}

			one := xmlquery.FindOne(doc, tt.xpath)
			if one == nil {
				t.Errorf("AddLayers() = %v,d want %v", "nil", tt.want)

			}

			if one != nil && one.Data != tt.want {
				t.Errorf("AddLayers() = %v, want %v", one.Data, tt.want)
			}

		})
	}
}

func TestCapabilities_AddOperationsMetadataWmts(t *testing.T) {

	c := &WmtsCapabilities{
		wmtsTemplates: templates,
	}

	getservice := func(inspire bool) Service {

		return Service{
			Dataset:            "eendataset",
			Organization:       "",
			Inspire:            inspire,
			UrlWmts:            "https://geodata.nationaalgeoregister.nl/kadastralekaart/wmts/v4_0",
			MetadataResourceId: "service-metadata-resource-id",
			Operations:         nil,
			Boundingbox:        "",
			Constraints:        nil,
			Features:           nil,
		}
	}

	dataset := Dataset{
		Name:               "eendataset",
		MetadataResourceId: "dataset-metadata-resource-id&><",
	}

	type want struct {
		result string
	}
	tests := []struct {
		name        string
		createError bool
		service     Service
		xpath       string
		want        *want
	}{
		{"want Inspire SpatialDataSetIdentifier", false, getservice(true), "/ows:OperationsMetadata/ows:ExtendedCapabilities/inspire_dls:ExtendedCapabilities/inspire_dls:SpatialDataSetIdentifier/inspire_common:Code/text()", &want{"dataset-metadata-resource-id&><"}},
		{"want Inspire MetadataUrl", false, getservice(true), "/ows:OperationsMetadata/ows:ExtendedCapabilities/inspire_dls:ExtendedCapabilities/inspire_common:MetadataUrl/inspire_common:URL/text()", &want{"http://www.nationaalgeoregister.nl/geonetwork/srv/eng/csw?service=CSW&version=2.0.2&request=GetRecordById&outputschema=http://www.isotc211.org/2005/gmd&elementsetname=full&id=service-metadata-resource-id"}},
		{"no Inspire SpatialDataSetIdentifier", false, getservice(false), "/ows:OperationsMetadata/ows:ExtendedCapabilities/inspire_dls:ExtendedCapabilities/inspire_dls:SpatialDataSetIdentifier/inspire_common:Code/text()", nil},
		{"no Inspire MetadataUrl", false, getservice(false), "/ows:OperationsMetadata/ows:ExtendedCapabilities/inspire_dls:ExtendedCapabilities/inspire_common:MetadataUrl/inspire_common:URL/text()", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//test for missing template
			if tt.createError {
				c.wmtsTemplates = template.Must(template.New("wfsTemplates").ParseGlob("./*"))
			}

			got, err := c.AddOperationsMetadata(dataset, tt.service).GenerateOperationsMetadata(TplOperationsMetadataProvider)
			if tt.createError && !(err == nil) {
				t.Errorf("AddCapabilityRequest() error = %v, want error %v", err != nil, tt.createError)
			}

			doc, err := xmlquery.Parse(strings.NewReader(got))
			if err != nil {
				log.Printf("error: %v", err)
			}

			one := xmlquery.FindOne(doc, tt.xpath)
			if one == nil && tt.want != nil {
				t.Errorf("AddCapabilityRequest() = %v, want %v", "nil", tt.want)

			}

			if one != nil && (tt.want == nil || one.Data != tt.want.result) {
				t.Errorf("AddCapabilityRequest() = %v, want %v", one.Data, tt.want)
			}

		})
	}
}

func TestWmtsCapabilities_AddTileMatrixSet(t *testing.T) {
	c := &WmtsCapabilities{wmtsTemplates: templates}
	dataType, err := c.AddTileMatrixSet().GenerateTileMatrixSet(TplTilematrixset28992)
	if err != nil {
		log.Printf("error: %v", err)
		reflect.DeepEqual(dataType, TplTilematrixset28992)
	}
}
