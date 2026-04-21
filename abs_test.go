// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats_test

// Import package testing
import (
	"testing" // testing

	"github.com/thorsphere/lpstats"
)

// TestAbsfp tests the result of Abs for a positive float64. If the result of Abs does not
// equal the positive float64, the test fails.
func TestAbsfp(t *testing.T) {
	var a float64 = 10
	testf(t, a /*a*/, a /*w*/, lpstats.Abs[float64])
}

// TestAbsfz tests the result of Abs for the zero float64. If the result of Abs does not
// equal 0, the test fails.
func TestAbsfz(t *testing.T) {
	var a float64 = 0
	testf(t, a /*a*/, a /*w*/, lpstats.Abs[float64])
}

// TestAbsfn tests the result of Abs for a negative float64. If the result of Abs does not
// equal the absolute value, the test fails.
func TestAbsfn(t *testing.T) {
	var a float64 = -10
	testf(t, a /*a*/, -a /*w*/, lpstats.Abs[float64])
}

// TestAbsip tests the result of Abs for a positive integer. If the result of Abs does not
// equal the positive integer, the test fails.
func TestAbsip(t *testing.T) {
	var a int64 = 10
	testi(t, a /*a*/, a /*w*/, lpstats.Abs[int64])
}

// TestAbsiz tests the result of Abs for the zero integer. If the result of Abs does not
// equal 0, the test fails.
func TestAbsiz(t *testing.T) {
	var a int64 = 0
	testi(t, a /*a*/, a /*w*/, lpstats.Abs[int64])
}

// TestAbsin tests the result of Abs for a negative integer. If the result of Abs does not
// equal the absolute value, the test fails.
func TestAbsin(t *testing.T) {
	var a int64 = -10
	testi(t, a /*a*/, -a /*w*/, lpstats.Abs[int64])
}
