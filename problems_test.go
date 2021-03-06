package optbench

import (
	"testing"
	"math"
)

type Evaluation func ([]float64) float64

func ApproxEqual(x float64, y float64) bool {
	return math.Abs(x - y) < 1e-8
}

func CheckVector(f Evaluation, xs []float64, v float64, t *testing.T) {
	if fx := f(xs); !ApproxEqual(fx, v) {
		t.Errorf("Expected f(%v) = %v, got %v.", xs, v, fx)
	}
}

// Test zeros at various dimensions for functions with a zero at [c, c, ..., c].
func CheckZeros(f Evaluation, c float64, t *testing.T) {
	for _, dim := range []int{1, 2, 3, 5, 8, 30, 51, 99} {
		xs := make([]float64, dim)
		for i := range xs { xs[i] = c }
		CheckVector(f, xs, 0.0, t)
	}
}

func TestSphere(t *testing.T) {
	f := Sphere
	CheckZeros(f, 0, t)
	// Test a few known, non-zero points.
	CheckVector(f, []float64{1, 1, 1}, 3, t)
	CheckVector(f, []float64{0.5, 0.0, 0.75, 0.1}, 0.8225, t)
}

func TestSchwefelsDoubleSum(t *testing.T) {
	f := SchwefelsDoubleSum
	CheckZeros(f, 0, t)
	xs := []float64{0.39, 0.19, 0.61, 0.15, 0.94, 0.48, 0.43, 0.4, 0.48, 0.53}
	CheckVector(f, xs, 77.3053, t)
	CheckVector(f, []float64{0.904, 0.457, 0.601}, 6.518981, t)
}

func TestRosenbrock(t *testing.T) {
	f := Rosenbrock
	CheckZeros(f, 1, t)
	CheckVector(f, make([]float64, 219), 218.0, t)
	CheckVector(f, []float64{0.43, 0.56, 0.8, 0.76, 0.63}, 40.059173, t)
}

func TestRastrigin(t *testing.T) {
	f := Rastrigin
	CheckZeros(f, 0, t)
	CheckVector(f, []float64{1, 1, 1}, 3.0, t)
	xs, v := []float64{0.08, 0.82, 0.8, 0.95, 0.59}, 35.39108443222979
	CheckVector(f, xs, v, t)
}

func TestSchwefel(t *testing.T) {
	f := Schwefel
	CheckZeros(f, -420.96874682399, t)
	CheckVector(f, make([]float64, 111), 46507.10048724015, t)
	xs, v := []float64{35.72, 3.24, 8.27, 58.05, -82.5}, 2118.2632974942176
	CheckVector(f, xs, v, t)
}

func TestAckley(t *testing.T) {
	f := Ackley
	CheckZeros(f, 0, t)
	xs, v := []float64{-35.97, 82.75, -15.0, 21.56, -88.64}, 21.631529305754
	CheckVector(f, xs, v, t)
}
