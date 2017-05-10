package regression

import (
	"math"
	"math/rand"
	"testing"

	"github.com/denisskin/gomath/vector"
	"github.com/stretchr/testify/assert"
)

func fnLine(x float64) float64 {
	return 3*x + 2
}
func fnPolynom2(x float64) float64 {
	return 1*x*x - 2*x + 3
}
func fnPolynom3(x float64) float64 {
	return x*x*x - 2*x*x + 3*x - 4
}
func fnRungeFunc(x float64) float64 {
	return 1 / (1 + x*x)
}

func TestLinearRegression(t *testing.T) {
	x0 := 6.666
	X := vector.Range(0, 10, 100)
	regressionFunc := LinearRegression

	// tests the regression function and returns deviation for x0
	regressionError := func(fn Func) (dev float64) {
		Y := X.Map(func(x float64) float64 {
			// add random perturbation (±3%)
			return fn(x) * (1 + 2*(rand.Float64()-0.5)*0.03)
		})
		fnReg := regressionFunc(X, Y)
		return math.Abs(fnReg(x0)/fn(x0) - 1)
	}

	assert.InDelta(t, 0.00, regressionError(fnLine), 0.01)
	assert.InDelta(t, 0.16, regressionError(fnPolynom2), 0.01)
	assert.InDelta(t, 0.42, regressionError(fnPolynom3), 0.01)
	assert.InDelta(t, 1.18, regressionError(fnRungeFunc), 0.01)
}

func TestNewInterpolationRegression(t *testing.T) {
	x0 := 6.666
	X := vector.Range(0, 10, 100)
	regressionFunc := NewInterpolationRegression(0, 0)

	// tests the regression function and returns deviation for x0
	regressionError := func(fn Func) (dev float64) {
		Y := X.Map(func(x float64) float64 {
			// add random perturbation (±3%)
			return fn(x) * (1 + 2*(rand.Float64()-0.5)*0.03)
		})
		fnReg := regressionFunc(X, Y)
		return math.Abs(fnReg(x0)/fn(x0) - 1)
	}

	assert.InDelta(t, 0.00, regressionError(fnLine), 0.01)
	assert.InDelta(t, 0.00, regressionError(fnPolynom2), 0.01)
	assert.InDelta(t, 0.00, regressionError(fnPolynom3), 0.01)
	assert.InDelta(t, 0.00, regressionError(fnRungeFunc), 0.01)
}
