package language

import (
	"log"
	"path/filepath"
)

func Detect(path string) (string, bool) {
	var lang string
	var ok bool

	switch filepath.Ext(path) {
	case ".go":
		lang, ok = "go", true
	case ".py":
		lang, ok = "python", true
	case ".js", ".ts":
		lang, ok = "javascript", true
	default:
		lang, ok = "", false
	}

	log.Printf("[pipeline] language detected: %s\n", lang)
	return lang, ok
}
