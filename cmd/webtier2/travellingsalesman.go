package main

import "math"

func isSquareMatrix(matrix [][]int32) bool {
	length := len(matrix)
	for _, row := range matrix {
		if len(row) != length {
			return false
		}
	}
	return true
}

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

func TravellingSalesmanNaive(matrix [][]int32) (cost int32, path []int) {
	if !isSquareMatrix(matrix) {
		return 0, nil
	}
	length := len(matrix)
	indices := make([]int, length)
	for i := range indices {
		indices[i] = i
	}

	minimumCost := int32(math.MaxInt32)
	minimumPath := make([]int, length)
	PermutateExceptFirstElement(func(path []int) {
		cost := int32(0)
		for i := 1; i < length; i++ {
			cost += matrix[path[i-1]][path[i]]
		}
		if cost < minimumCost {
			minimumCost = cost
			copy(minimumPath, path)
		}
	}, indices)
	return minimumCost, minimumPath
}
