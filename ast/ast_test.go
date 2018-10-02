package ast

import (
	"testing"

	"github.com/jcox250/monkey/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	expectedProgram := "let myVar = anotherVar;"
	if program.String() != expectedProgram {
		t.Errorf("program.String() wrong. got=%q, expected=%q",
			program.String(), expectedProgram)
	}
}
