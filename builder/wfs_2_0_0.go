package builder

import "encoding/xml"

type WFS_2_0_0 struct {
	XMLName xml.Name `xml:"wfs:WFS_Capabilities"`
	WFS_Namespaces
	ServiceIdentification ServiceIdentification `xml:"ows:ServiceIdentification"`
	ServiceProvider       ServiceProvider       `xml:"ows:ServiceProvider"`
	OperationsMetadata    OperationsMetadata    `xml:"ows:OperationsMetadata"`
	FeatureTypeList       FeatureTypeList       `xml:"FeatureTypeList"`
	FilterCapabilities    FilterCapabilities    `xml:"fes:Filter_Capabilities"`
}

type WFS_Namespaces struct {
	XmlnsGML           string `xml:"xmlns:gml,attr"`                      //http://www.opengis.net/gml/3.2
	XmlnsWFS           string `xml:"xmlns:wfs,attr"`                      //http://www.opengis.net/wfs/2.0
	XmlnsOWS           string `xml:"xmlns:ows,attr"`                      //http://www.opengis.net/ows/1.1
	XmlnsXlink         string `xml:"xmlns:xlink,attr"`                    //http://www.w3.org/1999/xlink
	XmlnsXSI           string `xml:"xmlns:xsi,attr"`                      //http://www.w3.org/2001/XMLSchema-instance
	XmlnsFes           string `xml:"xmlns:fes,attr"`                      //http://www.opengis.net/fes/2.0
	XmlnsInspireCommon string `xml:"xmlns:inspire_common,attr,omitempty"` //http://inspire.ec.europa.eu/schemas/common/1.0
	XmlnsInspireDls    string `xml:"xmlns:inspire_dls,attr,omitempty"`    //http://inspire.ec.europa.eu/schemas/inspire_dls/1.0
	XmlnsPrefix        string `xml:"xmlns:prefix,attr"`                   //namespace_uri placeholder
	Version            string `xml:"version,attr"`
	SchemaLocation     string `xml:"xsi:schemaLocation,attr"`
}

type ServiceIdentification struct {
	XMLName  xml.Name `xml:"ows:ServiceIdentification"`
	Title    string   `xml:"ows:Title"`
	Abstract string   `xml:"ows:Abstract"`
	Keywords struct {
		Keyword []string `xml:"ows:Keyword"`
	} `xml:"ows:Keywords"`
	ServiceType struct {
		CodeSpace string `xml:"codeSpace,attr"`
	} `xml:"ows:ServiceType"`
	ServiceTypeVersion string `xml:"ows:ServiceTypeVersion"`
	Fees               string `xml:"ows:Fees"`
	AccessConstraints  string `xml:"ows:AccessConstraints"`
}

type OperationsMetadata struct {
	XMLName   xml.Name `xml:"ows:OperationsMetadata"`
	Text      string   `xml:",chardata"`
	Operation []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
		DCP  struct {
			Text string `xml:",chardata"`
			HTTP struct {
				Text string `xml:",chardata"`
				Get  struct {
					Text string `xml:",chardata"`
					Type string `xml:"xlink:type,attr"`
					Href string `xml:"xlink:href,attr"`
				} `xml:"ows:Get"`
				Post struct {
					Text string `xml:",chardata"`
					Type string `xml:"xlink:type,attr"`
					Href string `xml:"xlink:href,attr"`
				} `xml:"ows:Post"`
			} `xml:"ows:HTTP"`
		} `xml:"ows:DCP"`
		Parameter []struct {
			Text          string `xml:",chardata"`
			Name          string `xml:"name,attr"`
			AllowedValues struct {
				Text  string   `xml:",chardata"`
				Value []string `xml:"ows:Value"`
			} `xml:"ows:AllowedValues"`
		} `xml:"ows:Parameter"`
	} `xml:"ows:Operation"`
	Parameter struct {
		Text          string `xml:",chardata"`
		Name          string `xml:"name,attr"`
		AllowedValues struct {
			Text  string   `xml:",chardata"`
			Value []string `xml:"ows:Value"`
		} `xml:"ows:AllowedValues"`
	} `xml:"ows:Parameter"`
	Constraint []struct {
		Text          string `xml:",chardata"`
		Name          string `xml:"name,attr"`
		NoValues      string `xml:"ows:NoValues"`
		DefaultValue  string `xml:"ows:DefaultValue"`
		AllowedValues struct {
			Text  string   `xml:",chardata"`
			Value []string `xml:"ows:Value"`
		} `xml:"ows:AllowedValues"`
	} `xml:"ows:Constraint"`
	ExtendedCapabilities *ExtendedCapabilities `xml:"ows:ExtendedCapabilities"`
}

type ExtendedCapabilities struct {
	Text                 string `xml:",chardata"`
	ExtendedCapabilities struct {
		Text        string `xml:",chardata"`
		MetadataUrl struct {
			Text      string `xml:",chardata"`
			Type      string `xml:"type,attr"`
			URL       string `xml:"inspire_common:URL"`
			MediaType string `xml:"inspire_common:MediaType"`
		} `xml:"inspire_common:MetadataUrl"`
		SupportedLanguages struct {
			Text            string `xml:",chardata"`
			DefaultLanguage struct {
				Text     string `xml:",chardata"`
				Language string `xml:"inspire_common:Language"`
			} `xml:"inspire_common:DefaultLanguage"`
		} `xml:"inspire_common:SupportedLanguages"`
		ResponseLanguage struct {
			Text     string `xml:",chardata"`
			Language string `xml:"inspire_common:Language"`
		} `xml:"inspire_common:ResponseLanguage"`
		SpatialDataSetIdentifier struct {
			Text string `xml:",chardata"`
			Code string `xml:"inspire_common:Code"`
		} `xml:"inspire_dls:SpatialDataSetIdentifier"`
	} `xml:"inspire_dls:ExtendedCapabilities"`
}

type FeatureTypeList struct {
}

type FilterCapabilities struct {
	Conformance struct {
		Constraint []struct {
			Name         string `xml:"name,attr" yaml:"name"`
			NoValues     string `xml:"ows:NoValues" yaml:"novalues"`
			DefaultValue string `xml:"ows:DefaultValue" yaml:"defaultvalue"`
		} `xml:"fes:Constraint" yaml:"constraint"`
	} `xml:"fes:Conformance" yaml:"conformance"`
	IDCapabilities struct {
		ResourceIdentifier struct {
			Name string `xml:"name,attr" yaml:"name" `
		} `xml:"fes:ResourceIdentifier" yaml:"resourceidentifier"`
	} `xml:"fes:Id_Capabilities" yaml:"idcapabilities"`
	ScalarCapabilities struct {
		LogicalOperators    string `xml:"LogicalOperators" yaml:"logicaloperators"`
		ComparisonOperators struct {
			ComparisonOperator []struct {
				Name string `xml:"name,attr"`
			} `xml:"fes:ComparisonOperator" yaml:"comparisonoperator"`
		} `xml:"fes:ComparisonOperators" yaml:"comparisonoperators"`
	} `xml:"fes:Scalar_Capabilities" yaml:"scalarcapabilities"`
	SpatialCapabilities struct {
		GeometryOperands struct {
			GeometryOperand []struct {
				Name string `xml:"name,attr"`
			} `xml:"fes:GeometryOperand"`
		} `xml:"fes:GeometryOperands"`
		SpatialOperators struct {
			SpatialOperator []struct {
				Name string `xml:"name,attr"`
			} `xml:"fes:SpatialOperator"`
		} `xml:"fes:SpatialOperators"`
	} `xml:"fes:Spatial_Capabilities"`
	// NO TemporalCapabilities!!!
	TemporalCapabilities *TemporalCapabilities `xml:"fes:Temporal_Capabilities" yaml:"temporalcapabilities"`
}

type TemporalCapabilities struct {
	TemporalOperands struct {
		TemporalOperand []struct {
			Name string `xml:"name,attr" yaml:"name"`
		} `xml:"fes:TemporalOperand" yaml:"temporaloperand"`
	} `xml:"fes:TemporalOperands" yaml:"temporaloperands"`
	TemporalOperators struct {
		TemporalOperator []struct {
			Name string `xml:"name,attr,omitempty" yaml:"name"`
		} `xml:"fes:TemporalOperator" yaml:"temporaloperator"`
	} `xml:"fes:TemporalOperators" yaml:"temporaloperators"`
}
