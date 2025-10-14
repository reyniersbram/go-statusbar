package xlib

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type C_Object[T any] interface {
	c() T
}

type (
	Display C.Display
	Window  C.Window
)

func (dpy *Display) c() *C.Display {
	return (*C.Display)(dpy)
}

func (w Window) c() C.Window {
	return C.Window(w)
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

func XDefaultScreen(dpy *Display) int {
	return int(C.XDefaultScreen(dpy.c()))
}

func XDefaultRootWindow(dpy *Display) Window {
	root := C.XDefaultRootWindow((*C.Display)(dpy))
	return Window(root)
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

func XFlush(dpy *Display) {
	C.XFlush(dpy.c())
}

func XCreateSimpleWindow(
	dpy *Display,
	w Window,
	x, y int,
	width, height uint,
	border_width uint,
	border uint,
	background uint,
) Window {
	return Window(C.XCreateSimpleWindow(
		dpy.c(),
		w.c(),
		C.int(x),
		C.int(y),
		C.uint(width),
		C.uint(height),
		C.uint(border_width),
		C.ulong(border),
		C.ulong(background),
	))
}

func XClearWindow(dpy *Display, w Window) {
	C.XClearWindow(dpy.c(), w.c())
}

func XMapWindow(dpy *Display, w Window) {
	C.XMapWindow(dpy.c(), w.c())
}

const ExposureMask = C.ExposureMask
const PropertyChangeMask = C.PropertyChangeMask
const Expose = C.Expose
const PropertyNotify = C.PropertyNotify

func XSelectInput(dpy *Display, w Window, mask int) {
	C.XSelectInput(dpy.c(), w.c(), C.long(mask))
}

type Event struct {
	Type int
}

func XNextEvent(dpy *Display) *Event {
	var ev C.XEvent
	C.XNextEvent(dpy.c(), &ev)
	evType := (*C.XAnyEvent)(unsafe.Pointer(&ev))._type
	return &Event{
		Type: int(evType),
	}
}

func XFetchName(dpy *Display, w Window) (string, error) {
	var name *C.char
	status := C.XFetchName(dpy.c(), w.c(), &name)
	if status == 0 || name == nil {
		return "", fmt.Errorf("Failed to fetch WM_NAME: %d", status)
	}
	defer C.XFree(unsafe.Pointer(name))
	return C.GoString(name), nil
}

func XLoadFont(dpy *Display, font string) C.Font {
	return C.XLoadFont(dpy.c(), C.CString(font))
}

func XCreateGC(dpy *Display, w Window, valuemask int) C.GC {
	return C.XCreateGC(dpy.c(), w.c(), C.ulong(valuemask), nil)
}

func XSetFont(dpy *Display, gc C.GC, font C.Font) {
	C.XSetFont(dpy.c(), gc, font)
}

func XSetForeground(dpy *Display, gc C.GC, foreground uint) {
	C.XSetForeground(dpy.c(), gc, C.ulong(foreground))
}

func XDrawString(dpy *Display, w Window, gc C.GC, x, y int, text string) {
	C.XDrawString(dpy.c(), w.c(), gc, C.int(x), C.int(y), C.CString(text), C.int(len(text)))
}
