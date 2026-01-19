package io

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func WriteComments(
	path string,
	fset *token.FileSet,
	file *ast.File,
	comments map[string]string,
) error {

	ast.Inspect(file, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		// já tem comentário? pula
		if fn.Doc != nil {
			return true
		}

		comment, exists := comments[fn.Name.Name]
		if !exists {
			return true
		}

		fn.Doc = &ast.CommentGroup{
			List: []*ast.Comment{
				{Text: comment},
			},
		}

		return true
	})

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return printer.Fprint(f, fset, file)
}
