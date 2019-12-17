package wms130

import (
	"encoding/xml"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func GetBase() Wms130 {
	wms130 := Wms130{}
	base, err := ioutil.ReadFile("./wms130/wms130.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wms130); err != nil {
		log.Fatalf("error: %v", err)
	}
	return wms130
}

type Wms130 struct {
	XMLName        xml.Name `xml:"WMS_Capabilities"`
	WMS_Namespaces `yaml:"namespaces"`
	Service        Service    `xml:"Service" yaml:"service"`
	Capability     Capability `xml:"Capability" yaml:"capability"`
}

type WMS_Namespaces struct {
	XmlnsWMS           string `xml:"xmlns,attr" yaml:"wms"`                                    //http://www.opengis.net/wms
	XmlnsSLD           string `xml:"xmlns:sld,attr" yaml:"sld"`                                //http://www.opengis.net/sld
	XmlnsXlink         string `xml:"xmlns:xlink,attr" yaml:"xlink"`                            //http://www.w3.org/1999/xlink
	XmlnsXSI           string `xml:"xmlns:xsi,attr" yaml:"xsi"`                                //http://www.w3.org/2001/XMLSchema-instance
	XmlnsInspireCommon string `xml:"xmlns:inspire_common,attr,omitempty" yaml:"inspirecommon"` //http://inspire.ec.europa.eu/schemas/common/1.0
	XmlnsInspireVs     string `xml:"xmlns:inspire_vs,attr,omitempty" yaml:"inspirevs"`         //http://inspire.ec.europa.eu/schemas/inspire_vs/1.0
	Version            string `xml:"version,attr" yaml:"version"`
	SchemaLocation     string `xml:"xsi:schemaLocation,attr" yaml:"schemalocation"`
}

type Service struct {
	Name        string `xml:"Name" yaml:"name"`
	Title       string `xml:"Title" yaml:"title"`
	Abstract    string `xml:"Abstract" yaml:"abstract"`
	KeywordList struct {
		Keyword []string `xml:"Keyword" yaml:"keyword"`
	} `xml:"KeywordList" yaml:"keywordlist"`
	OnlineResource struct {
		Xlink string `xml:"xlink,attr" yaml:"xlink"`
		Href  string `xml:"href,attr" yaml:"href"`
	} `xml:"OnlineResource" yaml:"onlineresource"`
	ContactInformation struct {
		ContactPersonPrimary struct {
			ContactPerson       string `xml:"ContactPerson" yaml:"contactperson"`
			ContactOrganization string `xml:"ContactOrganization" yaml:"contactorganization"`
		} `xml:"ContactPersonPrimary" yaml:"contactpersonprimary"`
		ContactPosition string `xml:"ContactPosition" yaml:"contactposition"`
		ContactAddress  struct {
			AddressType     string `xml:"AddressType" yaml:"addresstype"`
			Address         string `xml:"Address" yaml:"address"`
			City            string `xml:"City" yaml:"city"`
			StateOrProvince string `xml:"StateOrProvince" yaml:"stateorprovince"`
			PostCode        string `xml:"PostCode" yaml:"postalcode"`
			Country         string `xml:"Country" yaml:"country"`
		} `xml:"ContactAddress" yaml:"contactaddress"`
		ContactVoiceTelephone        string `xml:"ContactVoiceTelephone" yaml:"contactvoicetelephone"`
		ContactFacsimileTelephone    string `xml:"ContactFacsimileTelephone" yaml:"contactfacsimiletelephone"`
		ContactElectronicMailAddress string `xml:"ContactElectronicMailAddress" yaml:"contactelectronicmailaddress"`
	} `xml:"ContactInformation"`
	Fees              string `xml:"Fees" yaml:"fees"`
	AccessConstraints string `xml:"AccessConstraints" yaml:"accessconstraints"`
	MaxWidth          string `xml:"MaxWidth" yaml:"maxwidth"`
	MaxHeight         string `xml:"MaxHeight" yaml:"maxheight"`
}

type Capability struct {
	Request struct {
		GetCapabilities struct {
			Format  string  `xml:"Format" yaml:"format"`
			DCPType DCPType `xml:"DCPType" yaml:"dcptype"`
		} `xml:"GetCapabilities" yaml:"format"`
		GetMap struct {
			Format  []string `xml:"Format" yaml:"format"`
			DCPType DCPType  `xml:"DCPType" yaml:"dcptype"`
		} `xml:"GetMap" yaml:"format"`
		GetFeatureInfo struct {
			Format  []string `xml:"Format" yaml:"format"`
			DCPType DCPType  `xml:"DCPType" yaml:"dcptype"`
		} `xml:"GetFeatureInfo" yaml:"getfeatureinfo"`
	} `xml:"Request" yaml:"request"`
	Exception struct {
		Format []string `xml:"Format" yaml:"format"`
	} `xml:"Exception" yaml:"exception"`
	ExtendedCapabilities *WMS_1_3_0_ExtendedCapabilities `xml:"ExtendedCapabilities"`
	Layer                struct {
		Queryable   string `xml:"queryable,attr"`
		Name        string `xml:"Name"`
		Title       string `xml:"Title"`
		Abstract    string `xml:"Abstract"`
		KeywordList struct {
			Keyword []string `xml:"Keyword"`
		} `xml:"KeywordList"`
		CRS                     []string                `xml:"CRS"`
		EXGeographicBoundingBox EXGeographicBoundingBox `xml:"EX_GeographicBoundingBox"`
		BoundingBox             []BoundingBox           `xml:"BoundingBox"`
		Style                   Style                   `xml:"Style"`
		Layer                   []struct {
			Queryable string  `xml:"queryable,attr"`
			Opaque    string  `xml:"opaque,attr"`
			Cascaded  string  `xml:"cascaded,attr"`
			Name      string  `xml:"Name"`
			Title     string  `xml:"Title"`
			Abstract  string  `xml:"Abstract"`
			Style     []Style `xml:"Style"`
			Layer     []struct {
				Queryable   string `xml:"queryable,attr"`
				Opaque      string `xml:"opaque,attr"`
				Cascaded    string `xml:"cascaded,attr"`
				Name        string `xml:"Name"`
				Title       string `xml:"Title"`
				Abstract    string `xml:"Abstract"`
				KeywordList struct {
					Keyword []string `xml:"Keyword"`
				} `xml:"KeywordList"`
				EXGeographicBoundingBox EXGeographicBoundingBox `xml:"EX_GeographicBoundingBox"`
				BoundingBox             []BoundingBox           `xml:"BoundingBox"`
				AuthorityURL            struct {
					Name           string `xml:"name,attr"`
					OnlineResource struct {
						Xlink string `xml:"xlink,attr"`
						Href  string `xml:"href,attr"`
					} `xml:"OnlineResource"`
				} `xml:"AuthorityURL"`
				Identifier struct {
					Authority string `xml:"authority,attr"`
				} `xml:"Identifier"`
				MetadataURL struct {
					Type           string `xml:"type,attr"`
					Format         string `xml:"Format"`
					OnlineResource struct {
						Xlink string `xml:"xlink,attr"`
						Type  string `xml:"type,attr"`
						Href  string `xml:"href,attr"`
					} `xml:"OnlineResource"`
				} `xml:"MetadataURL"`
				Style []Style `xml:"Style"`
			} `xml:"Layer"`
			KeywordList struct {
				Keyword []string `xml:"Keyword"`
			} `xml:"KeywordList"`
			EXGeographicBoundingBox EXGeographicBoundingBox `xml:"EX_GeographicBoundingBox"`
			BoundingBox             []BoundingBox           `xml:"BoundingBox"`
			AuthorityURL            struct {
				Name           string `xml:"name,attr"`
				OnlineResource struct {
					Xlink string `xml:"xlink,attr"`
					Href  string `xml:"href,attr"`
				} `xml:"OnlineResource"`
			} `xml:"AuthorityURL"`
			Identifier struct {
				Authority string `xml:"authority,attr"`
			} `xml:"Identifier"`
			MetadataURL struct {
				Type           string `xml:"type,attr"`
				Format         string `xml:"Format"`
				OnlineResource struct {
					Xlink string `xml:"xlink,attr"`
					Type  string `xml:"type,attr"`
					Href  string `xml:"href,attr"`
				} `xml:"OnlineResource"`
			} `xml:"MetadataURL"`
		} `xml:"Layer"`
	} `xml:"Layer"`
}

type WMS_1_3_0_ExtendedCapabilities struct {
	MetadataUrl struct {
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
	} `xml:"inspire_vs:ResponseLanguage" yaml:"responselanguage"`
}

type EXGeographicBoundingBox struct {
	WestBoundLongitude string `xml:"westBoundLongitude"`
	EastBoundLongitude string `xml:"eastBoundLongitude"`
	SouthBoundLatitude string `xml:"southBoundLatitude"`
	NorthBoundLatitude string `xml:"northBoundLatitude"`
}

type BoundingBox struct {
	CRS  string `xml:"CRS,attr"`
	Minx string `xml:"minx,attr"`
	Miny string `xml:"miny,attr"`
	Maxx string `xml:"maxx,attr"`
	Maxy string `xml:"maxy,attr"`
}

type Style struct {
	Name      string `xml:"Name"`
	Title     string `xml:"Title"`
	LegendURL struct {
		Width          string `xml:"width,attr"`
		Height         string `xml:"height,attr"`
		Format         string `xml:"Format"`
		OnlineResource struct {
			Xlink string `xml:"xlink,attr"`
			Type  string `xml:"type,attr"`
			Href  string `xml:"href,attr"`
		} `xml:"OnlineResource"`
	} `xml:"LegendURL"`
}

type DCPType struct {
	HTTP struct {
		Get struct {
			OnlineResource struct {
				Xlink string `xml:"xmlns:xlink,attr" yaml:"xlink"`
				Href  string `xml:"href,attr" yaml:"href"`
			} `xml:"OnlineResource" yaml:"onlineresource"`
		} `xml:"Get" yaml:"Get"`
		Post *Post `xml:"Post" yaml:"post"`
	} `xml:"HTTP" yaml:"http"`
}

type Post struct {
	OnlineResource struct {
		Xlink string `xml:"xmlns:xlink,attr" yaml:"xlink"`
		Href  string `xml:"href,attr" yaml:"href"`
	} `xml:"OnlineResource" yaml:"onlineresource"`
}
