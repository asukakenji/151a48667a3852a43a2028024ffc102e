package bruteforce

import (
	"reflect"
	"testing"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/constant"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"
)

func DeepEqualOneOf(gotPath []int, expectedPaths [][]int) bool {
	for _, expectedPath := range expectedPaths {
		if reflect.DeepEqual(gotPath, expectedPath) {
			return true
		}
	}
	return false
}

var matrices = [][][]int{
	// https://youtu.be/-JjA4BLQyqE
	{
		{constant.Infinity, 1, 15, 6},
		{2, constant.Infinity, 7, 3},
		{9, 6, constant.Infinity, 12},
		{10, 4, 8, constant.Infinity},
	},
	// https://youtu.be/aQB_Y9D5pdw
	{
		{constant.Infinity, 10, 15, 20},
		{5, constant.Infinity, 9, 10},
		{6, 13, constant.Infinity, 12},
		{8, 8, 9, constant.Infinity},
	},
	// https://youtu.be/vNqE_LDTsa0
	{
		{constant.Infinity, 7, 6, 8, 4},
		{7, constant.Infinity, 8, 5, 6},
		{6, 8, constant.Infinity, 9, 7},
		{8, 5, 9, constant.Infinity, 8},
		{4, 6, 7, 8, constant.Infinity},
	},
	// https://youtu.be/FJkT_dRjX94
	// https://youtu.be/KzWC-t1y8Ac
	{
		{constant.Infinity, 20, 30, 10, 11},
		{15, constant.Infinity, 16, 4, 2},
		{3, 5, constant.Infinity, 2, 4},
		{19, 6, 18, constant.Infinity, 3},
		{16, 4, 7, 16, constant.Infinity},
	},
	// https://youtu.be/-cLsEHP0qt0
	// https://youtu.be/nN4K8xA8ShM
	// https://youtu.be/LjvdXKsvUpU
	// https://youtu.be/hjZDDz3r1es
	{
		{constant.Infinity, 10, 8, 9, 7},
		{10, constant.Infinity, 10, 5, 6},
		{8, 10, constant.Infinity, 8, 9},
		{9, 5, 8, constant.Infinity, 6},
		{7, 6, 9, 6, constant.Infinity},
	},
	// Example
	{
		{constant.Infinity, 15518, 9667},
		{15223, constant.Infinity, 8333},
		{10329, 8464, constant.Infinity},
	},
}

func TestTravellingSalesmanTour(t *testing.T) {
	cases := []struct {
		m             [][]int
		expectedCost  int
		expectedPaths [][]int
	}{
		{
			matrices[0], 21, [][]int{
				{0, 1, 3, 2}, // A -> B -> D -> C -> A
			},
		},
		{
			matrices[1], 35, [][]int{
				{0, 1, 3, 2}, // A -> B -> D -> C -> A
			},
		},
		{
			matrices[2], 30, [][]int{
				{0, 2, 3, 1, 4}, // A -> C -> D -> B -> E -> A
				{0, 4, 1, 3, 2}, // A -> E -> B -> D -> C -> A (symmetric)
			},
		},
		{
			matrices[3], 28, [][]int{
				{0, 3, 1, 4, 2}, // A -> D -> B -> E -> C -> A
			},
		},
		{
			matrices[4], 34, [][]int{
				{0, 2, 3, 1, 4}, // A -> C -> D -> B -> E -> A
				{0, 4, 1, 3, 2}, // A -> E -> B -> D -> C -> A (symmetric)
			},
		},
		{
			matrices[5], 33354, [][]int{
				{0, 2, 1},
			},
		},
	}
	for i, c := range cases {
		m := matrix.NewSquareMatrix(c.m)
		gotCost, gotPath := TravellingSalesmanTour(m)
		if gotCost != c.expectedCost || !DeepEqualOneOf(gotPath, c.expectedPaths) {
			t.Errorf("TestTravellingSalesmanTour Case #%d: Got: (%d, %v), Expected: (%d, %v)", i, gotCost, gotPath, c.expectedCost, c.expectedPaths)
		}
	}
}

func TestTravellingSalesmanPath(t *testing.T) {
	cases := []struct {
		m             [][]int
		expectedCost  int
		expectedPaths [][]int
	}{
		{
			matrices[0], 12, [][]int{
				{0, 1, 3, 2}, // A -> B -> D -> C
			},
		},
		{
			matrices[1], 29, [][]int{
				{0, 1, 3, 2}, // A -> B -> D -> C
			},
		},
		{
			matrices[2], 24, [][]int{
				{0, 4, 2, 1, 3}, // A -> E -> C -> B -> D
				{0, 2, 4, 1, 3}, // A -> C -> E -> B -> D
				//{0, 2, 3, 1, 4}, // A -> C -> D -> B -> E (cost = 26, symmetric in tour)
				{0, 4, 1, 3, 2}, // A -> E -> B -> D -> C
			},
		},
		{
			matrices[3], 25, [][]int{
				{0, 3, 1, 4, 2}, // A -> D -> B -> E -> C
				{0, 3, 4, 2, 1}, // A -> D -> E -> C -> B
			},
		},
		{
			matrices[4], 26, [][]int{
				//{0, 2, 3, 1, 4}, // A -> C -> D -> B -> E (cost = 27, symmetric in tour)
				{0, 4, 1, 3, 2}, // A -> E -> B -> D -> C
			},
		},
		{
			matrices[5], 18131, [][]int{
				{0, 2, 1},
			},
		},
	}
	for i, c := range cases {
		m := matrix.NewSquareMatrix(c.m)
		gotCost, gotPath := TravellingSalesmanPath(m)
		if gotCost != c.expectedCost || !DeepEqualOneOf(gotPath, c.expectedPaths) {
			t.Errorf("TestTravellingSalesmanPath Case #%d: Got: (%d, %v), Expected: (%d, %v)", i, gotCost, gotPath, c.expectedCost, c.expectedPaths)
		}
	}
}

func TestTravellingSalesmanAllTours(t *testing.T) {
	cases := []struct {
		m             [][]int
		expectedCost  int
		expectedPaths [][]int
	}{
		{
			matrices[0], 21, [][]int{
				{0, 1, 3, 2}, // A -> B -> D -> C -> A
			},
		},
		{
			matrices[1], 35, [][]int{
				{0, 1, 3, 2}, // A -> B -> D -> C -> A
			},
		},
		{
			matrices[2], 30, [][]int{
				{0, 2, 3, 1, 4}, // A -> C -> D -> B -> E -> A
				{0, 4, 1, 3, 2}, // A -> E -> B -> D -> C -> A (symmetric)
			},
		},
		{
			matrices[3], 28, [][]int{
				{0, 3, 1, 4, 2}, // A -> D -> B -> E -> C -> A
			},
		},
		{
			matrices[4], 34, [][]int{
				{0, 2, 3, 1, 4}, // A -> C -> D -> B -> E -> A
				{0, 4, 1, 3, 2}, // A -> E -> B -> D -> C -> A (symmetric)
			},
		},
		{
			matrices[5], 33354, [][]int{
				{0, 2, 1},
			},
		},
	}
	for i, c := range cases {
		m := matrix.NewSquareMatrix(c.m)
		gotCost, gotPaths := TravellingSalesmanAllTours(m)
		if gotCost != c.expectedCost || !reflect.DeepEqual(gotPaths, c.expectedPaths) {
			t.Errorf("TestTravellingSalesmanAllTours Case #%d: Got: (%d, %v), Expected: (%d, %v)", i, gotCost, gotPaths, c.expectedCost, c.expectedPaths)
		}
	}
}

func TestTravellingSalesmanAllPaths(t *testing.T) {
	cases := []struct {
		m             [][]int
		expectedCost  int
		expectedPaths [][]int
	}{
		{
			matrices[0], 12, [][]int{
				{0, 1, 3, 2}, // A -> B -> D -> C
			},
		},
		{
			matrices[1], 29, [][]int{
				{0, 1, 3, 2}, // A -> B -> D -> C
			},
		},
		{
			matrices[2], 24, [][]int{
				{0, 4, 2, 1, 3}, // A -> E -> C -> B -> D
				{0, 2, 4, 1, 3}, // A -> C -> E -> B -> D
				//{0, 2, 3, 1, 4}, // A -> C -> D -> B -> E (cost = 26, symmetric in tour)
				{0, 4, 1, 3, 2}, // A -> E -> B -> D -> C
			},
		},
		{
			matrices[3], 25, [][]int{
				{0, 3, 1, 4, 2}, // A -> D -> B -> E -> C
				{0, 3, 4, 2, 1}, // A -> D -> E -> C -> B
			},
		},
		{
			matrices[4], 26, [][]int{
				//{0, 2, 3, 1, 4}, // A -> C -> D -> B -> E (cost = 27, symmetric in tour)
				{0, 4, 1, 3, 2}, // A -> E -> B -> D -> C
			},
		},
		{
			matrices[5], 18131, [][]int{
				{0, 2, 1},
			},
		},
	}
	for i, c := range cases {
		m := matrix.NewSquareMatrix(c.m)
		gotCost, gotPaths := TravellingSalesmanAllPaths(m)
		if gotCost != c.expectedCost || !reflect.DeepEqual(gotPaths, c.expectedPaths) {
			t.Errorf("TestTravellingSalesmanAllPaths Case #%d: Got: (%d, %v), Expected: (%d, %v)", i, gotCost, gotPaths, c.expectedCost, c.expectedPaths)
		}
	}
}
