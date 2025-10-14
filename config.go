package main

import (
	"time"

	"github.com/reyniersbram/go-statusbar/components"
)

var Components = []components.Component{
	components.NewBattery("BAT0", time.Minute),
	components.NewDate("Mon 02 Jan 2006 15:04", 5*time.Second),
}
