// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package lpstats_test

import (
	"testing"

	"github.com/thorsphere/lpstats"
	"github.com/thorsphere/tserr"
)

// TestCopyStrPtr tests the returned value of CopyStrPtr for a non-nil pointer to a string.
func TestCopyStrPtr(t *testing.T) {
	// Test with a non-nil pointer
	ptr := lpstats.PtrStr("GDP Growth Rate")
	// Call CopyFloatPtr with the pointer and store the result
	result := lpstats.CopyStrPtr(ptr)
	// If the result does not equal the expected value, the test fails
	if *result != *ptr {
		t.Error(tserr.EqualStr(&tserr.EqualStrArgs{Var: "CopyFloatPtr", Actual: *result, Want: *ptr}))
	}
	// If the result is not a deep copy of the pointer, the test fails
	if result == ptr {
		t.Error(tserr.NotEqual(&tserr.NotEqualArgs{X: "Pointer to a string", Y: "Deep copy of the pointer to a string"}))
	}
}

// TestCopyStrPtrNil tests the returned value of CopyStrPtr for a nil pointer.
func TestCopyStrPtrNil(t *testing.T) {
	// Test with a nil pointer
	var ptr *string = nil
	// Call CopyFloatPtr with the pointer and store the result
	result := lpstats.CopyStrPtr(ptr)
	// If the result does not equal the expected value, the test fails
	if result != nil {
		t.Error(tserr.NilExpected("CopyStrPtr"))
	}
}
