package main

import (
	"fmt"

	"github.com/amonvix/go-doc-agent/internal/semantic"
)

func printFunction(fn semantic.Function) {
	fmt.Println("▶ Function:", fn.Name)

	if fn.File != "" {
		fmt.Println("  File:", fn.File)
	}

	if len(fn.Params) > 0 {
		fmt.Println("  Params:")
		for _, p := range fn.Params {
			fmt.Printf("    - %s %s\n", p.Name, p.Type)
		}
	}

	if len(fn.Returns) > 0 {
		fmt.Println("  Returns:")
		for _, r := range fn.Returns {
			fmt.Printf("    - %s\n", r.Type)
		}
	}

	fmt.Println()
}
