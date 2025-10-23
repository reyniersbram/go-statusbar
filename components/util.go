package components

import (
	"bytes"
	"log"
	"text/template"
)

func ExecuteTemplate(tmpl template.Template, data any) string {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		log.Println("template execution error: ", err)
		return "<error>"
	}
	return buf.String()
}
