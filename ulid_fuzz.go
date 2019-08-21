// +build gofuzz

package ulid

import (
	"bytes"
	"encoding/binary"
)

func FuzzNew(input []byte) int {
	var ms uint64
	binary.Read(bytes.NewReader(input), binary.LittleEndian, &ms)
	var entropy []byte
	if len(input) > 8 {
		entropy = input[7:]
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

func FuzzParse(input []byte) int {
	id, err := Parse(string(input))
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

func FuzzParseStrict(input []byte) int {
	id, err := ParseStrict(string(input))
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
