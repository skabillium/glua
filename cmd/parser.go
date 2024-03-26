package main

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
)

type PrecendenceEntry struct {
	Prefix     func()
	Infix      func()
	Precedence int
}

var PrecedenceRules = map[int]PrecendenceEntry{
	LParen: {},
}

type Parser struct {
	lexer *Lexer
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) advance() *Token {
	token, err := p.lexer.readToken()
	if err != nil {
		panic(err)
	}

	return token
}

func (p *Parser) match(kind int) bool {
	token, err := p.lexer.readToken()
	if err != nil {
		panic(err)
	}

	return token.kind == kind
}
