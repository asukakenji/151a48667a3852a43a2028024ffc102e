package bitstring

import "testing"

func TestLeftShift(t *testing.T) {
	cases := []struct {
		bs       BitString64
		amount   uint
		expected BitString64
	}{
		{0, 0, 0},
		{0, 1, 0},
		{0, 2, 0},
		{1, 0, 1},
		{1, 1, 2},
		{1, 2, 4},
		{2, 0, 2},
		{2, 1, 4},
		{2, 2, 8},
		{3, 0, 3},
		{3, 1, 6},
		{3, 2, 12},
	}
	for _, c := range cases {
		got := c.bs.LeftShift(c.amount)
		if got != c.expected {
			t.Errorf("BitString(0x%016x).LeftShift(%d) = 0x%016x, expected 0x%016x", c.bs, c.amount, got, c.expected)
		}
	}
}

func TestRightShift(t *testing.T) {
	cases := []struct {
		bs       BitString64
		amount   uint
		expected BitString64
	}{
		{0, 0, 0},
		{0, 1, 0},
		{0, 2, 0},
		{1, 0, 1},
		{1, 1, 0},
		{1, 2, 0},
		{2, 0, 2},
		{2, 1, 1},
		{2, 2, 0},
		{3, 0, 3},
		{3, 1, 1},
		{3, 2, 0},
	}
	for _, c := range cases {
		got := c.bs.RightShift(c.amount)
		if got != c.expected {
			t.Errorf("BitString(0x%016x).RightShift(%d) = 0x%016x, expected 0x%016x", c.bs, c.amount, got, c.expected)
		}
	}
}

func TestIsSet(t *testing.T) {
	cases := []struct {
		bs       BitString64
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
		got := c.bs.IsSet(c.index)
		if got != c.expected {
			t.Errorf("BitString(0x%016x).IsSet(%d) = %t, expected %t", c.bs, c.index, got, c.expected)
		}
	}
}

func TestSet(t *testing.T) {
	cases := []struct {
		bs       BitString64
		index    uint
		expected BitString64
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
		got := c.bs.Set(c.index)
		if got != c.expected {
			t.Errorf("BitString(0x%016x).Set(%d) = 0x%016x, expected 0x%016x", c.bs, c.index, got, c.expected)
		}
	}
}

func TestReset(t *testing.T) {
	cases := []struct {
		bs       BitString64
		index    uint
		expected BitString64
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
		got := c.bs.Reset(c.index)
		if got != c.expected {
			t.Errorf("BitString(0x%016x).Reset(%d) = 0x%016x, expected 0x%016x", c.bs, c.index, got, c.expected)
		}
	}
}

func TestUnsetRightmostOne(t *testing.T) {
	cases := []struct {
		bs       BitString64
		expected BitString64
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
		got := c.bs.UnsetRightmostOne()
		if got != c.expected {
			t.Errorf("BitString(0x%016x).UnsetRightmostOne() = 0x%016x, expected 0x%016x", c.bs, got, c.expected)
		}
	}
}

func TestSetRightmostZero(t *testing.T) {
	cases := []struct {
		bs       BitString64
		expected BitString64
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
		got := c.bs.SetRightmostZero()
		if got != c.expected {
			t.Errorf("BitString(0x%016x).SetRightmostZero() = 0x%016x, expected 0x%016x", c.bs, got, c.expected)
		}
	}
}

func TestNextValueWithSameBitCount(t *testing.T) {
	cases := []struct {
		bs       BitString64
		expected BitString64
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
		got := c.bs.NextValueWithSameBitCount()
		if got != c.expected {
			t.Errorf("BitString(0x%016x).TestNextValueWithSameBitCount() = 0x%016x, expected 0x%016x", c.bs, got, c.expected)
		}
	}
}

func TestInsertZero(t *testing.T) {
	cases := []struct {
		bs       BitString64
		index    uint
		expected BitString64
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
		{177, 0, 354},
		{177, 1, 353},
		{177, 2, 353},
		{177, 3, 353},
		{177, 4, 353},
		{177, 5, 337},
		{177, 6, 305},
		{177, 7, 305},
		{177, 8, 177},
	}
	for _, c := range cases {
		got := c.bs.InsertZero(c.index)
		if got != c.expected {
			t.Errorf("BitString(0x%016x).TestInsertZero(%d) = 0x%016x, expected 0x%016x", c.bs, c.index, got, c.expected)
		}
	}
}
