// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
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
