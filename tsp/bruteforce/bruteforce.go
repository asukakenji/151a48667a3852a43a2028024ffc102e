package bruteforce

import (
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/constant"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/tsp"
)

func permutateExceptFirstElement(f func([]int), d int, s []int) {
	if d == 0 {
		f(s)
		return
	}
	for i := d; i > 0; i-- {
		s[i], s[d] = s[d], s[i]
		permutateExceptFirstElement(f, d-1, s)
		s[i], s[d] = s[d], s[i]
	}
}

// len(s) must not be 0.
func PermutateExceptFirstElement(f func([]int), s []int) {
	// if len(s) == 0 {
	// 	f(s)
	// 	return
	// }
	permutateExceptFirstElement(f, len(s)-1, s)
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
		if tsp.IsLessThan(cost, minimumCost) {
			copy(minimumPath, path)
			minimumCost = cost
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

func travellingSalesmanAll(m matrix.Matrix, costFunc func(matrix.Matrix, []int) int) (cost int, path [][]int) {
	size, _ := m.Size()
	indices := make([]int, size)
	for i := range indices {
		indices[i] = i
	}

	minimumCost := constant.MaxInt
	var minimumPaths [][]int
	PermutateExceptFirstElement(func(path []int) {
		cost := costFunc(m, path)
		if tsp.IsLessThanEqualTo(cost, minimumCost) {
			minimumPath := make([]int, size)
			copy(minimumPath, path)
			if tsp.IsLessThan(cost, minimumCost) {
				minimumPaths = [][]int{minimumPath}
			} else {
				minimumPaths = append(minimumPaths, minimumPath)
			}
			minimumCost = cost
		}
	}, indices)
	return minimumCost, minimumPaths
}

func TravellingSalesmanAllPaths(m matrix.Matrix) (cost int, paths [][]int) {
	return travellingSalesmanAll(m, tsp.PathCost)
}

func TravellingSalesmanAllTours(m matrix.Matrix) (cost int, paths [][]int) {
	return travellingSalesmanAll(m, tsp.TourCost)
}
