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
	ptr := lpstats.PtrFloat(3.14)
	// The expected result is "3.1" since FmtFloatPtr formats the float with one decimal place
	w := "3.1"
	// Call FmtFloatPtr with the pointer and store the result
	result := lpstats.FmtFloatPtr(ptr, 1)
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
	result := lpstats.FmtFloatPtr(ptr,1)
	// If the result does not equal the expected value, the test fails
	if result != w {
		t.Error(tserr.EqualStr(&tserr.EqualStrArgs{Var: "FmtFloatPtr", Actual: result, Want: w}))
	}
}

// TestFloatPtrEqual tests the returned value of NearEqualFloatPtr for two non-nil pointers to equal float64 values.
// It fails if NearEqualFloatPtr does not return true for two non-nil pointers to equal values.
func TestFloatPtrEqual(t *testing.T) {
	// Test with two non-nil pointers to equal values
	ptr1 := lpstats.PtrFloat(3.14)
	ptr2 := lpstats.PtrFloat(3.14)
	// The expected result is true since both pointers point to equal values
	w := true
	// Call testfb with the pointers, expected value, and NearEqualFloatPtr function to test for equality
	testfb(t, ptr1, ptr2, w, lpstats.NearEqualFloatPtr)
}

// TestFloatPtrEqualNil tests the returned value of NearEqualFloatPtr for two nil pointers.
// It fails if NearEqualFloatPtr does not return true for two nil pointers.
func TestFloatPtrEqualNil1(t *testing.T) {
	// Test with two nil pointers
	var ptr1 *float64 = nil
	var ptr2 *float64 = nil
	// The expected result is true since both pointers are nil
	w := true
	// Call testfb with the pointers, expected value, and NearEqualFloatPtr function to test for equality
	testfb(t, ptr1, ptr2, w, lpstats.NearEqualFloatPtr)
}

// TestFloatPtrEqualNil tests the returned value of NearEqualFloatPtr for one nil pointer and one non-nil pointer.
// It fails if NearEqualFloatPtr does not return false for one nil pointer and one non-nil pointer.
func TestFloatPtrEqualNil2(t *testing.T) {
	// Test with one nil pointer and one non-nil pointer
	ptr1 := lpstats.PtrFloat(3.14)
	var ptr2 *float64 = nil
	// The expected result is false since the pointers point to different values
	w := false
	// Call testfb with the pointers, expected value, and NearEqualFloatPtr function to test for equality
	testfb(t, ptr1, ptr2, w, lpstats.NearEqualFloatPtr)
	// Call testfb with the pointers, expected value, and NearEqualFloatPtr function to test for equality in the other order
	testfb(t, ptr2, ptr1, w, lpstats.NearEqualFloatPtr)
}

// TestFloatPtrNotEqual tests the returned value of NearEqualFloatPtr for two non-nil pointers to unequal values.
// It fails if NearEqualFloatPtr does not return false for two non-nil pointers to unequal values.
func TestFloatPtrNotEqual(t *testing.T) {
	// Test with two non-nil pointers to unequal values
	ptr1 := lpstats.PtrFloat(3.14)
	ptr2 := lpstats.PtrFloat(2.71)
	// The expected result is false since the pointers point to unequal values
	w := false
	// Call testfb with the pointers, expected value, and NearEqualFloatPtr function to test for inequality
	testfb(t, ptr1, ptr2, w, lpstats.NearEqualFloatPtr)
}

// TestCopyFloatPtr tests the returned value of CopyFloatPtr for a non-nil pointer to a float64.
func TestCopyFloatPtr(t *testing.T) {
	// Test with a non-nil pointer
	ptr := lpstats.PtrFloat(3.14)
	// Call CopyFloatPtr with the pointer and store the result
	result := lpstats.CopyFloatPtr(ptr)
	// If the result does not equal the expected value, the test fails
	if *result != *ptr {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "CopyFloatPtr", Actual: *result, Want: *ptr}))
	}
	// If the result is not a deep copy of the pointer, the test fails
	if result == ptr {
		t.Error(tserr.NotEqual(&tserr.NotEqualArgs{X: "Pointer to a float64", Y: "Deep copy of the pointer to a float64"}))
	}
}

// TestCopyFloatPtrNil tests the returned value of CopyFloatPtr for a nil pointer.
func TestCopyFloatPtrNil(t *testing.T) {
	// Test with a nil pointer
	var ptr *float64 = nil
	// Call CopyFloatPtr with the pointer and store the result
	result := lpstats.CopyFloatPtr(ptr)
	// If the result does not equal the expected value, the test fails
	if result != nil {
		t.Error(tserr.NilExpected("CopyFloatPtr"))
	}
}
