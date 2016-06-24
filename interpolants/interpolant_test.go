package interpolants

import (
	"github.com/denisskin/gomath/vector"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestFloaterHormannInterpolateRunge(t *testing.T) {

	// Runge function
	fn := func(x float64) float64 {
		return 1 / (1 + x*x)
	}

	// interpolation
	X := vector.Range(-1, 1, 5) // {-1, -0.5, 0, 0.5, 1}
	Y := X.Map(fn)
	fn1 := NewFloaterHormannInterpolator(X, Y, 0)

	// experiments
	x, err := 0.00, 0.0000001 // the exact value with minimal divergence
	assert.True(t, math.Abs(fn(x)/fn1(x)-1) < err)

	x, err = 0.13, 0.01 // interpolation  with 1% divergence
	assert.True(t, math.Abs(fn(x)/fn1(x)-1) < err)

	x, err = 1.10, 0.03 // extrapolation with 3% divergence
	assert.True(t, math.Abs(fn(x)/fn1(x)-1) < err)

}
