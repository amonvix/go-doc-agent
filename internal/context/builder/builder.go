package builder

import (
	"go/parser"
	"go/token"
	"os"

	"github.com/amonvix/go-doc-agent/internal/context"
	"github.com/amonvix/go-doc-agent/internal/io"
	"github.com/amonvix/go-doc-agent/internal/language"
	goparser "github.com/amonvix/go-doc-agent/internal/language/go/parser"
)

func Build(path string) (*context.Project, error) {

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	var files []io.SourceFile

	if info.IsDir() {
		files, err = io.ScanDirectory(path)
	} else {
		file, err := io.ScanFile(path)
		if err != nil {
			return nil, err
		}
		files = []io.SourceFile{*file}
	}

	if err != nil {
		return nil, err
	}

	project := &context.Project{
		Path: path,
	}

	for _, file := range files {

		lang, ok := language.Detect(file.Path)
		if !ok {
			continue
		}

		switch lang {

		case "go":

			fset := token.NewFileSet()

			astFile, err := parser.ParseFile(
				fset,
				file.Path,
				nil,
				parser.ParseComments,
			)
			if err != nil {
				return nil, err
			}

			functions := goparser.ExtractFunctions(astFile)

			for i := range functions {
				functions[i].FilePath = file.Path
			}

			project.Functions = append(project.Functions, functions...)
		}
	}

	return project, nil
}
