package config

import (
	"pdok-capabilities-gen/wcs201"
	"pdok-capabilities-gen/wfs200"
	"pdok-capabilities-gen/wms130"
	"pdok-capabilities-gen/wmts100"
)

// Config is the base struct for the config yaml
type Config struct {
	Global   Global   `yaml:"global"`
	Services Services `yaml:"services"`
}

// Global contains a collection of base var that are globally usable for the generation of the capabilities
type Global struct {
	Prefix            string `yaml:"prefix"`
	Namespace         string `yaml:"namespace"`
	Onlineresourceurl string `yaml:"onlineresourceurl"`
	Path              string `yaml:"path"`
	Version           string `yaml:"version"`
	Empty             string
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
	Filename string        `yaml:"filename"`
	Wfs200   wfs200.Wfs200 `yaml:"definition"`
}

// The WMS130Config service struct
type WMS130Config struct {
	Filename string        `yaml:"filename"`
	Wms130   wms130.Wms130 `yaml:"definition"`
}

// The WMTS100Config service struct
type WMTS100Config struct {
	Filename string          `yaml:"filename"`
	Wmts100  wmts100.Wmts100 `yaml:"definition"`
}

// The WCS201Config service struct
type WCS201Config struct {
	Filename string        `yaml:"filename"`
	Wcs201   wcs201.Wcs201 `yaml:"definition"`
}
