package comment

import "go/ast"

func GenerateComments(file *ast.File) (map[ast.Node]string, error) {
	comments := make(map[ast.Node]string)

	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.FuncDecl:
			if node.Doc == nil {
				comments[node] = "// " + node.Name.Name + " TODO: add description"
			}
		}
		return true
	})

	return comments, nil
}
