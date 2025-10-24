# go-statusbar

A lightweight, extensible status bar for X11 window managers that updates the
root window name (compatible with WMs like **dwm**, **xmonad** and **bspwm**).

## Features:

- **Fully configurable in Go**: define your own components and layouts using Go
code.
- **Template-based output**: format component output using Go's
[text/template](https://pkg.go.dev/text/template) syntax.
- **Built-in system components**: includes ready-to-use modules like battery status,
  date/time, and more.
- **Concurrent design**: each component runs in its own goroutine for efficient and
  responsive updates.
- **Minimal dependencies**: no external configuration format required.

## Installation

### Requirements

- Go 1.23+ (`cgo` is required for the Xlib bindings)
- X11 environment with a compatible window manager
- (Optional) A Nerd Font for icons

```sh
git clone git@github.com:reyniersbram/go-statusbar.git
cd go-statusbar
go install
```

After installation, the `go-statusbar` binary will be available in your
`$GOPATH/bin` or `$HOME/go/bin` directory.

## Usage

Add the following line to your `.xinitrc` or window manager autostart script::

```sh
go-statusbar &
```

This will start `go-statusbar` in the background and continuously update your
X11 root window title with the status bar output.

## Configuration

*coming soon...*

## Components Overview

`go-statusbar` ships with several built-in components, including:

- **Battery**: monitors capacity and charging status via
`/sys/class/power_supply/`
- **DateTime**: displays the current date and time
- *(More components coming soon)*
- *(More components can easily be added by implementing the `Component`
interface)*

Each component's output can be customized via Go templates, allowing for rich,
dynamic formatting.

## Design Philosophy

1. **Simplicity**: small, simple, dependency-free, easy to hack
2. **Extensibility**: users can define their own components, templates, and
   template functions.
3. **Performance**: non-blocking design using goroutines for each component
   ensures responsive updates.

## Inspiration

This project is inspired by [dwmstatus](https://git.suckless.org/dwmstatus/), a
minimalist status bar written in C.  
`go-statusbar` provides a similar lightweight design, but in pure Go, making it
an easier starting point for developers who are comfortable with Go but are not
that familiar with C.  
By leveraging goroutines, it also avoids blocking and freezing issues that can
occur when components take a long time to refresh their data.
