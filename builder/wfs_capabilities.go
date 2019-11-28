package builder

import (
	"bytes"
	"fmt"
	"github.com/antchfx/xmlquery"
	"io"
	"log"
	"strings"
	"text/template"
)

const TplServiceidentification = "service-identification.tpl"
const TplFeatures200 = "feature_list_2_0_0.tpl"
const TplFeatures110 = "feature_list_1_1_0.tpl"
const TplOperationsMetadataProvider = "operations_metadata_2_0_0.tpl"
const TplOperationsMetadataProvider1 = "operations_metadata_1_1_0.tpl"
const TplServiceProvider = "service-provider.tpl"
const TplFilterCapabilities = "filter_capabilities_2_0_0.tpl"
const TplFilterCapabilities1 = "filter_capabilities_1_1_0.tpl"
const TplWfs200 = "wfs_2_0_0.tpl"
const TplWfs110 = "wfs_1_1_0.tpl"

type WfsCapabilities struct {
	wfsTemplates          Executor
	Version               string
	Dataset               Dataset
	Service               Service
	Inspire               bool
	ServiceIdentification Identification
	ServiceProvider       Organization
	Features              []Feature
	ServiceDef            ServiceDef
}

func (c *WfsCapabilities) Build(writer io.Writer) error {
	return c.buildWfs(c.Version, writer)
}

func NewWfsCapabilities(wfs *template.Template, def ServiceDef, version string) *WfsCapabilities {
	c := &WfsCapabilities{wfsTemplates: TemplateExecutor{wfs}, ServiceDef: def, Version: version}
	return c
}

func (c *WfsCapabilities) buildWfs(version string, writer io.Writer) error {
	switch version {
	case "2.0.0":
		return c.buildWfs2(writer)
	case "1.1.0":
		return c.buildWfs1(writer)
	}

	return fmt.Errorf(fmt.Sprintf("Invalid version : %s", version))
}

func (c *WfsCapabilities) buildWfs1(writer io.Writer) error {

	capabilitiesMap := make(map[string]interface{})

	serviceIdentification, err := c.GenerateServiceIdentification()
	if capabilitiesMap[TplServiceidentification] = serviceIdentification; err != nil {
		return err
	}

	serviceProvider, err := c.GenerateServiceProvider()
	if capabilitiesMap[TplServiceProvider] = serviceProvider; err != nil {
		return err
	}

	operationsMetadata, err := c.GenerateOperationsMetadata(TplOperationsMetadataProvider1)
	if capabilitiesMap[TplOperationsMetadataProvider1] = operationsMetadata; err != nil {
		return err
	}

	features, err := c.GenerateFeatures(TplFeatures110)
	if capabilitiesMap[TplFeatures110] = features; err != nil {
		return err
	}

	filters, err := c.GenerateFilters(TplFilterCapabilities1)
	if capabilitiesMap[TplFilterCapabilities1] = filters; err != nil {
		return err
	}

	capabilitiesMap["dataset"] = c.Dataset

	capabilities, err := c.generate(TplWfs110, capabilitiesMap)
	if err != nil {
		return err
	}

	node, err := xmlquery.Parse(strings.NewReader(capabilities))
	if err != nil { // not valid xml
		log.Fatalf("error: %v", err)
	}

	_, err = writer.Write([]byte(node.OutputXML(false)))

	if err != nil {
		return err
	}

	return nil
}

func (c *WfsCapabilities) buildWfs2(writer io.Writer) error {

	capabilitiesMap := make(map[string]interface{})

	serviceIdentification, err := c.GenerateServiceIdentification()
	if capabilitiesMap[TplServiceidentification] = serviceIdentification; err != nil {
		return err
	}

	serviceProvider, err := c.GenerateServiceProvider()
	if capabilitiesMap[TplServiceProvider] = serviceProvider; err != nil {
		return err
	}

	operationsMetadata, err := c.GenerateOperationsMetadata(TplOperationsMetadataProvider)
	if capabilitiesMap[TplOperationsMetadataProvider] = operationsMetadata; err != nil {
		return err
	}

	features, err := c.GenerateFeatures(TplFeatures200)
	if capabilitiesMap[TplFeatures200] = features; err != nil {
		return err
	}

	filters, err := c.GenerateFilters(TplFilterCapabilities)
	if capabilitiesMap[TplFilterCapabilities] = filters; err != nil {
		return err
	}

	capabilitiesMap["dataset"] = c.Dataset

	capabilities, err := c.generate(TplWfs200, capabilitiesMap)
	if err != nil {
		return err
	}

	if !IsValidXML([]byte(capabilities)) { // not valid xml
		log.Fatalf("Invalid XML : %s\n", capabilities)
	}

	_, err = writer.Write([]byte(capabilities))
	if err != nil {
		return err
	}

	return nil
}

func (c *WfsCapabilities) AddServiceIdentification(serviceIdentification Identification) *WfsCapabilities {
	c.ServiceIdentification = serviceIdentification
	return c
}

func (c *WfsCapabilities) GenerateServiceIdentification() (string, error) {
	c.ServiceIdentification.Version = c.Version
	return c.generate(TplServiceidentification, c.ServiceIdentification)
}

func (c *WfsCapabilities) AddFeatures(dataset Dataset, features []Feature, def ServiceDef) *WfsCapabilities {

	c.Dataset = dataset
	c.Features = features
	c.ServiceDef = def
	return c
}

func (c *WfsCapabilities) GenerateFeatures(template string) (string, error) {
	data := make(map[string]interface{})

	data["features"] = c.Features
	data["dataset"] = c.Dataset
	data["crs"] = c.ServiceDef.Crs
	data["boundingbox"] = c.ServiceDef.Boundingbox

	return c.generate(template, data)
}

func (c *WfsCapabilities) AddDataset(dataset Dataset) *WfsCapabilities {
	c.Dataset = dataset
	return c
}

func (c *WfsCapabilities) AddServiceProvider(organization Organization) *WfsCapabilities {
	c.ServiceProvider = organization
	return c
}

func (c *WfsCapabilities) GenerateServiceProvider() (string, error) {
	return c.generate(TplServiceProvider, c.ServiceProvider)
}

func (c *WfsCapabilities) AddOperationsMetadata(dataset Dataset, service Service) *WfsCapabilities {
	c.Dataset = dataset
	c.Service = service
	return c
}

func (c *WfsCapabilities) GenerateOperationsMetadata(template string) (string, error) {

	data := make(map[string]interface{})

	data["service"] = c.Service
	data["dataset"] = c.Dataset

	return c.generate(template, data)
}

func (c *WfsCapabilities) AddFilters() *WfsCapabilities {
	// no op
	return c
}

func (c *WfsCapabilities) GenerateFilters(template string) (string, error) {

	return c.generate(template, []Filter{})
}

func (c WfsCapabilities) generate(templateName string, data interface{}) (string, error) {
	var out bytes.Buffer
	err := c.wfsTemplates.ExecuteTemplate(&out, templateName, data)
	return out.String(), err
}
