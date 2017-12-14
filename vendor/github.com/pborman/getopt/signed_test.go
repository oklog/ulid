// Copyright 2013 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"fmt"
	"strings"
	"testing"
)

var signedNumberTests = []struct {
	where string
	in    []string
	l     SignedLimit
	out   int64
	err   string
}{
	{
		where: loc(),
	},
	{
		loc(),
		[]string{"test", "-n", "1010"},
		SignedLimit{Base: 2, Bits: 5},
		10,
		"",
	},
	{
		loc(),
		[]string{"test", "-n", "1010"},
		SignedLimit{Base: 2, Bits: 4},
		0,
		"test: value out of range: 1010\n",
	},
	{
		loc(),
		[]string{"test", "-n", "-1000"},
		SignedLimit{Base: 2, Bits: 4},
		-8,
		"",
	},
	{
		loc(),
		[]string{"test", "-n", "3"},
		SignedLimit{Min: 4, Max: 6},
		0,
		"test: value out of range (<4): 3\n",
	},
	{
		loc(),
		[]string{"test", "-n", "4"},
		SignedLimit{Min: 4, Max: 6},
		4,
		"",
	},
	{
		loc(),
		[]string{"test", "-n", "5"},
		SignedLimit{Min: 4, Max: 6},
		5,
		"",
	},
	{
		loc(),
		[]string{"test", "-n", "6"},
		SignedLimit{Min: 4, Max: 6},
		6,
		"",
	},
	{
		loc(),
		[]string{"test", "-n", "7"},
		SignedLimit{Min: 4, Max: 6},
		0,
		"test: value out of range (>6): 7\n",
	},
}

func TestSigneds(t *testing.T) {
	for x, tt := range signedNumberTests {
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		reset()
		n := Signed('n', 0, &tt.l)
		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if *n != tt.out {
			t.Errorf("%s: got %v, want %v", tt.where, *n, tt.out)
		}
	}
}
