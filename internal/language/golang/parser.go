package golang

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type FileInfo struct {
	Functions []string
	Types     []string
}

func ParseFile(path string) (*FileInfo, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	info := &FileInfo{}

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			info.Functions = append(info.Functions, x.Name.Name)
		case *ast.TypeSpec:
			info.Types = append(info.Types, x.Name.Name)
		}
		return true
	})

	return info, nil
}
