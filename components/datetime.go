package components

import (
	"text/template"
	"time"
)

const dateTemplateName = "date"

// DateTime represents the current date.
//
// The following template variables are available:
//
//   - Now: The current datetime
//   - formatDate: A helper function to format time using Go time layout syntax.
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
//	tmpl := "{{formatDate .Now \"Mon 02 Jan 15:04\"}}"
//	date := components.NewDateTime(5 * time.Second, tmpl)
func NewDateTime(
	duration time.Duration,
	tmplString string,
) *DateTime {
	funcs := template.FuncMap{
		"formatDate": func(t time.Time, layout string) string {
			return t.Format(layout)
		},
	}
	tmpl := template.Must(
		template.New(dateTemplateName).
			Funcs(funcs).
			Parse(tmplString),
	)
	date := &DateTime{
		duration: duration,
		tmpl:     tmpl,
	}
	return date
}

func (d DateTime) GetDuration() time.Duration {
	return d.duration
}

// Now returns the current datetime.
func (d DateTime) Now() time.Time {
	return time.Now()
}

func (d DateTime) String() string {
	return ExecuteTemplate(*d.tmpl, d)
}

func (d DateTime) Refresh() bool {
	return true
}
