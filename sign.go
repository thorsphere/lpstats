// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
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
