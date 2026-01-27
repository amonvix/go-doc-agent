package goadapter

import (
	"go/ast"
	"strings"

	"github.com/amonvix/go-doc-agent/internal/semantic/model"
)

func Build(astFile *ast.File) *model.Project {

	project := &model.Project{}

	// imports → dependencies
	for _, imp := range astFile.Imports {
		project.Dependencies = append(
			project.Dependencies,
			model.Dependency{
				Package: strings.Trim(imp.Path.Value, `"`),
				Type:    model.DependencyUnknown,
			},
		)
	}

	ast.Inspect(astFile, func(n ast.Node) bool {

		switch node := n.(type) {

		case *ast.FuncDecl:

			fn := model.Function{
				Name: node.Name.Name,
				Metadata: model.Metadata{
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
						fn.Params = append(fn.Params, name.Name)
					}
				}
			}

			// returns
			if node.Type.Results != nil {
				for _, r := range node.Type.Results.List {
					fn.Returns = append(fn.Returns, exprToString(r.Type))
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
				current.Calls = append(current.Calls, model.Call{
					Name:   fun.Sel.Name,
					Target: exprToString(fun.X),
				})

			case *ast.Ident:
				current.Calls = append(current.Calls, model.Call{
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
