package internal

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func WriteComments(path string, fset *token.FileSet, file *ast.File, comments map[ast.Node]string) error {
	for node, comment := range comments {
		switch n := node.(type) {
		case *ast.FuncDecl:
			n.Doc = &ast.CommentGroup{
				List: []*ast.Comment{
					{Text: comment},
				},
			}
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return printer.Fprint(f, fset, file)
}
