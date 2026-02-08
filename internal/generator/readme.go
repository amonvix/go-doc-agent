package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/amonvix/go-doc-agent/internal/language/golang"
)

func GenerateFolderREADME(dir string, files []string) error {
	readmePath := filepath.Join(dir, "README.md")

	f, err := os.Create(readmePath)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "# Package %s\n\n", filepath.Base(dir))
	// fmt.Fprintln(f, "This folder contains the following Go files:\n")

	for _, file := range files {
		info, err := golang.ParseFile(file)
		if err != nil {
			continue
		}

		fmt.Fprintf(f, "## %s\n", filepath.Base(file))

		if len(info.Types) > 0 {
			fmt.Fprintln(f, "**Types:**")
			for _, t := range info.Types {
				fmt.Fprintf(f, "- %s\n", t)
			}
		}

		if len(info.Functions) > 0 {
			fmt.Fprintln(f, "\n**Functions:**")
			for _, fn := range info.Functions {
				fmt.Fprintf(f, "- %s\n", fn)
			}
		}

		fmt.Fprintln(f)
	}

	return nil
}
