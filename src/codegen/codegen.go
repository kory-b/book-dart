package codegen

import (
	"bytes"
	"fmt"

	"book-dart/src/ast"
)

func GenerateCode(node ast.Node) string {
    var out bytes.Buffer
    generate(node, &out)
    return out.String()
}

func generate(node ast.Node, out *bytes.Buffer) {
    switch node := node.(type) {
    case *ast.Program:
        for _, stmt := range node.Statements {
            generate(stmt, out)
        }
    case *ast.AssignStatement:
        out.WriteString(fmt.Sprintf("%s = ", node.Name.Value))
        generate(node.Value, out)
        out.WriteString(";\n")
    case *ast.Identifier:
        out.WriteString(node.Value)
    case *ast.IntegerLiteral:
        out.WriteString(fmt.Sprintf("%d", node.Value))
    default:
        // Handle other node types as needed
    }
}
