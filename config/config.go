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
	Wfs200  WFS200Config  `yaml:"wfs200"`
	Wms130  WMS130Config  `yaml:"wms130"`
	Wmts100 WMTS100Config `yaml:"wmts100"`
	Wcs201  WCS201Config  `yaml:"wcs201"`
}

// The WFS200Config service struct
type WFS200Config struct {
	Filename              string                       `yaml:"filename"`
	ServiceIdentification wfs200.ServiceIdentification `yaml:"serviceidentification"`
	FeatureTypeList       wfs200.FeatureTypeList       `yaml:"featuretypelist"`
	ExtendedCapabilities  wfs200.ExtendedCapabilities  `yaml:"extendedcapabilities"`
}

// The WMS130Config service struct
type WMS130Config struct {
	Filename   string            `yaml:"filename"`
	Service    wms130.Service    `yaml:"service"`
	Capability wms130.Capability `yaml:"capability"`
}

// The WMTS100Config service struct
type WMTS100Config struct {
	Filename              string                        `yaml:"filename"`
	ServiceIdentification wmts100.ServiceIdentification `yaml:"serviceidentification"`
	Contents              wmts100.Contents              `yaml:"contents"`
	ServiceMetadataURL    wmts100.ServiceMetadataURL    `yaml:"servicemetadataurl"`
}

// The WCS201Config service struct
type WCS201Config struct {
	Filename              string                       `yaml:"filename"`
	ServiceIdentification wcs201.ServiceIdentification `yaml:"serviceidentification"`
}
