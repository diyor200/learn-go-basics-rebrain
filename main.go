package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, "test.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			fmt.Printf("return statement found on line %d:\n\t", fileSet.Position(ret.Pos()).Line)
		}
		return true
	})

	fmt.Println(Analys("test.go"))

}

type AnalysResult struct {
	DeclCount    int
	CallCount    int
	AssignCount  int
	ImportsCount int
}

func Analys(filepath string) (*AnalysResult, error) {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, filepath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var result AnalysResult

	// count imports
	for i := 0; i < len(node.Imports); i++ {
		result.ImportsCount++
	}

	ast.Inspect(node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.AssignStmt:
			result.AssignCount++
		case *ast.CallExpr:
			result.CallCount++
		case *ast.GenDecl:
			if t.Tok != token.IMPORT {
				result.DeclCount++
			}
		}

		return true
	})

	return &result, nil
}
