package services

import (
	"math"
	"testing"
)

func assertEvalOutput(t *testing.T, expression string, expectedResult float64, expctedError error) {
	result, err := Eval(expression)
	if err != expctedError || math.Abs(result-expectedResult) > 1e-11 {
		t.Errorf(
			`Eval("%v") = (%v, %v), expected (%v, %v)`,
			expression,
			result, err,
			expectedResult, expctedError,
		)
	}

}

func TestEvalAddition(t *testing.T) {
	assertEvalOutput(t, "1+2", 3, nil)
}

func TestEvalSubtraction(t *testing.T) {
	assertEvalOutput(t, "1-2", -1, nil)
}

func TestEvalMultiplication(t *testing.T) {
	assertEvalOutput(t, "3*2", 6, nil)
}

func TestEvalDivision(t *testing.T) {
	assertEvalOutput(t, "3/2", 1.5, nil)
}

func TestEvalExponention(t *testing.T) {
	assertEvalOutput(t, "3^2", 9, nil)
}

func TestEvalSquareRoot(t *testing.T) {
	assertEvalOutput(t, "sqrt9", 3, nil)
}

func TestEvalSquareRootPower(t *testing.T) {
	assertEvalOutput(t, "sqrt9^4", 81, nil)
}

func TestEvalPercentage(t *testing.T) {
	assertEvalOutput(t, "61%", .61, nil)
}

func TestEvalMultiple(t *testing.T) {
	assertEvalOutput(t, "3+4-9*2/3^2+sqrt900%", 8, nil)
}

func TestEvalParens(t *testing.T) {
	assertEvalOutput(t, "(3+2)*(2-3)", -5, nil)
}

func TestEvalNestedParens(t *testing.T) {
	assertEvalOutput(t, "(3+(-40%))*(sqrt(81)-3)", 2.6*6, nil)
}
