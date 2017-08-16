package heldkarp

import (
	"reflect"
	"testing"

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
		{-1, 1, 15, 6},
		{2, -1, 7, 3},
		{9, 6, -1, 12},
		{10, 4, 8, -1},
	},
	// https://youtu.be/aQB_Y9D5pdw
	{
		{-1, 10, 15, 20},
		{5, -1, 9, 10},
		{6, 13, -1, 12},
		{8, 8, 9, -1},
	},
	// https://youtu.be/vNqE_LDTsa0
	{
		{-1, 7, 6, 8, 4},
		{7, -1, 8, 5, 6},
		{6, 8, -1, 9, 7},
		{8, 5, 9, -1, 8},
		{4, 6, 7, 8, -1},
	},
	// https://youtu.be/FJkT_dRjX94
	// https://youtu.be/KzWC-t1y8Ac
	{
		{-1, 20, 30, 10, 11},
		{15, -1, 16, 4, 2},
		{3, 5, -1, 2, 4},
		{19, 6, 18, -1, 3},
		{16, 4, 7, 16, -1},
	},
	// https://youtu.be/-cLsEHP0qt0
	// https://youtu.be/nN4K8xA8ShM
	// https://youtu.be/LjvdXKsvUpU
	// https://youtu.be/hjZDDz3r1es
	{
		{-1, 10, 8, 9, 7},
		{10, -1, 10, 5, 6},
		{8, 10, -1, 8, 9},
		{9, 5, 8, -1, 6},
		{7, 6, 9, 6, -1},
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
	}
	for i, c := range cases {
		m := matrix.NewSquareMatrix(c.m)
		gotCost, gotPath := TravellingSalesmanTour(m)
		if gotCost != c.expectedCost || !DeepEqualOneOf(gotPath, c.expectedPaths) {
			t.Errorf("TestTravellingSalesmanTour Case #%d: Got: (%d, %v), Expected: (%d, %v)", i, gotCost, gotPath, c.expectedCost, c.expectedPaths)
		}
	}
}
