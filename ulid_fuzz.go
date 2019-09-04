// +build gofuzz

package ulid

import (
	"bytes"
	"encoding/binary"
)

// FuzzNew tests ULID construction.
func FuzzNew(fuzz []byte) int {
	ms, entropy := extractFuzzTimestamp(fuzz)
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

// FuzzNewMonotonic tests ULID construction with monotonic entropy.
func FuzzNewMonotonic(fuzz []byte) int {
	if len(fuzz) < (8 * 10) {
		return -1
	}
	timestamps := fuzz[0 : 8*10]
	var entropy []byte
	if len(fuzz) > (8 * 10) {
		entropy = fuzz[8*10:]
	} else {
		entropy = []byte{}
	}
	monotonic := Monotonic(bytes.NewReader(entropy), 0)
	var ms uint64
	for range [10]struct{}{} {
		ms, timestamps = extractFuzzTimestamp(timestamps)
		id, err := New(ms, monotonic)
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

func extractFuzzTimestamp(fuzz []byte) (ms uint64, rest []byte) {
	binary.Read(bytes.NewReader(fuzz), binary.LittleEndian, &ms)
	if len(fuzz) > 8 {
		rest = fuzz[7:]
	} else {
		rest = []byte{}
	}
	return
}
