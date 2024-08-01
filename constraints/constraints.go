// Package constraints provides type constraints for Go generics.
// Modified from https://github.com/TheAlgorithms/Go/blob/master/constraints/constraints.go.
// The original code is licensed under the MIT License.
package constraints

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}

type Ordered interface {
	Integer | ~string | Float
}
