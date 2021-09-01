package validate

import (
	xsdvalidate "bitbucket.org/jazzserve/go-xsd-validate"
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"ogc-capabilities-generator/config"
	"strings"
	"text/template"
)

type schemaImport struct {
	Namespace      string `xml:"namespace,attr"`
	SchemaLocation string `xml:"schemaLocation,attr"`
}

type schema struct {
	XMLName            xml.Name       `xml:"schema"`
	ElementFormDefault string         `xml:"elementFormDefault,attr"`
	Xmlns              string         `xml:"xmlns,attr"`
	Import             []schemaImport `xml:"import"`
}

func generateXsd(schemaLocations string, global config.Global) (*schema, error) {
	schemaTemplate := template.Must(template.New("schemaTemplate").Parse(schemaLocations))
	var schemaLocationsBuffer bytes.Buffer
	err := schemaTemplate.Execute(&schemaLocationsBuffer, global)
	if err != nil {
		return nil, fmt.Errorf("cannot parse schemaLocations")
	}

	schemas := strings.Split(strings.TrimSpace(schemaLocationsBuffer.String()), " ")
	schemaLength := len(schemas)
	if schemaLength > 0 && schemaLength%2 > 0 {
		return nil, fmt.Errorf("schemaLocations has length %d, we expect each schemalocation to have a namespace", schemaLength)
	}

	var imports []schemaImport
	for i := 0; i <= len(schemas)-2; i += 2 {
		namespace := schemas[i]
		schemaLocation := schemas[i+1]
		if len(namespace) < 10 || len(schemaLocation) < 4 || schemaLocation[len(schemaLocation)-4:] != ".xsd" {
			return nil, fmt.Errorf("we expect each schemalocation to have a namespace, found namespace %s and schemalocation %s", namespace, schemaLocation)
		}
		imports = append(imports, schemaImport{namespace, schemaLocation})
	}

	return &schema{
		ElementFormDefault: "qualified",
		Xmlns:              "http://www.w3.org/2001/XMLSchema",
		Import:             imports,
	}, nil
}

func ValidateCapabilities(config *config.Config, capabilities []byte, schemaLocations string) {
	err := xsdvalidate.Init()
	if err != nil {
		log.Fatalf("Error initializing validation with error: %s", err)
	}

	xsd, err := generateXsd(schemaLocations, config.Global)
	if err != nil {
		log.Fatalf("Error initializing validation with error: %s", err)
	}

	defer xsdvalidate.Cleanup()
	var xsdHandler *xsdvalidate.XsdHandler
	if len(xsd.Import) <= 0 {
		log.Fatalf("Error parsing schemas, no xsd schema found, error: %s", err)
	} else if len(xsd.Import) == 1 {
		xsdUrl := xsd.Import[0].SchemaLocation
		xsdHandler, err = xsdvalidate.NewXsdHandlerUrl(xsdUrl, xsdvalidate.ParsErrVerbose)
		if err != nil {
			log.Fatalf("Error parsing xsd from url: %s error: %s", xsdUrl, err)
		}
	} else {
		schema, err := xml.MarshalIndent(xsd, "", " ")
		if err != nil {
			log.Fatalf("Error unmarshaling xsd schema with error: %s", err)
		}
		xsdHandler, err = xsdvalidate.NewXsdHandlerMem(schema, xsdvalidate.ParsErrVerbose)
		if err != nil {
			log.Fatalf("Error parsing xsd schema with error: %s", err)
		}
	}

	err = xsdHandler.ValidateMem(capabilities, xsdvalidate.ValidErrDefault)
	xsdHandler.Free()
	if err != nil {
		log.Fatalf("Error validating capabilities error: %s", err)
	}
}
