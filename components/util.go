package components

import (
	"maps"
	"bytes"
	"log"
	"text/template"
)

// mergeFuncMaps merges two template.FuncMaps into one. The second map will
// overwrite values of the first map in case of name collisions.
func mergeFuncMaps(f1, f2 template.FuncMap) template.FuncMap {
	out := make(template.FuncMap, len(f1) + len(f2))
	maps.Copy(out, f1)
	maps.Copy(out, f2)
	return out
}

func executeTemplate(tmpl template.Template, data any) string {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		log.Println("template execution error: ", err)
		return "<error>"
	}
	return buf.String()
}
