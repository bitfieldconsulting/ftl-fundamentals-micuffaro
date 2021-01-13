package calculator

import (
	"errors"
)

// Add adds two integers
func Add(a, b int) int {
	return a + b
}

// Subtract subtracts two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply multiplies two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide adds two integers
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a / b, nil
}
