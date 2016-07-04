package interpolants

import "math"

// NewFloaterHormannInterpolator creates interpolator function by the method barycentric
// rational interpolants introduced by Floater and Hormann in 2007
// are “blends” of polynomial interpolants of fixed degree d.
// See: http://homeweb2.unifr.ch/kleing/pub/resources/Klein_final.pdf
func NewFloaterHormannInterpolator(X, Y []float64, d int) Interpolator {

	n := len(Y)

	if X == nil {
		X = make([]float64, n)
		for i := 0; i < n; i++ {
			X[i] = float64(i)
		}
	}

	if n != len(X) {
		panic("Numbers of elements X and Y not equal")
	}
	if d <= 0 {
		d = (4*n - 1) / 9
		if d > 7 {
			d = 7
		}
		if d >= n {
			d = n - 1
		}
	}
	if d <= 0 || d >= n {
		panic("Enexpected parametr d! (Sholud be 0 < d < n)")
	}

	w := make([]float64, n)
	for i := 0; i < n-d; i++ {
		w[i] = 1
	}

	sgn := 1.
	for j := d; j > 0; j-- {
		wi := 0.
		for i := 0; i <= n-j; i++ {
			_wi := w[i]
			w[i] = 0
			if i > 0 {
				w[i] += math.Abs(wi) / (X[i] - X[i-1])
			}
			if i < n-1 {
				w[i] += math.Abs(_wi) / (X[i+1] - X[i])
			}
			wi = _wi
			sgn = -sgn
			w[i] *= sgn
		}
	}

	return func(x float64) float64 {
		sx := 0.
		sy := 0.
		for k, yk := range Y {
			dx := x - X[k]
			if dx == 0 {
				return yk
			}
			v := w[k] / dx
			sx += v
			sy += v * yk
		}
		return sy / sx
	}
}
