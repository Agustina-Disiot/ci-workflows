package math

import "math"

func Add(x, y int) (res int) {
	return x + y
}

func Subtract(x, y int) (res int) {
	return x - y
}

func CircleArea(r float64) (res float64) {
	return math.Pi * r * r
}
