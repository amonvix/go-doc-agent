package generator

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/amonvix/go-doc-agent/internal/context"
)

func GenerateComments(funcs []context.Function) (map[string]string, error) {

	tplPath := filepath.Join(
		"templates",
		"go",
		"comment.tmpl",
	)

	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)

	for _, fn := range funcs {

		if fn.HasDoc {
			continue
		}

		var buf bytes.Buffer

		err := tpl.Execute(&buf, fn)
		if err != nil {
			return nil, err
		}

		result[fn.Name] = buf.String()
	}

	return result, nil
}
