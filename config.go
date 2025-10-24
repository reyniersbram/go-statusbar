package main

import (
	"text/template"
	"time"

	"github.com/reyniersbram/go-statusbar/components"
)

func batteryIcon(b components.Battery) string {
	switch {
	case b.ChargingStatus == "Charging":
		return "\U000f0084"
	case b.ChargingStatus == "Not charging":
		return "\U000f1211"
	case b.ChargingStatus == "Unknown":
		return "\U000f0091"
	case b.Capacity < 10:
		return "\U000f007a"
	case b.Capacity < 20:
		return "\U000f007b"
	case b.Capacity < 30:
		return "\U000f007c"
	case b.Capacity < 40:
		return "\U000f007d"
	case b.Capacity < 50:
		return "\U000f007e"
	case b.Capacity < 60:
		return "\U000f007f"
	case b.Capacity < 70:
		return "\U000f0080"
	case b.Capacity < 80:
		return "\U000f0081"
	case b.Capacity < 90:
		return "\U000f0082"
	default:
		return "\U000f0079"
	}
}

var Components = []Component{
	components.NewBatteryWithFuncMap(
		time.Minute,
		"BAT0",
		"{{ batteryIcon . }} {{.Capacity}}%",
		template.FuncMap{
			"batteryIcon": batteryIcon,
		},
	),
	components.NewDateTime(
		5*time.Second,
		"\U000f00f0 {{ .Now.Format \"Mon 02 Jan 15:04\" }}",
	),
}
