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

type ServiceProvider struct {
}

type OperationsMetadata struct {
}

type FeatureTypeList struct {
}

type FilterCapabilities struct {
	Conformance struct {
		Constraint []struct {
			Name         string `xml:"name,attr"`
			NoValues     string `xml:"ows:NoValues"`
			DefaultValue string `xml:"ows:DefaultValue"`
		} `xml:"fes:Constraint"`
	} `xml:"fes:Conformance"`
	IDCapabilities struct {
		ResourceIdentifier struct {
			Name string `xml:"name,attr"`
		} `xml:"fes:ResourceIdentifier"`
	} `xml:"fes:Id_Capabilities"`
	ScalarCapabilities struct {
		LogicalOperators    string `xml:"LogicalOperators"`
		ComparisonOperators struct {
			ComparisonOperator []struct {
				Name string `xml:"name,attr"`
			} `xml:"fes:ComparisonOperator"`
		} `xml:"fes:ComparisonOperators"`
	} `xml:"fes:Scalar_Capabilities"`
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
	// // NO TemporalCapabilities!!!
	// TemporalCapabilities struct {
	// 	TemporalOperands struct {
	// 		TemporalOperand []struct {
	// 			Name string `xml:"name,attr"`
	// 		} `xml:"fes:TemporalOperand"`
	// 	} `xml:"fes:TemporalOperands"`
	// 	TemporalOperators struct {
	// 		TemporalOperator struct {
	// 			Name string `xml:"name,attr"`
	// 		} `xml:"fes:TemporalOperator"`
	// 	} `xml:"fes:TemporalOperators"`
	// } `xml:"fes:Temporal_Capabilities"`
}
