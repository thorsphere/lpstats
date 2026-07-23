# lpstats

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorsphere/lpstats)](https://pkg.go.dev/mod/github.com/thorsphere/lpstats)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorsphere/lpstats)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorsphere/lpstats)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorsphere/lpstats)
[![CodeFactor](https://www.codefactor.io/repository/github/thorsphere/lpstats/badge)](https://www.codefactor.io/repository/github/thorsphere/lpstats)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorsphere/lpstats)

A lightweight Go package for simple mathematical and statistical calculations. It provides generic functions for computing mean, variance, and other common statistical operations with support for both integer and floating-point types.

## Features

- **Simple API**: No configuration required, just function calls
- **Type Flexible**: Support for integer and floating-point types via Go generics
- **Tested**: Comprehensive unit test coverage
- **Minimal Dependencies**: Only depends on the [Go Standard Library](https://pkg.go.dev/std) and [tserr](https://github.com/thorsphere/tserr) for error handling

## Installation

Install the package using `go get`:

```bash
go get github.com/thorsphere/lpstats
```

Import in your Go code:

```go
import "github.com/thorsphere/lpstats"
```

## Requirements

- Go 1.26 or higher (uses generic types)

## Type Constraints

Most functions use Go generics with type constraints defined in [constraints.go](constraints.go). The main constraints are:

- **`Number`**: All number types (int, uint, float, etc.)  
  Example: `func Square[T Number](x T) float64`

- **`Signed`**: Signed number types (int, float)  
  Example: `func Sign[T Signed](a T) T`

- **`Integer`**: Signed and unsigned integer types  
  Example: `func Abs[T Number](a T) T`

- **`Sinteger`**: Signed integer types (int, int8, int16, etc.)  
  Example: `func EqualS[T Sinteger](x, y []T) error`

- **`Uinteger`**: Unsigned integer types (uint, uint8, uint16, etc.)  
  Example: `func VarianceN[T Uinteger](n T) float64`

- **`Floating`**: Floating-point types (float32, float64)  
  Example: `func NearEqual[T Floating](a, b, maxDiff T) bool`

## API Reference

### Math

- **Absolute value**: `Abs()` - returns the absolute value of a number
- **Sign**: `Sign()` - returns the sign of a number
- **Square**: `Square()` - returns the square of a number
- **Sum**: `Sum()` - calculates the sum of values

### Statistics

- **Arithmetic mean**: `ArithmeticMean()` - calculates the arithmetic mean of a discrete value array
- **Expected value (uniform)**: `ExpectedValueU()` - calculates the expected value for a uniform distribution in interval [a,b]
- **Variance**: `Variance()` - calculates the population variance of a discrete value array
- **Variance (uniform)**: `VarianceU()` - calculates the variance of a uniform distribution in interval [a,b]
- **N-sided die variance**: `VarianceN()` - calculates the variance of an n-sided die

### Comparison

- **Equal for slices**: `EqualS()` - equality check for slices of signed integers
- **Near equal**: `NearEqual()` - approximate equality for floating-point types
- **Equal for float pointers**: `NearEqualFloatPtr()` - approximate equality check for pointers to float values
- **Equal for string pointers**: `EqualStrPtr()` - nil-safe equality check for pointers to strings

### Pointer & String Utilities

- **Pointer to number**: `PtrFloat()` - returns a pointer to a number
- **Copy number pointer**: `CopyFloatPtr()` - returns a copy of a pointer to a number; nil safe
- **Format float pointer**: `FmtFloatPtr()` - formats a pointer to a number as a string with a specified number of decimal places (clamped); returns `"nil"` for nil pointers
- **Pointer to string**: `PtrStr()` - returns a pointer to a string
- **Copy string pointer**: `CopyStrPtr()` - returns a copy of a pointer to a string; nil safe

> **Note:** Functions like `ArithmeticMean()`, `Sum()`, and `Variance()` return an error when the input slice is empty.


## Quick Start

### Example: Calculate Mean and Variance

```go
package main

import (
	"fmt"

	"github.com/thorsphere/lpstats"
)

func main() {
	// Work with different numeric types
	intSlice := []int{1, 2, 3, 4, 5, 6}
	floatSlice := []float64{1.1, 2.1, 3.1, 4.1, 5.1, 6.1}
	dieSize := uint(6)

	// Calculate arithmetic mean
	intMean, _ := lpstats.ArithmeticMean(intSlice)
	floatMean, _ := lpstats.ArithmeticMean(floatSlice)

	// Calculate population variance
	intVariance, _ := lpstats.Variance(intSlice)
	floatVariance, _ := lpstats.Variance(floatSlice)

	fmt.Printf("Integer slice - Mean: %.2f, Variance: %.2f\n", intMean, intVariance)
	fmt.Printf("Float slice   - Mean: %.2f, Variance: %.2f\n", floatMean, floatVariance)

	// Variance for uniform distribution (e.g., dice)
	diceVariance := lpstats.VarianceN(dieSize)
	fmt.Printf("Variance of a %d-sided die: %.4f\n", dieSize, diceVariance)

	// Expected value and variance for a uniform distribution in [1, 6]
	fmt.Printf("Expected value [1,6]: %.2f\n", lpstats.ExpectedValueU(1, 6))
	fmt.Printf("Variance [1,6]: %.2f\n", lpstats.VarianceU(1, 6))
}
```

Run in [Go Playground](https://go.dev/play/p/WuhNz-TENCw)

## Known Limitations

- Calculations use floating-point arithmetic and are not suitable for arbitrary precision fixed-point decimal arithmetic
- For very large datasets, consider memory and performance implications

## Documentation & Resources

- [Go Package Documentation](https://pkg.go.dev/github.com/thorsphere/lpstats) — Complete API reference
- [Open Source Insights](https://deps.dev/go/github.com%2Fthorsphere%2Flpstats) — Dependency analysis

## ⚖️ License & Commercial Usage

Copyright (c) 2023-2026 thorsphere. All rights reserved.

This project is licensed under the **Functional Source License v1.1 (FSL-1.1-ALv2)**. 

* The use, modification, and redistribution of this Go package is completely free for private, educational, non-commercial, and internal purposes. 
* If you are a company or institution looking to use this package in a commercial product, service, or business environment, you must secure a commercial license.
* Each version of this software automatically converts to the fully open-source Apache License, Version 2.0 on the second anniversary of its release.

For full details, please see the [LICENSE](LICENSE) file.

### 💼 Commercial Licensing & Inquiries

To purchase a commercial license or discuss support options, please reach out directly:

* 📩 **Contact:** business at thorsphere dot com
* 💬 **Response Time:** Usually within a couple of business days.

*Please include your company name and a brief overview of your use case so I can provide the right licensing details.*
