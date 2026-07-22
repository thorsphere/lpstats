// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package lpstats

// Import the fmt package for formatting strings
import (
	"fmt"
)

const (
	// MaxPrecision is the maximum number of decimal places for formatting.
	maxPrecision = 10
)

// FmtFloatPtr formats a pointer to a number as a string with the specified precision.
// It returns "nil" if the pointer is nil.
func FmtFloatPtr[T Floating](a *T, prec int) string {
	// Return "nil" if a is nil
	if a == nil {
		return "nil"
	}
	// Set the precision to 0 if it is less than 0
	if prec < 0 {
		prec = 0
	} else if prec > maxPrecision { // Set the precision to MaxPrecision if it is greater than MaxPrecision
		prec = maxPrecision
	}
	// Format the value pointed to by a with the given precision and return it as a string
	return fmt.Sprintf("%.*f", prec, *a)
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

// PtrFloat returns a pointer to a number.
func PtrFloat[T Floating](a T) *T {
	return &a
}

// CopyFloatPtr returns a copy of a pointer to a number.
// It returns nil if p is nil.
func CopyFloatPtr[T Floating](p *T) *T {
	// Return nil if p is nil
	if p == nil {
		return nil
	}
	// Create a new value of type T and copy the value pointed to by p into it
	v := *p
	// Return a pointer to the new value
	return &v
}
