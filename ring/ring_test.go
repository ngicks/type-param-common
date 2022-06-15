// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ring

import (
	"testing"
)

func makeN(n int) Ring[int] {
	r := New[int](n)
	for i := 1; i <= n; i++ {
		r.Set(i)
		r = r.Next()
	}
	return r
}

func sumN(n int) int { return (n*n + n) / 2 }

func TestLink1(t *testing.T) {
	r1a := makeN(1)
	var r1b Ring[int]
	r2a := r1a.Link(r1b)
	if r2a != r1a {
		t.Errorf("a) 2-element link failed")
	}

	r2b := r2a.Link(r2a.Next())
	if r2b != r2a.Next() {
		t.Errorf("b) 2-element link failed")
	}
}
