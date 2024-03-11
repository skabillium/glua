package main

import (
	"errors"
	"unicode"
)

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

	Name
	String
	Number
	Line
	Semi
	Eof

	Plus
	Minus
	Star
	Slash
	Percent
	Caret
	Amp
	Tilde
	Pipe
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

func (lex *Lexer) NextToken() {
}

func (lex *Lexer) readToken() (*Token, error) {
	lex.skipWhitespace()
	if lex.position >= len(lex.source) {
		return NewToken(Eof, len(lex.source), 1), nil
	}

	c := lex.currentChar()
	switch c {
	}

	if isName(c) {
		return lex.readName(), nil
	} else if unicode.IsDigit(c) {
		return lex.readNumber(), nil
	}

	return nil, errors.New("unexpected character")
}

func (lex *Lexer) skipWhitespace() {
	if lex.isAtEnd() {
		return
	}

	c := lex.currentChar()

	for c == ' ' || c == '\t' {
		lex.position++
		if lex.position >= len(lex.source) {
			break
		}

		c = lex.currentChar()
	}
}

func (lex *Lexer) readName() *Token {
	start := lex.position
	length := 1

	lex.position++
	for !lex.isAtEnd() && isAlphanumeric(lex.currentChar()) {
		length++
		lex.position++
	}

	name := lex.source[start : start+length]
	kind, ok := keywords[name]
	if !ok {
		return NewToken(Name, start, length)
	}

	return NewToken(kind, start, length)
}

func (lex *Lexer) readNumber() *Token {
	start := lex.position
	length := 1

	lex.position++
	// TODO: Support decimals
	for !lex.isAtEnd() && unicode.IsDigit(lex.currentChar()) {
		lex.position++
		length++
	}

	return NewToken(Number, start, length)
}

func (lex *Lexer) currentChar() rune {
	return rune(lex.source[lex.position])
}

// func (lex *Lexer) peek() rune {}

func (lex *Lexer) isAtEnd() bool {
	return lex.position >= len(lex.source)
}

func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_'
}

func isName(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'

}
