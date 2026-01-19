package generator

import (
	"bytes"
	"text/template"

	"github.com/amonvix/go-doc-agent/internal/context"
)

func GenerateComments(funcs []context.Function) (map[string]string, error) {
	result := make(map[string]string)

	tpl, err := template.ParseFiles("templates/go/go_comment.tmpl")
	if err != nil {
		return nil, err
	}

	for _, fn := range funcs {
		if fn.HasDoc {
			continue
		}

		var buf bytes.Buffer
		if err := tpl.Execute(&buf, fn); err != nil {
			return nil, err
		}

		result[fn.Name] = buf.String()
	}

	return result, nil
}
