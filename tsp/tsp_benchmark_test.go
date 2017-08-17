package tsp_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/constant"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/tsp/bruteforce"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/tsp/heldkarp"
)

const (
	// To keep any distance between 2 points in the square less than 255
	SquareSize = 181
)

var (
	seed = time.Now().UTC().UnixNano()
	src  = rand.NewSource(seed)
	rng  = rand.New(src)
)

func GeneratePoint() (x, y int) {
	x = rng.Intn(SquareSize)
	y = rng.Intn(SquareSize)
	return x, y
}

func GenerateSquareMatrix(size int) matrix.Matrix {
	type point struct {
		x int
		y int
	}
	points := make([]point, size)
	for i := 0; i < size; i++ {
		x, y := GeneratePoint()
		points[i] = point{x, y}
	}
	m := make([][]int, size)
	for row := range m {
		m[row] = make([]int, size)
		for col := range m[row] {
			if row == col {
				m[row][col] = constant.Infinity
			}
			x1, y1 := points[row].x, points[row].y
			x2, y2 := points[col].x, points[col].y
			dx := x2 - x1
			dy := y2 - y1
			m[row][col] = int(math.Sqrt(float64(dx*dx+dy*dy)) + 0.5)
		}
	}
	return matrix.NewSquareMatrix(m)
}

var (
	m2  = GenerateSquareMatrix(2)
	m3  = GenerateSquareMatrix(3)
	m4  = GenerateSquareMatrix(4)
	m5  = GenerateSquareMatrix(5)
	m6  = GenerateSquareMatrix(6)
	m7  = GenerateSquareMatrix(7)
	m8  = GenerateSquareMatrix(8)
	m9  = GenerateSquareMatrix(9)
	m10 = GenerateSquareMatrix(10)
	m11 = GenerateSquareMatrix(11)
	m12 = GenerateSquareMatrix(12)
	m13 = GenerateSquareMatrix(13)
	m14 = GenerateSquareMatrix(14)
	m15 = GenerateSquareMatrix(15)
	m16 = GenerateSquareMatrix(16)
	m17 = GenerateSquareMatrix(17)
	m18 = GenerateSquareMatrix(18)
	m19 = GenerateSquareMatrix(19)
)

func benchmarkTravellingSalesman(b *testing.B, tspFunc func(matrix.Matrix) (int, []int), m matrix.Matrix) {
	for i, count := 0, b.N; i < count; i++ {
		tspFunc(m)
	}
}

func BenchmarkBruteForce2(b *testing.B) {
	benchmarkTravellingSalesman(b, bruteforce.TravellingSalesmanTour, m2)
}

func BenchmarkBruteForce3(b *testing.B) {
	benchmarkTravellingSalesman(b, bruteforce.TravellingSalesmanTour, m3)
}

func BenchmarkBruteForce4(b *testing.B) {
	benchmarkTravellingSalesman(b, bruteforce.TravellingSalesmanTour, m4)
}

func BenchmarkBruteForce5(b *testing.B) {
	benchmarkTravellingSalesman(b, bruteforce.TravellingSalesmanTour, m5)
}

func BenchmarkBruteForce6(b *testing.B) {
	benchmarkTravellingSalesman(b, bruteforce.TravellingSalesmanTour, m6)
}

func BenchmarkBruteForce7(b *testing.B) {
	benchmarkTravellingSalesman(b, bruteforce.TravellingSalesmanTour, m7)
}

func BenchmarkBruteForce8(b *testing.B) {
	benchmarkTravellingSalesman(b, bruteforce.TravellingSalesmanTour, m8)
}

func BenchmarkBruteForce9(b *testing.B) {
	benchmarkTravellingSalesman(b, bruteforce.TravellingSalesmanTour, m9)
}

func BenchmarkBruteForce10(b *testing.B) {
	benchmarkTravellingSalesman(b, bruteforce.TravellingSalesmanTour, m10)
}

func BenchmarkBruteForce11(b *testing.B) {
	benchmarkTravellingSalesman(b, bruteforce.TravellingSalesmanTour, m11)
}

func BenchmarkBruteForce12(b *testing.B) {
	benchmarkTravellingSalesman(b, bruteforce.TravellingSalesmanTour, m12)
}

func BenchmarkHeldKarp2(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m2)
}

func BenchmarkHeldKarp3(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m3)
}

func BenchmarkHeldKarp4(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m4)
}

func BenchmarkHeldKarp5(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m5)
}

func BenchmarkHeldKarp6(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m6)
}

func BenchmarkHeldKarp7(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m7)
}

func BenchmarkHeldKarp8(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m8)
}

func BenchmarkHeldKarp9(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m9)
}

func BenchmarkHeldKarp10(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m10)
}

func BenchmarkHeldKarp11(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m11)
}

func BenchmarkHeldKarp12(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m12)
}

func BenchmarkHeldKarp13(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m13)
}

func BenchmarkHeldKarp14(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m14)
}

func BenchmarkHeldKarp15(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m15)
}

func BenchmarkHeldKarp16(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m16)
}

func BenchmarkHeldKarp17(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m17)
}

func BenchmarkHeldKarp18(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m18)
}

func BenchmarkHeldKarp19(b *testing.B) {
	benchmarkTravellingSalesman(b, heldkarp.TravellingSalesmanTour, m19)
}
