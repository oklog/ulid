// Copyright 2017 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestGeneric(t *testing.T) {
	const (
		shortTest = iota
		longTest
		bothTest
	)
	for _, tt := range []struct {
		where string
		kind  int
		val   interface{}
		str   string
		def   interface{}
		in    []string
		err   string
	}{
		// Do all four tests for string, the rest can mostly just use
		// shortTest (the 0 value).
		{
			where: loc(),
			kind:  shortTest,
			val:   "42",
			str:   "42",
			in:    []string{"test", "-s", "42"},
		},
		{
			where: loc(),
			kind:  longTest,
			val:   "42",
			str:   "42",
			in:    []string{"test", "--long", "42"},
		},
		{
			where: loc(),
			kind:  bothTest,
			val:   "42",
			str:   "42",
			in:    []string{"test", "--both", "42"},
		},
		{
			where: loc(),
			kind:  bothTest,
			val:   "42",
			str:   "42",
			in:    []string{"test", "-b", "42"},
		},
		{
			where: loc(),
			val:   "42",
			def:   "42",
			str:   "42",
			in:    []string{"test"},
		},
		{
			where: loc(),
			val:   "42",
			def:   "43",
			str:   "42",
			in:    []string{"test", "-s", "42"},
		},

		{
			where: loc(),
			val:   true,
			str:   "true",
			in:    []string{"test", "-s"},
		},
		{
			where: loc(),
			val:   true,
			def:   true,
			str:   "true",
			in:    []string{"test"},
		},
		{
			where: loc(),
			kind:  longTest,
			val:   false,
			str:   "false",
			in:    []string{"test", "--long=false"},
		},
		{
			where: loc(),
			kind:  longTest,
			val:   false,
			def:   true,
			str:   "false",
			in:    []string{"test", "--long=false"},
		},

		{
			where: loc(),
			val:   int(42),
			str:   "42",
			in:    []string{"test", "-s", "42"},
		},
		{
			where: loc(),
			val:   int8(42),
			str:   "42",
			in:    []string{"test", "-s", "42"},
		},
		{
			where: loc(),
			val:   int16(42),
			str:   "42",
			in:    []string{"test", "-s", "42"},
		},
		{
			where: loc(),
			val:   int32(42),
			str:   "42",
			in:    []string{"test", "-s", "42"},
		},
		{
			where: loc(),
			val:   int64(42),
			str:   "42",
			in:    []string{"test", "-s", "42"},
		},

		{
			where: loc(),
			val:   uint(42),
			str:   "42",
			in:    []string{"test", "-s", "42"},
		},
		{
			where: loc(),
			val:   uint8(42),
			str:   "42",
			in:    []string{"test", "-s", "42"},
		},
		{
			where: loc(),
			val:   uint16(42),
			str:   "42",
			in:    []string{"test", "-s", "42"},
		},
		{
			where: loc(),
			val:   uint32(42),
			str:   "42",
			in:    []string{"test", "-s", "42"},
		},
		{
			where: loc(),
			val:   uint64(42),
			str:   "42",
			in:    []string{"test", "-s", "42"},
		},

		{
			where: loc(),
			val:   float32(4.2),
			str:   "4.2",
			in:    []string{"test", "-s", "4.2"},
		},
		{
			where: loc(),
			val:   float64(4.2),
			str:   "4.2",
			in:    []string{"test", "-s", "4.2"},
		},

		{
			where: loc(),
			val:   time.Duration(time.Second * 42),
			def:   time.Second * 2,
			str:   "42s",
			in:    []string{"test", "-s", "42s"},
		},
		{
			where: loc(),
			val:   time.Duration(time.Second * 42),
			def:   time.Second * 2,
			str:   "42s",
			in:    []string{"test", "-s42s"},
		},
		{
			where: loc(),
			val:   time.Duration(time.Second * 2),
			def:   time.Second * 2,
			in:    []string{"test", "-s42"},
			str:   "2s",
			err:   "test: time: missing unit in duration 42",
		},

		{
			where: loc(),
			val:   []string{"42", "."},
			str:   "42,.",
			def:   []string{"one", "two", "three"},
			in:    []string{"test", "-s42", "-s."},
		},
		{
			where: loc(),
			val:   []string{"42", "."},
			str:   "42,.",
			def:   []string{"one", "two", "three"},
			in:    []string{"test", "-s42,."},
		},
		{
			where: loc(),
			val:   []string{"one", "two", "three"},
			def:   []string{"one", "two", "three"},
			str:   "one,two,three",
			in:    []string{"test"},
		},
	} {
		reset()
		var opt Option
		val := reflect.New(reflect.TypeOf(tt.val)).Interface()
		if tt.def != nil {
			reflect.ValueOf(val).Elem().Set(reflect.ValueOf(tt.def))
		}
		switch tt.kind {
		case shortTest:
			opt = Flag(val, 's')
		case longTest:
			opt = FlagLong(val, "long", 0)
		case bothTest:
			opt = FlagLong(val, "both", 'b')
		}
		_ = opt
		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
			continue
		}
		got := reflect.ValueOf(val).Elem().Interface()
		want := reflect.ValueOf(tt.val).Interface()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if str := opt.String(); str != tt.str {
			t.Errorf("%s: got string %q, want %q", tt.where, str, tt.str)
		}
	}
}

func TestGenericDup(t *testing.T) {
	defer func() {
		stderr = os.Stderr
		exit = os.Exit
	}()

	reset()
	var v1, v2 string
	type myPanic struct{}
	var errbuf bytes.Buffer
	stderr = &errbuf
	_, file, line, _ := runtime.Caller(0)
	Flag(&v1, 's')
	line++ // line is now the line number of the first call to Flag.

	exit = func(i int) { panic(myPanic{}) }
	defer func() {
		p := recover()
		if _, ok := p.(myPanic); ok {
			err := errbuf.String()
			if !strings.Contains(err, "-s already declared") || !strings.Contains(err, fmt.Sprintf("%s:%d", file, line)) {
				t.Errorf("unexpected error: %q\nshould contain \"-s already declared\" and \"%s:%d\"", err, file, line)
			}
		} else if p == nil {
			t.Errorf("Second call to Flag did not fail")
		} else {
			t.Errorf("panic %v", p)
		}
	}()
	Flag(&v2, 's')
}

func TestGenericDupNested(t *testing.T) {
	defer func() {
		stderr = os.Stderr
		exit = os.Exit
	}()

	reset()
	type myPanic struct{}
	var errbuf bytes.Buffer
	stderr = &errbuf
	_, file, line, _ := runtime.Caller(0)
	String('s', "default")
	line++ // line is now the line number of the first call to Flag.

	exit = func(i int) { panic(myPanic{}) }
	defer func() {
		p := recover()
		if _, ok := p.(myPanic); ok {
			err := errbuf.String()
			if !strings.Contains(err, "-s already declared") || !strings.Contains(err, fmt.Sprintf("%s:%d", file, line)) {
				t.Errorf("unexpected error: %q\nshould contain \"-s already declared\" and \"%s:%d\"", err, file, line)
			}
		} else if p == nil {
			t.Errorf("Second call to Flag did not fail")
		} else {
			t.Errorf("panic %v", p)
		}
	}()
	String('s', "default")
}
