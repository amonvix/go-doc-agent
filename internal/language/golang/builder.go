package golang

import (
	"go/ast"
	"strings"

	"github.com/amonvix/go-doc-agent/internal/semantic"
)

func Build(astFile *ast.File) *semantic.Project {

	project := &semantic.Project{}

	// imports → dependencies
	for _, imp := range astFile.Imports {
		project.Dependencies = append(
			project.Dependencies,
			semantic.Dependency{
				Package: strings.Trim(imp.Path.Value, `"`),
				Type:    semantic.DependencyUnknown,
			},
		)
	}

	ast.Inspect(astFile, func(n ast.Node) bool {

		switch node := n.(type) {

		case *ast.FuncDecl:

			fn := semantic.Function{
				Name: node.Name.Name,
				Metadata: semantic.Metadata{
					Language: "go",
				},
				IsPure: true,
			}

			// method detection
			if node.Recv != nil && len(node.Recv.List) > 0 {
				fn.IsMethod = true
			}

			// params
			if node.Type.Params != nil {
				for _, p := range node.Type.Params.List {
					for _, name := range p.Names {
						fn.Params = append(fn.Params, semantic.Param{
							Name: name.Name,
							Type: semantic.TypeInfo{
								Name: exprToString(p.Type),
							},
						})
					}
				}
			}

			// returns
			if node.Type.Results != nil {
				for _, r := range node.Type.Results.List {
					fn.Returns = append(fn.Returns, semantic.Return{
						Type: semantic.TypeInfo{
							Name: exprToString(r.Type),
						},
					})
				}
			}

			project.Functions = append(project.Functions, fn)

		case *ast.CallExpr:

			if len(project.Functions) == 0 {
				return true
			}

			current := &project.Functions[len(project.Functions)-1]

			switch fun := node.Fun.(type) {

			case *ast.SelectorExpr:
				current.Calls = append(current.Calls, semantic.Call{
					Name:   fun.Sel.Name,
					Target: exprToString(fun.X),
				})

			case *ast.Ident:
				current.Calls = append(current.Calls, semantic.Call{
					Name: fun.Name,
				})
			}
		}

		return true
	})

	return project
}

func exprToString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	default:
		return "unknown"
	}
}
