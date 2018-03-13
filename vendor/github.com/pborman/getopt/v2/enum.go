// Copyright 2017 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"errors"
	"fmt"
	"sync"
)

type enumValue string

var (
	enumValuesMu sync.Mutex
	enumValues   = make(map[*enumValue]map[string]struct{})
)

func (s *enumValue) Set(value string, opt Option) error {
	enumValuesMu.Lock()
	es, ok := enumValues[s]
	enumValuesMu.Unlock()
	if !ok || es == nil {
		return errors.New("this option has no values")
	}
	if _, ok := es[value]; !ok {
		return errors.New("invalid value: " + value)
	}
	*s = enumValue(value)
	return nil
}

func (s *enumValue) String() string {
	return string(*s)
}

// Enum creates an option that can only be set to one of the enumerated strings
// passed in values.  Passing nil or an empty slice results in an option that
// will always fail.  If not "", value is the default value of the enum.  If
// value is not listed in values then Enum will produce an error on standard
// error and then exit the program with a status of 1.
func Enum(name rune, values []string, value string, helpvalue ...string) *string {
	return CommandLine.Enum(name, values, value, helpvalue...)
}

func (s *Set) Enum(name rune, values []string, value string, helpvalue ...string) *string {
	var p enumValue
	p.define(values, value, &option{short: name})
	s.FlagLong(&p, "", name, helpvalue...)
	return (*string)(&p)
}

func EnumLong(name string, short rune, values []string, value string, helpvalue ...string) *string {
	return CommandLine.EnumLong(name, short, values, value, helpvalue...)
}

func (s *Set) EnumLong(name string, short rune, values []string, value string, helpvalue ...string) *string {
	var p enumValue
	p.define(values, value, &option{short: short, long: name})
	s.FlagLong(&p, name, short, helpvalue...)
	return (*string)(&p)
}

func (e *enumValue) define(values []string, def string, opt Option) {
	m := make(map[string]struct{})
	for _, v := range values {
		m[v] = struct{}{}
	}
	enumValuesMu.Lock()
	enumValues[e] = m
	enumValuesMu.Unlock()
	if def != "" {
		if err := e.Set(def, nil); err != nil {
			fmt.Fprintf(stderr, "setting default for %s: %v\n", opt.Name(), err)
			exit(1)
		}
	}
}
