// Copyright 2017 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import "time"

// Duration creates an option that parses its value as a time.Duration.
func Duration(name rune, value time.Duration, helpvalue ...string) *time.Duration {
	CommandLine.FlagLong(&value, "", name, helpvalue...)
	return &value
}

func (s *Set) Duration(name rune, value time.Duration, helpvalue ...string) *time.Duration {
	s.FlagLong(&value, "", name, helpvalue...)
	return &value
}

func DurationLong(name string, short rune, value time.Duration, helpvalue ...string) *time.Duration {
	CommandLine.FlagLong(&value, name, short, helpvalue...)
	return &value
}

func (s *Set) DurationLong(name string, short rune, value time.Duration, helpvalue ...string) *time.Duration {
	s.FlagLong(&value, name, short, helpvalue...)
	return &value
}
