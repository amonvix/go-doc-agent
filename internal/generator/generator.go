package generator

import "github.com/amonvix/go-doc-agent/internal/semantic"

func Generate(project *semantic.Project) (*DocBundle, error) {

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
