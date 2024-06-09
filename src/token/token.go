package token

type Type string

type Token struct {
    Type    Type
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"
    IDENT   = "IDENT"
    INT     = "INT"
    SQUOTE  = "'"
    DQUOTE  = "\""
    ASSIGN  = "="
    PLUS    = "+"
)
