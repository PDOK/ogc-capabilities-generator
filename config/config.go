package config

import (
	wcs201_response "github.com/pdok/ogc-specifications/pkg/wcs201/response"
	wfs200_response "github.com/pdok/ogc-specifications/pkg/wfs200/response"
	wms130_response "github.com/pdok/ogc-specifications/pkg/wms130/response"
	wmts100_response "github.com/pdok/ogc-specifications/pkg/wmts100/response"
)

// Config is the base struct for the config yaml
type Config struct {
	Global   Global   `yaml:"global"`
	Services Services `yaml:"services"`
}

// Global contains a collection of base var that are globally usable for the generation of the capabilities
type Global struct {
	Prefix                    string `yaml:"prefix"`
	Namespace                 string `yaml:"namespace"`
	Onlineresourceurl         string `yaml:"onlineresourceurl"`
	Path                      string `yaml:"path"`
	Version                   string `yaml:"version"`
	AdditionalSchemaLocations string `yaml:"additionalschemalocations"`
	Empty                     string
}

// Services contain a single service struct for every service type
type Services struct {
	WFS200Config  WFS200Config  `yaml:"wfs200"`
	WMS130Config  WMS130Config  `yaml:"wms130"`
	WMTS100Config WMTS100Config `yaml:"wmts100"`
	WCS201Config  WCS201Config  `yaml:"wcs201"`
}

// The WFS200Config service struct
type WFS200Config struct {
	Filename string                          `yaml:"filename"`
	Wfs200   wfs200_response.GetCapabilities `yaml:"definition"`
}

// The WMS130Config service struct
type WMS130Config struct {
	Filename string                          `yaml:"filename"`
	Wms130   wms130_response.GetCapabilities `yaml:"definition"`
}

// The WMTS100Config service struct
type WMTS100Config struct {
	Filename string                           `yaml:"filename"`
	Wmts100  wmts100_response.GetCapabilities `yaml:"definition"`
}

// The WCS201Config service struct
type WCS201Config struct {
	Filename string                          `yaml:"filename"`
	Wcs201   wcs201_response.GetCapabilities `yaml:"definition"`
}
