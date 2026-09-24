package services

import (
	"fmt"
	"testing"
)

func assertLexerOutput(t *testing.T, input string, expectedTokens []Token) {
	index := 0
	lexer := newLexer(input)

	asExpected := true
	errorMessage := "\n"
	for ; ; index++ {
		token, err := lexer.next()
		if err != nil {
			asExpected = false
			errorMessage += fmt.Sprintf("Token %v - Unexpected lexer error: %v", index, err)
		} else if index >= len(expectedTokens) {
			asExpected = false
			errorMessage += fmt.Sprintf("Token %v = %v, out of expected tokens\n", index, token)
		} else if token != expectedTokens[index] {
			asExpected = false
			errorMessage += fmt.Sprintf("Token %v = %v, expected %v\n", index, token, expectedTokens[index])
		} else {
			errorMessage += fmt.Sprintf("Token %v = %v\n", index, token)
		}
		if token.kind == tokEOF {
			index++
			break
		}
	}
	for ; index < len(expectedTokens); index++ {
		asExpected = false
		errorMessage += fmt.Sprintf("Token %v = nil, expected %v\n", index, expectedTokens[index])
	}

	if !asExpected {
		t.Error(errorMessage)
	}
}

func TestTokenizeNumber(t *testing.T) {
	assertLexerOutput(t, "12300", []Token{
		{tokNumber, "12300", 12300},
		{tokEOF, "", 0},
	})
}

func TestTokenizeDecimalNumber(t *testing.T) {
	assertLexerOutput(t, "12300.123", []Token{
		{tokNumber, "12300.123", 12300.123},
		{tokEOF, "", 0},
	})
}

func TestTokenizeScientificNotationNumber(t *testing.T) {
	assertLexerOutput(t, "1e10", []Token{
		{tokNumber, "1e10", 1e10},
		{tokEOF, "", 0},
	})
	assertLexerOutput(t, "1.5e-3", []Token{
		{tokNumber, "1.5e-3", 1.5e-3},
		{tokEOF, "", 0},
	})
}

func TestTokenizeAddition(t *testing.T) {
	assertLexerOutput(t, "1+2", []Token{
		{tokNumber, "1", 1},
		{tokPlus, "+", 0},
		{tokNumber, "2", 2},
		{tokEOF, "", 0},
	})
}

func TestTokenizeSubtruction(t *testing.T) {
	assertLexerOutput(t, "1-2", []Token{
		{tokNumber, "1", 1},
		{tokMinus, "-", 0},
		{tokNumber, "2", 2},
		{tokEOF, "", 0},
	})
}

func TestTokenizeMultiplication(t *testing.T) {
	assertLexerOutput(t, "1*2", []Token{
		{tokNumber, "1", 1},
		{tokStar, "*", 0},
		{tokNumber, "2", 2},
		{tokEOF, "", 0},
	})
}

func TestTokenizeDivision(t *testing.T) {
	assertLexerOutput(t, "1/2", []Token{
		{tokNumber, "1", 1},
		{tokSlash, "/", 0},
		{tokNumber, "2", 2},
		{tokEOF, "", 0},
	})
}

func TestTokenizeExponentiation(t *testing.T) {
	assertLexerOutput(t, "3^4", []Token{
		{tokNumber, "3", 3},
		{tokCaret, "^", 0},
		{tokNumber, "4", 4},
		{tokEOF, "", 0},
	})
}

func TestTokenizeParens(t *testing.T) {
	assertLexerOutput(t, "2*(4+5)", []Token{
		{tokNumber, "2", 2},
		{tokStar, "*", 0},
		{tokLParen, "(", 0},
		{tokNumber, "4", 4},
		{tokPlus, "+", 0},
		{tokNumber, "5", 5},
		{tokRParen, ")", 0},
		{tokEOF, "", 0},
	})
}

func TestTokenizePercent(t *testing.T) {
	assertLexerOutput(t, "3%", []Token{
		{tokNumber, "3", 3},
		{tokPercent, "%", 0},
		{tokEOF, "", 0},
	})
}

func TestTokenizeSquareRoot(t *testing.T) {
	assertLexerOutput(t, "sqrt3", []Token{
		{tokSqrt, "sqrt", 0},
		{tokNumber, "3", 3},
		{tokEOF, "", 0},
	})
	assertLexerOutput(t, "√8", []Token{
		{tokSqrt, "√", 0},
		{tokNumber, "8", 8},
		{tokEOF, "", 0},
	})
}
