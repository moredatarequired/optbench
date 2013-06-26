package optbench

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