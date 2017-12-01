package main

import (
	cryptorand "crypto/rand"
	"flag"
	"fmt"
	mathrand "math/rand"
	"os"
	"strings"
	"time"

	"github.com/oklog/ulid"
)

const rfc3339ms = "2006-01-02T15:04:05.999Z"

func main() {
	var (
		format = flag.String("format", "rfc3339", "rfc3339, unix, ms (decoding)")
		local  = flag.Bool("local", false, "use local time instead of UTC (decoding)")
		fast   = flag.Bool("fast", false, "use non-crypto-grade entropy (encoding)")
		zero   = flag.Bool("zero", false, "fix entropy to zero (encoding)")
	)
	flag.Parse()

	var formatFunc func(time.Time) string
	switch strings.ToLower(*format) {
	case "rfc3339":
		formatFunc = func(t time.Time) string { return t.Format(rfc3339ms) }
	case "unix":
		formatFunc = func(t time.Time) string { return fmt.Sprint(t.Unix()) }
	case "ms":
		formatFunc = func(t time.Time) string { return fmt.Sprint(t.UnixNano() / 1e6) }
	default:
		fmt.Fprintf(os.Stderr, "invalid -format %s\n", *format)
		os.Exit(1)
	}

	switch len(flag.Args()) {
	case 0:
		generate(*local, *fast, *zero)
	default:
		parse(flag.Args()[0], formatFunc)
	}
}

func generate(local, fast, zero bool) {
	now := time.Now()
	if !local {
		now = now.UTC()
	}
	ts := ulid.Timestamp(now)

	entropy := cryptorand.Reader
	if fast {
		seed := time.Now().UnixNano()
		source := mathrand.NewSource(seed)
		entropy = mathrand.New(source)
	}
	if zero {
		entropy = zeroReader{}
	}

	id, err := ulid.New(ts, entropy)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s\n", id)
}

func parse(s string, f func(time.Time) string) {
	id, err := ulid.Parse(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	var (
		msec = id.Time()
		sec  = msec / 1e3
		rem  = msec % 1e3
		nsec = rem * 1e6
		t    = time.Unix(int64(sec), int64(nsec))
	)
	fmt.Fprintf(os.Stderr, "%s\n", f(t))
}

type zeroReader struct{}

func (zeroReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 0
	}
	return len(p), nil
}
