package bitstring

// index starts with 0.
func IsBitSetAtIndex(x uint64, index uint) bool {
	return x&(1<<index) != 0
}

// Hacker's Delight: Section 2-1
func SetBitAtIndex(x uint64, index uint) uint64 {
	return x | (1 << index)
}

// Hacker's Delight: Section 2-1
func ResetBitAtIndex(x uint64, index uint) uint64 {
	return x &^ (1 << index)
}

// Hacker's Delight: Section 2-1
func SetRightmostZero(x uint64) uint64 {
	return x | (x + 1)
}

// Hacker's Delight: Section 2-1
func UnsetRightmostOne(x uint64) uint64 {
	return x & (x - 1)
}

// Hacker's Delight: Figure 2-1
func NextValueWithSameBitCount(x uint64) uint64 {
	smallest := x & -x
	ripple := x + smallest
	ones := x ^ ripple
	ones = (ones >> 2) / smallest
	return ripple | ones
}

// Hacker's Delight: Figure 5-14 (ntz4a)
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

// De Bruijn Sequence
// https://stackoverflow.com/a/32178577/142239
func NumberOfTrailingZerosForPowerOfTwo(x uint64) uint {
	return tableForNumberOfTrailingZerosForPowerOfTwo[uint((x*0x7EDD5E59A4E28C2)>>58)]
}

func InsertZero(x uint64, index uint) uint64 {
	maskR := (uint64(1) << index) - 1
	maskL := ^maskR
	return ((x & maskL) << 1) | (x & maskR)
}

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
