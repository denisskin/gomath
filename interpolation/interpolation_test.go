package interpolation

import (
	"math"
	"testing"

	"github.com/denisskin/gomath/vector"
	"github.com/stretchr/testify/assert"
)

func TestFloaterHormannInterpolation(t *testing.T) {

	// Runge function
	fn := func(x float64) float64 {
		return 1 / (1 + x*x)
	}

	// points
	X := vector.Range(-1, 1, 5) // {-1, -0.5, 0, 0.5, 1}
	Y := X.Map(fn)

	// interpolation
	fn1 := FloaterHormannInterpolation(X, Y, 0)

	// experiments
	x := 0.00 // the exact values are equal
	assert.True(t, fn(x) == fn1(x))

	x = 0.13 // interpolation  with 1% divergence
	assert.True(t, math.Abs(fn(x)/fn1(x)-1) < 0.01)

	x = 1.10 // extrapolation with 3% divergence
	assert.True(t, math.Abs(fn(x)/fn1(x)-1) < 0.03)
}
