package main

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/reyniersbram/go-statusbar/components"
	"github.com/reyniersbram/go-statusbar/internal/xlib"
)

func buildStatusLine(items []components.Component) string {
	var parts []string
	for _, item := range items {
		parts = append(parts, item.String())
	}
	return strings.Join(parts, " | ")
}

func loopComponent(component components.Component, notify chan struct{}) {
	tick := time.Tick(component.GetFrequency())
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
	components := []components.Component{
		components.NewBattery("BAT0", time.Minute),
		components.NewDate("Mon 02 Jan 2006 15:04", 5*time.Second),
	}
	for _, component := range components {
		component.Refresh()
	}
	statusline := buildStatusLine(components)
	xlib.XStoreName(dpy, root, statusline)

	notify := make(chan struct{}, len(components))
	for _, component := range components {
		go loopComponent(component, notify)
	}

	laststatus := statusline
	update := throttle(func() {
		statusline = buildStatusLine(components)
		if statusline != laststatus {
			xlib.XStoreName(dpy, root, statusline)
			laststatus = statusline
		}
	}, time.Second)
	for range notify {
		update()
	}
}
