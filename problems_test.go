package optbench

import (
	"testing"
)

// Test zeros at various dimensions for functions with a zero at [c, c, ..., c].
func CheckZeros(c float64, f func ([]float64) float64, t *testing.T) {
	for _, dim := range []int{1, 2, 3, 5, 8, 30, 51, 99} {
		xs := make([]float64, dim)
		for i := range xs { xs[i] = c }
		if fx := f(xs); fx != 0 {
			t.Errorf("Expected f(%x) = 0, got %v.", xs, fx)
		}
	}
}

func TestSphere(t *testing.T) {
	f := Sphere
	CheckZeros(0, f, t)
	// Test a few known, non-zero points.
	if fx := f([]float64{1, 1, 1}); fx != 3 {
		t.Errorf("Expected 0, got %v.", fx)
	}
	if fx := f([]float64{0.5, 0.0, 0.75, 0.1}); fx != 0.8225 {
		t.Errorf("Expected 0, got %v.", fx)
	}
}

func TestSchwefelsDoubleSum(t *testing.T) {
	f := SchwefelsDoubleSum
	CheckZeros(0, f, t)
	xs, v := []float64{0.39, 0.19, 0.61, 0.15, 0.94, 0.48, 0.43, 0.4, 0.48, 0.53}, 77.3053
	if fx := f(xs); fx != v { t.Errorf("Expected %v, got %v.", v, fx) }
	xs, v = []float64{0.904, 0.457, 0.601}, 6.518981
	if fx := f(xs); fx != v { t.Errorf("Expected %v, got %v.", v, fx) }
}

func TestRosenbrock(t *testing.T) {
	f := Rosenbrock
	CheckZeros(1, f, t)
	xs, v := make([]float64, 219), 218.0
	if fx := f(xs); fx != v { t.Errorf("Expected %v, got %v.", v, fx) }
	xs, v = []float64{0.43, 0.56, 0.8, 0.76, 0.63}, 40.059173
	if fx := f(xs); fx != v { t.Errorf("Expected %v, got %v.", v, fx) }
}
