// Copyright 2013 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

var durationTests = []struct {
	where string
	in    []string
	d     time.Duration
	dur   time.Duration
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
		[]string{"test", "-d", "1s", "--duration", "2s"},
		time.Second, 2 * time.Second,
		"",
	},
	{
		loc(),
		[]string{"test", "-d1s", "-d2s"},
		2 * time.Second, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-d1"},
		17, 42,
		"test: time: missing unit in duration 1\n",
	},
	{
		loc(),
		[]string{"test", "--duration", "foo"},
		17, 42,
		"test: time: invalid duration foo\n",
	},
}

func TestDuration(t *testing.T) {
	for x, tt := range durationTests {
		reset()
		d := Duration('d', 17)
		opt := DurationLong("duration", 0, 42)
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if got, want := *d, tt.d; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := *opt, tt.dur; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
	}
}
