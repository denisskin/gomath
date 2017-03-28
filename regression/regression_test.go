package regression

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"github.com/denisskin/gomath/vector"
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
func fnRunge(x float64) float64 {
	return 1 / (1 + x*x)
}

// testRegression tests regression function and returns deviation
func testRegression(regression RegressionFunc, fn Func) (dev float64) {

	X := vector.Range(0, 10, 100)

	Y := X.Map(func(x float64) float64 {
		// add random perturbation (±3%)
		return fn(x) * (1 + 2*(rand.Float64()-0.5)*0.03)
	})

	fnReg := regression(X, Y)

	x0 := X.Avg()

	return math.Abs(fnReg(x0)/fn(x0)-1)
}

func TestLinearRegression(t *testing.T) {
	assert.True(t, testRegression(LinearRegression, fnLine) < 0.01)
	assert.True(t, testRegression(LinearRegression, fnPolynom2) > 0.40)
	assert.True(t, testRegression(LinearRegression, fnPolynom3) > 1.20)
	assert.True(t, testRegression(LinearRegression, fnRunge) > 2.90)
}

func TestNewInterpolationRegression(t *testing.T) {
	interpolationRegression := NewInterpolationRegression(0, 0)

	assert.True(t, testRegression(interpolationRegression, fnLine) < 0.01)
	assert.True(t, testRegression(interpolationRegression, fnPolynom2) < 0.01)
	assert.True(t, testRegression(interpolationRegression, fnPolynom3) < 0.01)
	assert.True(t, testRegression(interpolationRegression, fnRunge) < 0.01)
}
