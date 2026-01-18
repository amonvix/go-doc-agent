package generator

import "go/ast"

func GenerateComments(file *ast.File) (map[ast.Node]string, error) {
	comments := make(map[ast.Node]string)

	ast.Inspect(file, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		if fn.Doc == nil {
			comments[fn] =
				"// " + fn.Name.Name + " TODO: add description"
		}

		return true
	})

	return comments, nil
}
