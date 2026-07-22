// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package lpstats

// Number constraints types to all number types
type Number interface {
	Integer | Floating
}

// Signed constraints types to signed number types
type Signed interface {
	Floating | Sinteger
}

// Integer constraints types to signed and unsigned integer types
type Integer interface {
	Sinteger | Uinteger
}

// Sinteger constraints types to signed integer types
type Sinteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Uinteger constraints types to unsigned integer types
type Uinteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Floating constraints types to floating types
type Floating interface {
	~float32 | ~float64
}
