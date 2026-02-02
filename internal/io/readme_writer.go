package io

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/amonvix/go-doc-agent/internal/generator"
)

func WriteReadme(
	readme *generator.ReadmeDoc,
	templatesDir string,
) error {

	outputPath := "README.generated.md"

	tmpl, err := template.ParseGlob(
		filepath.Join(templatesDir, "*.tmpl"),
	)
	if err != nil {
		return err
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.ExecuteTemplate(f, "header.tmpl", readme)
}
