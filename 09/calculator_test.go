package calculator_test

import (
	"calculator"
	"testing"
)

func TestSubtract(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 0},
		{a: 2, b: 2, want: 0},
		{a: 6, b: 7, want: -1},
		{a: 0, b: 0, want: 0},
		{a: -1, b: -1, want: 0},
		{a: 3, b: -8, want: 11},
		{a: -6, b: 7, want: -13},
		{a: 0, b: -4, want: 4},
	}
	for _, testCase := range testCases {
		got := calculator.Subtract(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Add(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 1},
		{a: 2, b: 2, want: 4},
		{a: 6, b: 7, want: 42},
		{a: 0, b: 0, want: 0},
		{a: -1, b: -1, want: 1},
		{a: 3, b: -8, want: -24},
		{a: -6, b: 7, want: -42},
		{a: 0, b: -4, want: 0},
	}
	for _, testCase := range testCases {
		got := calculator.Multiply(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Multiply(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 1},
		{a: 2, b: 2, want: 1},
		{a: 6, b: 7, want: 0},
		{a: -1, b: -1, want: 1},
		{a: 3, b: -8, want: 0},
		{a: -6, b: 7, want: 0},
	}
	for _, testCase := range testCases {
		got := calculator.Divide(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Divide(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}
