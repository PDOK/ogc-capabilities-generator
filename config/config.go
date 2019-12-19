package config

import (
	"pdok-capabilities-gen/wfs200"
	"pdok-capabilities-gen/wms130"
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
	Wfs200 WFS200Config `yaml:"wfs200"`
	Wms130 WMS130Config `yaml:"wms130"`
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
