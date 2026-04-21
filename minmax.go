// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Max returns the maximum of a and b. If a is higher than b, it returns a.
// Otherwise, it returns b. Marked as deprecated due to introduction of the
// native Go max function with Go version 1.21. Will be removed in future
// releases.
func Max[T Number](a, b T) T {
	return max(a, b)
}

// Min returns the minimum of a and b. If a is lower than b, it returns a.
// Otherwise, it returns b. Marked as deprecated due to introduction of the
// native Go min function with Go version 1.21. Will be removed in future
// releases.
func Min[T Number](a, b T) T {
	return min(a, b)
}
