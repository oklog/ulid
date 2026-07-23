package main

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"fmt"
	mathrand "math/rand/v2"
	"os"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
	getopt "github.com/pborman/getopt/v2"
)

const (
	defaultms = "Mon Jan 02 15:04:05.999 MST 2006"
	rfc3339ms = "2006-01-02T15:04:05.000Z07:00"
)

func main() {
	// Completely obnoxious.
	getopt.HelpColumn = 50
	getopt.DisplayWidth = 140

	fs := getopt.New()
	var (
		format = fs.StringLong("format", 'f', "default", "when parsing, show times in this format: default, rfc3339, unix, ms", "<format>")
		local  = fs.BoolLong("local", 'l', "when parsing, show local time instead of UTC")
		quick  = fs.BoolLong("quick", 'q', "when generating, use non-crypto-grade entropy")
		zero   = fs.BoolLong("zero", 'z', "when generating, fix entropy to all-zeroes")
		help   = fs.BoolLong("help", 'h', "print this help text")
	)
	if err := fs.Getopt(os.Args, nil); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if *help {
		fs.PrintUsage(os.Stderr)
		os.Exit(0)
	}

	var formatFunc func(time.Time) string
	switch strings.ToLower(*format) {
	case "default":
		formatFunc = func(t time.Time) string { return t.Format(defaultms) }
	case "rfc3339":
		formatFunc = func(t time.Time) string { return t.Format(rfc3339ms) }
	case "unix":
		formatFunc = func(t time.Time) string { return fmt.Sprint(t.Unix()) }
	case "ms":
		formatFunc = func(t time.Time) string { return fmt.Sprint(t.UnixNano() / 1e6) }
	default:
		fmt.Fprintf(os.Stderr, "invalid --format %s\n", *format)
		os.Exit(1)
	}

	switch args := fs.Args(); len(args) {
	case 0:
		generate(*quick, *zero)
	default:
		parse(args[0], *local, formatFunc)
	}
}

func generate(quick, zero bool) {
	entropy := cryptorand.Reader
	if quick {
		seed := [32]byte{}
		binary.LittleEndian.PutUint64(seed[:8], uint64(time.Now().UnixNano()))
		entropy = mathrand.NewChaCha8(seed)
	}
	if zero {
		entropy = zeroReader{}
	}

	id, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s\n", id)
}

func parse(s string, local bool, f func(time.Time) string) {
	id, err := ulid.Parse(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	t := ulid.Time(id.Time())
	if !local {
		t = t.UTC()
	}
	fmt.Fprintf(os.Stderr, "%s\n", f(t))
}

type zeroReader struct{}

func (zeroReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 0
	}
	return len(p), nil
}
