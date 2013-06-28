package optbench

import (
	"math"
)

// Produce f(x0) + f(x1) + ... + f(xn).
func SumOver(xs []float64, f func (float64) float64) float64 {
	var t float64
	for _, x := range xs { t += f(x) }
	return t
}

func Sphere(xs []float64) float64 {
	f := func (x float64) float64 { return x*x }
	return SumOver(xs, f)
}

func SchwefelsDoubleSum(xs []float64) float64 {
	var t float64
	ident := func (x float64) float64 { return x }
	for i := range xs {
		v := SumOver(xs[:i+1], ident)
		t += v*v
	}
	return t
}

func Rosenbrock(xs []float64) float64 {
	var t float64
	for i := 0; i < len(xs) - 1; i++ {
		a, b := xs[i+1] - xs[i]*xs[i], xs[i] - 1
		t += 100*a*a + b*b
	}
	return t
}

func Rastrigin(xs []float64) float64 {
	f := func (x float64) float64 { return x*x - 10*math.Cos(2*math.Pi*x) }
	return SumOver(xs, f) + 10.0 * float64(len(xs))
}

func Schwefel(xs []float64) float64 {
	m := 418.9828872724338
	f := func (x float64) float64 {
		return x * math.Sin(math.Pow(math.Abs(x), 0.5))
	}
	return SumOver(xs, f) + m * float64(len(xs))
}

func Ackley(xs []float64) float64 {
	p := float64(len(xs))
	f1 := func (x float64) float64 { return x*x }
	a := math.Exp(-0.2 * math.Pow(SumOver(xs, f1) / p, 0.5))
	f2 := func (x float64) float64 { return math.Cos(2*math.Pi*x) }
	b := math.Exp(SumOver(xs, f2) / p)
	return 20 + math.E - 20*a - b
}
