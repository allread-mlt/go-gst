package gst

/*
#cgo pkg-config: gstreamer-1.0
#cgo CFLAGS: -Wno-deprecated-declarations -g -Wall
#include <gst/gst.h>
#include "gst.go.h"
*/
import "C"

// Clock is a go wrapper around a GstClock.
type Clock struct{ *Object }

// Instance returns the underlying GstClock instance.
func (c *Clock) Instance() *C.GstClock { return C.toGstClock(c.unsafe()) }