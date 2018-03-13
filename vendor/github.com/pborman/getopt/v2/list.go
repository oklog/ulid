// Copyright 2017 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

// List creates an option that returns a slice of strings.  The parameters
// passed are converted from a comma separated value list into a slice.
// Subsequent occurrences append to the list.
func List(name rune, helpvalue ...string) *[]string {
	p := []string{}
	CommandLine.Flag(&p, name, helpvalue...)
	return &p
}

func (s *Set) List(name rune, helpvalue ...string) *[]string {
	p := []string{}
	s.Flag(&p, name, helpvalue...)
	return &p
}

func ListLong(name string, short rune, helpvalue ...string) *[]string {
	p := []string{}
	CommandLine.FlagLong(&p, name, short, helpvalue...)
	return &p
}

func (s *Set) ListLong(name string, short rune, helpvalue ...string) *[]string {
	p := []string{}
	s.FlagLong(&p, name, short, helpvalue...)
	return &p
}
