package bcd

import (
	"bytes"
	"testing"
)

type encDecTest struct {
	enc []byte
	dec []byte
}

var encDecTests = []encDecTest{
	{[]byte{}, []byte{}},
	{[]byte{0x01}, []byte{0, 1}},
	{[]byte{0x01, 0x23}, []byte{0, 1, 2, 3}},
	{[]byte{0x01, 0x23, 0x45}, []byte{0, 1, 2, 3, 4, 5}},
	{[]byte{0x01, 0x23, 0x45, 0x67}, []byte{0, 1, 2, 3, 4, 5, 6, 7}},
	{[]byte{0x01, 0x23, 0x45, 0x67, 0x89}, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
}

func TestEncode(t *testing.T) {
	for i, test := range encDecTests {
		dst := make([]byte, EncodedLen(len(test.dec)))
		n := Encode(dst, test.dec)
		if n != len(dst) {
			t.Errorf("#%d: bad return value: got: %d want: %d", i, n, len(dst))
		}
		if bytes.Equal(dst, test.enc) != true {
			t.Errorf("#%d: got: %#v want: %#v", i, dst, test.enc)
		}
	}
}

func TestDecode(t *testing.T) {
	for i, test := range encDecTests {
		dst := make([]byte, DecodedLen(len(test.enc)))
		n, err := Decode(dst, test.enc)
		if err != nil {
			t.Errorf("#%d: bad return value: got: %d want: %d", i, n, len(dst))
		}
		if bytes.Equal(dst, test.dec) != true {
			t.Errorf("#%d: got: %#v want: %#v", i, dst, test.dec)
		}
	}
}
