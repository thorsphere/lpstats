// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Import package tserr
import (
	"github.com/thorsphere/tserr" // tserr
)

// Variance returns the variance of discrete value array x as float64. It returns zero and an error
// if x is empty or if the calculation of the arithmetic mean fails.
func Variance[T Number](x []T) (float64, error) {
	// Return an error if x is empty
	if len(x) == 0 {
		return float64(0), tserr.Empty("x")
	}
	var (
		mean, vari float64
		err        error
	)
	// Retrieve arithmetic mean of x
	if mean, err = ArithmeticMean(x); err != nil {
		// Return zero and an error if arithmetic mean fails
		return float64(0), tserr.Op(&tserr.OpArgs{Op: "ArithmeticMean", Fn: "x", Err: err})
	}
	// Calculate and return the arithmetic mean
	for i := range x {
		vari += Square(float64(x[i]) - mean)
	}
	return vari / float64(len(x)), nil
}

// VarianceU returns the variance of the uniform distribution in the interval [a,b] as float64. If b is greater than a
// it returns the variance of the uniform distribution in the interval [b,a].
func VarianceU[T Number](a, b T) float64 {
	// l for lower bound and u for upper bound of interval [l,u]
	var l, u T
	// Define interval of the uniform distribution
	if a > b {
		// Interval is [b,a]
		l = b
		u = a
	} else {
		// Interval is [a,b]
		l = a
		u = b
	}
	// Calculate and return the variance
	return float64(Square(u-l)) / float64(12)
}

// VarianceN returns the Variance of an n-sided die as float64.
func VarianceN[T Uinteger](n T) float64 {
	// Calculate and return the variance
	return (Square(n) - float64(1)) / float64(12)
}
