package main

import (
	"fmt"
	"os"

	internal "github.com/amonvix/go-doc-agent/internal/comment"
)

// main TODO: add description
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: commentingo <file.go>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	astFile, fset, err := internal.ParseFile(filePath)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}

	comments, err := internal.GenerateComments(astFile)
	if err != nil {
		fmt.Println("Error generating comments:", err)
		os.Exit(1)
	}

	err = internal.WriteComments(filePath, fset, astFile, comments)
	if err != nil {
		fmt.Println("Error writing comments:", err)
		os.Exit(1)
	}

	fmt.Println("Comments added successfully.")
}
