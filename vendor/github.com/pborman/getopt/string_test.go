// Copyright 2013 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import "testing"

var stringTests = []struct {
	where  string
	in     []string
	sout   string
	optout string
	err    string
}{
	{
		loc(),
		[]string{},
		"one",
		"two",
		"",
	},
	{
		loc(),
		[]string{"test", "-s"},
		"one",
		"two",
		"test: missing parameter for -s\n",
	},
	{
		loc(),
		[]string{"test", "--opt"},
		"one",
		"two",
		"test: missing parameter for --opt\n",
	},
	{
		loc(),
		[]string{"test", "-svalue", "--opt=option"},
		"value",
		"option",
		"",
	},
	{
		loc(),
		[]string{"test", "-s", "value", "--opt", "option"},
		"value",
		"option",
		"",
	},
	{
		loc(),
		[]string{"test", "-swrong", "--opt=wrong", "-s", "value", "--opt", "option"},
		"value",
		"option",
		"",
	},
}

func TestString(t *testing.T) {
	for _, tt := range stringTests {
		reset()
		s := String('s', "one")
		opt := StringLong("opt", 0, "two")

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if *s != tt.sout {
			t.Errorf("%s: got s = %q, want %q", tt.where, *s, tt.sout)
		}
		if *opt != tt.optout {
			t.Errorf("%s: got opt = %q, want %q", tt.where, *opt, tt.optout)
		}
	}
}
