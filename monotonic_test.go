package ulid

import (
	"bytes"
	"encoding/hex"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestMonotonic(t *testing.T) {
	for i := int64(0); i < 1000; i++ {
		r := rand.New(rand.NewSource(int64(i)))
		now := Timestamp(time.Unix(r.Int63n(1000), r.Int63n(1000)))
		mkulid := Monotonic(r)

		var got []string
		for i := 0; i < 1+r.Intn(5); i++ {
			id, err := mkulid(now)
			if err != nil {
				panic(err)
			}
			got = append(got, id.String())
		}
		now++ // advance time
		for i := 0; i < 1+r.Intn(5); i++ {
			id, err := mkulid(now)
			if err != nil {
				panic(err)
			}
			got = append(got, id.String())
		}

		want := make([]string, len(got))
		copy(want, got)
		sort.Strings(want)

		if !reflect.DeepEqual(want, got) {
			for i, w := range want {
				t.Errorf("want[%d]=%v", i, w)
			}
			for i, g := range got {
				t.Errorf("got[%d]=%v", i, g)
			}
			t.Fatal(i)
		}
	}
}

func Test_incrementByte(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want []byte
	}{
		{"first digit", []byte{0x00, 0x00}, []byte{0x00, 0x01}},
		{"first digit", []byte{0x00, 0x01}, []byte{0x00, 0x02}},
		{"first digit", []byte{0x00, 0x02}, []byte{0x00, 0x03}},
		{"first digit", []byte{0x00, 0x03}, []byte{0x00, 0x04}},
		{"first digit", []byte{0x00, 0x04}, []byte{0x00, 0x05}},
		{"first digit", []byte{0x00, 0x05}, []byte{0x00, 0x06}},
		{"first digit", []byte{0x00, 0x06}, []byte{0x00, 0x07}},
		{"first digit", []byte{0x00, 0x07}, []byte{0x00, 0x08}},
		{"first digit", []byte{0x00, 0x08}, []byte{0x00, 0x09}},
		{"first digit", []byte{0x00, 0x09}, []byte{0x00, 0x0A}},
		{"first digit", []byte{0x00, 0x0A}, []byte{0x00, 0x0B}},
		{"first digit", []byte{0x00, 0x0B}, []byte{0x00, 0x0C}},
		{"first digit", []byte{0x00, 0x0C}, []byte{0x00, 0x0D}},
		{"first digit", []byte{0x00, 0x0D}, []byte{0x00, 0x0E}},
		{"first digit", []byte{0x00, 0x0E}, []byte{0x00, 0x0F}},
		{"first digit", []byte{0x00, 0x0F}, []byte{0x00, 0x10}},

		{"second digit", []byte{0x00, 0x00}, []byte{0x00, 0x01}},
		{"second digit", []byte{0x00, 0x10}, []byte{0x00, 0x11}},
		{"second digit", []byte{0x00, 0x20}, []byte{0x00, 0x21}},
		{"second digit", []byte{0x00, 0x30}, []byte{0x00, 0x31}},
		{"second digit", []byte{0x00, 0x40}, []byte{0x00, 0x41}},
		{"second digit", []byte{0x00, 0x50}, []byte{0x00, 0x51}},
		{"second digit", []byte{0x00, 0x60}, []byte{0x00, 0x61}},
		{"second digit", []byte{0x00, 0x70}, []byte{0x00, 0x71}},
		{"second digit", []byte{0x00, 0x80}, []byte{0x00, 0x81}},
		{"second digit", []byte{0x00, 0x90}, []byte{0x00, 0x91}},
		{"second digit", []byte{0x00, 0xA0}, []byte{0x00, 0xA1}},
		{"second digit", []byte{0x00, 0xB0}, []byte{0x00, 0xB1}},
		{"second digit", []byte{0x00, 0xC0}, []byte{0x00, 0xC1}},
		{"second digit", []byte{0x00, 0xD0}, []byte{0x00, 0xD1}},
		{"second digit", []byte{0x00, 0xE0}, []byte{0x00, 0xE1}},
		{"second digit", []byte{0x00, 0xF0}, []byte{0x00, 0xF1}},

		{"second digit with carry", []byte{0x00, 0x0F}, []byte{0x00, 0x10}},
		{"second digit with carry", []byte{0x00, 0x1F}, []byte{0x00, 0x20}},
		{"second digit with carry", []byte{0x00, 0x2F}, []byte{0x00, 0x30}},
		{"second digit with carry", []byte{0x00, 0x3F}, []byte{0x00, 0x40}},
		{"second digit with carry", []byte{0x00, 0x4F}, []byte{0x00, 0x50}},
		{"second digit with carry", []byte{0x00, 0x5F}, []byte{0x00, 0x60}},
		{"second digit with carry", []byte{0x00, 0x6F}, []byte{0x00, 0x70}},
		{"second digit with carry", []byte{0x00, 0x7F}, []byte{0x00, 0x80}},
		{"second digit with carry", []byte{0x00, 0x8F}, []byte{0x00, 0x90}},
		{"second digit with carry", []byte{0x00, 0x9F}, []byte{0x00, 0xA0}},
		{"second digit with carry", []byte{0x00, 0xAF}, []byte{0x00, 0xB0}},
		{"second digit with carry", []byte{0x00, 0xBF}, []byte{0x00, 0xC0}},
		{"second digit with carry", []byte{0x00, 0xCF}, []byte{0x00, 0xD0}},
		{"second digit with carry", []byte{0x00, 0xDF}, []byte{0x00, 0xE0}},
		{"second digit with carry", []byte{0x00, 0xEF}, []byte{0x00, 0xF0}},
		{"second digit with carry", []byte{0x00, 0xFF}, []byte{0x01, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name+"-"+hex.EncodeToString(tt.in), func(t *testing.T) {
			got := incrementBytes(tt.in)
			if !bytes.Equal(tt.want, got) {
				t.Errorf("want=%x", tt.want)
				t.Errorf(" got=%x", got)
			}
		})
	}
}
