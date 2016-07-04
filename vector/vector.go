package vector

import (
	"fmt"
	"math"
)

type Vector []float64

// NewVector creates a Vector with default value.
func NewVector(size int, defVal float64) Vector {
	v := make(Vector, size)
	if defVal != 0 {
		for i, _ := range v {
			v[i] = defVal
		}
	}
	return v
}

// Range creates a Vector containing a range of elements.
func Range(start, end float64, size int) Vector {
	v := make(Vector, size)
	n := size - 1
	for i := 0; i < n; i++ {
		v[i] = start + float64(i)/float64(n)*(end-start)
	}
	v[n] = end
	return v
}

// Size returns the length of the vector.
func (v Vector) Size() int {
	return len(v)
}

// String returns vector as string.
func (v Vector) String() string {
	return fmt.Sprintf("%v", []float64(v))
}

// Clone creates a copy of current vector.
func (v Vector) Clone() Vector {
	v1 := make(Vector, len(v))
	copy(v1, v)
	return v1
}

// Push adds an element to the end of the vector.
func (v Vector) Push(x float64) {
	v = append(v, x)
}

// Pop removes an element from the end of the vector.
func (v Vector) Pop() (x float64) {
	x, v = v[len(v)-1], v[:len(v)-1]
	return
}

// Shift adds an element to the beginning of the vector.
func (v Vector) Shift() (x float64) {
	x, v = v[0], v[1:]
	return
}

// Unshift removes an element from the beginning of the vector.
func (v Vector) Unshift(x float64) {
	v = append(Vector{x}, v...)
}

// Map applies the callback function to the elements of the vector and returns new Vector with results.
func (v Vector) Map(fn func(float64) float64) Vector {
	v1 := make(Vector, len(v))
	for i, x := range v {
		v1[i] = fn(x)
	}
	return v1
}

// Add adds a and b element-wise, placing the result in the receiver.
// Add will panic if the two vectors do not have the same length.
func (v Vector) Add(v1 Vector) {
	if len(v) != len(v1) {
		panic("Enexpected number of elements")
	}
	for i, x := range v1 {
		v[i] += x
	}
}

// Sum returns the sum of values
func (v Vector) Sum() (sum float64) {
	for _, x := range v {
		sum += x
	}
	return
}

const nan = math.NaN()

// Avg returns the average value of vector values
func (v Vector) Avg() float64 {
	if n := len(v); n > 0 {
		return v.Sum() / n
	}
	return nan
}

// StdDev returns the population standard deviation
func (v Vector) StdDev() float64 {
	if n := len(v); n > 0 {
		sx, sx2 := 0., 0.
		for _, x := range v {
			sx += x
			sx2 += x * x
		}
		mx := sx / n
		return math.Sqrt(sx2/n - mx*mx)
	}
	return nan
}

// LinearRegression returns regression
func (v Vector) LinearRegression(k, b, R2 float64) {
	n := len(v)
	sxy := 0.0
	sx, sx2 := 0.0, 0.0
	sy, sy2 := 0.0, 0.0

	for x, y := range v {
		sx += x
		sy += y
		sx2 += x * x
		sy2 += y * y
		sxy += x * y
	}
	k = (sxy*n - sx*sy) / (sx2*n - sx*sx)
	b = (sy - k*sx) / n
	R2 = (n*sxy - sx*sy) / math.Sqrt((n*sx2-sx*sx)*(n*sy2-sy*sy))
	return
}

//
func LinearRegression(X, Y Vector) (k, b, R2 float64) {
	n := len(X)
	if len(Y) != n {
		panic("Enexpected number of elements")
	}
	sxy := 0.0
	sx, sx2 := 0.0, 0.0
	sy, sy2 := 0.0, 0.0

	for i, x := range X {
		y := Y[i]
		sx += x
		sy += y
		sx2 += x * x
		sy2 += y * y
		sxy += x * y
	}
	k = (sxy*n - sx*sy) / (sx2*n - sx*sx)
	b = (sy - k*sx) / n
	R2 = (n*sxy - sx*sy) / math.Sqrt((n*sx2-sx*sx)*(n*sy2-sy*sy))
	return
}
