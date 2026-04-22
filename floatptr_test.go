// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats_test

// Import necessary packages for testing and the lpstats package
import (
	"testing" // testing

	"github.com/thorsphere/lpstats" // lpstats
	"github.com/thorsphere/tserr"   // tserr
)

// TestFmtFloatPtr tests the returned value of FmtFloatPtr for a non-nil pointer to a float64.
// It fails if FmtFloatPtr does not return the wanted value.
func TestFmtFloatPtr(t *testing.T) {
	// Test with a non-nil pointer
	f := 3.14
	ptr := &f
	// The expected result is "3.1" since FmtFloatPtr formats the float with one decimal place
	w := "3.1"
	// Call FmtFloatPtr with the pointer and store the result
	result := lpstats.FmtFloatPtr(ptr)
	// If the result does not equal the expected value, the test fails
	if result != w {
		t.Error(tserr.EqualStr(&tserr.EqualStrArgs{Var: "FmtFloatPtr", Actual: result, Want: w}))
	}
}

// TestFmtFloatPtrNil tests the returned value of FmtFloatPtr for a nil pointer.
// It fails if FmtFloatPtr does not return "nil" for a nil pointer.
func TestFmtFloatPtrNil(t *testing.T) {
	// Test with a nil pointer
	var ptr *float64 = nil
	// The expected result is "nil" since FmtFloatPtr returns "nil" for nil pointers
	w := "nil"
	// Call FmtFloatPtr with the pointer and store the result
	result := lpstats.FmtFloatPtr(ptr)
	// If the result does not equal the expected value, the test fails
	if result != w {
		t.Error(tserr.EqualStr(&tserr.EqualStrArgs{Var: "FmtFloatPtr", Actual: result, Want: w}))
	}
}

// TestFloatPtrEqual tests the returned value of FloatPtrEqual for two non-nil pointers to equal float64 values.
// It fails if FloatPtrEqual does not return true for two non-nil pointers to equal values.
func TestFloatPtrEqual(t *testing.T) {
	// Test with two non-nil pointers to equal values
	f1 := 3.14
	f2 := 3.14
	ptr1 := &f1
	ptr2 := &f2
	// The expected result is true since both pointers point to equal values
	w := true
	// Call testfb with the pointers, expected value, and FloatPtrEqual function to test for equality
	testfb(t, ptr1, ptr2, w, lpstats.FloatPtrEqual)
}

// TestFloatPtrEqualNil tests the returned value of FloatPtrEqual for two nil pointers.
// It fails if FloatPtrEqual does not return true for two nil pointers.
func TestFloatPtrEqualNil1(t *testing.T) {
	// Test with two nil pointers
	var ptr1 *float64 = nil
	var ptr2 *float64 = nil
	// The expected result is true since both pointers are nil
	w := true
	// Call testfb with the pointers, expected value, and FloatPtrEqual function to test for equality
	testfb(t, ptr1, ptr2, w, lpstats.FloatPtrEqual)
}

// TestFloatPtrEqualNil tests the returned value of FloatPtrEqual for one nil pointer and one non-nil pointer.
// It fails if FloatPtrEqual does not return false for one nil pointer and one non-nil pointer.
func TestFloatPtrEqualNil2(t *testing.T) {
	// Test with one nil pointer and one non-nil pointer
	f1 := 3.14
	ptr1 := &f1
	var ptr2 *float64 = nil
	// The expected result is false since the pointers point to different values
	w := false
	// Call testfb with the pointers, expected value, and FloatPtrEqual function to test for equality
	testfb(t, ptr1, ptr2, w, lpstats.FloatPtrEqual)
	// Call testfb with the pointers, expected value, and FloatPtrEqual function to test for equality in the other order
	testfb(t, ptr2, ptr1, w, lpstats.FloatPtrEqual)
}

// TestFloatPtrNotEqual tests the returned value of FloatPtrEqual for two non-nil pointers to unequal values.
// It fails if FloatPtrEqual does not return false for two non-nil pointers to unequal values.
func TestFloatPtrNotEqual(t *testing.T) {
	// Test with two non-nil pointers to unequal values
	f1 := 3.14
	f2 := 2.71
	ptr1 := &f1
	ptr2 := &f2
	// The expected result is false since the pointers point to unequal values
	w := false
	// Call testfb with the pointers, expected value, and FloatPtrEqual function to test for inequality
	testfb(t, ptr1, ptr2, w, lpstats.FloatPtrEqual)
}
