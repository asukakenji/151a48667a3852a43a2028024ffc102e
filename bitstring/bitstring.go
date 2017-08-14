package bitstring

/*
Ones returns a bit string with n rightmost contiguous 1-bits.

Example

Here are some examples:

	Ones(0) // 0 (00000000)
	Ones(1) // 1 (00000001)
	Ones(2) // 3 (00000011)
	Ones(3) // 7 (00000111)
*/
func Ones(n uint) uint64 {
	return (1 << n) - 1
}

/*
IsBitSetAtIndex checks whether the specified bit in x is set (equals 1) or not (equals 0).

The bits are indexed from right to left. The first (least significant) bit has an index of 0.

Example

Here are some examples:

	IsBitSetAtIndex(9, 0) // true
	IsBitSetAtIndex(9, 1) // false
	IsBitSetAtIndex(9, 2) // false
	IsBitSetAtIndex(9, 3) // true
*/
func IsBitSetAtIndex(x uint64, index uint) bool {
	return x&(1<<index) != 0
}

/*
SetBitAtIndex sets the specified bit in x to 1.

The bits are indexed from right to left. The first (least significant) bit has an index of 0.

Example

Here are some examples:

	SetBitAtIndex(9, 0) // 9 (00001001) ->  9 (00001001)
	SetBitAtIndex(9, 1) // 9 (00001001) -> 11 (00001011)
	SetBitAtIndex(9, 2) // 9 (00001001) -> 13 (00001101)
	SetBitAtIndex(9, 3) // 9 (00001001) ->  9 (00001001)

Reference

Hacker's Delight: Section 2-1
*/
func SetBitAtIndex(x uint64, index uint) uint64 {
	return x | (1 << index)
}

/*
ResetBitAtIndex resets the specified bit in x to 0.

The bits are indexed from right to left. The first (least significant) bit has an index of 0.

Example

Here are some examples:

	ResetBitAtIndex(9, 0) // 9 (00001001) -> 8 (00001000)
	ResetBitAtIndex(9, 1) // 9 (00001001) -> 9 (00001001)
	ResetBitAtIndex(9, 2) // 9 (00001001) -> 9 (00001001)
	ResetBitAtIndex(9, 3) // 9 (00001001) -> 1 (00000001)

Reference

Hacker's Delight: Section 2-1
*/
func ResetBitAtIndex(x uint64, index uint) uint64 {
	return x &^ (1 << index)
}

/*
SetRightmostZero sets the rightmost 0-bit in x to 1.

Example

Here are some examples:

	SetRightmostZero(3) // 3 (00000011) -> 7 (00000111)
	SetRightmostZero(5) // 5 (00000101) -> 7 (00000111)
	SetRightmostZero(6) // 6 (00000110) -> 7 (00000111)

Reference

Hacker's Delight: Section 2-1
*/
func SetRightmostZero(x uint64) uint64 {
	return x | (x + 1)
}

/*
ResetRightmostOne resets the rightmost 1-bit in x to 0.

Example

Here are some examples:

	ResetRightmostOne(3) // 3 (00000011) -> 2 (00000010)
	ResetRightmostOne(5) // 5 (00000101) -> 4 (00000100)
	ResetRightmostOne(6) // 6 (00000110) -> 4 (00000100)

Reference

Hacker's Delight: Section 2-1
*/
func ResetRightmostOne(x uint64) uint64 {
	return x & (x - 1)
}

/*
NextNumberWithSameBitCount returns the next number greater than x
having the same number of 1-bit as x.

If x is 0, it panics.

Example

Here are some examples:

	NextNumberWithSameBitCount(3) // 3 (00000011) ->  5 (00000101)
	NextNumberWithSameBitCount(5) // 5 (00000101) ->  6 (00000110)
	NextNumberWithSameBitCount(6) // 6 (00000110) ->  9 (00001001)
	NextNumberWithSameBitCount(9) // 9 (00001001) -> 10 (00001010)

Reference

Hacker's Delight: Figure 2-1
*/
func NextNumberWithSameBitCount(x uint64) uint64 {
	smallest := x & -x
	ripple := x + smallest
	ones := x ^ ripple
	ones = (ones >> 2) / smallest
	return ripple | ones
}

/*
NumberOfTrailingZeros returns the number of trailing 0-bits in x.

Reference

Hacker's Delight: Figure 5-14 (ntz4a)
*/
func NumberOfTrailingZeros(x uint64) uint {
	if x == 0 {
		return 64
	}
	n := uint(63)
	var a uint

	b := uint(uint32(x))
	if b != 0 {
		n = n - 32
		a = b
	} else {
		a = uint(uint32(x >> 32))
	}

	b = a << 16
	if b != 0 {
		n = n - 16
		a = b
	}

	b = a << 8
	if b != 0 {
		n = n - 8
		a = b
	}

	b = a << 4
	if b != 0 {
		n = n - 4
		a = b
	}

	b = a << 2
	if b != 0 {
		n = n - 2
		a = b
	}

	return n - ((a << 1) >> 31)
}

var tableForNumberOfTrailingZerosForPowerOfTwo = []uint{
	63, 0, 58, 1, 59, 47, 53, 2,
	60, 39, 48, 27, 54, 33, 42, 3,
	61, 51, 37, 40, 49, 18, 28, 20,
	55, 30, 34, 11, 43, 14, 22, 4,
	62, 57, 46, 52, 38, 26, 32, 41,
	50, 36, 17, 19, 29, 10, 13, 21,
	56, 45, 25, 31, 35, 16, 9, 12,
	44, 24, 15, 8, 23, 7, 6, 5,
}

/*
Reference

De Bruijn Sequence
https://stackoverflow.com/a/32178577/142239
*/
func NumberOfTrailingZerosForPowerOfTwo(x uint64) uint {
	return tableForNumberOfTrailingZerosForPowerOfTwo[uint((x*0x7EDD5E59A4E28C2)>>58)]
}

/*
InsertZero inserts a 0-bit at the specified position in x.

The bits are indexed from right to left. The first (least significant) bit has an index of 0.

Example

Here are some examples:

	InsertZero(6, 0) // 6 (00000110) -> 12 (00001100)
	InsertZero(6, 1) // 6 (00000110) -> 12 (00001100)
	InsertZero(6, 2) // 6 (00000110) -> 10 (00001010)
	InsertZero(6, 3) // 6 (00000110) ->  6 (00000110)
*/
func InsertZero(x uint64, index uint) uint64 {
	// TODO: Write another implementation using shift first and benchmark both.
	maskR := (uint64(1) << index) - 1
	maskL := ^maskR
	return ((x & maskL) << 1) | (x & maskR)
}

/*
ToString returns the binary representation of x as string.

Example

Here are some examples:

	ToString(3) // "11"
	ToString(5) // "101"
	ToString(6) // "110"
	ToString(9) // "1001"
*/
func ToString(x uint64) string {
	s := ""
	for x != 0 {
		if x&1 == 0 {
			s = "0" + s
		} else {
			s = "1" + s
		}
		x >>= 1
	}
	return s
}
