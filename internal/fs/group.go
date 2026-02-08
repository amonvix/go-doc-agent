package fs

import "path/filepath"

func GroupByDir(files []string) map[string][]string {
	grouped := make(map[string][]string)
	for _, f := range files {
		dir := filepath.Dir(f)
		grouped[dir] = append(grouped[dir], f)
	}
	return grouped
}
