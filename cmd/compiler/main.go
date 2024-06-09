package main

import (
    "fmt"
    "book-dart/pkg/lexer"
    "book-dart/pkg/parser"
)

func main() {
    input := "x = 5"
    l := lexer.New(input)
    p := parser.New(l)

    program := p.ParseProgram()
    fmt.Printf("%+v\n", program)
}
