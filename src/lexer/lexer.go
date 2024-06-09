package lexer

import (
	"book-dart/src/token"
	"unicode"
)

type Lexer struct {
    input        string
    position     int
    readPosition int
    ch           byte
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token
    l.skipWhitespace()
    
    switch l.ch {
    case '=':
        tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
    case '+':
        tok = token.Token{Type: token.PLUS, Literal: string(l.ch)}
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type = token.IDENT
            return tok
        } else if isDigit(l.ch) {
            tok.Literal = l.readNumber()
            tok.Type = token.INT
            return tok
        } else {
            tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
        }
    }
    
    l.readChar()
    return tok
}

func (l *Lexer) skipWhitespace() {
    for unicode.IsSpace(rune(l.ch)) {
        l.readChar()
    }
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func isLetter(ch byte) bool {
    return unicode.IsLetter(rune(ch))
}

func isDigit(ch byte) bool {
    return unicode.IsDigit(rune(ch))
}