package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/reyniersbram/go-statusbar/internal/xlib"
)

// Component represents a component in the status bar.
// Each module is responsible for holding and updating its own data. The Refresh
// function will update the data of the component. The String function should
// only convert this data to a string.
type Component interface {
	fmt.Stringer
	// Refresh updates the component's internal data.
	// It reports whether the update was successful.
	Refresh() bool
	// GetDuration returns how many seconds should be between each refresh of the
	// component
	GetDuration() time.Duration
}

func buildStatusLine(items []Component) string {
	var parts []string
	for _, item := range items {
		parts = append(parts, item.String())
	}
	return strings.Join(parts, " | ")
}

func loopComponent(component Component, notify chan struct{}) {
	tick := time.Tick(component.GetDuration())
	for range tick {
		if component.Refresh() {
			notify <- struct{}{}
		} else {
			log.Println("Component was not able to refresh")
			return
		}
	}
}

func throttle(f func(), duration time.Duration) func() {
	var mux sync.Mutex
	waiting := false
	return func() {
		mux.Lock()
		defer mux.Unlock()
		if waiting {
			return
		}
		waiting = true
		go func() {
			time.Sleep(duration)
			mux.Lock()
			waiting = false
			mux.Unlock()
		}()
		f()
	}
}

func main() {
	dpy, err := xlib.XOpenDisplay()
	if err != nil {
		log.Fatal(err)
	}
	defer xlib.XCloseDisplay(dpy)
	root := xlib.XDefaultRootWindow(dpy)
	xlib.XStoreName(dpy, root, "initializing...")
	for _, component := range Components {
		component.Refresh()
	}
	statusline := buildStatusLine(Components)
	xlib.XStoreName(dpy, root, statusline)

	notify := make(chan struct{}, len(Components))
	for _, component := range Components {
		go loopComponent(component, notify)
	}

	lastStatus := statusline
	update := throttle(func() {
		statusline = buildStatusLine(Components)
		if statusline != lastStatus {
			xlib.XStoreName(dpy, root, statusline)
			lastStatus = statusline
		}
	}, time.Second)
	for range notify {
		update()
	}
}
