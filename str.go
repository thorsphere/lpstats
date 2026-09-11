// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package lpstats

import (
	"fmt"
	"strings"

	"github.com/thorsphere/tserr"
)

// PtrStr returns a pointer to a string.
func PtrStr(a string) *string {
	return &a
}

// CopyStrPtr returns a copy of a pointer to a string.
// It returns nil if p is nil.
func CopyStrPtr(p *string) *string {
	// Return nil if p is nil
	if p == nil {
		return nil
	}
	// Create a new value of type string and copy the value pointed to by p into it
	v := *p
	// Return a pointer to the new value
	return &v
}

// EqualStrPtr returns whether two pointers to strings are equal.
// It returns true if both pointers are nil, or if both pointers are non-nil and the values they point to are equal.
// It returns false if one pointer is nil and the other is not, or if both pointers are non-nil and the values
// they point to are not equal.
func EqualStrPtr(a, b *string) bool {
	// If both pointers are nil, they are considered equal
	if a == nil && b == nil {
		return true
	}
	// If one pointer is nil and the other is not, they are not equal
	if a == nil || b == nil {
		return false
	}
	// If both pointers are non-nil, compare the values they point to
	return *a == *b
}

// CountArgs returns the number of arguments that must be passed to a function
// of the fmt package (e.g. Sprintf) to fill all verbs in the format string s.
// A literal %% is not counted. Flags, width and precision (e.g. %-08.3f) belong
// to one verb. A width or precision of * consumes an additional argument.
// Explicit argument indexes of the form [n] are taken into account; CountArgs
// returns the highest one-based argument index referenced by s. It returns an
// error if s contains a malformed verb, i.e. a verb that fmt would render as
// %!(NOVERB) or %!(BADINDEX).
func CountArgs(s string) (int, error) {
	var (
		n          int  // highest one-based argument index referenced so far
		argNum     int  // zero-based index of the next argument, as tracked by fmt
		afterIndex bool // the previous item in the format was an index like [3]
	)
	// invalid returns an InvalidFormat error for format string s.
	// The detail describes the format problem; format may contain verbs
	// filled by the optional arguments a.
	invalid := func(format string, a ...any) error {
		return tserr.InvalidFormat(&tserr.InvalidFormatArgs{
			F:      "format string",
			Value:  s,
			Detail: fmt.Sprintf(format, a...),
		})
	}
	// index parses an explicit argument index [n] starting at s[i] == '['.
	// It returns the one-based index and the position of the byte after ']'.
	index := func(i int) (int, int, error) {
		// j is the position of the byte after '['. v is the digit value.
		j, v := i+1, 0
		// Consume digits from the argument index in string s.
		for j < len(s) && s[j] >= '0' && s[j] <= '9' {
			// Calculate the digit value.
			v = v*10 + int(s[j]-'0')
			// Advance to the next byte.
			j++
		}
		// Check for malformed argument index.
		switch {
		// malformed, because j == i+1 (i.e. no digits), j >= len(s) (i.e. no ']'),
		// or s[j] != ']' (i.e. not the end of the argument index).
		case j == i+1 || j >= len(s) || s[j] != ']':
			return 0, 0, invalid("malformed argument index at offset %d", i)
		// malformed, because v == 0 (i.e. the argument index is 0).
		case v == 0:
			return 0, 0, invalid("argument index 0 at offset %d", i)
		}
		// Return the argument index and the position of the byte after ']'.
		return v, j + 1, nil
	}
	// Iterate over the format string s.
	for i := 0; i < len(s); i++ {
		// Skip non-% characters.
		if s[i] != '%' {
			continue
		}
		// Examine the byte after '%'
		i++
		// If i >= len(s), there is a trailing % without a verb at the end of the format string.
		if i >= len(s) {
			return 0, invalid("trailing '%%' without a verb at offset %d", i-1)
		}
		// A literal %% consumes no argument.
		if s[i] == '%' {
			continue
		}
		// Skip flags.
		for i < len(s) && strings.IndexByte("+-# 0", s[i]) >= 0 {
			i++
		}
		// Explicit argument index before the width, e.g. %[2]d or %[1]*d.
		afterIndex = false // Is set to true if the item has an argument index.
		// If the verb starts with an explicit argument index, parse it.
		if i < len(s) && s[i] == '[' {
			// Parse the argument index.
			v, j, e := index(i)
			// Return an error if the argument index is malformed.
			if e != nil {
				return 0, e
			}
			// The index itself consumes no argument and does not update n;
			// the referenced argument is consumed by the following '*' or by
			// the verb itself, e.g. in %[3]d the verb d consumes the third
			// argument and updates n there.
			// argNum is the zero-based index of the parsed argument, as
			// tracked by fmt.
			// afterIndex is true because the item has an argument index.
			// i is set to the position of the byte after the argument index.
			argNum, afterIndex, i = v-1, true, j
		}
		// Width; a width of * consumes an additional argument.
		if i < len(s) && s[i] == '*' {
			// Update the highest argument index referenced so far.
			n = max(n, argNum+1)
			// Increment argnum.
			argNum++
			// Skip the '*'.
			i++
			// Set afterIndex to false because the item consumes an additional argument.
			afterIndex = false
		} else {
			// Otherwise, skip digits; w is the position of the first
			// width digit, if a width is present.
			w := i
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				i++
			}
			// A width after an argument index, e.g. "%[3]2d", is
			// rendered as %!(BADINDEX) by fmt. A missing width, e.g.
			// "%[3]d", is valid.
			if afterIndex && i > w {
				return 0, invalid("width after argument index at offset %d", w)
			}
		}
		// Precision; a precision of * consumes an additional argument.
		// The i+1 < len(s) check mirrors fmt: in "%." the '.' is the verb
		// itself, not the start of a precision.
		if i+1 < len(s) && s[i] == '.' {
			// Skip the '.'.
			i++
			// If the previous item has an argument index, return an error.
			if afterIndex {
				// "%[3].2d": fmt renders %!(BADINDEX).
				return 0, invalid("precision after argument index at offset %d", i)
			}
			// The precision may have its own index, e.g. %.[2]*f.
			if i < len(s) && s[i] == '[' {
				// Parse the argument index.
				v, j, e := index(i)
				// Return an error if the argument index is malformed.
				if e != nil {
					return 0, e
				}
				// The index itself consumes no argument and does not update n;
				// the referenced argument is consumed by the following '*' or by
				// the verb itself, e.g. in %[3]d the verb d consumes the third
				// argument and updates n there.
				// argNum is the zero-based index of the parsed argument, as
				// tracked by fmt.
				// afterIndex is true because the item has an argument index.
				// i is set to the position of the byte after the argument index.
				argNum, afterIndex, i = v-1, true, j
			}
			// A precision of * consumes an additional argument,
			// e.g. %.*f or %.[2]*f.
			if i < len(s) && s[i] == '*' {
				// Update the highest argument index referenced so far.
				n = max(n, argNum+1)
				// Increment argnum.
				argNum++
				// Skip the '*'.
				i++
				// Set afterIndex to false because the item consumes an additional argument.
				afterIndex = false
			} else { // Otherwise, skip digits.
				for i < len(s) && s[i] >= '0' && s[i] <= '9' {
					i++
				}
			}
		}
		// The verb may have its own explicit argument index [n], which
		// appears after the flags, the width and the precision, e.g.
		// %5[3]d, %*[2]d or %.*[3]f. In %5[3]d, the width is 5 and the
		// verb d formats the third argument. In %*[2]d, the width is
		// taken from the first argument and the verb d formats the
		// second argument. In %.*[3]f, the precision is taken from the
		// first argument and the verb f formats the third argument.
		// The index is only parsed if the previous item (the width or
		// the precision) does not already have an index, because fmt
		// parses at most one index per position. If the previous item
		// has an index, the '[' is the verb character itself: fmt
		// renders %[1][2]d as %![(int=1)2]d, i.e. the verb '[' consumes
		// the first argument and "2]d" is literal text.
		// If the format string ends before the verb, e.g. "%5", the
		// check for a missing verb below returns an error.
		if !afterIndex && i < len(s) && s[i] == '[' {
			// Parse the argument index.
			v, j, e := index(i)
			// Return an error if the argument index is malformed.
			if e != nil {
				return 0, e
			}
			// argNum is the index of the parsed argument, as tracked by fmt.
			// i is set to the position of the byte after the argument index.
			argNum, i = v-1, j
		}
		// If i >= len(s), the format string ended after flags, argument
		// index, width or precision without a verb character.
		if i >= len(s) {
			// "%5" or "%-": fmt renders %!(NOVERB).
			return 0, invalid("missing verb at offset %d", i)
		}
		// The verb character itself; it is skipped by the loop's i++.
		// A verb of % consumes no argument, e.g. the second % in "%5%".
		if s[i] != '%' {
			// Update the highest argument index referenced so far.
			n = max(n, argNum+1)
			// Increment argnum.
			argNum++
		}
	}
	return n, nil
}
