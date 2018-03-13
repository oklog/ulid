// Copyright 2013 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"fmt"
	"strings"
	"testing"
)

var boolTests = []struct {
	where string
	in    []string
	f     bool
	fc    int
	opt   bool
	optc  int
	err   string
}{
	{
		loc(),
		[]string{},
		false, 0,
		false, 0,
		"",
	},
	{
		loc(),
		[]string{"test", "-f", "--opt"},
		true, 1,
		true, 1,
		"",
	},
	{
		loc(),
		[]string{"test", "--f", "--opt"},
		true, 1,
		true, 1,
		"",
	},
	{
		loc(),
		[]string{"test", "-ff", "-f", "--opt", "--opt"},
		true, 3,
		true, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "--opt", "--opt=false"},
		false, 0,
		false, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-f", "false"},
		true, 1,
		false, 0,
		"",
	},
	{
		loc(),
		[]string{"test", "-f=false"},
		true, 1,
		false, 0,
		"test: unknown option: -=\n",
	},
	{
		loc(),
		[]string{"test", "-f", "false"},
		true, 1,
		false, 0,
		"",
	},
}

func TestBool(t *testing.T) {
	for x, tt := range boolTests {
		reset()
		f := Bool('f')
		opt := BoolLong("opt", 0)
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if got, want := *f, tt.f; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := *opt, tt.opt; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := GetCount('f'), tt.fc; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := GetCount("opt"), tt.optc; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
	}
}
