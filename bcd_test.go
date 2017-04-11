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
	{[]byte{0x01}, []byte{0x00, 0x01}},
	{[]byte{0x01, 0x23}, []byte{0x00, 0x01, 0x02, 0x03}},
	{[]byte{0x01, 0x23, 0x45}, []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05}},
	{[]byte{0x01, 0x23, 0x45, 0x67}, []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}},
	{[]byte{0x01, 0x23, 0x45, 0x67, 0x89}, []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09}},
}

func TestEncode(t *testing.T) {
	for i, test := range encDecTests {
		dst := make([]byte, EncodedLen(len(test.dec)))
		n, err := Encode(dst, test.dec)
		if err != nil {
			t.Errorf("#%d: bad return value: got: %d want: %d", i, n, len(dst))
		} else if bytes.Equal(dst, test.enc) != true {
			t.Errorf("#%d: got: %#v want: %#v", i, dst, test.enc)
		}
	}
}

func TestDecode(t *testing.T) {
	for i, test := range encDecTests {
		dst := make([]byte, DecodedLen(len(test.enc)))
		n := Decode(dst, test.enc)
		if n != len(dst) {
			t.Errorf("#%d: bad return value: got: %d want: %d", i, n, len(dst))
		}
		if bytes.Equal(dst, test.dec) != true {
			t.Errorf("#%d: got: %#v want: %#v", i, dst, test.dec)
		}
	}
}
