// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE-go file.

package main

import "runtime"

type big [10 << 20]byte

func f(x interface{}, start int64) {
	x1, x := x, nil
	if delta := inuse() - start; delta < 9<<20 {
		println("after alloc: expected delta at least 9 MB, got", delta>>20, "MB")
	}
	g(x1)
	if delta := inuse() - start; delta > 1<<20 {
		println("after drop: expected delta below 1 MB, got", delta>>20, "MB")
	}
}

func main() {
	x := inuse()
	f(new(big), x)
}

func inuse() int64 {
	runtime.GC()
	var st runtime.MemStats
	runtime.ReadMemStats(&st)
	return int64(st.Alloc)
}

//go:noinline
func g(interface{}) {}
