package main

import (
	"errors"
	"unicode"
)

// TODO: Tokenize '#'

const (
	// Keywords
	And = iota
	Break
	Do
	Else
	ElseIf
	End
	False
	For
	Function
	Goto
	If
	In
	Local
	Nil
	Not
	Or
	Repeat
	Return
	Then
	True
	Until
	While

	Comment
	Name
	String
	Number
	Line
	Semi

	Plus
	Minus
	Star
	Slash
	Percent
	Caret
	Amp
	Tilde
	Pipe
	Hash
	LtLt
	GtGt
	DoubleSlash
	EqEq
	Neq
	Lte
	Gte
	Lt
	Gt
	Eq
	LParen
	RParen
	LCurly
	RCurly
	LBracket
	RBracket
	DoubleCol
	Col
	Comma
	Dot
	DotDot
	DotDotDot
	Eof
)

type Token struct {
	kind   int
	start  int
	length int
}

func NewToken(kind int, start int, length int) *Token {
	return &Token{kind: kind, start: start, length: length}
}

type Lexer struct {
	source   string
	position int
}

var keywords = map[string]int{
	"and":      And,
	"break":    Break,
	"do":       Do,
	"else":     Else,
	"elseif":   ElseIf,
	"end":      End,
	"false":    False,
	"for":      For,
	"function": Function,
	"goto":     Goto,
	"if":       If,
	"in":       In,
	"local":    Local,
	"nil":      Nil,
	"not":      Not,
	"or":       Or,
	"repeat":   Repeat,
	"return":   Return,
	"then":     Then,
	"true":     True,
	"until":    Until,
	"while":    While,
}

func NewLexer(source string) *Lexer {
	return &Lexer{source: source, position: 0}
}

func (lex *Lexer) readToken() (*Token, error) {
	lex.skipWhitespace()
	if lex.position >= len(lex.source) {
		return NewToken(Eof, len(lex.source), 1), nil
	}

	c := lex.currentChar()
	switch c {
	case '\n':
		return lex.makeToken(Line, 1), nil
	case ';':
		return lex.makeToken(Semi, 1), nil
	case '+':
		return lex.makeToken(Plus, 1), nil
	case '*':
		return lex.makeToken(Star, 1), nil
	case '%':
		return lex.makeToken(Percent, 1), nil
	case '^':
		return lex.makeToken(Caret, 1), nil
	case '&':
		return lex.makeToken(Amp, 1), nil
	case '|':
		return lex.makeToken(Pipe, 1), nil
	case '(':
		return lex.makeToken(LParen, 1), nil
	case ')':
		return lex.makeToken(RParen, 1), nil
	case '{':
		return lex.makeToken(LCurly, 1), nil
	case '}':
		return lex.makeToken(RCurly, 1), nil
	case '[':
		return lex.makeToken(LBracket, 1), nil
	case ']':
		return lex.makeToken(RBracket, 1), nil
	case ',':
		return lex.makeToken(Comma, 1), nil
	case '#':
		return lex.makeToken(Hash, 1), nil
	case '"', '\'':
		return lex.readString(c), nil
	// Multi-character tokens
	case '-':
		if lex.match("--") {
			return lex.readComment(), nil
		}
		return lex.makeToken(Minus, 1), nil
	case '<':
		if lex.match("<<") {
			return lex.makeToken(LtLt, 2), nil
		} else if lex.match("<=") {
			return lex.makeToken(Lte, 2), nil
		}
		return lex.makeToken(Lt, 1), nil
	case '>':
		if lex.match(">>") {
			return lex.makeToken(GtGt, 2), nil
		} else if lex.match(">=") {
			return lex.makeToken(Gte, 2), nil
		}
		return lex.makeToken(Gt, 1), nil
	case '~':
		if lex.match("~=") {
			return lex.makeToken(Neq, 2), nil
		}
		return lex.makeToken(Tilde, 1), nil
	case '=':
		if lex.match("==") {
			return lex.makeToken(EqEq, 2), nil
		}
		return lex.makeToken(Eq, 1), nil
	case ':':
		if lex.match("::") {
			return lex.makeToken(DoubleCol, 2), nil
		}
		return lex.makeToken(Col, 1), nil
	case '/':
		if lex.match("//") {
			return lex.makeToken(DoubleSlash, 2), nil
		}
		return lex.makeToken(Slash, 1), nil
	case '.':
		if unicode.IsDigit(lex.peekChar()) {
			return lex.readNumber(false), nil
		}

		if lex.match("...") {
			return lex.makeToken(DotDotDot, 3), nil
		} else if lex.match("..") {
			return lex.makeToken(DotDot, 2), nil
		}

		return lex.makeToken(Dot, 1), nil
	default:
		if isName(c) {
			return lex.readName(), nil
		} else if unicode.IsDigit(c) {
			return lex.readNumber(true), nil
		}
	}

	return nil, errors.New("unexpected character: " + string(c))
}

func (lex *Lexer) skipWhitespace() {
	if lex.isAtEnd() {
		return
	}

	c := lex.currentChar()

	for c == ' ' || c == '\t' {
		lex.advance()
		if lex.position >= len(lex.source) {
			break
		}

		c = lex.currentChar()
	}
}

func (lex *Lexer) readName() *Token {
	start := lex.position
	length := 1

	lex.advance()
	for !lex.isAtEnd() && isAlphanumeric(lex.currentChar()) {
		length++
		lex.advance()
	}

	name := lex.source[start : start+length]
	kind, ok := keywords[name]
	if !ok {
		return NewToken(Name, start, length)
	}

	return NewToken(kind, start, length)
}

func (lex *Lexer) readNumber(allowDecimals bool) *Token {
	start := lex.position
	length := 1

	lex.advance()
	for !lex.isAtEnd() && (unicode.IsDigit(lex.currentChar()) || allowDecimals && lex.currentChar() == '.') {
		if lex.currentChar() == '.' {
			allowDecimals = false
		}
		lex.advance()
		length++
	}

	return NewToken(Number, start, length)
}

// TODO: This could return an error
func (lex *Lexer) readString(term rune) *Token {
	start := lex.position
	length := 1

	// var buffer strings.Builder

	for !lex.isAtEnd() {
		c := lex.nextChar()
		length++
		if c == term {
			lex.advance()
			break
		}
	}

	return &Token{kind: String, start: start, length: length}
}

func (lex *Lexer) readComment() *Token {
	start := lex.position
	length := 2
	lex.advanceBy(2)

	for !lex.isAtEnd() && lex.currentChar() != '\n' {
		lex.advance()
		length++
	}

	return &Token{kind: Comment, start: start, length: length}
}

func (lex *Lexer) readOperator() *Token {
	return &Token{kind: LtLt, start: lex.position, length: 1}
}

func (lex *Lexer) currentChar() rune {
	return rune(lex.source[lex.position])
}

func (lex *Lexer) nextChar() rune {
	lex.advance()
	return lex.currentChar()
}

func (lex *Lexer) peekChar() rune {
	return rune(lex.source[lex.position+1])
}

func (lex *Lexer) advance() {
	lex.position++
}

func (lex *Lexer) advanceBy(incr int) {
	lex.position += incr
}

func (lex *Lexer) match(str string) bool {
	if lex.position+len(str) > len(lex.source) {
		return false
	}

	return str == lex.source[lex.position:lex.position+len(str)]
}

func (lex *Lexer) makeToken(kind int, length int) *Token {
	token := &Token{kind: kind, start: lex.position, length: length}
	lex.advanceBy(length)
	return token
}

func (lex *Lexer) isAtEnd() bool {
	return lex.position >= len(lex.source)
}

func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_'
}

func isName(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'

}
