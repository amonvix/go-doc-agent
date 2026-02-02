package generator

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/amonvix/go-doc-agent/internal/semantic"
)

func GenerateComments(
	funcs []semantic.Function,
) (map[string]string, error) {

	results := make(map[string]string)

	tmplPath := filepath.Join(
		"templates",
		"comments",
		"function.tmpl",
	)

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return nil, fmt.Errorf("load function template: %w", err)
	}

	for _, fn := range funcs {

		var buf bytes.Buffer

		if err := tmpl.Execute(&buf, fn); err != nil {
			return nil, fmt.Errorf(
				"generate comment for %s: %w",
				fn.Name,
				err,
			)
		}

		results[fn.Name] = buf.String()
	}

	return results, nil
}
