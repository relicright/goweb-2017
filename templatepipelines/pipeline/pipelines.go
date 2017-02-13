package pipeline

import (
"math"
"text/template"
)

func Double(x int) int {
	return x + x
}

func Square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func SqRoot(x float64) float64  {
	return math.Sqrt(x)
}

// FuncMap var used to map out what functions will be passed into the template.
var Fm2 = template.FuncMap{
	"db": Double,
	"fsq": Square,
	"fsqrt": SqRoot,
}
