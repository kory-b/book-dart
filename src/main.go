package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"book-dart/src/codegen"
	"book-dart/src/lexer"
	"book-dart/src/parser"
)

func main() {
    inputFile := flag.String("input", "", "Input file to compile (default: stdin)")
    flag.Parse()

    var input string

    if *inputFile != "" {
        data, err := os.ReadFile(*inputFile)
        if err != nil {
            log.Fatalf("Error reading input file: %v", err)
        }
        input = string(data)
    } else {
        data, err := io.ReadAll(os.Stdin)
        if err != nil {
            log.Fatalf("Error reading from stdin: %v", err)
        }
        input = string(data)
    }

    l := lexer.New(input)
    p := parser.New(l)
    program := p.ParseProgram()

    code := codegen.GenerateCode(program)
    
    fmt.Println(code)
}
