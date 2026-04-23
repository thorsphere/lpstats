# lpstats

A lightweight Go package for simple mathematical and statistical calculations. Provides generic functions for computing mean, variance, and other common statistical operations with support for both integer and floating-point types.

[![Go Report Card](https://goreportcard.com/badge/github.com/thorsphere/lpstats)](https://goreportcard.com/report/github.com/thorsphere/lpstats)
[![CodeFactor](https://www.codefactor.io/repository/github/thorsphere/lpstats/badge)](https://www.codefactor.io/repository/github/thorsphere/lpstats)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorsphere/lpstats)

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorsphere/lpstats)](https://pkg.go.dev/mod/github.com/thorsphere/lpstats)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorsphere/lpstats)
![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/thorsphere/lpstats)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorsphere/lpstats)
![GitHub last commit](https://img.shields.io/github/last-commit/thorsphere/lpstats)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/thorsphere/lpstats)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/thorsphere/lpstats)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorsphere/lpstats)
![GitHub](https://img.shields.io/github/license/thorsphere/lpstats)

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

- **`Sinteger`**: Signed integer types (int, int8, int16, etc.)  
  Example: `func EqualS[T Sinteger](x, y T) error`

- **`Uinteger`**: Unsigned integer types (uint, uint8, uint16, etc.)  
  Example: `func VarianceN[T Uinteger](n T) float64`

## API Reference

### Supported Functions

- **Absolute value**: `Abs()` - returns the absolute value of a number
- **Arithmetic mean**: `ArithmeticMean()` - calculates the mean of a discrete value array
- **Equal for slices**: `EqualS()` - equality check for slices of signed integers
- **Near equal**: `NearEqual()` - approximate equality for float types
- **Sign**: `Sign()` - returns the sign of a number
- **Square**: `Square()` - returns the square of a number
- **Sum**: `Sum()` - calculates the sum of values
- **Mean**: `Mean()` - calculates the mean of a discrete value array
- **Variance**: `Variance()` - calculates the variance of a discrete value array
- **Expected value**: `VarianceN()` - calculates variance for a uniform distribution
- **Format float pointer**: `FmtFloatPtr()` - formats a pointer to a number as a string with one decimal place
- **Equal for float pointers**: `NearEqualFloatPtr()` - approximate equality check for pointers to float values

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

	// Calculate variance
	intVariance, _ := lpstats.Variance(intSlice)
	floatVariance, _ := lpstats.Variance(floatSlice)

	fmt.Printf("Integer slice - Mean: %.2f, Variance: %.2f\n", intMean, intVariance)
	fmt.Printf("Float slice - Mean: %.2f, Variance: %.2f\n", floatMean, floatVariance)

	// Variance for uniform distribution (e.g., dice)
	diceVariance := lpstats.VarianceN(dieSize)
	fmt.Printf("Variance of a %d-sided die: %.4f\n", dieSize, diceVariance)
}
```

## Known Limitations

- Calculations use floating-point arithmetic and are not suitable for arbitrary precision fixed-point decimal arithmetic
- For very large datasets, consider memory and performance implications

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Security

For security issues, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the GNU Affero General Public License v3.0. See [LICENSE](LICENSE) for details.

## Resources

- [API Documentation (pkg.go.dev)](https://pkg.go.dev/github.com/thorsphere/lpstats)
- [Go Report Card](https://goreportcard.com/report/github.com/thorsphere/lpstats)
- [Code Quality Analysis](https://www.codefactor.io/repository/github/thorsphere/lpstats)
- [Dependency Analysis (Open Source Insights)](https://deps.dev/go/github.com%2Fthorsphere%2Flpstats)
