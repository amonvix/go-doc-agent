package adapter

import "github.com/amonvix/go-doc-agent/internal/context"

type Adapter interface {
	Analyze(ctx *context.Project) error
	Language() string
}
