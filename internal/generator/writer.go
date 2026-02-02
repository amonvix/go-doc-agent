package generator

import (
	"os"
	"path/filepath"
	"text/template"
)

func WriteTemplate(
	outputPath string,
	templatePath string,
	data any,
) error {

	tmpl, err := template.New(filepath.Base(templatePath)).ParseFiles(templatePath)

	if err != nil {
		return err
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}
