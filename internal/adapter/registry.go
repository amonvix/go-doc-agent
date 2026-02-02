package adapter

import (
	"fmt"

	"github.com/amonvix/go-doc-agent/internal/language"
)

var registry = map[language.ID]Adapter{}

func Register(lang language.ID, a Adapter) {
	registry[lang] = a
}

func Select(lang language.ID) (Adapter, error) {
	a, ok := registry[lang]
	if !ok {
		return nil, fmt.Errorf("no adapter registered for language %s", lang)
	}
	return a, nil
}
