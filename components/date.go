package components

import (
	"time"
)

type Date struct {
	Ticker
	layout string
}

func NewDate(layout string, duration time.Duration) *Date {
	date := &Date{
		Ticker: Ticker{Frequency: duration},
		layout: layout,
	}
	return date
}

func (d Date) String() string {
	if d.layout == "" {
		return time.Now().Format(time.UnixDate)
	}
	return time.Now().Format(d.layout)
}

func (d Date) Refresh() bool {
	return true
}
