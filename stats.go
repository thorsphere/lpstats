// Package lpstats provides mathematical and statistical utilities for Go.
//
// This package offers a collection of generic functions for performing common
// mathematical operations and statistical calculations on numeric types
// (integers and floats). Type constraints are defined in constraints.go.
//
// Basic Math:
//   - Abs: Absolute value
//   - Sign: Sign determination
//   - Square: Squaring a number
//
// Aggregate Operations:
//   - Sum: Sum of elements (returns error if slice is empty)
//   - ArithmeticMean: Arithmetic mean of a discrete value array (returns error if slice is empty)
//   - ExpectedValueU: Expected value for a uniform distribution in interval [a,b]
//   - Variance: Population variance of a discrete value array (returns error if slice is empty)
//   - VarianceU: Variance of a uniform distribution in interval [a,b]
//   - VarianceN: Variance of an n-sided die
//
// Comparison:
//   - EqualS: Equality check for slices of signed integers
//   - NearEqual: Near equality for floating-point types
//   - NearEqualFloatPtr: Near equality check for pointers to floating-point values
//   - EqualStrPtr: Nil-safe equality check for pointers to strings
//
// Pointer & String Utilities:
//   - FmtFloatPtr: Format a pointer to a number as a string with a specified precision (clamped)
//   - PtrFloat: Returns a pointer to a number
//   - CopyFloatPtr: Returns a copy of a pointer to a number (nil if nil)
//   - PtrStr: Returns a pointer to a string
//   - CopyStrPtr: Returns a copy of a pointer to a string (nil if nil)
//
// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package lpstats