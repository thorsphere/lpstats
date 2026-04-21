// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats_test

// Import package testing and the lpstats package for testing the Sum function.
import (
	"testing" // testing

	"github.com/thorsphere/lpstats" // lpstats
	"github.com/thorsphere/tserr"   // tserr
)

// TestSumE tests if Sum returns an error for an empty slice
func TestSumE(t *testing.T) {
	if e := testfa(t, nil /*a*/, 0 /*w*/, lpstats.Sum[int]); e == nil {
		t.Error(tserr.NilFailed("Sum"))
	}
}

// TestSumi tests the returned value of Sum for a slice of integers.
// It fails if Sum does not return the wanted value.
func TestSumi(t *testing.T) {
	if e := testfa(t, []int{1, 2, 3, 4, 5, 6} /*a*/, 21 /*w*/, lpstats.Sum[int]); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Sum", Fn: "slice of integers", Err: e}))
	}
}

// TestSumf tests the returned value of Sum for a slice of float64.
// It fails if Sum does not return the wanted value.
func TestSumf(t *testing.T) {
	if e := testfa(t, []float64{1.1, 2.1, 3.1, 4.1, 5.1, 6.1} /*a*/, 21.6 /*w*/, lpstats.Sum[float64]); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Sum", Fn: "slice of integers", Err: e}))
	}
}
