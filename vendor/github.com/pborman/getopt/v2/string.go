// Copyright 2017 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

// String returns a value option that stores is value as a string.  The
// initial value of the string is passed in value.
func String(name rune, value string, helpvalue ...string) *string {
	CommandLine.Flag(&value, name, helpvalue...)
	return &value
}

func (s *Set) String(name rune, value string, helpvalue ...string) *string {
	s.Flag(&value, name, helpvalue...)
	return &value
}

func StringLong(name string, short rune, value string, helpvalue ...string) *string {
	CommandLine.FlagLong(&value, name, short, helpvalue...)
	return &value
}

func (s *Set) StringLong(name string, short rune, value string, helpvalue ...string) *string {
	s.FlagLong(&value, name, short, helpvalue...)
	return &value
}
