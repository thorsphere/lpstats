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

// TestNEf tests NearEqual to return false for a difference of a and b greater than the
// provided maximum difference. It fails if NearEqual returns true.
func TestNEf(t *testing.T) {
	if lpstats.NearEqual(0.01, 0.03, -0.01) {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "NearEqual", Actual: "true", Want: "false"}))
	}
}

// TestNEe tests NearEqual to return true for a difference of a and b equal to the
// provided maximum difference. It fails if NearEqual returns false.
func TestNEe(t *testing.T) {
	if !lpstats.NearEqual(0.01, 0.02, -0.01) {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "NearEqual", Actual: "false", Want: "true"}))
	}
}

// TestNEt tests NearEqual to return true for a difference of a and b lower than the
// provided maximum difference. It fails if NearEqual returns false.
func TestNEt(t *testing.T) {
	if !lpstats.NearEqual(0.01, 0.02, -0.02) {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "NearEqual", Actual: "false", Want: "true"}))
	}
}

// TestESl tests EqualS for two slices of integers with different length.
// The test fails, if EqualS returns nil.
func TestESl(t *testing.T) {
	// Allocate and initialize int slices y1 and y2 with different lengths sLen and sLen-1
	y1, y2 := make([]int, sLen), make([]int, sLen-1)
	// The test fails, if EqualS returns nil
	if e := lpstats.EqualS(y1, y2); e == nil {
		t.Error(tserr.NilFailed("EqualS"))
	}
}

// TestESv tests EqualS for two slices of integers with different values.
// The test fails, if EqualS returns nil.
func TestESv(t *testing.T) {
	// Allocate and initialize int slices y1 and y2 with length sLen
	y1, y2 := make([]int, sLen), make([]int, sLen)
	// Set y1 and y2 last value to different values 1 and 2
	y1[sLen-1], y2[sLen-1] = 1, 2
	// The test fails, if EqualS returns nil
	if e := lpstats.EqualS(y1, y2); e == nil {
		t.Error(tserr.NilFailed("EqualS"))
	}
}

// TestESn tests EqualS for nil slices of integers.
// The test fails, if EqualS returns nil.
func TestESn(t *testing.T) {
	// Allocate and initialize int slice y1 with length sLen and set int slice y2 to nil
	var y1, y2 []int = make([]int, sLen), nil
	// The test fails, if EqualS returns nil
	if e := lpstats.EqualS(y1, y2); e == nil {
		t.Error(tserr.NilFailed("EqualS"))
	}
	// The test fails, if EqualS returns nil
	if e := lpstats.EqualS(y2, y1); e == nil {
		t.Error(tserr.NilFailed("EqualS"))
	}
}

// TestES tests EqualS for two slices of integers with same the length and the same values.
// The test fails, if EqualS returns an error.
func TestES(t *testing.T) {
	// Allocate and initialize int slices y1 and y2 with length sLen
	y1, y2 := make([]int, sLen), make([]int, sLen)
	// Set y1 and y2 last value to the same value 1
	y1[sLen-1], y2[sLen-1] = 1, 1
	// The test fails, if EqualS returns an error
	if e := lpstats.EqualS(y1, y2); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "EqualS", Fn: "y1 and y2", Err: e}))
	}
}
