package codegen

import (
	"testing"

	"book-dart/src/ast"
	"book-dart/src/token"
)

func TestGenerateCode(t *testing.T) {
    program := &ast.Program{
        Statements: []ast.Statement{
            &ast.AssignStatement{
                Token: token.Token{Type: token.ASSIGN, Literal: "="},
                Name:  &ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
                Value: &ast.IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
            },
        },
    }

    code := GenerateCode(program)
    expectedCode := "x = 5;\n"

    if code != expectedCode {
        t.Fatalf("code wrong. expected=%q, got=%q", expectedCode, code)
    }
}
