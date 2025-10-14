package main

import (
	"fmt"
	"log"

	"github.com/reyniersbram/go-statusbar/internal/xlib"
)

func main() {
	display, err := xlib.XOpenDisplay()
	if err != nil {
		log.Fatal(err)
	}
	defer xlib.XCloseDisplay(display)

	root := xlib.XDefaultRootWindow(display)
	window := xlib.XCreateSimpleWindow(
		display,
		root,
		100,
		100,
		400,
		100,
		10,
		0x00ff00, // green
		0xff0000, // red
	)
	xlib.XSelectInput(display, window, xlib.ExposureMask)
	xlib.XSelectInput(display, root, xlib.PropertyChangeMask)

	font := xlib.XLoadFont(display, "fixed")
	gc := xlib.XCreateGC(display, window, 0)
	xlib.XSetFont(display, gc, font)
	xlib.XSetForeground(display, gc, 0x0000ff)

	xlib.XMapWindow(display, window)
	xlib.XFlush(display)

	for {
		ev := xlib.XNextEvent(display)
		switch ev.Type {
		case xlib.Expose:
			xlib.XClearWindow(display, window)
			name, err := xlib.XFetchName(display, root)
			fmt.Println(name)
			if err != nil {
				log.Println("Could not read WM_NAME: ", err)
				continue
			}
			xlib.XDrawString(display, window, gc, 20, 50, name)
		case xlib.PropertyNotify:
			log.Println("WM_NAME changed")
		}
	}
}
