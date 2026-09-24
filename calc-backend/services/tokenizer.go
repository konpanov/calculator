package services

import (
	"fmt"
	"strconv"
	"unicode"
)

type TokenKind int

const (
	tokNumber TokenKind = iota
	tokPlus
	tokMinus
	tokStar
	tokSlash
	tokPercent
	tokCaret
	tokLParen
	tokRParen
	tokComma
	tokSqrt
	tokEOF
)

type Token struct {
	kind TokenKind
	text string
	num  float64
}

type Lexer struct {
	input []rune
	pos   int
}

func newLexer(s string) *Lexer {
	return &Lexer{input: []rune(s)}
}

func (l *Lexer) peekRune() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

func (l *Lexer) next() (Token, error) {
	for l.pos < len(l.input) && unicode.IsSpace(l.input[l.pos]) {
		l.pos++
	}
	if l.pos >= len(l.input) {
		return Token{kind: tokEOF}, nil
	}

	c := l.input[l.pos]

	switch {
	case unicode.IsDigit(c) || (c == '.' && l.pos+1 < len(l.input) && unicode.IsDigit(l.input[l.pos+1])):
		start := l.pos
		seenDot := false
		for l.pos < len(l.input) {
			ch := l.input[l.pos]
			if unicode.IsDigit(ch) {
				l.pos++
				continue
			}
			if ch == '.' && !seenDot {
				seenDot = true
				l.pos++
				continue
			}

			// scientific notation e.g. 1e10, 1.5e-3
			if (ch == 'e' || ch == 'E') && l.pos+1 < len(l.input) {
				next := l.input[l.pos+1]
				if unicode.IsDigit(next) || ((next == '+' || next == '-') && l.pos+2 < len(l.input) && unicode.IsDigit(l.input[l.pos+2])) {
					l.pos++ // consume e/E
					if l.input[l.pos] == '+' || l.input[l.pos] == '-' {
						l.pos++
					}
					continue
				}
			}
			break
		}
		text := string(l.input[start:l.pos])
		v, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return Token{}, fmt.Errorf("invalid number %q: %w", text, err)
		}
		return Token{kind: tokNumber, text: text, num: v}, nil
	case c == '√':
		l.pos++
		return Token{kind: tokSqrt, text: "√"}, nil
	case c == 's':
		start := l.pos
		sqrtText := []rune("sqrt")
		for l.pos < len(l.input) {
			ch := l.input[l.pos]
			relPos := l.pos - start
			if relPos == len(sqrtText) || ch != sqrtText[relPos] {
				break
			}
			l.pos++
		}
		text := string(l.input[start:l.pos])
		if len(text) != len(sqrtText) {
			return Token{}, fmt.Errorf("invalid identifier %q", text)
		}
		return Token{kind: tokSqrt, text: text}, nil
	case c == '+':
		l.pos++
		return Token{kind: tokPlus, text: "+"}, nil
	case c == '-' || c == '−':
		l.pos++
		return Token{kind: tokMinus, text: "-"}, nil
	case c == '*' || c == '×':
		l.pos++
		return Token{kind: tokStar, text: "*"}, nil
	case c == '/' || c == '÷':
		l.pos++
		return Token{kind: tokSlash, text: "/"}, nil
	case c == '%':
		l.pos++
		return Token{kind: tokPercent, text: "%"}, nil
	case c == '^':
		l.pos++
		return Token{kind: tokCaret, text: "^"}, nil
	case c == '(':
		l.pos++
		return Token{kind: tokLParen, text: "("}, nil
	case c == ')':
		l.pos++
		return Token{kind: tokRParen, text: ")"}, nil
	case c == ',' || c == '.':
		l.pos++
		return Token{kind: tokComma, text: ","}, nil
	default:
		return Token{}, fmt.Errorf("unexpected character %q at position %d", c, l.pos)
	}
}
