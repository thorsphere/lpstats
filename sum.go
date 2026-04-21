// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Import package tserr
import "github.com/thorsphere/tserr" // tserr

// Sum returns the sum of all elements in array x as float64.
func Sum[T Number](x []T) (float64, error) {
	// Return an error if x is empty
	if len(x) == 0 {
		return float64(0), tserr.Empty("x")
	}
	var sum float64
	for i := 0; i < len(x); i++ {
		sum += float64(x[i])
	}
	return sum, nil
}
