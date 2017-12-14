// Copyright 2013 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"testing"
)

var breakupTests = []struct {
	in  string
	max int
	out []string
}{
	{"", 8, []string{}},
	{"a fox", 8, []string{"a fox"}},
	{"a foxhound is sly", 2, []string{"a", "foxhound", "is", "sly"}},
	{"a foxhound is sly", 5, []string{"a", "foxhound", "is", "sly"}},
	{"a foxhound is sly", 6, []string{"a", "foxhound", "is sly"}},
	{"a foxhound is sly", 7, []string{"a", "foxhound", "is sly"}},
	{"a foxhound is sly", 8, []string{"a", "foxhound", "is sly"}},
	{"a foxhound is sly", 9, []string{"a", "foxhound", "is sly"}},
	{"a foxhound is sly", 10, []string{"a foxhound", "is sly"}},
}

func TestBreakup(t *testing.T) {
	for x, tt := range breakupTests {
		out := breakup(tt.in, tt.max)
		if badSlice(out, tt.out) {
			t.Errorf("#%d: got %v, want %v", x, out, tt.out)
		}
	}
}
