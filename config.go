package main

import (
	"time"

	"github.com/reyniersbram/go-statusbar/components"
)

var Components = []Component{
	components.NewBattery(
		time.Minute,
		"BAT0",
		"{{.Icon}} {{.Capacity}}%",
	),
	components.NewDateTime(
		5*time.Second,
		"\U000f00f0 {{formatDate .Now \"Mon 02 Jan 15:04\"}}",
	),
}
