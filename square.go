// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Square returns the square of x as float64
func Square[T Number](x T) float64 {
	y := float64(x)
	return y * y
}
