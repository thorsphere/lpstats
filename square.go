// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package lpstats

// Square returns the square of x as float64
func Square[T Number](x T) float64 {
	y := float64(x)
	return y * y
}
