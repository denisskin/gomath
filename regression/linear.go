package regression

func NewLinearRegression() RegressionFunc {
	return LinearRegression
}

func LinearRegression(X, Y []float64) Func {

	n := float64(len(X))
	var sx, sy, sx2, sy2, sxy float64
	for i, x := range X {
		y := Y[i]

		sx += x
		sy += y
		sx2 += x * x
		sy2 += y * y
		sxy += x * y
	}
	// todo: R2  := (n*sxy - sx*sy) / math.Sqrt((n*sx2-sx*sx)*(n*sy2-sy*sy))

	k := (sxy*n - sx*sy) / (sx2*n - sx*sx)
	b := (sy - k*sx) / n

	return func(x float64) float64 {
		return k*x + b
	}
}
