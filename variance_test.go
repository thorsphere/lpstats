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

// TestVarianceN tests the returned value of VarianceN for a 6-sided die. If
// the result does not equal to 2.92 with the precision of two decimal places, the test fails.
func TestVarianceN(t *testing.T) {
	testf(t, 6 /*a*/, 2.92 /*w*/, lpstats.VarianceN[uint])
}

// TestVarianceUi tests the returned value of VarianceU for an integer interval. If
// the result does not equal the expected value, the test fails.
func TestVarianceUi(t *testing.T) {
	testf2(t, 10 /*a*/, 20 /*b*/, 8.33 /*w*/, lpstats.VarianceU[int])
}

// TestVarianceUf tests the returned value of VarianceU for an float64 interval. If
// the result does not equal the expected value, the test fails.
func TestVarianceUf(t *testing.T) {
	testf2(t, 20.1 /*a*/, 10.1 /*b*/, 8.33 /*w*/, lpstats.VarianceU[float64])
}

// TestVarianceE tests if Variance returns an error for an empty slice
func TestVarianceE(t *testing.T) {
	if e := testfa(t, nil /*a*/, 0 /*w*/, lpstats.Variance[int]); e == nil {
		t.Error(tserr.NilFailed("Variance"))
	}
}

// TestVariancei tests the returned value of Variance for a slice of integers.
// It fails if Sum does not return the wanted value.
func TestVariancei(t *testing.T) {
	if e := testfa(t, []int{1, 2, 3, 4, 5, 6} /*a*/, 2.92 /*w*/, lpstats.Variance[int]); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Variance", Fn: "slice of integers", Err: e}))
	}
}

// TestVariancef tests the returned value of Variance for a slice of float64.
// It fails if Sum does not return the wanted value.
func TestVariancef(t *testing.T) {
	if e := testfa(t, []float64{1, 2, 3, 4, 5, 6} /*a*/, 2.92 /*w*/, lpstats.Variance[float64]); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Variance", Fn: "slice of integers", Err: e}))
	}
}
