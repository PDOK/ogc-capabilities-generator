package builder

import (
	"encoding/xml"
	"io"
	"text/template"
)

type TemplateExecutor struct {
	Templates *template.Template
}

func (t TemplateExecutor) ExecuteTemplate(wr io.Writer, templateName string, data interface{}) error {
	return t.Templates.ExecuteTemplate(wr, templateName, data)
}

func (t TemplateExecutor) Lookup(templateName string) *template.Template {
	return t.Templates.Lookup(templateName)
}
func IsValidXML(data []byte) bool {
	err := xml.Unmarshal(data, new(interface{}))

	return err == nil
}
