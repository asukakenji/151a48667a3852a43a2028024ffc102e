package bitstring

import "testing"

func TestOnes(t *testing.T) {
	cases := []struct {
		n        uint
		expected uint64
	}{
		{0, 0x00},
		{1, 0x01},
		{2, 0x03},
		{3, 0x07},
		{4, 0x0f},
		{5, 0x1f},
		{6, 0x3f},
		{7, 0x7f},
		{8, 0xff},
		{9, 0x01ff},
		{10, 0x03ff},
		{11, 0x07ff},
		{12, 0x0fff},
		{13, 0x1fff},
		{14, 0x3fff},
		{15, 0x7fff},
		{16, 0xffff},
		{17, 0x01ffff},
		{18, 0x03ffff},
		{19, 0x07ffff},
		{20, 0x0fffff},
		{21, 0x1fffff},
		{22, 0x3fffff},
		{23, 0x7fffff},
		{24, 0xffffff},
		{25, 0x01ffffff},
		{26, 0x03ffffff},
		{27, 0x07ffffff},
		{28, 0x0fffffff},
		{29, 0x1fffffff},
		{30, 0x3fffffff},
		{31, 0x7fffffff},
		{32, 0xffffffff},
		{33, 0x01ffffffff},
		{34, 0x03ffffffff},
		{35, 0x07ffffffff},
		{36, 0x0fffffffff},
		{37, 0x1fffffffff},
		{38, 0x3fffffffff},
		{39, 0x7fffffffff},
		{40, 0xffffffffff},
		{41, 0x01ffffffffff},
		{42, 0x03ffffffffff},
		{43, 0x07ffffffffff},
		{44, 0x0fffffffffff},
		{45, 0x1fffffffffff},
		{46, 0x3fffffffffff},
		{47, 0x7fffffffffff},
		{48, 0xffffffffffff},
		{49, 0x01ffffffffffff},
		{50, 0x03ffffffffffff},
		{51, 0x07ffffffffffff},
		{52, 0x0fffffffffffff},
		{53, 0x1fffffffffffff},
		{54, 0x3fffffffffffff},
		{55, 0x7fffffffffffff},
		{56, 0xffffffffffffff},
		{57, 0x01ffffffffffffff},
		{58, 0x03ffffffffffffff},
		{59, 0x07ffffffffffffff},
		{60, 0x0fffffffffffffff},
		{61, 0x1fffffffffffffff},
		{62, 0x3fffffffffffffff},
		{63, 0x7fffffffffffffff},
		{64, 0xffffffffffffffff},
	}
	for _, c := range cases {
		got := Ones(c.n)
		if got != c.expected {
			t.Errorf(
				"Ones(%d) = 0x%016x, expected 0x%016x",
				c.n, got, c.expected,
			)
		}
	}
}

func TestIsBitSetAtIndex(t *testing.T) {
	cases := []struct {
		x        uint64
		index    uint
		expected bool
	}{
		{0, 0, false},
		{0, 1, false},
		{0, 2, false},
		{1, 0, true},
		{1, 1, false},
		{1, 2, false},
		{2, 0, false},
		{2, 1, true},
		{2, 2, false},
		{3, 0, true},
		{3, 1, true},
		{3, 2, false},
	}
	for _, c := range cases {
		got := IsBitSetAtIndex(c.x, c.index)
		if got != c.expected {
			t.Errorf(
				"IsBitSetAtIndex(0x%016x, %d) = %t, expected %t",
				c.x, c.index, got, c.expected,
			)
		}
	}
}

func TestSetBitAtIndex(t *testing.T) {
	cases := []struct {
		x        uint64
		index    uint
		expected uint64
	}{
		{0, 0, 1},
		{0, 1, 2},
		{0, 2, 4},
		{1, 0, 1},
		{1, 1, 3},
		{1, 2, 5},
		{2, 0, 3},
		{2, 1, 2},
		{2, 2, 6},
		{3, 0, 3},
		{3, 1, 3},
		{3, 2, 7},
	}
	for _, c := range cases {
		got := SetBitAtIndex(c.x, c.index)
		if got != c.expected {
			t.Errorf(
				"SetBitAtIndex(0x%016x, %d) = 0x%016x, expected 0x%016x",
				c.x, c.index, got, c.expected,
			)
		}
	}
}

func TestResetBitAtIndex(t *testing.T) {
	cases := []struct {
		x        uint64
		index    uint
		expected uint64
	}{
		{0, 0, 0},
		{0, 1, 0},
		{0, 2, 0},
		{1, 0, 0},
		{1, 1, 1},
		{1, 2, 1},
		{2, 0, 2},
		{2, 1, 0},
		{2, 2, 2},
		{3, 0, 2},
		{3, 1, 1},
		{3, 2, 3},
	}
	for _, c := range cases {
		got := ResetBitAtIndex(c.x, c.index)
		if got != c.expected {
			t.Errorf(
				"ResetBitAtIndex(0x%016x, %d) = 0x%016x, expected 0x%016x",
				c.x, c.index, got, c.expected,
			)
		}
	}
}

func TestSetRightmostZero(t *testing.T) {
	cases := []struct {
		x        uint64
		expected uint64
	}{
		{0, 1},
		{1, 3},
		{2, 3},
		{3, 7},
		{4, 5},
		{5, 7},
		{6, 7},
		{7, 15},
		{8, 9},
		{9, 11},
		{10, 11},
		{11, 15},
		{12, 13},
		{13, 15},
		{14, 15},
		{15, 31},
		{16, 17},
	}
	for _, c := range cases {
		got := SetRightmostZero(c.x)
		if got != c.expected {
			t.Errorf(
				"SetRightmostZero(0x%016x) = 0x%016x, expected 0x%016x",
				c.x, got, c.expected,
			)
		}
	}
}

func TestResetRightmostOne(t *testing.T) {
	cases := []struct {
		x        uint64
		expected uint64
	}{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 2},
		{4, 0},
		{5, 4},
		{6, 4},
		{7, 6},
		{8, 0},
		{9, 8},
		{10, 8},
		{11, 10},
		{12, 8},
		{13, 12},
		{14, 12},
		{15, 14},
		{16, 0},
	}
	for _, c := range cases {
		got := ResetRightmostOne(c.x)
		if got != c.expected {
			t.Errorf(
				"ResetRightmostOne(0x%016x) = 0x%016x, expected 0x%016x",
				c.x, got, c.expected,
			)
		}
	}
}

func TestNextNumberWithSameBitCount(t *testing.T) {
	cases := []struct {
		x        uint64
		expected uint64
	}{
		{1, 2},
		{2, 4},
		{3, 5},
		{4, 8},
		{5, 6},
		{6, 9},
		{7, 11},
		{8, 16},
		{9, 10},
		{10, 12}, // a
		{11, 13}, // b
		{12, 17}, // c
		{13, 14}, // d
		{14, 19}, // e
		{15, 23}, // f
		{16, 32}, // 10
	}
	for _, c := range cases {
		got := NextNumberWithSameBitCount(c.x)
		if got != c.expected {
			t.Errorf(
				"NextNumberWithSameBitCount(0x%016x) = 0x%016x, expected 0x%016x",
				c.x, got, c.expected,
			)
		}
	}
}

func TestNumberOfTrailingZeros(t *testing.T) {
	cases := []struct {
		i        uint64
		expected uint
	}{
		{0x0000000000000000, 64},
		{0x0000000000000001, 0},
		{0x0000000000000002, 1},
		{0x0000000000000004, 2},
		{0x0000000000000008, 3},
		{0x0000000000000010, 4},
		{0x0000000000000020, 5},
		{0x0000000000000040, 6},
		{0x0000000000000080, 7},
		{0x0000000000000100, 8},
		{0x0000000000000200, 9},
		{0x0000000000000400, 10},
		{0x0000000000000800, 11},
		{0x0000000000001000, 12},
		{0x0000000000002000, 13},
		{0x0000000000004000, 14},
		{0x0000000000008000, 15},
		{0x0000000000010000, 16},
		{0x0000000000020000, 17},
		{0x0000000000040000, 18},
		{0x0000000000080000, 19},
		{0x0000000000100000, 20},
		{0x0000000000200000, 21},
		{0x0000000000400000, 22},
		{0x0000000000800000, 23},
		{0x0000000001000000, 24},
		{0x0000000002000000, 25},
		{0x0000000004000000, 26},
		{0x0000000008000000, 27},
		{0x0000000010000000, 28},
		{0x0000000020000000, 29},
		{0x0000000040000000, 30},
		{0x0000000080000000, 31},
		{0x0000000100000000, 32},
		{0x0000000200000000, 33},
		{0x0000000400000000, 34},
		{0x0000000800000000, 35},
		{0x0000001000000000, 36},
		{0x0000002000000000, 37},
		{0x0000004000000000, 38},
		{0x0000008000000000, 39},
		{0x0000010000000000, 40},
		{0x0000020000000000, 41},
		{0x0000040000000000, 42},
		{0x0000080000000000, 43},
		{0x0000100000000000, 44},
		{0x0000200000000000, 45},
		{0x0000400000000000, 46},
		{0x0000800000000000, 47},
		{0x0001000000000000, 48},
		{0x0002000000000000, 49},
		{0x0004000000000000, 50},
		{0x0008000000000000, 51},
		{0x0010000000000000, 52},
		{0x0020000000000000, 53},
		{0x0040000000000000, 54},
		{0x0080000000000000, 55},
		{0x0100000000000000, 56},
		{0x0200000000000000, 57},
		{0x0400000000000000, 58},
		{0x0800000000000000, 59},
		{0x1000000000000000, 60},
		{0x2000000000000000, 61},
		{0x4000000000000000, 62},
		{0x8000000000000000, 63},
	}
	for _, c := range cases {
		got := NumberOfTrailingZeros(c.i)
		if got != c.expected {
			t.Errorf(
				"NumberOfTrailingZeros(0x%016x) = %d, expected %d",
				c.i, got, c.expected,
			)
		}
	}
}

func TestNumberOfTrailingZerosForPowerOfTwo(t *testing.T) {
	cases := []struct {
		set      uint64
		expected uint
	}{
		{0x0000000000000000, 64},
		{0x0000000000000001, 0},
		{0x0000000000000002, 1},
		{0x0000000000000004, 2},
		{0x0000000000000008, 3},
		{0x0000000000000010, 4},
		{0x0000000000000020, 5},
		{0x0000000000000040, 6},
		{0x0000000000000080, 7},
		{0x0000000000000100, 8},
		{0x0000000000000200, 9},
		{0x0000000000000400, 10},
		{0x0000000000000800, 11},
		{0x0000000000001000, 12},
		{0x0000000000002000, 13},
		{0x0000000000004000, 14},
		{0x0000000000008000, 15},
		{0x0000000000010000, 16},
		{0x0000000000020000, 17},
		{0x0000000000040000, 18},
		{0x0000000000080000, 19},
		{0x0000000000100000, 20},
		{0x0000000000200000, 21},
		{0x0000000000400000, 22},
		{0x0000000000800000, 23},
		{0x0000000001000000, 24},
		{0x0000000002000000, 25},
		{0x0000000004000000, 26},
		{0x0000000008000000, 27},
		{0x0000000010000000, 28},
		{0x0000000020000000, 29},
		{0x0000000040000000, 30},
		{0x0000000080000000, 31},
		{0x0000000100000000, 32},
		{0x0000000200000000, 33},
		{0x0000000400000000, 34},
		{0x0000000800000000, 35},
		{0x0000001000000000, 36},
		{0x0000002000000000, 37},
		{0x0000004000000000, 38},
		{0x0000008000000000, 39},
		{0x0000010000000000, 40},
		{0x0000020000000000, 41},
		{0x0000040000000000, 42},
		{0x0000080000000000, 43},
		{0x0000100000000000, 44},
		{0x0000200000000000, 45},
		{0x0000400000000000, 46},
		{0x0000800000000000, 47},
		{0x0001000000000000, 48},
		{0x0002000000000000, 49},
		{0x0004000000000000, 50},
		{0x0008000000000000, 51},
		{0x0010000000000000, 52},
		{0x0020000000000000, 53},
		{0x0040000000000000, 54},
		{0x0080000000000000, 55},
		{0x0100000000000000, 56},
		{0x0200000000000000, 57},
		{0x0400000000000000, 58},
		{0x0800000000000000, 59},
		{0x1000000000000000, 60},
		{0x2000000000000000, 61},
		{0x4000000000000000, 62},
		{0x8000000000000000, 63},
	}
	for _, c := range cases {
		got := NumberOfTrailingZerosForPowerOfTwo(c.set)
		if got != c.expected {
			t.Errorf(
				"NumberOfTrailingZerosForPowerOfTwo(0x%016x) = %d, expected %d",
				c.set, got, c.expected,
			)
		}
	}
}

func TestInsertZero(t *testing.T) {
	cases := []struct {
		x        uint64
		index    uint
		expected uint64
	}{
		{0, 0, 0},
		{0, 1, 0},
		{0, 2, 0},
		{0, 3, 0},
		{1, 0, 2},
		{1, 1, 1},
		{1, 2, 1},
		{1, 3, 1},
		{2, 0, 4},
		{2, 1, 4},
		{2, 2, 2},
		{2, 3, 2},
		{3, 0, 6},
		{3, 1, 5},
		{3, 2, 3},
		{3, 3, 3},
		{177, 0, 354}, // 10110001 -> 101100010
		{177, 1, 353}, // 10110001 -> 101100001
		{177, 2, 353}, // 10110001 -> 101100001
		{177, 3, 353}, // 10110001 -> 101100001
		{177, 4, 353}, // 10110001 -> 101100001
		{177, 5, 337}, // 10110001 -> 101010001
		{177, 6, 305}, // 10110001 -> 100110001
		{177, 7, 305}, // 10110001 -> 100110001
		{177, 8, 177}, // 10110001 -> 010110001
	}
	for _, c := range cases {
		got := InsertZero(c.x, c.index)
		if got != c.expected {
			t.Errorf(
				"InsertZero(0x%016x, %d) = 0x%016x, expected 0x%016x",
				c.x, c.index, got, c.expected,
			)
		}
	}
}
