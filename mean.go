// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Import package tserr
import "github.com/thorsphere/tserr" // tserr

// ArithmeticMean returns the arithmetic mean of number array x as float64. It returns
// zero and an error if x is empty.
func ArithmeticMean[T Number](x []T) (float64, error) {
	// Return zero and an error if x is empty
	if len(x) <= 0 {
		return float64(0), tserr.Empty("x")
	}
	// calculate and return the arithmetic mean
	var (
		s float64
		e error
	)
	if s, e = Sum(x); e != nil {
		return float64(0), tserr.Op(&tserr.OpArgs{Op: "Sum", Fn: "x", Err: e})
	}
	return s / float64(len(x)), nil
}

// ExpectedValueU returns the expected value for the uniform distribution in interval [a,b] as float64.
func ExpectedValueU[T Number](a, b T) float64 {
	return (float64(a) + float64(b)) / float64(2)
}
