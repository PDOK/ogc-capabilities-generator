package builder

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"text/template"
)

const TplWmts100 = "wmts_1_0_0.tpl"
const TplTilematrixset28992 = "tilematrixset_28992.tpl"
const TplTilematrixset3857 = "tilematrixset_3857.tpl"
const TplTilematrixset25831 = "tilematrixset_25831.tpl"
const TplLayersWmts = "layer.tpl"

type WmtsCapabilities struct {
	wmtsTemplates         Executor
	Version               string
	Dataset               Dataset
	Service               Service
	ServiceIdentification Identification
	ServiceProvider       Organization
	Layers                []Layer
	ServiceDef            ServiceDef
}

func (c *WmtsCapabilities) Build(writer io.Writer) error {
	return c.BuildWmts(c.Version, writer)
}

func NewWmtsCapabilities(wmts *template.Template, def ServiceDef, serviceVersion string) *WmtsCapabilities {
	c := &WmtsCapabilities{wmtsTemplates: TemplateExecutor{wmts}, ServiceDef: def, Version: serviceVersion}
	return c
}

func (c *WmtsCapabilities) BuildWmts(version string, writer io.Writer) error {
	switch version {
	case "1.0.0":
		return c.buildWmts1(writer)
	}

	return errors.New(fmt.Sprintf("Invalid version : %s", version))
}

func (c *WmtsCapabilities) buildWmts1(writer io.Writer) error {

	capabilitiesMap := make(map[string]interface{})

	layers, err := c.GenerateLayers(TplLayersWmts)
	if capabilitiesMap[TplLayersWmts] = layers; err != nil {
		return err
	}

	tilematrixset28892, err := c.GenerateTileMatrixSet(TplTilematrixset28992)
	if capabilitiesMap[TplTilematrixset28992] = tilematrixset28892; err != nil {
		return err
	}

	tilematrixset3857, err := c.GenerateTileMatrixSet(TplTilematrixset3857)
	if capabilitiesMap[TplTilematrixset3857] = tilematrixset3857; err != nil {
		return err
	}

	tilematrixset25831, err := c.GenerateTileMatrixSet(TplTilematrixset25831)
	if capabilitiesMap[TplTilematrixset25831] = tilematrixset25831; err != nil {
		return err
	}

	capabilitiesMap["dataset"] = c.Dataset

	capabilities, err := c.generate(TplWmts100, capabilitiesMap)
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

func (c *WmtsCapabilities) AddTileMatrixSet() *WmtsCapabilities {
	return c
}

func (c *WmtsCapabilities) GenerateTileMatrixSet(template string) (string, error) {

	var tilematrixset []Tilematrixset
	return c.generate(template, tilematrixset)
}

func (c *WmtsCapabilities) AddDataset(dataset Dataset) *WmtsCapabilities {
	c.Dataset = dataset
	return c
}

func (c *WmtsCapabilities) AddOperationsMetadata(dataset Dataset, service Service) *WmtsCapabilities {
	c.Dataset = dataset
	c.Service = service
	return c
}

func (c *WmtsCapabilities) GenerateOperationsMetadata(template string) (string, error) {

	data := make(map[string]interface{})

	data["service"] = c.Service
	data["dataset"] = c.Dataset

	return c.generate(template, data)
}

func (c WmtsCapabilities) generate(templateName string, data interface{}) (string, error) {
	var out bytes.Buffer
	err := c.wmtsTemplates.ExecuteTemplate(&out, templateName, data)
	return string(out.Bytes()), err
}

func (c *WmtsCapabilities) AddLayers(dataset Dataset, layers []Layer, def ServiceDef, service Service) *WmtsCapabilities {

	c.Service = service
	c.Dataset = dataset
	c.Layers = layers
	c.ServiceDef = def
	return c
}

func (c *WmtsCapabilities) GenerateLayers(template string) (string, error) {
	data := make(map[string]interface{})

	data["layers"] = c.Layers
	data["dataset"] = c.Dataset
	data["crs"] = c.ServiceDef.Crs
	data["srs"] = c.ServiceDef.Srs
	data["boundingbox"] = c.ServiceDef.Boundingbox
	data["service"] = c.Service

	return c.generate(template, data)
}
