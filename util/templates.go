package util

import (
	"text/template"
)

func GetTemplates(path string) *template.Template {
	return template.Must(template.New("templates").Funcs(
		template.FuncMap{
			"escape":    template.HTMLEscapeString,
			"isDefault": func(i int) bool { return i == 0 },
		}).ParseGlob(path))
}
