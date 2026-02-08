package commentingo

import (
	"flag"
)

func Parse() string {
	input := flag.String("input", ".", "Path to project root")
	flag.Parse()
	return *input
}
