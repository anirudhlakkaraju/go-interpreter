package evaluator

import (
	"fmt"

	"github.com/anirudhlakkaraju/go-interpreter/interpreter/ch3/src/monkey/ast"
	"github.com/anirudhlakkaraju/go-interpreter/interpreter/ch3/src/monkey/object"
)

func Eval(node ast.Node) object.Object {

	switch node := node.(type) {
	case *ast.Program:
		fmt.Printf("Program\n")
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		fmt.Printf("ExpStmt\n")
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		fmt.Printf("Integer\n")
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
