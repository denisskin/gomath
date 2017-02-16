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
func (v *Vector) Push(x float64) {
	*v = append(*v, x)
}

// Pop removes an element from the end of the vector.
func (v Vector) Pop() (x float64) {
	x, v = v[len(v)-1], v[:len(v)-1]
	return
}

// Shift adds an element to the beginning of the vector.
func (v *Vector) Shift() (x float64) {
	pv := *v
	x, *v = pv[0], pv[1:]
	return
}

// Unshift removes an element from the beginning of the vector.
func (v *Vector) Unshift(x float64) {
	*v = append(Vector{x}, *v...)
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

// Avg returns the average value of vector values
func (v Vector) Avg() float64 {
	if n := len(v); n > 0 {
		return v.Sum() / float64(n)
	}
	return math.NaN()
}

// StdDev returns the population standard deviation
func (v Vector) StdDev() float64 {
	if n := len(v); n > 0 {
		sx, sx2 := 0., 0.
		for _, x := range v {
			sx += x
			sx2 += x * x
		}
		mx := sx / float64(n)
		return math.Sqrt(sx2/float64(n) - mx*mx)
	}
	return math.NaN()
}

// Min returns the minimum value of vector values
func (v Vector) Min() (min float64) {
	min = math.NaN()
	for i, v := range v {
		if i == 0 || v < min {
			min = v
		}
	}
	return
}

// Max returns the maximum value of vector values
func (v Vector) Max() (max float64) {
	max = math.NaN()
	for i, v := range v {
		if i == 0 || v > max {
			max = v
		}
	}
	return
}
