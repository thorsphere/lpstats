// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package lpstats

// Abs returns the absolute value for a.
// Note: The smallest value of an integer type does not have a
// matching positive value. The Abs function returns a
// negative value in this case.
func Abs[T Number](a T) T {
	// Return -a if a is negative
	if a < 0 {
		return -a
	}
	// Return +a otherwise
	return a
}
