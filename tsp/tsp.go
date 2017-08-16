package tsp

import (
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/constant"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"
)

func PathCost(m matrix.Matrix, path []int) int {
	cost := 0
	size := len(path)
	for i := 1; i < size; i++ {
		cost = AddCosts(cost, m.Get(path[i-1], path[i]))
	}
	return cost
}

func TourCost(m matrix.Matrix, path []int) int {
	return AddCosts(PathCost(m, path), m.Get(path[len(path)-1], path[0]))
}

func AddCosts(cost1, cost2 int) int {
	if cost1 == constant.Infinity || cost2 == constant.Infinity {
		return constant.Infinity
	}
	return cost1 + cost2
}

func IsLessThan(cost1, cost2 int) bool {
	if cost1 == constant.Infinity {
		return false
	}
	if cost2 == constant.Infinity {
		return true
	}
	return cost1 < cost2
}

func IsLessThanEqualTo(cost1, cost2 int) bool {
	if cost1 == constant.Infinity {
		return cost2 == constant.Infinity
	}
	if cost2 == constant.Infinity {
		return true
	}
	return cost1 <= cost2
}
