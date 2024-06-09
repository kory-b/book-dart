package ast

import "book-dart/src/token"

type Node interface {
    TokenLiteral() string
}

type Statement interface {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}

type Program struct {
    Statements []Statement
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    }
    return ""
}

type Identifier struct {
    Token token.Token
    Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
    return i.Token.Literal
}

type AssignStatement struct {
    Token token.Token
    Name  *Identifier
    Value Expression
}

func (as *AssignStatement) statementNode() {}
func (as *AssignStatement) TokenLiteral() string {
    return as.Token.Literal
}

type IntegerLiteral struct {
    Token token.Token
    Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
    return il.Token.Literal
}
