// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats_test

// Import package testing
import (
	"testing" // testing

	"github.com/thorsphere/lpstats" // lpstats
	"github.com/thorsphere/tserr"   // tserr
)

// TestExpectedValueUi tests the returned value of ExpectedValueU for integers as arguments.
// It fails if ExpectedValueU does not return the wanted value.
func TestExpectedValueUi(t *testing.T) {
	testf2(t, 10 /*a*/, 20 /*b*/, 15 /*w*/, lpstats.ExpectedValueU[int])
}

// TestExpectedValueUf tests the returned value of ExpectedValueU for float64 as arguments.
// It fails if ExpectedValueU does not return the wanted value.
func TestExpectedValueUf(t *testing.T) {
	testf2(t, 10.1 /*a*/, 20.1 /*b*/, 15.1 /*w*/, lpstats.ExpectedValueU[float64])
}

// TestArithmeticMeanE tests if ArithmeticMean returns an error for an empty slice
func TestArithmeticMeanE(t *testing.T) {
	if e := testfa(t, nil /*a*/, 0 /*w*/, lpstats.ArithmeticMean[int]); e == nil {
		t.Error(tserr.NilFailed("ArithmeticMean"))
	}
}

// TestArithmeticMeani tests the returned value of ArithmeticMean for a slice of integers.
// It fails if ArithmeticMean does not return the wanted value.
func TestArithmeticMeani(t *testing.T) {
	if e := testfa(t, []int{2500, 2700, 2400, 2300, 2550, 2650, 2750, 2450, 2600, 2400} /*a*/, 2530 /*w*/, lpstats.ArithmeticMean[int]); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ArithmeticMean", Fn: "slice of integers", Err: e}))
	}
}

// TestArithmeticMeanf tests the returned value of ArithmeticMean for a slice of float64.
// It fails if ArithmeticMean does not return the wanted value.
func TestArithmeticMeanf(t *testing.T) {
	if e := testfa(t, []float64{2500, 2700, 2400, 2300, 2550, 2650, 2750, 2450, 2600, 2400} /*a*/, 2530 /*w*/, lpstats.ArithmeticMean[float64]); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ArithmeticMean", Fn: "slice of integers", Err: e}))
	}
}
