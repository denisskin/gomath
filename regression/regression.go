package regression

// Func is abstract
type Func func(float64) float64

type RegressionFunc func(x, y []float64) Func

//type FuncN func([]float64) float64
//
//type RegressionFuncN func(x [][]float64, y []float64) FuncN
