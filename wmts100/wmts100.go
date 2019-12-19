package wmts100

import (
	"encoding/xml"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// GetBase function to get a filled "template" based on the wmts100.yaml config
func GetBase() Wmts100 {
	wmts100 := Wmts100{}
	base, err := ioutil.ReadFile("./wmts100/wmts100.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err = yaml.Unmarshal(base, &wmts100); err != nil {
		log.Fatalf("error: %v", err)
	}
	return wmts100
}

// Wmts100 base struct
type Wmts100 struct {
	XMLName               xml.Name `xml:"Capabilities"`
	Namespaces            `yaml:"namespaces"`
	ServiceIdentification ServiceIdentification `xml:"ServiceIdentification" yaml:"serviceidentification"`
	Contents              Contents              `xml:"Contents" yaml:"contents"`
	ServiceMetadataURL    ServiceMetadataURL    `xml:"ServiceMetadataURL" yaml:"servicemetadataurl"`
}

// Namespaces struct containing the namespaces needed for the XML document
type Namespaces struct {
	Xmlns          string `xml:"xmlns,attr" yaml:"xmlns"`       //http://www.opengis.net/wmts/1.0
	XmlnsOws       string `xml:"xmlns:ows,attr" yaml:"ows"`     //http://www.opengis.net/ows/1.1
	XmlnsXlink     string `xml:"xmlns:xlink,attr" yaml:"xlink"` //http://www.w3.org/1999/xlink
	XmlnsXSI       string `xml:"xmlns:xsi,attr" yaml:"xsi"`     //http://www.w3.org/2001/XMLSchema-instance
	XmlnsGml       string `xml:"xmlns:gml,attr" yaml:"gml"`     //http://www.opengis.net/gml
	Version        string `xml:"version,attr" yaml:"version"`
	SchemaLocation string `xml:"xsi:schemaLocation,attr" yaml:"schemalocation"`
}

type ServiceIdentification struct {
	Title              string `xml:"Title"`
	Abstract           string `xml:"Abstract"`
	ServiceType        string `xml:"ServiceType"`
	ServiceTypeVersion string `xml:"ServiceTypeVersion"`
	Fees               string `xml:"Fees"`
	AccessConstraints  string `xml:"AccessConstraints"`
}

type Contents struct {
	Layer struct {
		Title            string `xml:"Title"`
		Abstract         string `xml:"Abstract"`
		WGS84BoundingBox struct {
			LowerCorner string `xml:"LowerCorner"`
			UpperCorner string `xml:"UpperCorner"`
		} `xml:"WGS84BoundingBox"`
		Identifier string `xml:"Identifier"`
		Style      struct {
			Identifier string `xml:"Identifier"`
		} `xml:"Style"`
		Format            string `xml:"Format"`
		TileMatrixSetLink []struct {
			TileMatrixSet string `xml:"TileMatrixSet"`
		} `xml:"TileMatrixSetLink"`
		ResourceURL struct {
			Format       string `xml:"format,attr"`
			ResourceType string `xml:"resourceType,attr"`
			Template     string `xml:"template,attr"`
		} `xml:"ResourceURL"`
	} `xml:"Layer"`
	TileMatrixSet []struct {
		Identifier   string `xml:"Identifier"`
		SupportedCRS string `xml:"SupportedCRS"`
		TileMatrix   []struct {
			Identifier       string `xml:"Identifier"`
			ScaleDenominator string `xml:"ScaleDenominator"`
			TopLeftCorner    string `xml:"TopLeftCorner"`
			TileWidth        string `xml:"TileWidth"`
			TileHeight       string `xml:"TileHeight"`
			MatrixWidth      string `xml:"MatrixWidth"`
			MatrixHeight     string `xml:"MatrixHeight"`
		} `xml:"TileMatrix"`
	} `xml:"TileMatrixSet"`
}

type ServiceMetadataURL struct {
	Href string `xml:"href,attr"`
}
