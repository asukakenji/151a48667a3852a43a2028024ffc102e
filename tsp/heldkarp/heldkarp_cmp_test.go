package heldkarp

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/constant"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/tsp"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/tsp/bruteforce"
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

func TestTravellingSalesmanTour_Compare(t *testing.T) {
	for i := 0; i < 1024; i++ {
		size := 2 + rng.Intn(9)
		m := GenerateSquareMatrix(size)
		expectedCost, expectedPath := bruteforce.TravellingSalesmanTour(m)
		gotCost, gotPath := TravellingSalesmanTour(m)
		if gotCost != expectedCost {
			t.Errorf(
				"TravellingSalesmanTour(%v) = (%d, %v), expected (%d, %v)",
				m, gotCost, gotPath, expectedCost, expectedPath,
			)
		}
		if !reflect.DeepEqual(gotPath, expectedPath) {
			cost1 := tsp.TourCost(m, gotPath)
			cost2 := tsp.TourCost(m, expectedPath)
			if cost1 != cost2 {
				t.Errorf(
					"TravellingSalesmanTour(%v) = (%d, %v), expected (%d, %v)",
					m, gotCost, gotPath, expectedCost, expectedPath,
				)
			}
		}
	}
}
