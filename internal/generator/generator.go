package generator

import (
	"log"

	"github.com/amonvix/go-doc-agent/internal/semantic"
)

func Generate(project *semantic.Project) (*DocBundle, error) {
	log.Println("[generator] generating documentation bundle")

	bundle := &DocBundle{}

	bundle.Readme = buildReadme(project)

	for _, fn := range project.Functions {
		bundle.Comments = append(bundle.Comments, CommentDoc{
			FilePath: project.Name,
			Target:   fn.Name,
			Text:     buildFunctionSummary(fn),
		})
	}

	return bundle, nil
}
