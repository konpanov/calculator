package services

import (
	"fmt"
	"math"
)

// Parser parses and evaluates a single arithmetic expression.
type Parser struct {
	lex *Lexer
	cur Token
}

var ErrorDevisionByZero = fmt.Errorf("Cannot devide by zero")
var ErrorNegativeSqrt = fmt.Errorf("Cannot take square root of a negative value")

// NewParser creates a parser for the given expression string.
func NewParser(expr string) (*Parser, error) {
	p := &Parser{lex: newLexer(expr)}
	if err := p.advance(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Parser) advance() error {
	t, err := p.lex.next()
	if err != nil {
		return err
	}
	p.cur = t
	return nil
}

func (p *Parser) expect(k TokenKind, what string) error {
	if p.cur.kind != k {
		return fmt.Errorf("expected %s, got %q", what, p.cur.text)
	}
	return p.advance()
}

// Eval parses and evaluates the whole expression, erroring on trailing input.
func (p *Parser) Eval() (float64, error) {
	v, err := p.parseExpr()
	if err != nil {
		return 0, err
	}
	if p.cur.kind != tokEOF {
		return 0, fmt.Errorf("unexpected trailing token %q", p.cur.text)
	}
	return v, nil
}

// expr := term (("+"|"-") term)*
func (p *Parser) parseExpr() (float64, error) {
	v, err := p.parseTerm()
	if err != nil {
		return 0, err
	}
	for p.cur.kind == tokPlus || p.cur.kind == tokMinus {
		op := p.cur.kind
		if err := p.advance(); err != nil {
			return 0, err
		}
		rhs, err := p.parseTerm()
		if err != nil {
			return 0, err
		}
		if op == tokPlus {
			v += rhs
		} else {
			v -= rhs
		}
	}
	return v, nil
}

// term := unary (("*"|"/") unary)*
func (p *Parser) parseTerm() (float64, error) {
	v, err := p.parseUnary()
	if err != nil {
		return 0, err
	}
	for p.cur.kind == tokStar || p.cur.kind == tokSlash {
		op := p.cur.kind
		if err := p.advance(); err != nil {
			return 0, err
		}
		rhs, err := p.parseUnary()
		if err != nil {
			return 0, err
		}
		switch op {
		case tokStar:
			v *= rhs
		case tokSlash:
			if rhs == 0 {
				return 0, ErrorDevisionByZero
			}
			v /= rhs
		}
	}
	return v, nil
}

// unary := ("-"|"+")? power
func (p *Parser) parseUnary() (float64, error) {
	if p.cur.kind == tokSqrt {
		if err := p.advance(); err != nil {
			return 0, err
		}
		v, err := p.parseUnary()
		if v < 0 {
			return 0, ErrorNegativeSqrt
		}
		return math.Sqrt(v), err
	}
	if p.cur.kind == tokMinus {
		if err := p.advance(); err != nil {
			return 0, err
		}
		v, err := p.parseUnary()
		return -v, err
	}
	if p.cur.kind == tokPlus {
		if err := p.advance(); err != nil {
			return 0, err
		}
		return p.parseUnary()
	}
	return p.parsePecentage()
}

// percentage := power ("%")?
func (p *Parser) parsePecentage() (float64, error) {
	base, err := p.parsePower()
	if err != nil {
		return 0, err
	}
	if p.cur.kind == tokPercent {
		if err := p.advance(); err != nil {
			return 0, err
		}
		return base / 100, nil
	}
	return base, nil
}

// power := sqrt ("^" unary)?   (right-associative)
func (p *Parser) parsePower() (float64, error) {
	base, err := p.parsePrimary()
	if err != nil {
		return 0, err
	}
	if p.cur.kind == tokCaret {
		if err := p.advance(); err != nil {
			return 0, err
		}
		exp, err := p.parseUnary() // right-assoc + allows "2^-2"
		if err != nil {
			return 0, err
		}
		return math.Pow(base, exp), nil
	}
	return base, nil
}

// primary := NUMBER | "(" expr ")"
func (p *Parser) parsePrimary() (float64, error) {
	switch p.cur.kind {
	case tokNumber:
		v := p.cur.num
		if err := p.advance(); err != nil {
			return 0, err
		}
		return v, nil

	case tokLParen:
		if err := p.advance(); err != nil {
			return 0, err
		}
		v, err := p.parseExpr()
		if err != nil {
			return 0, err
		}
		if err := p.expect(tokRParen, "')'"); err != nil {
			return 0, err
		}
		return v, nil

	default:
		return 0, fmt.Errorf("unexpected token %q", p.cur.text)
	}
}

func Eval(expr string) (float64, error) {
	p, err := NewParser(expr)
	if err != nil {
		return 0, err
	}
	return p.Eval()
}
