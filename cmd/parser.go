package main

type (
	PrefixParseFunc  func()
	InfixParseFunc   func()
	PostfixParseFunc func()
)

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

type Parser struct {
	lexer    *Lexer
	vm       *VM
	current  *Token
	previous *Token

	prefixParseFuncs  map[int]PrefixParseFunc
	infixParseFuncs   map[int]InfixParseFunc
	postfixParseFuncs map[int]PostfixParseFunc
}

func NewParser(lexer *Lexer) *Parser {
	p := &Parser{
		lexer:             lexer,
		prefixParseFuncs:  map[int]PrefixParseFunc{},
		infixParseFuncs:   map[int]InfixParseFunc{},
		postfixParseFuncs: map[int]PostfixParseFunc{},
	}

	p.registerPrefix(Nil, p.parseNil)
	p.registerPrefix(Number, p.parseNumber)
	p.registerPrefix(True, p.parseBool)
	p.registerPrefix(False, p.parseBool)
	p.registerPrefix(String, p.parseString)
	p.registerPrefix(Not, p.parsePrefixExpr)
	p.registerPrefix(Hash, p.parsePrefixExpr)
	p.registerPrefix(Minus, p.parsePrefixExpr)
	p.registerPrefix(Tilde, p.parsePrefixExpr)

	p.registerInfix(Or, p.parseInfixExpr)
	p.registerInfix(And, p.parseInfixExpr)
	p.registerInfix(Lt, p.parseInfixExpr)
	p.registerInfix(Lte, p.parseInfixExpr)
	p.registerInfix(Gt, p.parseInfixExpr)
	p.registerInfix(Gte, p.parseInfixExpr)
	p.registerInfix(EqEq, p.parseInfixExpr)
	p.registerInfix(Neq, p.parseInfixExpr)
	p.registerInfix(Pipe, p.parseInfixExpr)
	p.registerInfix(Amp, p.parseInfixExpr)
	p.registerInfix(LtLt, p.parseInfixExpr)
	p.registerInfix(GtGt, p.parseInfixExpr)
	p.registerInfix(DotDot, p.parseInfixExpr)
	p.registerInfix(Plus, p.parseInfixExpr)
	p.registerInfix(Minus, p.parseInfixExpr)
	p.registerInfix(Slash, p.parseInfixExpr)
	p.registerInfix(DoubleSlash, p.parseInfixExpr)
	p.registerInfix(Percent, p.parseInfixExpr)
	p.registerInfix(Caret, p.parseInfixExpr)

	return p
}

func (p *Parser) parseNil() {
	p.vm.Write(OpNil)
}

func (p *Parser) parseNumber() {
	// p.vm.Write(OpNumber)
}

func (p *Parser) parseBool() {
	if p.current.kind == True {
		p.vm.Write(OpTrue)
		return
	}

	p.vm.Write(OpFalse)
}

func (p *Parser) parseString() {}

func (p *Parser) parsePrefixExpr() {}

func (p *Parser) parseInfixExpr() {}

func (p *Parser) registerPrefix(token int, fn PrefixParseFunc) {
	p.prefixParseFuncs[token] = fn
}

func (p *Parser) registerInfix(token int, fn InfixParseFunc) {
	p.infixParseFuncs[token] = fn
}

func (p *Parser) registerPostfix(token int, fn PostfixParseFunc) {
	p.postfixParseFuncs[token] = fn
}
