// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats_test

// Import package testing and tserr
import (
	"testing" // testing

	"github.com/thorsphere/lpstats" // lpstats
	"github.com/thorsphere/tserr"   // tserr
)

const (
	// maxDiff defines the maximum difference for near equal floats
	maxDiff float64 = 0.01
	// sLen defines the number of elements of a slice
	sLen int = 10
)

// testi tests function f to return wanted integer w for argument a. If the
// result of f for a does not equal w, the test fails.
func testi[T lpstats.Number](t *testing.T, a T, w int64, f func(T) int64) {
	if t == nil {
		panic("nil pointer")
	}
	// Retrieve f of a in r
	r := f(a)
	// If the result r does not equal the wanted result w, the test fails
	if r != w {
		t.Error(tserr.Equal(&tserr.EqualArgs{Var: "function of a", Actual: r, Want: w}))
	}
}

// testf tests function f to return wanted float64 w for argument a. If the
// result of f for a does not equal w with the precision of two decimal places, the test fails.
func testf[T lpstats.Number](t *testing.T, a T, w float64, f func(T) float64) {
	if t == nil {
		panic("nil pointer")
	}
	// Retrieve f of a in r
	r := f(a)
	// If the result r does not equal the wanted result w, the test fails
	if !lpstats.NearEqual(r, w, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "function of a", Actual: r, Want: w}))
	}
}

// testf2 tests function f to return wanted float64 w for arguments a and b. If the
// result of f for a and b does not equal w with the precision of two decimal places, the test fails.
func testf2[T lpstats.Number](t *testing.T, a, b T, w float64, f func(T, T) float64) {
	if t == nil {
		panic("nil pointer")
	}
	// Retrieve f of a and b in r
	r := f(a, b)
	// If the result r does not equal the wanted result w, the test fails
	if !lpstats.NearEqual(r, w, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "function of a and b", Actual: r, Want: w}))
	}
}

// testi2 tests function f to return wanted int64 w for arguments a and b. If the
// result of f for a and b does not equal w, the test fails.
func testi2[T lpstats.Number](t *testing.T, a, b T, w int64, f func(T, T) int64) {
	if t == nil {
		panic("nil pointer")
	}
	// Retrieve f of a and b in r
	r := f(a, b)
	// If the result r does not equal the wanted result w, the test fails
	if r != w {
		t.Error(tserr.Equal(&tserr.EqualArgs{Var: "function of a and b", Actual: r, Want: w}))
	}
}

// testfa tests function f to return wanted float64 w for slice a. If the
// result of f for slice a does not equal w with the precision of two decimal places, the test fails.
// It returns the error of f, if any
func testfa[T lpstats.Number](t *testing.T, a []T, w float64, f func([]T) (float64, error)) error {
	if t == nil {
		panic("nil pointer")
	}
	r, e := f(a)
	// If the result b does not equal the wanted result w, the test fails
	if !lpstats.NearEqual(r, w, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "function of a", Actual: r, Want: w}))
	}
	return e
}
