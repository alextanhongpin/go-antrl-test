package main

import (
	"fmt"

	"github.com/alextanhongpin/go-antlr-test/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	// Setup the input stream.
	in := antlr.NewInputStream("1 + 2 * 3")

	// Create the lexer.
	lexer := parser.NewCalcLexer(in)

	// Read all tokens.
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()],
			t.GetText(),
		)
	}
}
