// Copyright 2013 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"fmt"
	"strings"
	"testing"
)

var enumTests = []struct {
	where  string
	in     []string
	values []string
	out    string
	err    string
}{
	{
		loc(),
		nil,
		[]string{},
		"",
		"",
	},
	{
		loc(),
		[]string{"test", "-e", "val1"},
		[]string{"val1", "val2"},
		"val1",
		"",
	},
	{
		loc(),
		[]string{"test", "-e", "val1", "-e", "val2"},
		[]string{"val1", "val2"},
		"val2",
		"",
	},
	{
		loc(),
		[]string{"test", "-e", "val3"},
		[]string{"val1", "val2"},
		"",
		"test: invalid value: val3\n",
	},
}

func TestEnum(t *testing.T) {
	for x, tt := range enumTests {
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		reset()
		e := Enum('e', tt.values)
		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if *e != tt.out {
			t.Errorf("%s: got %v, want %v", tt.where, *e, tt.out)
		}
	}
}
