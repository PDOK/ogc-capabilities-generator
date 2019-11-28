package builder

import (
	"bytes"
	"fmt"
	"github.com/antchfx/xmlquery"
	_ "github.com/stretchr/testify/assert"

	"log"
	"pdok-capabilities-gen/util"
	"reflect"
	"strings"
	"testing"
)

var templatesWms = util.GetTemplates("../templates/wms/*")

func TestCapabilities_Build_wms1(t *testing.T) {
	c := &WmsCapabilities{
		wmsTemplates: &MockTemplateExecutor{[]string{}, make(map[string]interface{})},
	}

	c.AddDataset(Dataset{}).AddLayers(Dataset{}, []Layer{}, ServiceDef{}, Service{
		Dataset:            "",
		Organization:       "",
		Identification:     "",
		Inspire:            false,
		UrlWms:             "",
		MetadataResourceId: "",
		Operations:         nil,
		Boundingbox:        "",
		Constraints:        nil,
		Features:           nil,
	})
	t.Run("test Build", func(t *testing.T) {

		var buff bytes.Buffer
		err := c.BuildWms("1.3.0", &buff)
		if err != nil {
			log.Printf("error: %v", err)

		}

		println(string(buff.Bytes()))

	})

}

func TestCapabilities_BuildOrder_wms(t *testing.T) {

	m := &MockTemplateExecutor{[]string{}, make(map[string]interface{})}
	c := &WmsCapabilities{
		wmsTemplates: m,
	}
	c.AddDataset(Dataset{}).AddLayers(Dataset{}, []Layer{}, ServiceDef{}, Service{
		Dataset:            "",
		Organization:       "",
		Identification:     "",
		Inspire:            false,
		UrlWms:             "",
		MetadataResourceId: "",
		Operations:         nil,
		Boundingbox:        "",
		Constraints:        nil,
		Features:           nil,
	})

	c.Version = "1.3.0"

	orderTemplates := []string{"service_provider.tpl", "capability_request.tpl", "inspire_common.tpl", "wms_1_3_0.tpl"}
	orderInterface := []string{"map[string]interface {}", "map[string]interface {}", "map[string]interface {}", "map[string]interface {}"}

	t.Run("test Build", func(t *testing.T) {
		var buff bytes.Buffer
		err := c.Build(&buff)

		if err != nil {
			t.Errorf("Build yielded an error, not expected")
			return
		}
		d := m.Invoked

		if len(m.InvokeOrder) != len(orderTemplates) {
			t.Errorf("Wrong number of invoked templatesWms got %v, want %v", len(m.InvokeOrder), len(orderTemplates))
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

func TestCapabilities_BuildWrongVersion_wms(t *testing.T) {
	c := &WmsCapabilities{
		wmsTemplates: &MockTemplateExecutor{[]string{}, make(map[string]interface{})},
	}
	c.AddDataset(Dataset{}).AddLayers(Dataset{}, []Layer{}, ServiceDef{}, Service{
		Dataset:            "",
		Organization:       "",
		Identification:     "",
		Inspire:            false,
		UrlWms:             "",
		MetadataResourceId: "",
		Operations:         nil,
		Boundingbox:        "",
		Constraints:        nil,
		Features:           nil,
	})
	t.Run("test Build", func(t *testing.T) {
		var buff bytes.Buffer
		err := c.Build(&buff)
		if err != nil {
			log.Printf("error: %v", err)
		}
		println(string(buff.Bytes()))

	})

}

func TestCapabilities_AddLayers(t *testing.T) {

	c := &WmsCapabilities{
		wmsTemplates: templatesWms,
	}
	layer := []Layer{{
		Name:        "krw_oppervlaktewaterlichamen_lijnen_rws_act",
		Title:       "kadastralegrens",
		Abstract:    "123",
		Keywords:    []string{"Kadastrale percelen", "infoMapAccessService"},
		Boundingbox: "WGS84",
		Crs:         []string{"28992"},
		Srs:         nil,
		Styles:      nil,
		Layers:      nil,
		Nested:      "",
	},
	}

	crs := make(map[string]Crs)
	crs["28992"] = Crs{
		Name:  28992,
		Value: "28992",
	}

	bbox := make(map[string]Boundingbox)
	bbox["WGS84"] = Boundingbox{
		Coordinates: []string{"1", "2", "3", "4"},
		Srs:         0,
	}

	servicedef := ServiceDef{
		Datasets:       nil,
		Organizations:  nil,
		Identification: nil,
		Crs:            crs,
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
		{"want_keywork", "//Keyword/text()", "Kadastrale percelen"},
		{"want_abstract", "//Abstract/text()", "123"},
		{"want_name", "//Name/text()", "krw_oppervlaktewaterlichamen_lijnen_rws_act"},
		{"want_title", "//Title/text()", "kadastralegrens"},
		{"want_westBoundLongitude", "//westBoundLongitude/text()", "1"},
		{"want_eastBoundLongitude", "//eastBoundLongitude/text()", "3"},
		{"want_southBoundLatitude", "//southBoundLatitude/text()", "2"},
		{"want_northBoundLatitude", "//northBoundLatitude/text()", "4"},
		{"want_crsDefault", "//CRS/text()", "28992"},
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

func TestCapabilities_AddServiceProviderWms(t *testing.T) {

	c := &WmsCapabilities{
		wmsTemplates: templatesWms,
	}

	organization := Organization{
		Name:       "PDOK",
		Individual: "KlantContactCenter PDOK",
		Position:   "pointOfContact",
		City:       "Apeldoorn",
		Country:    "Nederland",
		Email:      "BeheerPDOK@kadaster.nl",
	}
	identification := Identification{
		Title:    "",
		Abstract: "",
		Keywords: nil,
		Version:  "",
	}

	tests := []struct {
		name  string
		xpath string
		want  string
	}{
		{"want PDOK", "//ContactOrganization/text()", "PDOK"},
		{"want individual", "//ContactPerson/text()", "KlantContactCenter PDOK"},
		{"want position", "//ContactPosition/text()", "pointOfContact"},
		{"want city", "//City/text()", "Apeldoorn"},
		{"want country", "//Country/text()", "Nederland"},
		{"want email", "//ContactElectronicMailAddress/text()", "BeheerPDOK@kadaster.nl"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := c.AddServiceProvider(organization, identification).GenerateServiceProvider(TplServiceProviderWms)
			if err != nil {
				t.Errorf("AddServiceProvider() error = %v, want error %v", err, false)
			}
			doc, err := xmlquery.Parse(strings.NewReader(got))

			one := xmlquery.FindOne(doc, tt.xpath)
			if one == nil {
				t.Errorf("AddServiceProvider() = %v, want %v", "nil", tt.want)

			}

			if one != nil && one.Data != tt.want {
				t.Errorf("AddServiceProvider() = %v, want %v", one.Data, tt.want)
			}

		})
	}
}

func TestCapabilities_AddDatasetWms(t *testing.T) {

	type args struct {
		dataset Dataset
	}
	tests := []struct {
		name string
		args args
		want *WmsCapabilities
	}{
		{
			name: "test 1",
			args: args{Dataset{
				Name:               "Pietje",
				MetadataResourceId: "Puk",
			}},
			want: &WmsCapabilities{Dataset: Dataset{
				Name:               "Pietje",
				MetadataResourceId: "Puk",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WmsCapabilities{
				Dataset: tt.args.dataset,
			}
			if got := c.AddDataset(tt.args.dataset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDataset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapabilities_AddCapabilityRequest(t *testing.T) {

	c := &WmsCapabilities{
		wmsTemplates: templatesWms,
	}

	service := Service{
		Dataset:            "",
		Organization:       "",
		Identification:     "",
		Inspire:            false,
		UrlWms:             "https://www.test.nl",
		MetadataResourceId: "",
		Operations:         nil,
		Boundingbox:        "",
		Constraints:        nil,
		Features:           nil,
		Layers:             nil,
	}

	dataset := Dataset{
		Name:               "eendataset",
		MetadataResourceId: "dataset-metadata-resource-id&><",
	}

	tests := []struct {
		name     string
		xpath    string
		wantFunc func(node *xmlquery.Node) error
	}{
		{"want_url", "/Request/GetCapabilities/DCPType/HTTP/Get/OnlineResource/@xlink:href",
			func(node *xmlquery.Node) error {
				data := "https://www.test.nl?SERVICE=WMS&language=dut&"
				if node.FirstChild.Data != data {
					return fmt.Errorf(data)
				}
				return nil
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := c.AddCapabilityRequest(dataset, service).GenerateCapabilityRequest(TplCapabilityRequest)
			if err != nil {
				log.Printf("error: %v", err)
			}

			doc, err := xmlquery.Parse(strings.NewReader(got))
			if err != nil {
				log.Printf("error: %v", err)
			}

			one := xmlquery.FindOne(doc, tt.xpath)
			if one == nil {
				t.Errorf("AddCapabilityRequest() Xpath '%s' did not yield any results ", tt.xpath)
			}

			if one != nil && tt.wantFunc(one) != nil {
				t.Errorf("AddCapabilityRequest() xpath: %s != %v ", tt.xpath, tt.wantFunc(one))
			}

		})
	}
}
