package commentingo

import (
	"fmt"
	"os"

	"github.com/amonvix/go-doc-agent/internal/generator"
	"github.com/amonvix/go-doc-agent/internal/io"
	"github.com/amonvix/go-doc-agent/internal/language/go/parser"
)

// main TODO: add description
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: commentingo <file.go>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	astFile, fset, err := parser.ParseFile(filePath)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}

	funcs := parser.ExtractFunctions(astFile)

	comments, err := generator.GenerateComments(funcs)
	if err != nil {
		fmt.Println("Error generating comments:", err)
		os.Exit(1)
	}

	err = io.WriteComments(filePath, fset, astFile, comments)
	if err != nil {
		fmt.Println("Error writing comments:", err)
		os.Exit(1)
	}

	fmt.Println("Comments added successfully.")
}
