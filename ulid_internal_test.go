// Copyright 2016 The Oklog Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ulid

import (
	"bytes"
	"errors"
	"testing"
)

// These tests replace process-wide entropy and must not run in parallel.
func setDefaultMonotonicEntropy(t *testing.T, entropy MonotonicEntropy) {
	t.Helper()

	defaultEntropy.mu.Lock()
	original := *defaultEntropy.MonotonicEntropy
	*defaultEntropy.MonotonicEntropy = entropy
	defaultEntropy.mu.Unlock()

	t.Cleanup(func() {
		defaultEntropy.mu.Lock()
		*defaultEntropy.MonotonicEntropy = original
		defaultEntropy.mu.Unlock()
	})
}

func TestMakeSharesDefaultEntropy(t *testing.T) {
	seed := make([]byte, 10)
	setDefaultMonotonicEntropy(t, MonotonicEntropy{
		Reader: bytes.NewReader(bytes.Repeat([]byte{0}, 20)),
		inc:    1,
	})

	id := Make()
	if !bytes.Equal(id[6:], seed) {
		t.Fatalf("Make entropy: got %v, want %v", id[6:], seed)
	}

	next, err := New(id.Time(), DefaultEntropy())
	if err != nil {
		t.Fatal(err)
	}

	want := id
	for i := len(want) - 1; i >= 6; i-- {
		want[i]++
		if want[i] != 0 {
			break
		}
	}
	if next != want {
		t.Fatalf("New after Make: got %v, want %v", next, want)
	}
}

func TestMakePanicsOnEntropyReadError(t *testing.T) {
	want := errors.New("entropy read failed")
	setDefaultMonotonicEntropy(t, MonotonicEntropy{
		Reader: errorEntropyReader{err: want},
		inc:    1,
	})

	defer func() {
		if got := recover(); got != want {
			t.Fatalf("Make panic: got %v, want %v", got, want)
		}
	}()
	_ = Make()
}

type errorEntropyReader struct {
	err error
}

func (r errorEntropyReader) Read([]byte) (int, error) {
	return 0, r.err
}
