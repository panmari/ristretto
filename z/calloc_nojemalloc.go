// Copyright 2020 The LevelDB-Go and Pebble Authors. All rights reserved. Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.

// +build !jemalloc

package z

import (
	"fmt"
)

// Provides versions of Calloc, CallocNoRef, etc when jemalloc is not available
// (eg: build without jemalloc tag).

// Calloc allocates a slice of size n.
func Calloc(n int) []byte {
	return make([]byte, n)
}

// CallocNoRef will not give you memory back without jemalloc.
func CallocNoRef(n int) []byte {
	// We do the add here just to stay compatible with a corresponding Free call.
	return nil
}

// Free does not do anything in this mode.
func Free(b []byte) {}

func StatsPrint() {
	fmt.Println("Using Go memory")
}

// ReadMemStats doesn't do anything since all the memory is being managed
// by the Go runtime.
func ReadMemStats(_ *MemStats) { return }
