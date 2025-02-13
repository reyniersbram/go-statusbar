package xlib

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
import "C"
import "errors"

type C_Object[T any] interface {
	c() T
}

type Display (C.Display)
type Window (C.Window)

func (dpy *Display) c() *C.Display {
	return (*C.Display)(dpy)
}

func (w Window) c() C.Window {
	return (C.Window)(w)
}

func XOpenDisplay() (*Display, error) {
	dpy := C.XOpenDisplay(nil)
	if dpy == nil {
		return nil, errors.New("could not open display")
	}
	return (*Display)(dpy), nil
}

func XCloseDisplay(dpy *Display) {
	C.XCloseDisplay(dpy.c())
}

func XDefaultRootWindow(dpy *Display) Window {
	root := C.XDefaultRootWindow((*C.Display)(dpy))
	return (Window)(root)

}

func XStoreName(dpy *Display, root Window, text string) {
	C.XStoreName(dpy.c(), root.c(), (C.CString)(text))
	C.XSync(dpy.c(), C.False)
}

func XSync(dpy *Display, discard bool) {
	if discard {
		C.XSync(dpy.c(), C.True)
	} else {
		C.XSync(dpy.c(), C.False)
	}
}

