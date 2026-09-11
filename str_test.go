// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package lpstats_test

import (
	"fmt"
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

// testCountArgs calls CountArgs for format string f and compares the returned
// number of arguments with the wanted number w. The test fails, if CountArgs
// returns an error or a number of arguments not equal to w.
func testCountArgs(t *testing.T, f string, w int) {
	// Panic if t is nil
	if t == nil {
		panic(tserr.NilPtr())
	}
	// Call CountArgs with the format string and store the result
	r, e := lpstats.CountArgs(f)
	// The test fails, if CountArgs returns an error
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "CountArgs", Fn: f, Err: e}))
	}
	// The test fails, if the returned number of arguments is not equal to w
	if r != w {
		t.Error(tserr.EqualInt(&tserr.EqualIntArgs{Var: "CountArgs(" + f + ")", Actual: int64(r), Want: int64(w)}))
	}
}

// TestCountArgs tests the returned value of CountArgs for valid format strings.
// It fails, if CountArgs does not return the wanted number of arguments.
func TestCountArgs(t *testing.T) {
	// Table of test cases with format string f and wanted number of
	// arguments w
	tests := []struct {
		f string // format string
		w int    // wanted number of arguments
	}{
		// No verbs
		{"", 0},
		{"no verbs", 0},
		{"100%%", 0},
		{"a%%b%%c", 0},
		{"%5%", 0}, // a verb of % consumes no argument
		// Simple verbs
		{"%d", 1},
		{"%v %v", 2},
		{"%s%%d", 1},
		// Flags, width and precision belong to one verb
		{"%-08.3f", 1},
		{"%+d %#x % d", 3},
		{"%9.2f", 1},
		{"%9.f", 1}, // precision 0
		// Width and precision of * consume an additional argument
		{"%*d", 2},
		{"%.*f", 2},
		{"%*.*f", 3},
		{"%-*d", 2},
		// Explicit argument indexes
		{"%[1]d", 1},
		{"%[1]d %[1]s", 1}, // the same argument is used twice
		{"%[2]d %[1]d", 2}, // reordered
		{"%[2]d %d", 3},    // after [2], the next verb uses argument 3
		{"%[1]*d", 2},      // width from argument 1, verb from argument 2
		{"%.[2]*f", 3},     // precision from argument 2, verb from argument 3
		{"%[1]*.[3]f", 3},  // width from argument 1, precision from argument 3
		{"%[3]*.[2]*[1]f", 3},
		{"%5[3]d", 3},      // index after the width
		{"%*[2]d", 2},      // index after a width of *
		{"%.*[3]f", 3},     // index after a precision of *
		{"%d %[2]s %d", 3}, // sequential, then index, then sequential
		// In %[1][2]d, the second '[' is the verb character itself,
		// because fmt parses at most one index per position. The verb
		// '[' consumes the first argument; "2]d" is literal text.
		{"%[1]d %[1][2]d", 1},
		// Width from argument 1, precision index [2] with literal
		// precision digits 5, verb from argument 2. Valid, because the
		// '*' resets afterIndex: the precision index [2] is parsed after
		// the '.', and precision digits after a precision index are
		// allowed (unlike width digits after a width index, %[3]2d).
		{"%[1]*.[2]5f", 2},
		{"100%% %d", 1},
		// Multibyte characters are not confused with verbs
		{"Grüße %d", 1},
		{"日本語%s", 1},
	}
	// Iterate over the test cases
	for _, tt := range tests {
		// Call testCountArgs with the format string and the wanted number
		testCountArgs(t, tt.f, tt.w)
	}
}

// TestCountArgsErr tests the returned value of CountArgs for invalid format strings.
// It fails, if CountArgs does not return an error.
func TestCountArgsErr(t *testing.T) {
	// Table of test cases with invalid format strings f
	tests := []struct {
		f string // invalid format string
	}{
		// Trailing % without a verb
		{"100%"},
		// Flags, width or precision without a verb
		{"%-"},
		{"%5"},
		{"%.3"},
		{"%08"},
		{"%[3]"},
		{"%.*"},
		// Malformed argument indexes
		{"%[d"},
		{"%[]d"},
		{"%[0]d"},
		{"%[123"},
		{"%.[d"},
		{"%5[d"},
		// Width or precision after an argument index
		{"%[3]2d"},
		{"%[3].2d"},
		{"%[3].d"},
		// Malformed index in every position
		{"%[1]d %[d"},
	}
	// Iterate over the test cases
	for _, tt := range tests {
		// Call CountArgs with the invalid format string
		_, e := lpstats.CountArgs(tt.f)
		// The test fails, if CountArgs does not return an error
		if e == nil {
			t.Error(tserr.NilFailed("CountArgs(" + tt.f + ")"))
		}
	}
}

// TestEqualStrPtrNilNil tests the returned value of EqualStrPtr for two nil pointers.
// It fails if EqualStrPtr does not return true for two nil pointers.
func TestEqualStrPtrNilNil(t *testing.T) {
	// Test with two nil pointers
	var a, b *string = nil, nil
	// The expected result is true, since both pointers are nil
	w := true
	// If EqualStrPtr does not return the expected value, the test fails
	if r := lpstats.EqualStrPtr(a, b); r != w {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "EqualStrPtr", Actual: fmt.Sprintf("%t", r), Want: fmt.Sprintf("%t", w)}))
	}
}

// TestEqualStrPtrNilOne tests the returned value of EqualStrPtr for one nil pointer
// and one non-nil pointer. It fails if EqualStrPtr does not return false.
func TestEqualStrPtrNilOne(t *testing.T) {
	// Test with one non-nil pointer and one nil pointer
	a := lpstats.PtrStr("GDP Growth Rate")
	var b *string = nil
	// The expected result is false, since one pointer is nil and the other is not
	w := false
	// If EqualStrPtr does not return the expected value for a non-nil and a nil
	// pointer, the test fails
	if r := lpstats.EqualStrPtr(a, b); r != w {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "EqualStrPtr", Actual: fmt.Sprintf("%t", r), Want: fmt.Sprintf("%t", w)}))
	}
	// If EqualStrPtr does not return the expected value for a nil and a non-nil
	// pointer, the test fails
	if r := lpstats.EqualStrPtr(b, a); r != w {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "EqualStrPtr", Actual: fmt.Sprintf("%t", r), Want: fmt.Sprintf("%t", w)}))
	}
}

// TestEqualStrPtrEqual tests the returned value of EqualStrPtr for two non-nil
// pointers to equal strings. It fails if EqualStrPtr does not return true.
func TestEqualStrPtrEqual(t *testing.T) {
	// Test with two non-nil pointers to equal strings
	a := lpstats.PtrStr("GDP Growth Rate")
	b := lpstats.PtrStr("GDP Growth Rate")
	// The expected result is true, since both pointers point to equal values
	w := true
	// If EqualStrPtr does not return the expected value, the test fails
	if r := lpstats.EqualStrPtr(a, b); r != w {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "EqualStrPtr", Actual: fmt.Sprintf("%t", r), Want: fmt.Sprintf("%t", w)}))
	}
}

// TestEqualStrPtrNotEqual tests the returned value of EqualStrPtr for two non-nil
// pointers to unequal strings. It fails if EqualStrPtr does not return false.
func TestEqualStrPtrNotEqual(t *testing.T) {
	// Test with two non-nil pointers to unequal strings
	a := lpstats.PtrStr("GDP Growth Rate")
	b := lpstats.PtrStr("Inflation Rate")
	// The expected result is false, since the pointers point to unequal values
	w := false
	// If EqualStrPtr does not return the expected value, the test fails
	if r := lpstats.EqualStrPtr(a, b); r != w {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "EqualStrPtr", Actual: fmt.Sprintf("%t", r), Want: fmt.Sprintf("%t", w)}))
	}
}
