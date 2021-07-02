// Copyright 2021 Philip Schlump. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package base45_test

import (
	"testing"

	"github.com/pschlump/base45"
)

func Test_EncodeDecodeMatch(t *testing.T) {

	tests := []struct {
		txt         string
		expectedDec string
		expectedEnc string
	}{
		{ // 0
			txt:         "http://example.com/testdir/b.html",
			expectedDec: "http://example.com/testdir/b.html",
			expectedEnc: `A9DIWE0G7S:5$9FQ$DTVD+%5+3E/:56$C6WE*EDP:50*5FWEI2`,
		},
		{ // 1
			txt:         `{"status":"success"}`,
			expectedDec: `{"status":"success"}`,
			expectedEnc: `MPF QEIEC7%EWE4:F4 $EKPCZQE9G4`,
		},
		{ // 2
			txt:         `abcde`,
			expectedDec: `abcde`,
			expectedEnc: `0ECJPCB2`,
		},
		{ // 3
			txt:         `abcdef`,
			expectedDec: `abcdef`,
			expectedEnc: `0ECJPC% C`,
		},
	}

	for ii, test := range tests {
		e := base45.Base45Encode([]byte(test.txt))
		by, _ := base45.Base45Decode(e)
		s := string(by)

		// fmt.Printf("in ->%s<- enc ->%s<-\n", test.txt, e)

		if test.txt != s {
			t.Errorf("Test %2d, after encode(decode(%s)) != ->%s<-\n", ii, s, test.txt)
		}
		if test.expectedDec != s {
			t.Errorf("Test %2d, expected ->%s<- got ->%s<-\n", ii, test.expectedDec, s)
		}
		if test.expectedEnc != string(e) {
			t.Errorf("Test %2d, expected ->%s<- got ->%s<-\n", ii, test.expectedEnc, e)
		}
	}
}

// TODO - test generating error with invalid input
