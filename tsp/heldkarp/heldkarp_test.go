package heldkarp

import (
	"testing"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"
)

func Test1(t *testing.T) {
	m := matrix.NewSquareMatrix([][]int{
		{-1, 1, 15, 6},
		{2, -1, 7, 3},
		{9, 6, -1, 12},
		{10, 4, 8, -1},
	})
	TravellingSalesmanTour(m)
}
