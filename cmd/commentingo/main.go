package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"

	"github.com/amonvix/go-doc-agent/internal/context"
	"github.com/amonvix/go-doc-agent/internal/context/builder"
	"github.com/amonvix/go-doc-agent/internal/generator"
	"github.com/amonvix/go-doc-agent/internal/io"
	"github.com/amonvix/go-doc-agent/internal/language/golang"
	"github.com/amonvix/go-doc-agent/internal/semantic"
	"github.com/amonvix/go-doc-agent/internal/semantic/analyzer"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: commentingo <path>")
		os.Exit(1)
	}

	path := context.NewPath(os.Args[1])

	// -----------------------------
	// 1. Build context project
	// -----------------------------
	ctxProject, err := builder.Build(path)
	if err != nil {
		fmt.Println("Build error:", err)
		os.Exit(1)
	}

	// -----------------------------
	// 2. Semantic adapter (defaults)
	// -----------------------------
	semanticProject := &semantic.Project{}
	goAdapter := &golang.Adapter{}

	for _, file := range ctxProject.Files {

		fset := token.NewFileSet()

		astFile, err := parser.ParseFile(
			fset,
			file.Path,
			file.Content,
			parser.ParseComments,
		)
		if err != nil {
			fmt.Println("Parser error:", err)
			os.Exit(1)
		}

		p, err := goAdapter.Build(astFile)
		if err != nil {
			fmt.Println("Adapter error:", err)
			os.Exit(1)
		}

		semanticProject.Functions =
			append(semanticProject.Functions, p.Functions...)
	}
	log.Println("[pipeline] starting analysis")

	// -----------------------------
	// 3. Semantic analyzer
	// -----------------------------
	analyzer.Analyze(semanticProject)

	// -----------------------------
	// 4. Generate documentation
	// -----------------------------
	bundle, err := generator.Generate(semanticProject)
	if err != nil {
		fmt.Println("Generator error:", err)
		os.Exit(1)
	}

	// -----------------------------
	// 5. Write README
	// -----------------------------
	err = io.WriteReadme(
		bundle.Readme,
		filepath.Join("templates", "readme"),
	)
	if err != nil {
		fmt.Println("README writer error:", err)
		os.Exit(1)
	}

	// -----------------------------
	// 6. Write semantic comments
	// -----------------------------
	err = io.WriteCommentsMarkdown(
		path,
		bundle.Comments,
	)
	if err != nil {
		fmt.Println("Comments writer error:", err)
		os.Exit(1)
	}

	fmt.Println("✅ Documentation generated successfully.")
}
