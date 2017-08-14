package tsp

import "github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"

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
