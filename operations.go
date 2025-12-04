package main

import (
	"fmt"
	"math"
)

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func sqrt(a float64) float64 {
	return math.Sqrt(a)
}

func mod(a, b float64) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return int(a) % int(b), nil
}

func square(a float64) float64 {
	return a * a
}

func abs(a float64) float64 {
	return math.Abs(a)
}
