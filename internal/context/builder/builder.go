package builder

import (
	"os"

	"github.com/amonvix/go-doc-agent/internal/context"
	"github.com/amonvix/go-doc-agent/internal/fs"
)

func Build(path context.Path) (*context.Project, error) {

	info, err := os.Stat(path.String())
	if err != nil {
		return nil, err
	}

	var files []fs.SourceFile

	if info.IsDir() {
		files, err = fs.ScanDirectory(path.String())
	} else {
		file, err := fs.ScanFile(path.String())
		if err != nil {
			return nil, err
		}
		files = []fs.SourceFile{*file}
	}

	if err != nil {
		return nil, err
	}

	return &context.Project{
		Path:  path.String(),
		Files: files,
	}, nil
}
