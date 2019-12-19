package wcs201

import (
	"encoding/xml"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// GetBase function to get a filled "template" based on the wmts100.yaml config
func GetBase() Wcs201 {
	wcs201 := Wcs201{}
	base, err := ioutil.ReadFile("./wcs201/wcs201.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wcs201); err != nil {
		log.Fatalf("error: %v", err)
	}
	return wcs201
}

type Wcs201 struct {
	XMLName               xml.Name `xml:"Capabilities"`
	Namespaces            `yaml:"namespaces"`
	ServiceIdentification ServiceIdentification `xml:"ServiceIdentification"`
	ServiceProvider       ServiceProvider       `xml:"ServiceProvider"`
	OperationsMetadata    OperationsMetadata    `xml:"OperationsMetadata"`
	ServiceMetadata       ServiceMetadata       `xml:"ServiceMetadata"`
	Contents              Contents              `xml:"Contents"`
}

// Namespaces struct containing the namespaces needed for the XML document
type Namespaces struct {
	XmlnsGML    string `xml:"xmlns:gml,attr" yaml:"gml"`       //http://www.opengis.net/gml/3.2
	XmlnsGMLcov string `xml:"xmlns:gmlcov,attr" yaml:"gmlcov"` //http://www.opengis.net/gmlcov/1.0
	XmlnsWCS    string `xml:"xmlns:wcs,attr" yaml:"wcs"`       //http://www.opengis.net/wcs/2.0
	XmlnsOWS    string `xml:"xmlns:ows,attr" yaml:"ows"`       //http://www.opengis.net/ows/1.1
	XmlnsOGC    string `xml:"xmlns:ogc,attr" yaml:"ogc"`       //http://www.opengis.net/ogc
	XmlnsXlink  string `xml:"xmlns:xlink,attr" yaml:"xlink"`   //http://www.w3.org/1999/xlink
	//XmlnsSWE                   string                `xml:"xmlns:swe,attr"`
	XmlnsXSI           string `xml:"xmlns:xsi,attr" yaml:"xsi"`                                //http://www.w3.org/2001/XMLSchema-instance
	XmlnsCrs           string `xml:"xmlns:crs,attr" yaml:"crs"`                                //http://www.opengis.net/wcs/crs/1.0
	XmlnsInt           string `xml:"xmlns:int,attr" yaml:"int"`                                //http://www.opengis.net/wcs/interpolation/1.0
	XmlnsInspireCommon string `xml:"xmlns:inspire_common,attr,omitempty" yaml:"inspirecommon"` //http://inspire.ec.europa.eu/schemas/common/1.0
	XmlnsInspireDls    string `xml:"xmlns:inspire_dls,attr,omitempty" yaml:"inspiredls"`       //http://inspire.ec.europa.eu/schemas/inspire_dls/1.0
	XmlnsPrefix        string `xml:"xmlns:{{.Prefix}},attr" yaml:"prefix"`                     //namespace_uri placeholder
	Version            string `xml:"version,attr" yaml:"version"`
	SchemaLocation     string `xml:"xsi:schemaLocation,attr" yaml:"schemalocation"`
}

type ServiceIdentification struct {
	Title    string `xml:"Title"`
	Abstract string `xml:"Abstract"`
	Keywords struct {
		Keyword []string `xml:"Keyword"`
	} `xml:"Keywords"`
	ServiceType struct {
		CodeSpace string `xml:"codeSpace,attr"`
	} `xml:"ServiceType"`
	ServiceTypeVersion []string `xml:"ServiceTypeVersion"`
	Profile            []string `xml:"Profile"`
	Fees               string   `xml:"Fees"`
	AccessConstraints  string   `xml:"AccessConstraints"`
}

type ServiceProvider struct {
	ProviderName string `xml:"ProviderName"`
	ProviderSite struct {
		Type string `xml:"type,attr"`
		Href string `xml:"href,attr"`
	} `xml:"ProviderSite"`
	ServiceContact struct {
		IndividualName string `xml:"IndividualName"`
		PositionName   string `xml:"PositionName"`
		ContactInfo    struct {
			Phone struct {
				Voice     string `xml:"Voice"`
				Facsimile string `xml:"Facsimile"`
			} `xml:"Phone"`
			Address struct {
				DeliveryPoint         string `xml:"DeliveryPoint"`
				City                  string `xml:"City"`
				AdministrativeArea    string `xml:"AdministrativeArea"`
				PostalCode            string `xml:"PostalCode"`
				Country               string `xml:"Country"`
				ElectronicMailAddress string `xml:"ElectronicMailAddress"`
			} `xml:"Address"`
			OnlineResource struct {
				Type string `xml:"type,attr"`
				Href string `xml:"href,attr"`
			} `xml:"OnlineResource"`
			HoursOfService      string `xml:"HoursOfService"`
			ContactInstructions string `xml:"ContactInstructions"`
		} `xml:"ContactInfo"`
		Role string `xml:"Role"`
	} `xml:"ServiceContact"`
}

type OperationsMetadata struct {
	Operation []struct {
		Name string `xml:"name,attr"`
		DCP  struct {
			HTTP struct {
				Get struct {
					Type string `xml:"type,attr"`
					Href string `xml:"href,attr"`
				} `xml:"Get"`
				Post struct {
					Type       string `xml:"type,attr"`
					Href       string `xml:"href,attr"`
					Constraint struct {
						Name          string `xml:"name,attr"`
						AllowedValues struct {
							Text  string `xml:",chardata"`
							Value string `xml:"Value"`
						} `xml:"AllowedValues"`
					} `xml:"Constraint"`
				} `xml:"Post"`
			} `xml:"HTTP"`
		} `xml:"DCP"`
	} `xml:"Operation"`
	ExtendedCapabilities ExtendedCapabilities `xml:"ExtendedCapabilities"`
}

type ExtendedCapabilities struct {
	ExtendedCapabilities struct {
		MetadataUrl struct {
			Type      string `xml:"type,attr"`
			URL       string `xml:"URL"`
			MediaType string `xml:"MediaType"`
		} `xml:"MetadataUrl"`
		SupportedLanguages struct {
			DefaultLanguage struct {
				Language string `xml:"Language"`
			} `xml:"DefaultLanguage"`
		} `xml:"SupportedLanguages"`
		ResponseLanguage struct {
			Language string `xml:"Language"`
		} `xml:"ResponseLanguage"`
		SpatialDataSetIdentifier struct {
			Code string `xml:"Code"`
		} `xml:"SpatialDataSetIdentifier"`
	} `xml:"ExtendedCapabilities"`
}

type ServiceMetadata struct {
	FormatSupported []string `xml:"formatSupported"`
	Extension       struct {
		InterpolationMetadata struct {
			InterpolationSupported []string `xml:"InterpolationSupported"`
		} `xml:"InterpolationMetadata"`
		CrsMetadata struct {
			CrsSupported []string `xml:"crsSupported"`
		} `xml:"CrsMetadata"`
	} `xml:"Extension"`
}

type Contents struct {
	CoverageSummary []struct {
		CoverageID      string `xml:"CoverageId"`
		CoverageSubtype string `xml:"CoverageSubtype"`
	} `xml:"CoverageSummary"`
}
