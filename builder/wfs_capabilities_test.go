package builder

import (
	"bytes"
	"fmt"
	"github.com/antchfx/xmlquery"
	_ "github.com/stretchr/testify/assert"
	"io"
	"log"
	"pdok-capabilities-gen/util"
	"reflect"
	"strings"
	"testing"
	"text/template"
)

var templates = util.GetTemplates("../templates/wfs/*")

type MockTemplateExecutor struct {
	InvokeOrder []string
	Invoked     map[string]interface{}
}

func (t *MockTemplateExecutor) ExecuteTemplate(wr io.Writer, templateName string, data interface{}) error {
	_, err := wr.Write([]byte(
		fmt.Sprintf("<mock>%s</mock>\n", templateName)))
	if err != nil {
		return err
	}
	t.InvokeOrder = append(t.InvokeOrder, templateName)
	t.Invoked[templateName] = data
	return nil
}
func (t *MockTemplateExecutor) Lookup(templateName string) *template.Template {
	return nil
}

func TestCapabilities_BuildWfs2(t *testing.T) {
	c := &WfsCapabilities{
		wfsTemplates: &MockTemplateExecutor{[]string{}, make(map[string]interface{})},
	}

	c.AddDataset(Dataset{}).AddFeatures(Dataset{}, []Feature{}, ServiceDef{}).AddServiceIdentification(Identification{}).AddOperationsMetadata(Dataset{}, Service{
		Dataset:            "",
		Organization:       "",
		Identification:     "",
		Inspire:            false,
		UrlWfs:             "",
		MetadataResourceId: "",
		Operations:         nil,
		Boundingbox:        "",
		Constraints:        nil,
		Features:           nil,
	})

	t.Run("test Build", func(t *testing.T) {
		var buff bytes.Buffer
		err := c.buildWfs("2.0.0", &buff)
		if err != nil {
			log.Printf("error: %v", err)
		}
		println(string(buff.Bytes()))

	})

}

func TestCapabilities_BuildWfs1(t *testing.T) {
	c := &WfsCapabilities{
		wfsTemplates: &MockTemplateExecutor{[]string{}, make(map[string]interface{})},
	}
	c.AddDataset(Dataset{}).AddFeatures(Dataset{}, []Feature{}, ServiceDef{}).AddServiceIdentification(Identification{}).AddOperationsMetadata(Dataset{}, Service{
		Dataset:            "",
		Organization:       "",
		Identification:     "",
		Inspire:            false,
		UrlWfs:             "",
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

func TestCapabilities_BuildOrder(t *testing.T) {

	m := &MockTemplateExecutor{[]string{}, make(map[string]interface{})}
	c := &WfsCapabilities{
		wfsTemplates: m,
	}

	c.AddDataset(Dataset{}).AddFeatures(Dataset{}, []Feature{}, ServiceDef{}).AddServiceIdentification(Identification{}).AddOperationsMetadata(Dataset{}, Service{
		Dataset:            "",
		Organization:       "",
		Identification:     "",
		Inspire:            false,
		UrlWfs:             "",
		MetadataResourceId: "",
		Operations:         nil,
		Boundingbox:        "",
		Constraints:        nil,
		Features:           nil,
	})

	orderTemplates := []string{"service-identification.tpl", "service-provider.tpl", "operations_metadata_1_1_0.tpl", "feature_list_1_1_0.tpl", "filter_capabilities_1_1_0.tpl", "wfs_1_1_0.tpl"}
	orderInterface := []string{"builder.Identification", "builder.Organization", "map[string]interface {}", "map[string]interface {}", "[]builder.Filter", "map[string]interface {}"}

	c.Version = "1.1.0"

	t.Run("test Build", func(t *testing.T) {
		var buff bytes.Buffer
		err := c.Build(&buff)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		d := m.Invoked

		if len(m.InvokeOrder) != len(orderTemplates) {
			t.Errorf("Wrong number of invoked Templates got %v, want %v", len(m.InvokeOrder), len(orderTemplates))
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

func TestCapabilities_BuildWrongVersion(t *testing.T) {
	c := &WfsCapabilities{
		wfsTemplates: &MockTemplateExecutor{[]string{}, make(map[string]interface{})},
	}
	c.AddDataset(Dataset{}).AddFeatures(Dataset{}, []Feature{}, ServiceDef{}).AddServiceIdentification(Identification{}).AddOperationsMetadata(Dataset{}, Service{
		Dataset:            "",
		Organization:       "",
		Identification:     "",
		Inspire:            false,
		UrlWfs:             "",
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
func TestCapabilities_AddServiceIdentification(t *testing.T) {

	c := &WfsCapabilities{
		wfsTemplates: TemplateExecutor{templates},
	}

	identification := Identification{
		Title:    "kadastralekaartv4",
		Abstract: "omschrijving van de dataset",
		Keywords: []string{"Kadaster", "kadastrale percelen"},
		Version:  "2.0.0",
	}

	tests := []struct {
		name  string
		xpath string
		want  string
	}{
		{"want title", "//ows:Title/text()", "kadastralekaartv4"},
		{"want abstract", "//ows:Abstract/text()", "omschrijving van de dataset"},
		{"want keywords1", "//ows:Keywords//ows:Keyword/text()", "Kadaster"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := c.AddServiceIdentification(identification).GenerateServiceIdentification()
			if err != nil {
				log.Printf("error: %v", err)
			}

			doc, err := xmlquery.Parse(strings.NewReader(got))
			if err != nil {
				log.Printf("error: %v", err)
			}

			one := xmlquery.FindOne(doc, tt.xpath)
			if one == nil {
				t.Errorf("AddServiceIdentification() = %v,d want %v", "nil", tt.want)

			}

			if one != nil && one.Data != tt.want {
				t.Errorf("AddServiceIdentification() = %v, want %v", one.Data, tt.want)
			}

		})
	}
}

func TestCapabilities_AddFeatures(t *testing.T) {

	c := &WfsCapabilities{
		wfsTemplates: templates,
	}

	features := []Feature{{
		Name:        "krw_oppervlaktewaterlichamen_lijnen_rws_act",
		Title:       "kadastralegrens",
		Abstract:    "Een Kadastrale Grens",
		Keywords:    []string{"Kadastrale percelen", "infoMapAccessService"},
		Boundingbox: "WGS84",
		Crs:         []string{"28992"},
	}}

	crs := make(map[string]Crs)
	crs["28992"] = Crs{
		Name:  28992,
		Value: "urn:ogc:def:crs:EPSG::28992",
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
		{"want_keywork", "//ows:Keyword/text()", "Kadastrale percelen"},
		{"want_abstract", "//Abstract/text()", "Een Kadastrale Grens"},
		{"want_name", "//Name/text()", "krw_oppervlaktewaterlichamen_lijnen_rws_act"},
		{"want_title", "//Title/text()", "kadastralegrens"},
		{"want_LowerCorner", "//ows:LowerCorner/text()", "1 2"},
		{"want_UpperCorner", "//ows:UpperCorner/text()", "3 4"},
		{"want_crsDefault", "//DefaultCRS/text()", "urn:ogc:def:crs:EPSG::28992"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := c.AddFeatures(dataset, features, servicedef).GenerateFeatures(TplFeatures110)

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

func TestCapabilities_AddServiceProvider(t *testing.T) {

	c := &WfsCapabilities{
		wfsTemplates: templates,
	}

	organization := Organization{
		Name:       "PDOK",
		Individual: "KlantContactCenter PDOK",
		Position:   "pointOfContact",
		City:       "Apeldoorn",
		Country:    "Nederland",
		Email:      "BeheerPDOK@kadaster.nl",
	}

	tests := []struct {
		name  string
		xpath string
		want  string
	}{
		{"want PDOK", "//ows:ProviderName/text()", "PDOK"},
		{"want individual", "//ows:IndividualName/text()", "KlantContactCenter PDOK"},
		{"want position", "//ows:PositionName/text()", "pointOfContact"},
		{"want city", "//ows:City/text()", "Apeldoorn"},
		{"want country", "//ows:Country/text()", "Nederland"},
		{"want email", "//ows:ElectronicMailAddress/text()", "BeheerPDOK@kadaster.nl"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := c.AddServiceProvider(organization).GenerateServiceProvider()
			if err != nil {
				return
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

func TestCapabilities_AddOperationsMetadata(t *testing.T) {

	c := &WfsCapabilities{
		wfsTemplates: templates,
	}

	getservice := func(inspire bool) Service {

		return Service{
			Dataset:            "eendataset",
			Organization:       "",
			Inspire:            inspire,
			UrlWfs:             "https://geodata.nationaalgeoregister.nl/kadastralekaart/wfs/v4_0",
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
				c.wfsTemplates = template.Must(template.New("wfsTemplates").ParseGlob("./*"))
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

func TestCapabilities_AddFilters(t *testing.T) {
	c := &WfsCapabilities{wfsTemplates: templates}
	dataType, err := c.AddFilters().GenerateFilters(TplFilterCapabilities)
	if err != nil {
		log.Printf("error: %v", err)
		reflect.DeepEqual(dataType, TplFilterCapabilities)
	}
}

func TestCapabilities_AddDataset(t *testing.T) {

	type args struct {
		dataset Dataset
	}
	tests := []struct {
		name string
		args args
		want *WfsCapabilities
	}{
		{
			name: "test 1",
			args: args{Dataset{
				Name:               "Pietje",
				MetadataResourceId: "Puk",
			}},
			want: &WfsCapabilities{Dataset: Dataset{
				Name:               "Pietje",
				MetadataResourceId: "Puk",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WfsCapabilities{
				Dataset: tt.args.dataset,
			}
			if got := c.AddDataset(tt.args.dataset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDataset() = %v, want %v", got, tt.want)
			}
		})
	}
}
