// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Import package fmt as well as tserr
import (
	"fmt" // fmt

	"github.com/thorsphere/tserr" // tserr
)

// NearEqual returns whether a and b near equal with the maximum absolute difference maxDiff.
// It returns true if the absolute value of the difference of a and b equals or is lower than maxDiff.
// It returns false if the absolute value of the difference of a and b is higher than maxDiff.
func NearEqual[T Floating](a, b, maxDiff T) bool {
	return Abs(a-b) <= Abs(maxDiff)
}

// EqualS returns an error if signed integer slices x and y are not equal in length and in their values.
// If signed integer slices x and y are equal in length and all of their values, EqualS returns nil.
// If x or y are nil, EqualS returns an error.
func EqualS[T Sinteger](x, y []T) error {
	// Return an error if x is nil
	if x == nil {
		return tserr.NotExistent("x")
	}
	// Return an error if y is nil
	if y == nil {
		return tserr.NotExistent("y")
	}
	// Return an error, if the length of x and y are not equal
	if len(x) != len(y) {
		return tserr.Equal(&tserr.EqualArgs{Var: "Length of y", Actual: int64(len(y)), Want: int64(len(x))})
	}
	// Iterate over range of x
	for i, v := range x {
		// Compare v = x[i] with y[i]
		if v != y[i] {
			// Return an error, if v = x[i] and y[i] are not equal
			return tserr.Equal(&tserr.EqualArgs{Var: fmt.Sprintf("y[%d]", i), Actual: int64(y[i]), Want: int64(v)})
		}
	}
	return nil
}
