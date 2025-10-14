package components

import (
	"bytes"
	"log"
	"text/template"
	"time"
)

type Date struct {
	Ticker
	layout string
	tmpl   *template.Template
}

func NewDate(layout string, duration time.Duration) *Date {
	tmpl := template.Must(
		template.New("date").
			Parse("{{.Icon}} {{.GetDate}}"))
	date := &Date{
		Ticker: Ticker{Duration: duration},
		layout: layout,
		tmpl:   tmpl,
	}
	return date
}

func (d Date) GetDate() string {
	if d.layout == "" {
		return time.Now().Format(time.UnixDate)
	}
	return time.Now().Format(d.layout)
}

func (d Date) String() string {
	var buf bytes.Buffer
	err := d.tmpl.Execute(&buf, d)
	if err != nil {
		log.Println("template execution error: ", err)
	}
	return buf.String()
}

func (d Date) Refresh() bool {
	return true
}

func (d Date) Icon() string {
	return "\U000f00f0"
}
