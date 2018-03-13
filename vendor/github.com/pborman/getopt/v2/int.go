// Copyright 2017 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

// Int creates an option that parses its value as an integer.
func Int(name rune, value int, helpvalue ...string) *int {
	return CommandLine.Int(name, value, helpvalue...)
}

func (s *Set) Int(name rune, value int, helpvalue ...string) *int {
	s.Flag(&value, name, helpvalue...)
	return &value
}

func IntLong(name string, short rune, value int, helpvalue ...string) *int {
	return CommandLine.IntLong(name, short, value, helpvalue...)
}

func (s *Set) IntLong(name string, short rune, value int, helpvalue ...string) *int {
	s.FlagLong(&value, name, short, helpvalue...)
	return &value
}

// Int16 creates an option that parses its value as a 16 bit integer.
func Int16(name rune, value int16, helpvalue ...string) *int16 {
	return CommandLine.Int16(name, value, helpvalue...)
}

func (s *Set) Int16(name rune, value int16, helpvalue ...string) *int16 {
	s.Flag(&value, name, helpvalue...)
	return &value
}

func Int16Long(name string, short rune, value int16, helpvalue ...string) *int16 {
	return CommandLine.Int16Long(name, short, value, helpvalue...)
}

func (s *Set) Int16Long(name string, short rune, value int16, helpvalue ...string) *int16 {
	s.FlagLong(&value, name, short, helpvalue...)
	return &value
}

// Int32 creates an option that parses its value as a 32 bit integer.
func Int32(name rune, value int32, helpvalue ...string) *int32 {
	return CommandLine.Int32(name, value, helpvalue...)
}

func (s *Set) Int32(name rune, value int32, helpvalue ...string) *int32 {
	s.Flag(&value, name, helpvalue...)
	return &value
}

func Int32Long(name string, short rune, value int32, helpvalue ...string) *int32 {
	return CommandLine.Int32Long(name, short, value, helpvalue...)
}

func (s *Set) Int32Long(name string, short rune, value int32, helpvalue ...string) *int32 {
	s.FlagLong(&value, name, short, helpvalue...)
	return &value
}

// Int64 creates an option that parses its value as a 64 bit integer.
func Int64(name rune, value int64, helpvalue ...string) *int64 {
	return CommandLine.Int64(name, value, helpvalue...)
}

func (s *Set) Int64(name rune, value int64, helpvalue ...string) *int64 {
	s.Flag(&value, name, helpvalue...)
	return &value
}

func Int64Long(name string, short rune, value int64, helpvalue ...string) *int64 {
	return CommandLine.Int64Long(name, short, value, helpvalue...)
}

func (s *Set) Int64Long(name string, short rune, value int64, helpvalue ...string) *int64 {
	s.FlagLong(&value, name, short, helpvalue...)
	return &value
}

// Uint creates an option that parses its value as an unsigned integer.
func Uint(name rune, value uint, helpvalue ...string) *uint {
	return CommandLine.Uint(name, value, helpvalue...)
}

func (s *Set) Uint(name rune, value uint, helpvalue ...string) *uint {
	s.Flag(&value, name, helpvalue...)
	return &value
}

func UintLong(name string, short rune, value uint, helpvalue ...string) *uint {
	return CommandLine.UintLong(name, short, value, helpvalue...)
}

func (s *Set) UintLong(name string, short rune, value uint, helpvalue ...string) *uint {
	s.FlagLong(&value, name, short, helpvalue...)
	return &value
}

// Uint16 creates an option that parses its value as a 16 bit unsigned integer.
func Uint16(name rune, value uint16, helpvalue ...string) *uint16 {
	return CommandLine.Uint16(name, value, helpvalue...)
}

func (s *Set) Uint16(name rune, value uint16, helpvalue ...string) *uint16 {
	s.Flag(&value, name, helpvalue...)
	return &value
}

func Uint16Long(name string, short rune, value uint16, helpvalue ...string) *uint16 {
	return CommandLine.Uint16Long(name, short, value, helpvalue...)
}

func (s *Set) Uint16Long(name string, short rune, value uint16, helpvalue ...string) *uint16 {
	s.FlagLong(&value, name, short, helpvalue...)
	return &value
}

// Uint32 creates an option that parses its value as a 32 bit unsigned integer.
func Uint32(name rune, value uint32, helpvalue ...string) *uint32 {
	return CommandLine.Uint32(name, value, helpvalue...)
}

func (s *Set) Uint32(name rune, value uint32, helpvalue ...string) *uint32 {
	s.Flag(&value, name, helpvalue...)
	return &value
}

func Uint32Long(name string, short rune, value uint32, helpvalue ...string) *uint32 {
	return CommandLine.Uint32Long(name, short, value, helpvalue...)
}

func (s *Set) Uint32Long(name string, short rune, value uint32, helpvalue ...string) *uint32 {
	s.FlagLong(&value, name, short, helpvalue...)
	return &value
}

// Uint64 creates an option that parses its value as a 64 bit unsigned integer.
func Uint64(name rune, value uint64, helpvalue ...string) *uint64 {
	return CommandLine.Uint64(name, value, helpvalue...)
}

func (s *Set) Uint64(name rune, value uint64, helpvalue ...string) *uint64 {
	s.Flag(&value, name, helpvalue...)
	return &value
}

func Uint64Long(name string, short rune, value uint64, helpvalue ...string) *uint64 {
	return CommandLine.Uint64Long(name, short, value, helpvalue...)
}

func (s *Set) Uint64Long(name string, short rune, value uint64, helpvalue ...string) *uint64 {
	s.FlagLong(&value, name, short, helpvalue...)
	return &value
}
