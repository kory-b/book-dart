package parser

import (
	"testing"

	"book-dart/src/ast"
	"book-dart/src/lexer"
)

func TestParseProgram(t *testing.T) {
    input := `x = 5`
    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    if program == nil {
        t.Fatalf("ParseProgram() returned nil")
    }

    if len(program.Statements) != 1 {
        t.Fatalf("program.Statements does not contain 1 statement. got=%d",
            len(program.Statements))
    }

    stmt := program.Statements[0]
    assignStmt, ok := stmt.(*ast.AssignStatement)
    if !ok {
        t.Fatalf("stmt is not *ast.AssignStatement. got=%T", stmt)
    }

    if assignStmt.Name == nil {
        t.Fatalf("assignStmt.Name is nil")
    }

    if assignStmt.Name.Value != "x" {
        t.Fatalf("assignStmt.Name.Value not 'x'. got=%s", assignStmt.Name.Value)
    }

    if assignStmt.Value == nil {
        t.Fatalf("assignStmt.Value is nil")
    }

    if assignStmt.Value.TokenLiteral() != "5" {
        t.Fatalf("assignStmt.Value.TokenLiteral not '5'. got=%s", assignStmt.Value.TokenLiteral())
    }
}
