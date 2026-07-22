// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package lpstats

// PtrStr returns a pointer to a string.
func PtrStr(a string) *string {
	return &a
}

// CopyStrPtr returns a copy of a pointer to a string.
// It returns nil if p is nil.
func CopyStrPtr(p *string) *string {
	// Return nil if p is nil
	if p == nil {
		return nil
	}
	// Create a new value of type string and copy the value pointed to by p into it
	v := *p
	// Return a pointer to the new value
	return &v
}

// EqualStrPtr returns whether two pointers to strings are equal.
// It returns true if both pointers are nil, or if both pointers are non-nil and the values they point to are equal.
// It returns false if one pointer is nil and the other is not, or if both pointers are non-nil and the values
// they point to are not equal.
func EqualStrPtr(a, b *string) bool {
	// If both pointers are nil, they are considered equal
	if a == nil && b == nil {
		return true
	}
	// If one pointer is nil and the other is not, they are not equal
	if a == nil || b == nil {
		return false
	}
	// If both pointers are non-nil, compare the values they point to
	return *a == *b
}
