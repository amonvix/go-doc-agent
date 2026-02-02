package generator

import (
	"path/filepath"

	"github.com/amonvix/go-doc-agent/internal/semantic"
)

func buildReadme(project *semantic.Project) *ReadmeDoc {
	return &ReadmeDoc{
		Title:       filepath.Base(project.Name),
		Description: "Auto-generated documentation by go-doc-agent.",
		Sections: []Section{
			{
				Name:    "Overview",
				Content: "This documentation was generated automatically.",
			},
		},
	}
}
