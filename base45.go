// Copyright 2021 Philip Schlump. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package base45 implements base45 encoding as specified by by the
// proposed standard: https://datatracker.ietf.org/doc/draft-faltstrom-base45/
package base45

import (
	"bytes"
	"fmt"
	"log"
)

// This character set matches with what can be encoded in a QR code
// efficiently.
var encCharSet = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:")

func init() {
	if len(encCharSet) != 45 {
		log.Fatal("Incorrect length for encCharSet")
	}
}

// Base45Encode takes a set of bytes and encodes it to a string
// in base45 with a character set that is efficient for QR code
// generation.
func Base45Encode(s []byte) (rv string) {

	var buffer bytes.Buffer

	// Walk the input 2 bytes at a time to make each pair into a 16 bit number.
	for i := 0; i+1 < len(s); i += 2 {
		v := (uint(s[i]) << 8) | uint(s[i+1])
		for j := 0; j < 3; j++ {
			buffer.WriteString(string(encCharSet[(int(v) % 45)]))
			v = v / 45
		}
	}

	// If the input is an odd number of bytes long, then deal with the last byte.
	if len(s)%2 == 1 {
		v := uint(s[len(s)-1]) // last byte
		for j := 0; j < 2; j++ {
			buffer.WriteString(string(encCharSet[(int(v) % 45)]))
			v = v / 45
		}
	}

	return buffer.String()
}

// Base45Decode takes a base45 encode string and converts it
// back into its original binary form.  If there is a character
// that is not in the base45 character set then you get an error
// back.
func Base45Decode(s string) (out []byte, err error) {

	var buffer bytes.Buffer

	// Walk the input three bytes at a time converting to a 16bit int and
	// appending that to the buffer.
	for i := 2; i < len(s); i += 3 {
		v := 0
		for j := 0; j < 3; j++ {
			iv := bytes.IndexByte(encCharSet, s[i-j])
			if iv < 0 {
				err = fmt.Errorf("Invalid character ->%c<- at position %d in ->%s<- , not in base45 encoding character set", s[i-j], i-j, s)
				return
			}
			v = v*45 + iv
		}
		buffer.WriteByte(byte((v >> 8) & 0xff))
		buffer.WriteByte(byte(v & 0xff))
	}

	// Finish the last two bytes if they exist
	if len(s)%3 > 0 {
		v := 0
		i := len(s) - 1
		for j := 0; j < 2; j++ {
			iv := bytes.IndexByte(encCharSet, s[i-j])
			if iv < 0 {
				err = fmt.Errorf("Invalid character ->%c<- at position %d in ->%s<- , not in base45 encoding character set", s[i-j], i-j, s)
				return
			}
			v = v*45 + iv
		}
		buffer.WriteByte(byte(v))
	}

	return buffer.Bytes(), nil
}
