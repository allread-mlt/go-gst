//lint:file-ignore U1000 Ignore all unused code, this is a work in progress

// +plugin:Name=gobin
// +plugin:Description=A bin element written in go
// +plugin:Version=v0.0.1
// +plugin:License=gst.LicenseLGPL
// +plugin:Source=go-gst
// +plugin:Package=examples
// +plugin:Origin=https://github.com/go-gst/go-gst
// +plugin:ReleaseDate=2021-01-18
//
// +element:Name=gobin
// +element:Rank=gst.RankNone
// +element:Impl=gobin
// +element:Subclass=gst.ExtendsBin
// +element:Interfaces=gst.InterfaceChildProxy
//
//go:generate gst-plugin-gen
package main

import "github.com/gotk3/gotk3/glib"

func main() {}

type gobin struct{}

func (g *gobin) New() glib.GoObjectSubclass { return &gobin{} }

func (g *gobin) ClassInit(klass *glib.ObjectClass) {}
