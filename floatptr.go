// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Import the fmt package for formatting strings
import "fmt"

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

// FloatPtrEqual compares two pointers to numbers for equality.
// It returns true if both pointers are nil or if they point to equal values, and false otherwise.
func FloatPtrEqual[T Floating](a, b *T) bool {
	// If both pointers are nil, they are considered equal
	if a == nil && b == nil {
		return true
	}
	// If one pointer is nil and the other is not, they are not equal
	if a == nil || b == nil {
		return false
	}
	// If both pointers are not nil, compare the values they point to
	return *a == *b
}
