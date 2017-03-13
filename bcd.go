// Package bcd implements binary coded decimal encoding and decoding.
package bcd

import "errors"

// ErrLength results from decoding an odd length slice.
var ErrLength = errors.New("bcd: odd length bcd string")

// EncodedLen returns the length of an encoding of n source bytes.
func EncodedLen(n int) int { return n / 2 }

// Encode encodes src into dst.
func Encode(dst, src []byte) int {
	for i := 0; i < len(src)/2; i++ {
		dst[i] = src[i*2]<<4 + src[i*2+1]
	}
	return len(src) / 2
}

// DecodedLen returns the length of an decoding of x source bytes.
func DecodedLen(x int) int { return x * 2 }

// Decode decodes src into dst.
func Decode(dst, src []byte) (int, error) {
	if len(src)*2%2 == 1 {
		return 0, ErrLength
	}
	for i := 0; i < len(src); i++ {
		dst[i*2] = src[i] >> 4
		dst[i*2+1] = src[i] & 0x0f
	}
	return len(src) * 2, nil
}
