package ulid

import (
	"errors"
	"io"
)

// Generator generates ULIDs.
type Generator func(uint64) (ULID, error)

// Monotonic returns a Generator of ULID that guarantees the ULIDs will be strictly
// monotonically increasing.
func Monotonic(entropy io.Reader) Generator {
	var (
		lastMS   uint64
		lastULID ULID
	)
	return func(ms uint64) (ULID, error) {
		var err error
		if ms > lastMS {
			lastMS = ms
			lastULID, err = New(ms, entropy)
			return lastULID, err
		}
		incrEntropy := incrementBytes(lastULID.Entropy())
		var dup ULID
		dup.SetTime(ms)
		if err := dup.SetEntropy(incrEntropy); err != nil {
			return dup, err
		}
		lastULID = dup
		lastMS = ms
		return dup, nil
	}
}

var errOverflow = errors.New("overflowed entropy while incrementing it")

func incrementBytes(in []byte) []byte {
	const (
		minByte byte = 0
		maxByte byte = 255
	)
	out := make([]byte, len(in))
	copy(out, in)

	leastSigByteIdx := len(out) - 1
	mostSigByteIdex := 0

	for i := leastSigByteIdx; i >= mostSigByteIdex; i-- {
		if out[i] == maxByte {
			out[i] = minByte
			continue
		}
		out[i]++
		return out
	}
	panic(errOverflow)
}
