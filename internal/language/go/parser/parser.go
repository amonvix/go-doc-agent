package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func ParseFile(path string) (*ast.File, *token.FileSet, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}
	return node, fset, nil
}
