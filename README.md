# Comment in Go

> **Automated Go source code commenter.**
> Add clear, idiomatic GoDoc comments to your `.go` files in seconds.

---

## Why

Writing good documentation is boring. Reading undocumented code is worse.
**Comment in Go** fixes that by parsing your code with Go's AST and injecting clean, convention‑friendly comments automatically.

No frameworks. No agents. No magic YAML. Just Go.

---

## What it does (Release 1)

- Parses a Go file using the official `go/ast` and `go/parser`
- Detects functions without doc comments
- Generates and injects GoDoc-style comments **in place**
- Rewrites the same file (like `gofmt` does)

---

## What’s next (Release 2)

- Read existing comments across a folder
- Infer project structure and intent
- Generate a `README.md` automatically

> Roadmap is intentional. Core first, intelligence later.

---

## Installation

```bash
go build -o commentingo ./cmd/commentingo
```

Or clone and run directly:

```bash
go run ./cmd/commentingo -- ./path/to/file.go
```

---

## Usage

```bash
commentingo ./examples/sample.go
```

That’s it. The file explained by comments in place.

---

## Example

Before:

```go
func Add(a int, b int) int {
    return a + b
}
```

After:

```go
// Add adds two integers and returns the result.
func Add(a int, b int) int {
    return a + b
}
```

---

## Project Structure

```text
cmd/
  commentingo/      # CLI entrypoint
internal/           # Core logic (parser, commenter, writer)
examples/           # Sample files
prompts/            # LLM prompts (future)
roadmap/            # Planned features (e.g. Python, README generator)
```

---

## Design Principles

- **Single responsibility** – do one thing well
- **Native tooling** – use Go AST, not regex
- **Zero configuration** – no setup ceremony
- **Product mindset** – built to be raised, not simple studied

---

## Why Go

Because documentation tools should be:

- fast
- predictable
- boring (in the best way)

Go is perfect for that.

---

## Status

- ✅ Release 1 – Comment Go files
- ⏳ Release 2 – README generator
- ⏳ Multi-file analysis
- ⏳ LLM-powered semantic comments

---

## Author

Built by **Amon (Daniel Pedroso)**

Engineer. Incident analyst. Builder of things that should exist.

---

## License

MIT. Do what you want. Just don’t ship garbage.
