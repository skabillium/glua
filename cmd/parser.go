package main

import "fmt"

// Operator precedence
const (
	PrecNone = iota
	PrecOr
	PrecAnd
	PrecComparison // < > <= >= ~= ==
	PrecBitOr
	PrecBitNot
	PrecBitAnd
	PrecBitShift // << >>
	PrecConcat
	PrecTerm   // + -
	PrecFactor // * / // %
	PrecUnary  // not # - ~
	PrecExp    // ^
	PrecCall
)

type PrecendenceEntry struct {
	Prefix     func(p *Parser)
	Infix      func(p *Parser)
	Precedence int
}

var PrecedenceRules = map[int]PrecendenceEntry{
	// PrecNone for literals
	Nil:   {Prefix: literal, Infix: nil, Precedence: PrecNone},
	False: {Prefix: literal, Infix: nil, Precedence: PrecNone},
	True:  {Prefix: literal, Infix: nil, Precedence: PrecNone},
	// PrecOr
	Or: {Prefix: nil, Infix: or, Precedence: PrecOr},
	// PrecAnd
	And: {Prefix: nil, Infix: and, Precedence: PrecOr},
	// PrecComparison
	Lt:   {Prefix: nil, Infix: binary, Precedence: PrecComparison},
	Gt:   {Prefix: nil, Infix: binary, Precedence: PrecComparison},
	Lte:  {Prefix: nil, Infix: binary, Precedence: PrecComparison},
	Gte:  {Prefix: nil, Infix: binary, Precedence: PrecComparison},
	Neq:  {Prefix: nil, Infix: binary, Precedence: PrecComparison},
	EqEq: {Prefix: nil, Infix: binary, Precedence: PrecComparison},
	// PrecBitOr
	Pipe: {Prefix: nil, Infix: binary, Precedence: PrecBitOr},
	// PrecBitNot
	Tilde: {Prefix: unary, Infix: nil, Precedence: PrecBitNot},
	// PrecBitAnd
	Amp: {Prefix: nil, Infix: binary, Precedence: PrecBitAnd},
	// PrecBitShift
	LtLt: {Prefix: nil, Infix: binary, Precedence: PrecBitShift},
	GtGt: {Prefix: nil, Infix: binary, Precedence: PrecBitShift},
	// PrecConcat
	DotDot: {Prefix: nil, Infix: binary, Precedence: PrecConcat},
	// PrecTerm
	Minus: {Prefix: unary, Infix: binary, Precedence: PrecTerm},
	Plus:  {Prefix: nil, Infix: binary, Precedence: PrecTerm},
	// PrecFactor
	Slash:       {Prefix: nil, Infix: binary, Precedence: PrecFactor},
	Star:        {Prefix: nil, Infix: binary, Precedence: PrecFactor},
	DoubleSlash: {Prefix: nil, Infix: binary, Precedence: PrecFactor},
	Percent:     {Prefix: nil, Infix: binary, Precedence: PrecFactor},
	// PrecUnary
	Not:  {Prefix: unary, Infix: nil, Precedence: PrecUnary},
	Hash: {Prefix: unary, Infix: nil, Precedence: PrecUnary},
	// PrecExp
	Caret: {Prefix: nil, Infix: binary, Precedence: PrecExp},
	// PrecCall
	LParen: {Prefix: grouping, Infix: call, Precedence: PrecCall},
}

type Parser struct {
	lexer *Lexer
	// vm       *VM
	current  *Token
	previous *Token
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) parseExpression() {
	p.parsePrecedence(PrecNone)
}

func (p *Parser) parsePrecedence(prec int) {
	p.advance()
	rule, found := PrecedenceRules[p.previous.kind]
	if !found || rule.Prefix == nil {
		panic(fmt.Sprintf("Rule not found for token %d\n", p.previous.kind))
	}

	for rule.Precedence <= PrecedenceRules[p.current.kind].Precedence {
		p.advance()
		rule := PrecedenceRules[p.previous.kind]
		rule.Infix(p)
	}
}

func (p *Parser) Compile() error {
	for !p.match(Eof) {
		p.parseExpression()
	}

	return nil
}

func binary(p *Parser) {}

func unary(p *Parser) {}

func grouping(p *Parser) {}

func call(p *Parser) {}

func or(p *Parser) {}

func and(p *Parser) {}

func literal(p *Parser) {}

func (p *Parser) advance() *Token {
	p.previous = p.current
	token, err := p.lexer.readToken()
	if err != nil {
		panic(err) // TODO: Handle errors more gracefully
	}

	return token
}

func (p *Parser) match(kind int) bool {
	if p.current.kind != kind {
		return false
	}
	p.advance()
	return true
}
