package context

import "go/ast"

type Function struct {
	Name   string
	HasDoc bool
	Node   *ast.FuncDecl
}
