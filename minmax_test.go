// Copyright (c) 2023-2026 thorsphere
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats_test

// Import standard library packages and the lpstats package for testing the Min and Max functions.
import (
	"testing" // Import the testing package for writing test cases.

	"github.com/thorsphere/lpstats" // Import the lpstats package to access the Min and Max functions.
)

// TestMaxf1 tests the result of Max for two float64 arguments. If the result of Max does not
// equal the maximum value, the test fails.
func TestMaxf1(t *testing.T) {
	testf2(t, 1.0, 2.0, 2.0, lpstats.Max[float64])
}

// TestMaxf2 tests the result of Max for two float64 arguments. If the result of Max does not
// equal the maximum value, the test fails.
func TestMaxf2(t *testing.T) {
	testf2(t, 2.0, 1.0, 2.0, lpstats.Max[float64])
}

// TestMaxf3 tests the result of Max for two equal float64 arguments. If the result of Max does not
// equal the same value, the test fails.
func TestMaxf3(t *testing.T) {
	testf2(t, 2.0, 2.0, 2.0, lpstats.Max[float64])
}

// TestMaxi1 tests the result of Max for two int64 arguments. If the result of Max does not
// equal the maximum value, the test fails.
func TestMaxi1(t *testing.T) {
	testi2(t, 1, 2, 2, lpstats.Max[int64])
}

// TestMaxi2 tests the result of Max for two int64 arguments. If the result of Max does not
// equal the maximum value, the test fails.
func TestMaxi2(t *testing.T) {
	testi2(t, 2, 1, 2, lpstats.Max[int64])
}

// TestMaxi3 tests the result of Max for two equal int64 arguments. If the result of Max does not
// equal the same value, the test fails.
func TestMaxi3(t *testing.T) {
	testi2(t, 2, 2, 2, lpstats.Max[int64])
}

// TestMinf1 tests the result of Min for two float64 arguments. If the result of Min does not
// equal the minimum value, the test fails.
func TestMinf1(t *testing.T) {
	testf2(t, 1.0, 2.0, 1.0, lpstats.Min[float64])
}

// TestMinf2 tests the result of Min for two float64 arguments. If the result of Min does not
// equal the minimum value, the test fails.
func TestMinf2(t *testing.T) {
	testf2(t, 2.0, 1.0, 1.0, lpstats.Min[float64])
}

// TestMinf3 tests the result of Min for two equal float64 arguments. If the result of Min does not
// equal the same value, the test fails.
func TestMinf3(t *testing.T) {
	testf2(t, 2.0, 2.0, 2.0, lpstats.Min[float64])
}

// TestMini1 tests the result of Min for two int64 arguments. If the result of Min does not
// equal the minimum value, the test fails.
func TestMini1(t *testing.T) {
	testi2(t, 1, 2, 1, lpstats.Min[int64])
}

// TestMini2 tests the result of Min for two int64 arguments. If the result of Min does not
// equal the minimum value, the test fails.
func TestMini2(t *testing.T) {
	testi2(t, 2, 1, 1, lpstats.Min[int64])
}

// TestMini3 tests the result of Min for two equal int64 arguments. If the result of Min does not
// equal the same value, the test fails.
func TestMini3(t *testing.T) {
	testi2(t, 2, 2, 2, lpstats.Min[int64])
}
