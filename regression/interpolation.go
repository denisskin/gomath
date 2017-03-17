package regression

import (
	"github.com/denisskin/gomath/interpolation"
	"github.com/denisskin/gomath/vector"
)

func NewInterpolationRegression(size, d int) RegressionFunc {
	return func(X, Y []float64) Func {
		return interpolationRegression(X, Y, size, d)
	}
}

func interpolationRegression(X, Y vector.Vector, size, d int) Func {
	if size == 0 {
		size = 10
	}
	if size >= X.Size() {
		return Func(interpolation.FloaterHormannInterpolation(X, Y, d))
	}

	x0 := X.Min()
	dx := (X.Max() - x0) / float64(size)

	nn := make([]int, size)
	xx := make([]float64, size)
	yy := make([]float64, size)

	for i, x := range X {
		j := int((x - x0) / dx)
		if j >= size {
			j = size - 1
		}
		nn[j]++
		xx[j] += x
		yy[j] += Y[i]
	}
	for j, n := range nn {
		xx[j] /= float64(n)
		yy[j] /= float64(n)
	}

	return Func(interpolation.FloaterHormannInterpolation(xx, yy, d))
}
