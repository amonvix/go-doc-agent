package language

import "path/filepath"

func Detect(path string) (string, bool) {
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
