// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package lpstats

// Sign returns -1  if a is negative. It returns +1
// if positive and 0 if a is zero.
func Sign[T Signed](a T) T {
	// Return 0 in case a equals zero
	if a == 0 {
		return 0
	}
	// Return -1 if a is negative
	if a < 0 {
		return -1
	}
	// Return +1 otherwise
	return 1
}
