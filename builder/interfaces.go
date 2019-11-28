package builder

import (
	"io"
	"text/template"
)

type Builder interface {
	Build(writer io.Writer) error
}

type Executor interface {
	ExecuteTemplate(wr io.Writer, templateName string, data interface{}) error
	Lookup(templateName string) *template.Template
}
