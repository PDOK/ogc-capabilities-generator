package builder

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"text/template"
)

const TplWms130 = "wms_1_3_0.tpl"
const TplServiceProviderWms = "service_provider.tpl"
const TplCapabilityRequest = "capability_request.tpl"
const TplLayers = "layer.tpl"
const TplInspireCommon = "inspire_common.tpl"

type WmsCapabilities struct {
	wmsTemplates          Executor
	Version               string
	Dataset               Dataset
	Service               Service
	Inspire               bool
	ServiceIdentification string
	ServiceProvider       Organization
	Layers                []Layer
	ServiceDef            string
}

func (c *WmsCapabilities) Build(writer io.Writer) error {
	return c.BuildWms(c.Version, writer)
}

func NewWmsCapabilities(wms *template.Template, def string, serviceVersion string) *WmsCapabilities {
	c := &WmsCapabilities{wmsTemplates: TemplateExecutor{wms}, ServiceDef: def, Version: serviceVersion}
	return c
}

func (c *WmsCapabilities) BuildWms(version string, writer io.Writer) error {
	switch version {

	case "1.3.0":
		return c.buildWms(writer)
	}

	return fmt.Errorf(fmt.Sprintf("Invalid version : %s", version))
}

func (c *WmsCapabilities) buildWms(writer io.Writer) error {

	capabilitiesMap := make(map[string]interface{})

	serviceProvider, err := c.GenerateServiceProvider(TplServiceProviderWms)
	if capabilitiesMap[TplServiceProviderWms] = serviceProvider; err != nil {
		return err
	}

	capabilityRequest, err := c.GenerateCapabilityRequest(TplCapabilityRequest)
	if capabilitiesMap[TplCapabilityRequest] = capabilityRequest; err != nil {
		return err
	}

	layers, err := c.GenerateLayers(TplLayers)
	if capabilitiesMap[TplLayers] = layers; err != nil {
		return err
	}

	inspirecommon, err := c.GenerateInspireCommon(TplInspireCommon)
	if capabilitiesMap[TplInspireCommon] = inspirecommon; err != nil {
		return err
	}

	capabilitiesMap["dataset"] = c.Dataset

	capabilities, err := c.generate(TplWms130, capabilitiesMap)
	if err != nil {
		return err
	}

	if !IsValidXML([]byte(capabilities)) { // not valid xml
		log.Fatalf("Invalid XML : %s", capabilities)
	}

	_, err = writer.Write([]byte(capabilities))
	if err != nil {
		return err
	}

	return nil
}

func (c *WmsCapabilities) AddLayers(dataset Dataset, layers []Layer, def string, service Service) *WmsCapabilities {

	c.Service = service
	c.Dataset = dataset
	c.Layers = layers
	c.ServiceDef = def
	return c
}

func (c *WmsCapabilities) generateNestedLayers(t *template.Template, parentLayer LayersStruct) (string, error) {

	parent := &parentLayer
	parent.Layer.Nested = ""
	for _, childLayer := range parentLayer.Layer.Layers {

		layersStruct := LayersStruct{
			Layer:      childLayer,
			ServiceDef: c.ServiceDef,
		}

		nested, err := c.generateNestedLayers(t, layersStruct)
		if err != nil {
			return "", err
		}
		parent.Layer.Nested += nested
	}

	var out bytes.Buffer
	err := t.Execute(&out, parentLayer)
	if err != nil {
		return " ", err
	}

	return out.String(), nil
}

type LayersStruct struct {
	Layer      Layer
	ServiceDef string
	Dataset    Dataset
	Service    Service
}

func (c *WmsCapabilities) GenerateLayers(templatename string) (string, error) {

	layerTemplate := c.wmsTemplates.Lookup(templatename)
	out := bytes.NewBuffer([]byte(""))

	for _, layer := range c.Layers {
		layersStruct := LayersStruct{
			Layer:      layer,
			ServiceDef: c.ServiceDef,
			Service:    c.Service,
			Dataset:    c.Dataset,
		}
		s, err := c.generateNestedLayers(layerTemplate, layersStruct)
		if err != nil {
			return "", nil
		}
		_, err = out.WriteString(s)
		if err != nil {
			return "", nil
		}
	}
	return out.String(), nil
}

func (c *WmsCapabilities) AddDataset(dataset Dataset) *WmsCapabilities {
	c.Dataset = dataset
	return c
}

func (c *WmsCapabilities) AddServiceProvider(organization Organization, identification string) *WmsCapabilities {
	c.ServiceProvider = organization
	c.ServiceIdentification = identification

	return c
}

func (c *WmsCapabilities) GenerateServiceProvider(template string) (string, error) {
	data := make(map[string]interface{})

	data["organization"] = c.ServiceProvider
	data["identification"] = c.ServiceIdentification

	return c.generate(template, data)
}

func (c *WmsCapabilities) AddCapabilityRequest(dataset Dataset, service Service) *WmsCapabilities {
	c.Dataset = dataset
	c.Service = service
	return c
}

func (c *WmsCapabilities) GenerateCapabilityRequest(template string) (string, error) {

	data := make(map[string]interface{})

	data["service"] = c.Service
	data["dataset"] = c.Dataset

	return c.generate(template, data)
}

func (c *WmsCapabilities) AddInspireCommon(service Service) *WmsCapabilities {
	c.Service = service
	return c
}

func (c *WmsCapabilities) GenerateInspireCommon(template string) (string, error) {

	data := make(map[string]interface{})

	data["service"] = c.Service
	return c.generate(template, data)
}

func (c WmsCapabilities) generate(templateName string, data interface{}) (string, error) {
	var out bytes.Buffer
	err := c.wmsTemplates.ExecuteTemplate(&out, templateName, data)
	return out.String(), err
}
