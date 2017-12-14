// Copyright 2017 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"fmt"
	"runtime"
	"strings"
)

// Value is the interface to the dynamic value stored in a flag.  Flags of type
// Value are declared using the Flag and FlagLong functions.
type Value interface {
	// Set converts value into the appropriate type and assigns it to the
	// receiver value.  Option details are provided via opt (such as the
	// flags name).
	//
	// Set is used to reset the value of an option to its default value
	// (which is stored in string form internally).
	Set(value string, opt Option) error

	// String returns the value of the flag as a string.
	String() string
}

var thisPackage string

// init initializes thisPackage to our full package with the trailing .
// included.
func init() {
        pc, _, _, ok := runtime.Caller(0)
        if !ok {
                return
        }
        f := runtime.FuncForPC(pc)
        if f == nil {
                return
        }
        thisPackage = f.Name()
        x := strings.LastIndex(thisPackage, "/")
        if x < 0 {
                return
        }
        y := strings.Index(thisPackage[x:], ".")
        if y < 0 {
                return
        }
        // thisPackage includes the trailing . after the package name.
        thisPackage = thisPackage[:x+y+1]
}

// calledFrom returns a string containing the file and linenumber of the first
// stack frame above us that is not part of this package and is not a test.
// This is used to determine where a flag was initialized.
func calledFrom() string {
	for i := 2; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			return ""
		}
		if !strings.HasSuffix(file, "_test.go") {
			f := runtime.FuncForPC(pc)
			if f != nil && strings.HasPrefix(f.Name(), thisPackage) {
				continue
			}
		}
		return fmt.Sprintf("%s:%d", file, line)
	}
}

func (s *Set) addFlag(p Value, name string, short rune, helpvalue ...string) Option {
	opt := &option{
		short:  short,
		long:   name,
		value:  p,
		defval: p.String(),
	}

	switch len(helpvalue) {
	case 2:
		opt.name = helpvalue[1]
		fallthrough
	case 1:
		opt.help = helpvalue[0]
	case 0:
	default:
		panic("Too many strings for String helpvalue")
	}
	if where := calledFrom(); where != "" {
		opt.where = where
	}
	if opt.short == 0 && opt.long == "" {
		fmt.Fprintf(stderr, opt.where+": no short or long option given")
		exit(1)
	}
	s.AddOption(opt)
	return opt
}
