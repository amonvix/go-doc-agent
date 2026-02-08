package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"

	"github.com/amonvix/go-doc-agent/internal/commentingo"
	"github.com/amonvix/go-doc-agent/internal/context"
	"github.com/amonvix/go-doc-agent/internal/context/builder"
	"github.com/amonvix/go-doc-agent/internal/fs"
	"github.com/amonvix/go-doc-agent/internal/generator"
	"github.com/amonvix/go-doc-agent/internal/io"
	"github.com/amonvix/go-doc-agent/internal/language/golang"
	"github.com/amonvix/go-doc-agent/internal/semantic"
	"github.com/amonvix/go-doc-agent/internal/semantic/analyzer"
)

func main() {

	jsonOutput := flag.Bool("json", false, "output result as JSON")
	prettyOutput := flag.Bool("pretty", false, "pretty-print JSON output")
	functionsView := flag.Bool("functions", false, "print function-by-function view")
	debugMode := flag.Bool("debug", false, "enable debug logging")

	flag.Parse()

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

	debugLog(*debugMode, "building context from path: %s", path)

	if *jsonOutput {
		var data []byte

		if *prettyOutput {
			data, err = json.MarshalIndent(ctxProject, "", "  ")
		} else {
			data, err = json.Marshal(ctxProject)
		}

		if err != nil {
			debugLog(*debugMode, "json marshal error: %v", err)
			os.Exit(1)
		}

		os.Stdout.Write(data)
		return
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
		debugLog(*debugMode, "parsed file: %s", file.Path)

		p, err := goAdapter.Build(astFile)
		if err != nil {
			fmt.Println("Adapter error:", err)
			os.Exit(1)
		}

		semanticProject.Functions =
			append(semanticProject.Functions, p.Functions...)

	}
	debugLog(*debugMode, "starting semantic analysis")

	// -----------------------------
	// 2.1. Search for .go files
	// -----------------------------

	root := commentingo.Parse()

	files, err := fs.WalkGoFiles(root)
	if err != nil {
		log.Fatal(err)
	}

	grouped := fs.GroupByDir(files)
	for dir, files := range grouped {
		if err := generator.GenerateFolderREADME(dir, files); err != nil {
			log.Println("error generating README for", dir, err)
		}
	}

	// -----------------------------
	// 3. Functions describer
	// -----------------------------

	if *functionsView {
		for _, fn := range semanticProject.Functions {
			commentingo.PrintFunction(fn)
		}
		return
	}
	debugLog(*debugMode, "detected %d functions", len(semanticProject.Functions))

	// -----------------------------
	// 4. Semantic analyzer
	// -----------------------------
	analyzer.Analyze(semanticProject)

	// -----------------------------
	// 5. Generate documentation
	// -----------------------------
	bundle, err := generator.Generate(semanticProject)
	if err != nil {
		fmt.Println("Generator error:", err)
		os.Exit(1)
	}

	// -----------------------------
	// 6. Write README
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
	// 7. Write semantic comments
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

func debugLog(enabled bool, format string, args ...any) {
	if !enabled {
		return
	}
	fmt.Fprintf(os.Stderr, "[go-doc-agent][debug] "+format+"\n", args...)
}
