package analyzer

import (
	"unicode"

	"github.com/amonvix/go-doc-agent/internal/semantic"
)

func DetectEntrypoint(fn *semantic.Function) {
	fn.IsEntryPoint =
		!fn.IsMethod &&
			len(fn.Name) > 0 &&
			unicode.IsUpper(rune(fn.Name[0]))
}
