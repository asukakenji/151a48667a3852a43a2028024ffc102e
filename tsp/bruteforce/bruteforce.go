package bruteforce

import (
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/constant"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/tsp"
)

func permutateExceptFirstElement(f func([]int), d int, s []int) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i > 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutateExceptFirstElement(f, d-1, s)
			s[i], s[d] = s[d], s[i]
		}
	}
}

func PermutateExceptFirstElement(f func([]int), s []int) {
	if len(s) == 0 {
		f(s)
	} else {
		permutateExceptFirstElement(f, len(s)-1, s)
	}
}

func travellingSalesman(m matrix.Matrix, costFunc func(matrix.Matrix, []int) int) (cost int, path []int) {
	size, _ := m.Size()
	indices := make([]int, size)
	for i := range indices {
		indices[i] = i
	}

	minimumCost := constant.MaxInt
	minimumPath := make([]int, size)
	PermutateExceptFirstElement(func(path []int) {
		cost := costFunc(m, path)
		if cost < minimumCost {
			minimumCost = cost
			copy(minimumPath, path)
		}
	}, indices)
	return minimumCost, minimumPath
}

func TravellingSalesmanPath(m matrix.Matrix) (cost int, path []int) {
	return travellingSalesman(m, tsp.PathCost)
}

func TravellingSalesmanTour(m matrix.Matrix) (cost int, path []int) {
	return travellingSalesman(m, tsp.TourCost)
}
