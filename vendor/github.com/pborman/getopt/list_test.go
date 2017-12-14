// Copyright 2013 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import "testing"

var listTests = []struct {
	where   string
	in      []string
	l, list []string
	err     string
}{
	{
		loc(),
		[]string{},
		nil, nil,
		"",
	},
	{
		loc(),
		[]string{"test", "-l", "one"},
		[]string{"one"}, nil,
		"",
	},
	{
		loc(),
		[]string{"test", "-lone", "-ltwo"},
		[]string{"one", "two"}, nil,
		"",
	},
	{
		loc(),
		[]string{"test", "--list", "one"},
		nil, []string{"one"},
		"",
	},
	{
		loc(),
		[]string{"test", "--list=one", "--list=two"},
		nil, []string{"one", "two"},
		"",
	},
	{
		loc(),
		[]string{"test", "--list=one,two"},
		nil, []string{"one", "two"},
		"",
	},
}

func TestList(t *testing.T) {
	for _, tt := range listTests {
		reset()
		l := List('l')
		list := ListLong("list", 0)

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if badSlice(*l, tt.l) {
			t.Errorf("%s: got s = %q, want %q", tt.where, *l, tt.l)
		}
		if badSlice(*list, tt.list) {
			t.Errorf("%s: got s = %q, want %q", tt.where, *list, tt.list)
		}
	}
}

func TestDefaultList(t *testing.T) {
	reset()
	list := []string{"d1", "d2", "d3"}
	ListVar(&list, 'l')
	parse([]string{"test"})

	want := []string{"d1", "d2", "d3"}
	if badSlice(list, want) {
		t.Errorf("got s = %q, want %q", list, want)
	}

	parse([]string{"test", "-l", "one"})
	want = []string{"one"}
	if badSlice(list, want) {
		t.Errorf("got s = %q, want %q", list, want)
	}

	parse([]string{"test", "-l", "two"})
	want = []string{"one", "two"}
	if badSlice(list, want) {
		t.Errorf("got s = %q, want %q", list, want)
	}
	Lookup('l').Reset()
	want = []string{"d1", "d2", "d3"}
	if badSlice(list, want) {
		t.Errorf("got s = %q, want %q", list, want)
	}
}
