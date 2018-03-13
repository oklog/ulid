// Copyright 2013 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"bytes"
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"
)

var errorString string

func reset() {
	CommandLine.shortOptions = make(map[rune]*option)
	CommandLine.longOptions = make(map[string]*option)
	CommandLine.options = nil
	CommandLine.args = nil
	CommandLine.program = ""
	errorString = ""
}

func parse(args []string) {
	err := CommandLine.Getopt(args, nil)
	if err != nil {
		b := &bytes.Buffer{}

		fmt.Fprintln(b, CommandLine.program+": "+err.Error())
		CommandLine.PrintUsage(b)
		errorString = b.String()
	}
}

func badSlice(a, b []string) bool {
	if len(a) != len(b) {
		return true
	}
	for x, v := range a {
		if b[x] != v {
			return true
		}
	}
	return false
}

func loc() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s:%d", path.Base(file), line)
}

func (o *option) Equal(opt *option) bool {
	if o.value != nil && opt.value == nil {
		return false
	}
	if o.value == nil && opt.value != nil {
		return false
	}
	if o.value != nil && o.value.String() != opt.value.String() {
		return false
	}

	oc := *o
	optc := *opt
	oc.value = nil
	optc.value = nil
	return reflect.DeepEqual(&oc, &optc)
}

func newStringValue(s string) *stringValue { return (*stringValue)(&s) }

func checkError(err string) string {
	switch {
	case err == errorString:
		return ""
	case err == "":
		return fmt.Sprintf("unexpected error %q", errorString)
	case errorString == "":
		return fmt.Sprintf("did not get expected error %q", err)
	case !strings.HasPrefix(errorString, err):
		return fmt.Sprintf("got error %q, want %q", errorString, err)
	}
	return ""
}
