package main

import "fmt"

func main() {
	source := "\"str'"

	lex := NewLexer(source)

	token, err := lex.readToken()
	if err != nil {
		panic(err)
	}

	fmt.Print(source[token.start : token.start+token.length])

	// for {
	// 	tok, err := lex.readToken()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if tok.kind == Eof {
	// 		fmt.Println("EOF")
	// 		break
	// 	}

	// 	fmt.Println("Token:", tok.kind, "Start:", tok.start, "Length:", tok.length)
	// }
}
