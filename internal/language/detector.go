package language

import (
	"log"
	"path/filepath"
)

func Detect(path string) (string, bool) {
	log.Printf("[pipeline] language detected: %s\n", lang)

	switch filepath.Ext(path) {

	case ".go":
		return "go", true

	case ".py":
		return "python", true

	case ".js", ".ts":
		return "javascript", true

	default:
		return "", false
	}
}
