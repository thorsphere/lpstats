// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Import the fmt package for formatting strings
import (
	"fmt"
)

// FmtFloatPtr formats a pointer to a number as a string with one decimal place.
// It returns "nil" if the pointer is nil.
func FmtFloatPtr[T Floating](a *T) string {
	// Return "nil" if a is nil
	if a == nil {
		return "nil"
	}
	// Format the value pointed to by a with one decimal place and return it as a string
	return fmt.Sprintf("%.1f", *a)
}

// NearEqualFloatPtr returns whether two pointers to numbers are equal within a specified maximum absolute difference maxDiff.
// It returns true if both pointers are nil, or if both pointers are non-nil and the absolute difference of the values they point to is less than or equal to maxDiff.
// It returns false if one pointer is nil and the other is not, or if both pointers are non-nil and the absolute difference of the values they point to is greater than maxDiff.
func NearEqualFloatPtr[T Floating](a, b *T, maxDiff T) bool {
	// If both pointers are nil, they are considered equal
	if a == nil && b == nil {
		return true
	}
	// If one pointer is nil and the other is not, they are not equal
	if a == nil || b == nil {
		return false
	}
	// If both pointers are non-nil, compare the values they point to using NearEqual with the specified maxDiff
	return NearEqual(*a, *b, maxDiff)
}
