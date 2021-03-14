package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestParse(t *testing.T) {
	fileSet := token.NewFileSet()
	packages, _ := parser.ParseDir(fileSet, "sample", nil, parser.AllErrors)
	ast.Print(nil, packages)
}
