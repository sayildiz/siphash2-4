package main

import (
	"encoding/binary"
	"fmt"
	"math/bits"
)

func main() {
	//primitve implementation of Siphash 2-4 with test values from paper

	//key from paper
	var k0 uint64 = 0x0706050403020100
	var k1 uint64 = 0x0f0e0d0c0b0a0908

	// initialize constants
	var v0 uint64 = k0 ^ 0x736f6d6570736575
	var v1 uint64 = k1 ^ 0x646f72616e646f6d
	var v2 uint64 = k0 ^ 0x6c7967656e657261
	var v3 uint64 = k1 ^ 0x7465646279746573

	// test message from paper
	message := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

	// split message into 8 byte words
	msg1 := binary.LittleEndian.Uint64(message[:8])
	msg2 := binary.LittleEndian.Uint64(message[8:])

	//compress msg1
	v3 ^= msg1
	sipRound(&v0, &v1, &v2, &v3)
	sipRound(&v0, &v1, &v2, &v3)
	v0 ^= msg1
	//compress msg2
	v3 ^= msg2
	sipRound(&v0, &v1, &v2, &v3)
	sipRound(&v0, &v1, &v2, &v3)
	v0 ^= msg2

	//finalize
	v2 ^= 0xff
	sipRound(&v0, &v1, &v2, &v3)
	sipRound(&v0, &v1, &v2, &v3)
	sipRound(&v0, &v1, &v2, &v3)
	sipRound(&v0, &v1, &v2, &v3)
	res := v0 ^ v1 ^ v2 ^ v3

	fmt.Printf("%x", res)
}

func sipRound(v0 *uint64, v1 *uint64, v2 *uint64, v3 *uint64) {
	*v0 += *v1
	*v1 = bits.RotateLeft64(*v1, 13)
	*v1 ^= *v0
	*v0 = bits.RotateLeft64(*v0, 32)

	*v2 += *v3
	*v3 = bits.RotateLeft64(*v3, 16)
	*v3 ^= *v2

	*v0 += *v3
	*v3 = bits.RotateLeft64(*v3, 21)
	*v3 ^= *v0

	*v2 += *v1
	*v1 = bits.RotateLeft64(*v1, 17)
	*v1 ^= *v2
	*v2 = bits.RotateLeft64(*v2, 32)
}
