package wfs2_0_0

import "encoding/xml"

type WFS_2_0_0 struct {
	XMLName               xml.Name `xml:"wfs:WFS_Capabilities"`
	WFS_Namespaces        `yaml:"namespaces"`
	ServiceIdentification ServiceIdentification `xml:"ows:ServiceIdentification" yaml:"serviceidentification"`
	ServiceProvider       ServiceProvider       `xml:"ows:ServiceProvider" yaml:"serviceprovider"`
	OperationsMetadata    OperationsMetadata    `xml:"ows:OperationsMetadata" yaml:"operationsmetadata"`
	FeatureTypeList       FeatureTypeList       `xml:"wfs:FeatureTypeList" yaml:"featuretypelist"`
	FilterCapabilities    FilterCapabilities    `xml:"fes:Filter_Capabilities" yaml:"filtercapabilities"`
}

type WFS_Namespaces struct {
	XmlnsGML           string `xml:"xmlns:gml,attr" yaml:"gml"`                                //http://www.opengis.net/gml/3.2
	XmlnsWFS           string `xml:"xmlns:wfs,attr" yaml:"wfs"`                                //http://www.opengis.net/wfs/2.0
	XmlnsOWS           string `xml:"xmlns:ows,attr" yaml:"ows"`                                //http://www.opengis.net/ows/1.1
	XmlnsXlink         string `xml:"xmlns:xlink,attr" yaml:"xlink"`                            //http://www.w3.org/1999/xlink
	XmlnsXSI           string `xml:"xmlns:xsi,attr" yaml:"xsi"`                                //http://www.w3.org/2001/XMLSchema-instance
	XmlnsFes           string `xml:"xmlns:fes,attr" yaml:"fes"`                                //http://www.opengis.net/fes/2.0
	XmlnsInspireCommon string `xml:"xmlns:inspire_common,attr,omitempty" yaml:"inspirecommon"` //http://inspire.ec.europa.eu/schemas/common/1.0
	XmlnsInspireDls    string `xml:"xmlns:inspire_dls,attr,omitempty" yaml:"inspiredls"`       //http://inspire.ec.europa.eu/schemas/inspire_dls/1.0
	XmlnsPrefix        string `xml:"xmlns:{{.Prefix}},attr" yaml:"prefix"`                     //namespace_uri placeholder
	Version            string `xml:"version,attr" yaml:"version"`
	SchemaLocation     string `xml:"xsi:schemaLocation,attr" yaml:"schemalocation"`
}

type ServiceIdentification struct {
	XMLName  xml.Name `xml:"ows:ServiceIdentification"`
	Title    string   `xml:"ows:Title" yaml:"title"`
	Abstract string   `xml:"ows:Abstract" yaml:"abstract"`
	Keywords struct {
		Keyword []string `xml:"ows:Keyword" yaml:"keyword"`
	} `xml:"ows:Keywords" yaml:"keywords"`
	ServiceType struct {
		Text      string `xml:",chardata" yaml:"text"`
		CodeSpace string `xml:"codeSpace,attr" yaml:"codespace"`
	} `xml:"ows:ServiceType"`
	ServiceTypeVersion string `xml:"ows:ServiceTypeVersion" yaml:"servicetypeversion"`
	Fees               string `xml:"ows:Fees" yaml:"fees"`
	AccessConstraints  string `xml:"ows:AccessConstraints" yaml:"accesscontraints"`
}

type ServiceProvider struct {
	XMLName      xml.Name `xml:"ows:ServiceProvider"`
	ProviderName string   `xml:"ows:ProviderName" yaml:"providername"`
	ProviderSite struct {
		Type string `xml:"xlink:type,attr" yaml:"type"`
		Href string `xml:"xlink:href,attr" yaml:"href"`
	} `xml:"ows:ProviderSite" yaml:"providersite"`
	ServiceContact struct {
		IndividualName string `xml:"ows:IndividualName" yaml:"individualname"`
		PositionName   string `xml:"ows:PositionName" yaml:"positionname"`
		ContactInfo    struct {
			Text  string `xml:",chardata"`
			Phone struct {
				Voice     string `xml:"ows:Voice" yaml:"voice"`
				Facsimile string `xml:"ows:Facsimile" yaml:"facsmile"`
			} `xml:"ows:Phone" yaml:"phone"`
			Address struct {
				DeliveryPoint         string `xml:"ows:DeliveryPoint" yaml:"deliverypoint"`
				City                  string `xml:"ows:City" yaml:"city"`
				AdministrativeArea    string `xml:"ows:AdministrativeArea" yaml:"administrativearea"`
				PostalCode            string `xml:"ows:PostalCode" yaml:"postalcode"`
				Country               string `xml:"ows:Country" yaml:"country"`
				ElectronicMailAddress string `xml:"ows:ElectronicMailAddress" yaml:"electronicmailaddress"`
			} `xml:"ows:Address" yaml:"address"`
			OnlineResource struct {
				Type string `xml:"xlink:type,attr" yaml:"type"`
				Href string `xml:"xlink:href,attr" yaml:"href"`
			} `xml:"ows:OnlineResource" yaml:"onlineresource"`
			HoursOfService      string `xml:"ows:HoursOfService" yaml:"hoursofservice"`
			ContactInstructions string `xml:"ows:ContactInstructions" yaml:"contactinstructions"`
		} `xml:"ows:ContactInfo" yaml:"contactinfo"`
		Role string `xml:"ows:Role" yaml:"role"`
	} `xml:"ows:ServiceContact" yaml:"servicecontact"`
}

type Post struct {
	Type string `xml:"xlink:type,attr" yaml:"type"`
	Href string `xml:"xlink:href,attr" yaml:"href"`
}

type OperationsMetadata struct {
	XMLName   xml.Name `xml:"ows:OperationsMetadata"`
	Operation []struct {
		Name string `xml:"name,attr"`
		DCP  struct {
			HTTP struct {
				Get struct {
					Type string `xml:"xlink:type,attr" yaml:"type"`
					Href string `xml:"xlink:href,attr" yaml:"href"`
				} `xml:"ows:Get" yaml:"get"`
				Post *Post `xml:"ows:Post",omitempty yaml:"post"`
			} `xml:"ows:HTTP" yaml:"http"`
		} `xml:"ows:DCP" yaml:"dcp"`
		Parameter []struct {
			Name          string `xml:"name,attr"`
			AllowedValues struct {
				Value []string `xml:"ows:Value"`
			} `xml:"ows:AllowedValues"`
		} `xml:"ows:Parameter"`
	} `xml:"ows:Operation"`
	Parameter struct {
		Name          string `xml:"name,attr" yaml:"name"`
		AllowedValues struct {
			Value []string `xml:"ows:Value" yaml:"value"`
		} `xml:"ows:AllowedValues" yaml:"allowedvalues"`
	} `xml:"ows:Parameter" yaml:"parameter"`
	Constraint []struct {
		Text          string         `xml:",chardata"`
		Name          string         `xml:"name,attr" yaml:"name"`
		NoValues      *string        `xml:"ows:NoValues" yaml:"novalue"`
		DefaultValue  *string        `xml:"ows:DefaultValue" yaml:"defaultvalue"`
		AllowedValues *AllowedValues `xml:"ows:AllowedValues" yaml:"allowedvalues"`
	} `xml:"ows:Constraint" yaml:"constraint"`
	ExtendedCapabilities *WFS_2_0_0_ExtendedCapabilities `xml:"ows:ExtendedCapabilities" yaml:"extendedcapabilities"`
}

type AllowedValues struct {
	Value []string `xml:"ows:Value" yaml:"value"`
}

type WFS_2_0_0_ExtendedCapabilities struct {
	ExtendedCapabilities struct {
		Text        string `xml:",chardata"`
		MetadataURL struct {
			Type      string `xml:"xsi:type,attr" yaml:"type"`
			URL       string `xml:"inspire_common:URL" yaml:"url"`
			MediaType string `xml:"inspire_common:MediaType" yaml:"mediatype"`
		} `xml:"inspire_common:MetadataUrl" yaml:"metadataurl"`
		SupportedLanguages struct {
			DefaultLanguage struct {
				Language string `xml:"inspire_common:Language" yaml:"language"`
			} `xml:"inspire_common:DefaultLanguage" yaml:"defaultlanguage"`
		} `xml:"inspire_common:SupportedLanguages" yaml:"supportedlanguages"`
		ResponseLanguage struct {
			Language string `xml:"inspire_common:Language" yaml:"language"`
		} `xml:"inspire_common:ResponseLanguage" yaml:"responselanguage"`
		SpatialDataSetIdentifier struct {
			Code string `xml:"inspire_common:Code" yaml:"code"`
		} `xml:"inspire_dls:SpatialDataSetIdentifier" yaml:"spatialdatasetidentifier"`
	} `xml:"inspire_dls:ExtendedCapabilities" yaml:"extendedcapabilities"`
}

type FeatureTypeList struct {
	XMLName     xml.Name `xml:"wfs:FeatureTypeList"`
	FeatureType []struct {
		Name     string `xml:"wfs:Name" yaml:"name"`
		Title    string `xml:"wfs:Title" yaml:"title"`
		Abstract string `xml:"wfs:Abstract" yaml:"abstract"`
		Keywords struct {
			Keyword []string `xml:"ows:Keyword" yaml:"keyword"`
		} `xml:"ows:Keywords" yaml:"keywords"`
		DefaultCRS    string   `xml:"wfs:DefaultCRS" yaml:"defaultcrs"`
		OtherCRS      []string `xml:"wfs:OtherCRS" yaml:"othercrs"`
		OutputFormats struct {
			Format []string `xml:"wfs:Format" yaml:"format"`
		} `xml:"wfs:OutputFormats" yaml:"outputformats"`
		WGS84BoundingBox struct {
			Dimensions  string `xml:"dimensions,attr" yaml:"dimensions"`
			LowerCorner string `xml:"ows:LowerCorner" yaml:"lowercorner"`
			UpperCorner string `xml:"ows:UpperCorner" yaml:"uppercorner"`
		} `xml:"ows:WGS84BoundingBox" yaml:"wgs84boundingbox"`
		MetadataURL struct {
			Href string `xml:"xlink:href,attr" yaml:"href"`
		} `xml:"wfs:MetadataURL" yaml:"metadataurl"`
	} `xml:"wfs:FeatureType" yaml:"featuretype"`
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
		LogicalOperators    string `xml:"fes:LogicalOperators" yaml:"logicaloperators"`
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
