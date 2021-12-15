// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (linux && cgo) || (darwin && cgo) || (freebsd && cgo)
// +build linux,cgo darwin,cgo freebsd,cgo

package plugin

/*
 */
import "C"

import (
	"sync"
)

func open(name string) (*Plugin, error) {
	panic("open")
}

func lookup(p *Plugin, symName string) (Symbol, error) {
	panic("lookup")
}

var (
	pluginsMu sync.Mutex
	plugins   map[string]*Plugin
)
