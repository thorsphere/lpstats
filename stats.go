// Package lpstats provides mathematical and statistical utilities for Go.
//
// This package offers a collection of functions for performing common mathematical
// operations and statistical calculations on numeric types (integers and floats).
//
// Core Functions:
//
// Basic Math:
//   - Abs: Absolute value
//   - Sign: Sign determination
//   - Square: Squaring a number
//
// Aggregate Operations:
//   - Sum: Sum of elements
//   - Mean: Arithmetic mean of a discrete value array
//   - Variance: Variance of a discrete value array
//   - VarianceU: Variance of a uniform distribution
//   - VarianceN: Variance of an n-sided die
//
// Utility Functions:
//   - Equal: Equality check for slices of signed integers
//   - NearEqual: Near equality for float types
//   - ExpectedValue: Expected value for a uniform distribution
//   - FmtFloatPtr: Format a pointer to a floating-point number as a string
//   - NearEqualFloatPtr: Near equality check for pointers to float types
//
// Most functions support both integer and floating-point numeric types.
//
// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats
