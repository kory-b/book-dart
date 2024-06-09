package parser

import (
    "fmt"
    "book-dart/pkg/ast"
    "book-dart/pkg/lexer"
    "book-dart/pkg/token"
)

type Parser struct {
    l      *lexer.Lexer
    curTok token.Token
    peekTok token.Token
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}
    p.nextToken()
    p.nextToken()
    return p
}

func (p *Parser) nextToken() {
    p.curTok = p.peekTok
    p.peekTok = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    program.Statements = []ast.Statement{}

    for p.curTok.Type != token.EOF {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }

    return program
}

func (p *Parser) parseStatement() ast.Statement {
    switch p.curTok.Type {
    case token.IDENT:
        return p.parseAssignStatement()
    default:
        return nil
    }
}

func (p *Parser) parseAssignStatement() *ast.AssignStatement {
    stmt := &ast.AssignStatement{Token: p.curTok, Name: &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}}

    if !p.expectPeek(token.ASSIGN) {
        return nil
    }

    p.nextToken()
    stmt.Value = p.parseExpression()

    return stmt
}

func (p *Parser) parseExpression() ast.Expression {
    switch p.curTok.Type {
    case token.IDENT:
        return &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}
    case token.INT:
        return &ast.IntegerLiteral{Token: p.curTok, Value: p.parseIntegerLiteral()}
    default:
        return nil
    }
}

func (p *Parser) parseIntegerLiteral() int64 {
    var value int64
    fmt.Sscanf(p.curTok.Literal, "%d", &value)
    return value
}

func (p *Parser) expectPeek(t token.Type) bool {
    if p.peekTok.Type == t {
        p.nextToken()
        return true
    }
    return false
}
