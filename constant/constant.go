package constant

const (
	// MaxUint is the maximum value that could be represented with a uint type.
	MaxUint = ^uint(0)

	// MaxInt is the maximum value that could be represented with an int type.
	MaxInt = int(MaxUint >> 1)

	// Infinity is the value that represents an unreachable distance.
	Infinity = -1
)
