package main

import (
	"fmt"
	"os"

	"github.com/amonvix/go-doc-agent/internal/context"
	"github.com/amonvix/go-doc-agent/internal/context/builder"
	"github.com/amonvix/go-doc-agent/internal/generator"
	"github.com/amonvix/go-doc-agent/internal/io"

	goparser "github.com/amonvix/go-doc-agent/internal/language/go/parser"
)

// main TODO: add description
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  commentingo <file.go>")
		fmt.Println("  commentingo <directory>")
		os.Exit(1)
	}

	input := os.Args[1]

	project, err := builder.Build(input)
	if err != nil {
		fmt.Println("Build error:", err)
		os.Exit(1)
	}

	grouped := context.GroupFunctionsByFile(project.Functions)

	for filePath, funcs := range grouped {

		astFile, fset, err := goparser.ParseFile(filePath)
		if err != nil {
			fmt.Println("Parse error:", err)
			continue
		}

		comments, err := generator.GenerateComments(funcs)
		if err != nil {
			fmt.Println("Generator error:", err)
			continue
		}

		err = io.WriteComments(filePath, fset, astFile, comments)
		if err != nil {
			fmt.Println("Writer error:", err)
			continue
		}

		fmt.Println("Commented:", filePath)
	}

	fmt.Println("Project documentation complete.")
	return
}
