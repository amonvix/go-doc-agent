package parser

import (
	"go/ast"

	"github.com/amonvix/go-doc-agent/internal/context"
)

func ExtractFunctions(file *ast.File) []context.Function {
	var funcs []context.Function

	ast.Inspect(file, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		funcs = append(funcs, context.Function{
			Name:   fn.Name.Name,
			HasDoc: fn.Doc != nil,
			Node:   fn,
		})

		return true
	})

	return funcs
}
