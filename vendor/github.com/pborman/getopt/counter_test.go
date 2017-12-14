// Copyright 2013 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"fmt"
	"strings"
	"testing"
)

var counterTests = []struct {
	where string
	in    []string
	c     int
	cnt   int
	err   string
}{
	{
		loc(),
		[]string{},
		0,
		0,
		"",
	},
	{
		loc(),
		[]string{"test", "-c", "--cnt"},
		1,
		1,
		"",
	},
	{
		loc(),
		[]string{"test", "-cc", "-c", "--cnt", "--cnt"},
		3,
		2,
		"",
	},
	{
		loc(),
		[]string{"test", "--c=17", "--cnt=42"},
		17,
		42,
		"",
	},
	{
		loc(),
		[]string{"test", "--cnt=false"},
		0, 0,
		"test: not a valid number: false\n",
	},
}

func TestCounter(t *testing.T) {
	for x, tt := range counterTests {
		reset()
		c := Counter('c')
		cnt := CounterLong("cnt", 0)
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if got, want := *c, tt.c; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := *cnt, tt.cnt; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
	}

	reset()
	c := 5
	opt := CounterVar(&c, 'c')
	parse([]string{"test", "-c"})
	if c != 6 {
		t.Errorf("got %d, want 6", c)
	}
	if opt.Count() != 1 {
		t.Errorf("got %d, want 1", c)
	}
	Reset()
	if c != 5 {
		t.Errorf("got %d, want 5", c)
	}
}
