// Copyright 2017 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"fmt"
	"strings"
	"testing"
)

var intTests = []struct {
	where string
	in    []string
	i     int
	int   int
	err   string
}{
	{
		loc(),
		[]string{},
		17, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i", "1", "--int", "2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "--int=2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "-i2"},
		2, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i=1"},
		17, 42,
		"test: not a valid number: =1\n",
	},
	{
		loc(),
		[]string{"test", "-i0x20"},
		0x20, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i010"},
		8, 42,
		"",
	},
}

func TestInt(t *testing.T) {
	for x, tt := range intTests {
		reset()
		i := Int('i', 17)
		opt := IntLong("int", 0, 42)
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if got, want := *i, tt.i; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := *opt, tt.int; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
	}
}

var int16Tests = []struct {
	where string
	in    []string
	i     int16
	int16 int16
	err   string
}{
	{
		loc(),
		[]string{},
		17, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i", "1", "--int16", "2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "--int16=2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "-i2"},
		2, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i=1"},
		17, 42,
		"test: not a valid number: =1\n",
	},
	{
		loc(),
		[]string{"test", "-i0x20"},
		0x20, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i010"},
		8, 42,
		"",
	},
}

func TestInt16(t *testing.T) {
	for x, tt := range int16Tests {
		reset()
		i := Int16('i', 17)
		opt := Int16Long("int16", 0, 42)
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if got, want := *i, tt.i; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := *opt, tt.int16; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
	}
}

var int32Tests = []struct {
	where string
	in    []string
	i     int32
	int32 int32
	err   string
}{
	{
		loc(),
		[]string{},
		17, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i", "1", "--int32", "2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "--int32=2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "-i2"},
		2, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i=1"},
		17, 42,
		"test: not a valid number: =1\n",
	},
	{
		loc(),
		[]string{"test", "-i0x20"},
		0x20, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i010"},
		8, 42,
		"",
	},
}

func TestInt32(t *testing.T) {
	for x, tt := range int32Tests {
		reset()
		i := Int32('i', 17)
		opt := Int32Long("int32", 0, 42)
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if got, want := *i, tt.i; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := *opt, tt.int32; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
	}
}

var int64Tests = []struct {
	where string
	in    []string
	i     int64
	int64 int64
	err   string
}{
	{
		loc(),
		[]string{},
		17, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i", "1", "--int64", "2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "--int64=2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "-i2"},
		2, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i=1"},
		17, 42,
		"test: not a valid number: =1\n",
	},
	{
		loc(),
		[]string{"test", "-i0x20"},
		0x20, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i010"},
		8, 42,
		"",
	},
}

func TestInt64(t *testing.T) {
	for x, tt := range int64Tests {
		reset()
		i := Int64('i', 17)
		opt := Int64Long("int64", 0, 42)
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if got, want := *i, tt.i; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := *opt, tt.int64; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
	}
}

var uintTests = []struct {
	where string
	in    []string
	i     uint
	uint  uint
	err   string
}{
	{
		loc(),
		[]string{},
		17, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i", "1", "--uint", "2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "--uint=2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "-i2"},
		2, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i=1"},
		17, 42,
		"test: not a valid number: =1\n",
	},
	{
		loc(),
		[]string{"test", "-i0x20"},
		0x20, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i010"},
		8, 42,
		"",
	},
}

func TestUint(t *testing.T) {
	for x, tt := range uintTests {
		reset()
		i := Uint('i', 17)
		opt := UintLong("uint", 0, 42)
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if got, want := *i, tt.i; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := *opt, tt.uint; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
	}
}

var uint16Tests = []struct {
	where  string
	in     []string
	i      uint16
	uint16 uint16
	err    string
}{
	{
		loc(),
		[]string{},
		17, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i", "1", "--uint16", "2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "--uint16=2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "-i2"},
		2, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i=1"},
		17, 42,
		"test: not a valid number: =1\n",
	},
	{
		loc(),
		[]string{"test", "-i0x20"},
		0x20, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i010"},
		8, 42,
		"",
	},
}

func TestUint16(t *testing.T) {
	for x, tt := range uint16Tests {
		reset()
		i := Uint16('i', 17)
		opt := Uint16Long("uint16", 0, 42)
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if got, want := *i, tt.i; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := *opt, tt.uint16; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
	}
}

var uint32Tests = []struct {
	where  string
	in     []string
	i      uint32
	uint32 uint32
	err    string
}{
	{
		loc(),
		[]string{},
		17, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i", "1", "--uint32", "2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "--uint32=2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "-i2"},
		2, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i=1"},
		17, 42,
		"test: not a valid number: =1\n",
	},
	{
		loc(),
		[]string{"test", "-i0x20"},
		0x20, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i010"},
		8, 42,
		"",
	},
}

func TestUint32(t *testing.T) {
	for x, tt := range uint32Tests {
		reset()
		i := Uint32('i', 17)
		opt := Uint32Long("uint32", 0, 42)
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if got, want := *i, tt.i; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := *opt, tt.uint32; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
	}
}

var uint64Tests = []struct {
	where  string
	in     []string
	i      uint64
	uint64 uint64
	err    string
}{
	{
		loc(),
		[]string{},
		17, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i", "1", "--uint64", "2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "--uint64=2"},
		1, 2,
		"",
	},
	{
		loc(),
		[]string{"test", "-i1", "-i2"},
		2, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i=1"},
		17, 42,
		"test: not a valid number: =1\n",
	},
	{
		loc(),
		[]string{"test", "-i0x20"},
		0x20, 42,
		"",
	},
	{
		loc(),
		[]string{"test", "-i010"},
		8, 42,
		"",
	},
}

func TestUint64(t *testing.T) {
	for x, tt := range uint64Tests {
		reset()
		i := Uint64('i', 17)
		opt := Uint64Long("uint64", 0, 42)
		if strings.Index(tt.where, ":-") > 0 {
			tt.where = fmt.Sprintf("#%d", x)
		}

		parse(tt.in)
		if s := checkError(tt.err); s != "" {
			t.Errorf("%s: %s", tt.where, s)
		}
		if got, want := *i, tt.i; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
		if got, want := *opt, tt.uint64; got != want {
			t.Errorf("%s: got %v, want %v", tt.where, got, want)
		}
	}
}
