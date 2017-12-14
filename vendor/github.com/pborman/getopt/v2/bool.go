// Copyright 2017 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

// Bool creates a flag option that is a bool.  Bools normally do not take a
// value however one can be assigned by using the long form of the option:
//
//  --option=true
//  --o=false
//
// The value is case insensitive and one of true, false, t, f, on, off, t and 0.
func Bool(name rune, helpvalue ...string) *bool {
	var b bool
	CommandLine.Flag(&b, name, helpvalue...)
	return &b
}

func BoolLong(name string, short rune, helpvalue ...string) *bool {
	var p bool
	CommandLine.FlagLong(&p, name, short, helpvalue...)
	return &p
}

func (s *Set) Bool(name rune, helpvalue ...string) *bool {
	var b bool
	s.Flag(&b, name, helpvalue...)
	return &b
}

func (s *Set) BoolLong(name string, short rune, helpvalue ...string) *bool {
	var p bool
	s.FlagLong(&p, name, short, helpvalue...)
	return &p
}
