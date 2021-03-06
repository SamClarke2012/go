// run

// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

var a uint8

func main() {
	b := int8(func() int32 { return -1 }())
	a = uint8(b)
	if int32(a) != 255 {
		// Failing case prints 'got 255 expected 255'
		println("got", a, "expected 255")
	}
}
