package parser

import (
	"go/ast"

	"github.com/amonvix/go-doc-agent/internal/context"
)

func ExtractFunctions(file *ast.File) []context.Function {
	var functions []context.Function

	ast.Inspect(file, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		f := context.Function{
			Name:   fn.Name.Name,
			HasDoc: fn.Doc != nil,
		}

		// parameters
		if fn.Type.Params != nil {
			for _, p := range fn.Type.Params.List {
				typ := exprToString(p.Type)

				for _, name := range p.Names {
					f.Params = append(f.Params, context.Param{
						Name: name.Name,
						Type: typ,
					})
				}
			}
		}

		// returns
		if fn.Type.Results != nil {
			for _, r := range fn.Type.Results.List {
				f.Returns = append(f.Returns, context.Return{
					Type: exprToString(r.Type),
				})
			}
		}

		functions = append(functions, f)
		return true
	})

	return functions
}
