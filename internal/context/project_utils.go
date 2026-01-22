package context

func GroupFunctionsByFile(funcs []Function) map[string][]Function {

	grouped := make(map[string][]Function)

	for _, fn := range funcs {
		grouped[fn.FilePath] = append(grouped[fn.FilePath], fn)
	}

	return grouped
}
