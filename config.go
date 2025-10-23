package main

import (
	"time"

	"github.com/reyniersbram/go-statusbar/components"
)

var Components = []Component{
	components.NewBattery("BAT0", time.Minute),
	components.NewDate("Mon 02 Jan 15:04", 5*time.Second),
}
