package builder

import (
	"pdok-capabilities-gen/wfs2_0_0"
	"pdok-capabilities-gen/wms1_3_0"
)

type Config struct {
	Global   Global   `yaml:"global"`
	Services Services `yaml:"services"`
}

type Global struct {
	Prefix            string `yaml:"prefix"`
	Namespace         string `yaml:"namespace"`
	Onlineresourceurl string `yaml:"onlineresourceurl"`
	Path              string `yaml:"path"`
	Version           string `yaml:"version"`
	Empty             string
}

type Services struct {
	WFS_2_0_0 WFS_2_0_0_config `yaml:"wfs_2_0_0"`
	WMS_1_3_0 WMS_1_3_0_config `yaml:"wms_1_3_0"`
}

type WFS_2_0_0_config struct {
	ServiceIdentification wfs2_0_0.ServiceIdentification          `yaml:"serviceidentification"`
	FeatureTypeList       wfs2_0_0.FeatureTypeList                `yaml:"featuretypelist"`
	ExtendedCapabilities  wfs2_0_0.WFS_2_0_0_ExtendedCapabilities `yaml:"extendedcapabilities"`
}

type WMS_1_3_0_config struct {
	Service    wms1_3_0.Service    `yaml:"service"`
	Capability wms1_3_0.Capability `yaml:"capability"`
}
