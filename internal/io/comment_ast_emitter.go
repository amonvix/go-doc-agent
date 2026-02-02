package io

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func WriteComments(
	filePath string,
	comments map[string]string,
) error {
	log.Println("[writer] writing outputs")

	fset := token.NewFileSet()

	file, err := parser.ParseFile(
		fset,
		filePath,
		nil,
		parser.ParseComments,
	)
	if err != nil {
		return err
	}

	ast.Inspect(file, func(n ast.Node) bool {

		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		comment, exists := comments[fn.Name.Name]
		if !exists {
			return true
		}

		// já tem comentário → pula
		if fn.Doc != nil {
			return true
		}

		fn.Doc = &ast.CommentGroup{
			List: []*ast.Comment{
				{
					Text: "// " + comment,
				},
			},
		}

		return true
	})

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return printer.Fprint(f, fset, file)
}
