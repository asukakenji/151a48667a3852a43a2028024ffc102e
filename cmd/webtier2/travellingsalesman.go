package main

import (
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/constant"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"
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

func PathCost(m matrix.Matrix, path []int) int {
	cost := 0
	size := len(path)
	for i := 1; i < size; i++ {
		cost += m.Get(path[i-1], path[i])
	}
	return cost
}

func TourCost(m matrix.Matrix, path []int) int {
	return PathCost(m, path) + m.Get(path[len(path)-1], path[0])
}

func TravellingSalesmanNaive(m matrix.Matrix) (cost int, path []int) {
	size, _ := m.Size()
	indices := make([]int, size)
	for i := range indices {
		indices[i] = i
	}

	minimumCost := constant.MaxInt
	minimumPath := make([]int, size)
	PermutateExceptFirstElement(func(path []int) {
		cost := PathCost(m, path)
		if cost < minimumCost {
			minimumCost = cost
			copy(minimumPath, path)
		}
	}, indices)
	return minimumCost, minimumPath
}
