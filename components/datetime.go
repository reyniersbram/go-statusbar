package components

import (
	"text/template"
	"time"
)

const dateTemplateName = "date"

var defaultDateTimeFuncMap = template.FuncMap{
}

// DateTime represents the current date.
//
// The following template variables are available:
//
//   - Now: The current datetime
type DateTime struct {
	duration time.Duration
	tmpl     *template.Template
}

// NewDateTime initializes a new DateTime component. The tmplString is a Go
// text/template string that can include Date fields to format its display
// output.
//
// Example:
//
//	tmpl := "{{ .Now.Format \"Mon 02 Jan 15:04\" }}"
//	date := components.NewDateTime(5 * time.Second, tmpl)
func NewDateTime(
	duration time.Duration,
	tmplString string,
) DateTime {
	return NewDateTimeWithFuncMap(duration, tmplString, template.FuncMap{})
}

// NewDateTimeWithFuncMap initializes a new DateTime component.
// The provided funcMap merges with the default function map. If both maps
// contain the same key, the funcMap overrides the default.
func NewDateTimeWithFuncMap(
	duration time.Duration,
	tmplString string,
	funcMap template.FuncMap,
) DateTime {
	tmpl := template.Must(
		template.New(dateTemplateName).
			Funcs(mergeFuncMaps(defaultDateTimeFuncMap, funcMap)).
			Parse(tmplString),
	)
	date := DateTime{
		duration: duration,
		tmpl:     tmpl,
	}
	return date
}

// Now returns the current datetime.
func (d DateTime) Now() time.Time {
	return time.Now()
}

func (d DateTime) String() string {
	return executeTemplate(*d.tmpl, d)
}

func (d DateTime) GetDuration() time.Duration {
	return d.duration
}

func (d DateTime) Refresh() bool {
	return true
}
