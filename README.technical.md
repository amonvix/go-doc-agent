# go-doc-agent — Technical Documentation ☕🧠

This document describes the internal architecture, design decisions, and execution flow of the **go-doc-agent** system.

It is intended for engineers who want to understand **how the system works internally**, not just what it produces.

If you drink coffee while reading source code, this file is for you.

---

## 🎯 Design Goals

The project was built around the following non-negotiable goals:

- Avoid language-specific coupling
- Support future multi-language expansion
- Isolate parsing logic from documentation output
- Prevent documentation logic from depending on AI
- Enable deterministic and testable behavior
- Keep extensibility possible without refactoring the core
- Ensure architecture clarity even at large scale

---

## 🧠 Architectural Philosophy

The system follows a **pipeline-based architecture**.

Each stage:

- has a single responsibility  
- communicates only through defined models  
- does not access internal logic of other layers  

No layer is allowed to shortcut the pipeline.

This guarantees:

- low coupling  
- high cohesion  
- predictable execution flow  
- safe extensibility  
- maintainable growth  

In short: no spaghetti systems.

---

## 🧩 High-Level Execution Flow

Project Path
↓
Directory Scanner
↓
Language Detection
↓
Language Parser (AST)
↓
Structure Extractor
↓
Semantic Analyzer
↓
Meaning Output Model
↓
Writer Engine
↓
Generated Comments / Docs / README

---

Each arrow represents a strict data contract — not a function call shortcut.

---

## 📂 Project Structure

```text
├── README.md
├── README.technical.md
├── cmd
│   └── commentingo
│       └── main.go
├── docs
│   ├── README.md
│   ├── explanation.md
│   ├── pipeline.md
│   └── tricks.md
├── examples
│   ├── api
│   │   ├── handler.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── mixed.go
│   ├── no_comments.go
│   ├── readme
│   │   ├── README.generated.md
│   │   └── project-structure.txt
│   ├── sample
│   │   ├── input.go
│   │   └── output.go
│   └── sample.go
├── go.mod
├── internal
│   ├── context
│   │   ├── builder
│   │   │   └── builder.go
│   │   ├── extractor.go
│   │   ├── model.go
│   │   └── project_utils.go
│   ├── semantic
│   │   ├── analyzer
│   │   ├── adapter
│   │   └── model
│   ├── generator
│   │   ├── commenter.go
│   │   └── readme_generator.go
│   ├── io
│   │   ├── scanner.go
│   │   ├── writer.go
│   │   └── readme_writer.go
│   └── language
│       ├── detector.go
│       └── go
│           └── parser
├── templates
│   ├── comments
│   ├── config
│   └── readme
└── prompts
    └── go_comment_prompt.txt

```
---

## 🔍 Core Components

---

### 1️⃣ Directory Scanner (`internal/io`)

Responsible for:

- walking the project directory tree
- identifying readable source files
- ignoring unsupported formats
- loading file content safely

This layer has:

**no AST knowledge**
**language awareness**
**no semantic responsibility**

- It deals exclusively with filesystem operations.

---

### 2️⃣ Language Detection (`internal/language`)

Determines for each file:

- supported language
- parser availability

This allows future extensions such as:

- Python
- Java
- TypeScript
- Rust

without modifying the core engine.

---

### 3️⃣ Parser Layer (`internal/parser`)

Parsers implement a shared interface:

```go
type Parser interface {
    Parse(file File) (*context.FileContext, error)
}
```

Each language owns:

- its AST logic

- syntax rules

- language-specific construct

The core system never interacts with AST directly.

### 4️⃣ Go AST Parser

The Go implementation uses:

- go/parser

- go/ast

- go/token

Responsibilities include:

- extracting functions and methods

- identifying receivers

- detecting exported vs private symbols

- preserving comments

- mapping source positions

The output is normalized into the unified context model.


### 5️⃣ Unified Context Model (internal/context)

This is the heart of the system.

It represents the codebase independently of language.

Example:

type Function struct {
    Name        string
    Receiver    string
    Parameters  []Parameter
    Returns     []Return
    IsExported  bool
    Comments    []string
}

Writers never care whether data came from:

- Go AST

- Python AST

- Tree-sitter

- LLM output

- Static analysis

Only the context matters.

### 6️⃣ Semantic Analyzer (internal/semantic)

The semantic analyzer answers:

“What does this structure mean?”

# It determines:

- architectural role (handler, service, repository…)

- system layer (API, domain, infrastructure…)

- dependencies (database, network, filesystem…)

- intent (CRUD, orchestration, mapping…)

- detected behavior flags

- optional confidence score

This transforms structure into meaning.

### 7️⃣ Meaning Output Model

The analyzer produces a normalized semantic result:

Function:
- role: repository
- layer: persistence
- intent: data-access
- dependencies: database
- behavior: CRUD
- confidence: high

This model becomes the single source of truth for all documentation output.

### 8️⃣ Writer Layer (internal/generator)

Writers consume semantic meaning — never source code.

# They generate:

- GoDoc comments

- Markdown documentation

- README files

- future formats (HTML, JSON, diagrams)

# Writers are:

- language-agnostic

- deterministic

- output-focused

They translate meaning into documentation.

🤖 AI Integration (Optional)

AI is treated as:

- a comment generator

it is not:

- not a parser

- not a decision engine

- The system remains functional without AI.

# This prevents:

- vendor lock-in

- non-deterministic builds

- dependency on external services

AI is a plugin — not a foundation.

### 🔒 Why AST-Based Parsing


Regex-based documentation tools fail because they:

- break on formatting

- cannot understand scope

- cannot identify receivers

- misinterpret nested logic

AST parsing provides:

- syntactic correctness

- structural certainty

- safe refactoring support

- future-proof parsing

## 🧪 Error Handling Strategy

- Errors are isolated per file

- one malformed file does not break the project

- parsing failures are reported, not fatal

- partial results are allowed

This mirrors real-world CI/CD behavior.


## 🚀 Extending the System

To add a new language:

1. Implement a parser adapter

2. Map AST data into the context model

3. Register the language

4. Reuse all existing writers

No refactoring required.

🧭 Summary

go-doc-agent is not a script.

It is a documentation engine designed with:

- strict separation of concerns

- deterministic execution flow

- extensibility as a first-class feature

- production-grade architecture

- semantic understanding.

Documentation should not be written manually.

It should be derived from truth — the source code itself.