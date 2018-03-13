// Copyright 2013 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"fmt"
	"strings"
	"testing"
)

var uint32Tests = []struct {
	where  string
	in     []string
	i      uint32
	uint32 uint32
	err    string
}{
	{
		loc(),
		[]string{},
		17, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i", "1", "--uint32", "2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "--uint32=2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "-i2"},
		2, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i=1"},
		17, 42,
		"test: not a valid number: =1\n",
	},
	{
		loc(),
		[]string{"test", "-i0x20"},
		0x20, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i010"},
		8, 42,
		"",
	},
}

func TestUint32(t *testing.T) {
	for x, tt := range uint32Tests {
		reset()
		i := Uint32('i', 17)
		opt := Uint32Long("uint32", 0, 42)
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if got, want := *i, tt.i; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := *opt, tt.uint32; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
	}
}
