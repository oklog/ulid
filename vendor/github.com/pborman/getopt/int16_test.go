// Copyright 2013 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"fmt"
	"strings"
	"testing"
)

var int16Tests = []struct {
	where string
	in    []string
	i     int16
	int16 int16
	err   string
}{
	{
		loc(),
		[]string{},
		17, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i", "1", "--int16", "2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "--int16=2"},
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

func TestInt16(t *testing.T) {
	for x, tt := range int16Tests {
		reset()
		i := Int16('i', 17)
		opt := Int16Long("int16", 0, 42)
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
		if got, want := *opt, tt.int16; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
	}
}
