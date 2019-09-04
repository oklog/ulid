// +build gofuzz

package ulid

import (
	"bytes"
	"encoding/binary"
)

// FuzzNew tests ULID construction with fuzzed timestamp and entropy.
func FuzzNew(fuzz []byte) int {
	var ms uint64
	binary.Read(bytes.NewReader(fuzz), binary.LittleEndian, &ms)
	var entropy []byte
	if len(fuzz) > 8 {
		entropy = fuzz[7:]
	}
	id, err := New(ms, bytes.NewReader(entropy))
	if err != nil {
		return 0
	}
	id.Bytes()
	id.Entropy()
	id.Time()
	_, err = id.MarshalText()
	if err != nil {
		return 0
	}
	_, err = id.MarshalBinary()
	if err != nil {
		return 0
	}
	return 1
}

// FuzzParse tests ULID parsing.
func FuzzParse(fuzz []byte) int {
	id, err := Parse(string(fuzz))
	if err != nil {
		return 0
	}
	id.Bytes()
	id.Entropy()
	id.Time()
	_, err = id.MarshalText()
	if err != nil {
		return 0
	}
	_, err = id.MarshalBinary()
	if err != nil {
		return 0
	}
	return 1
}

// FuzzParse tests strict ULID parsing.
func FuzzParseStrict(fuzz []byte) int {
	id, err := ParseStrict(string(fuzz))
	if err != nil {
		return 0
	}
	id.Bytes()
	id.Entropy()
	id.Time()
	_, err = id.MarshalText()
	if err != nil {
		return 0
	}
	_, err = id.MarshalBinary()
	if err != nil {
		return 0
	}
	return 1
}
