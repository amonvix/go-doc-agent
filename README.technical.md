# go-doc-agent — Technical Documentation

This document describes the internal architecture, design decisions and execution flow of the **go-doc-agent** project.

It is intended for engineers who want to understand **how the system works internally**, not just what it does.

---

## 🎯 Design Goals

The project was designed with the following goals:

- Avoid language-specific coupling
- Support future multi-language expansion
- Isolate parsing logic from output generation
- Prevent documentation logic from depending on AI
- Enable deterministic, testable behavior
- Keep the system extensible without refactoring the core

---

## 🧠 Architectural Philosophy

The system follows a **pipeline-based architecture**.

Each stage has a single responsibility and communicates through a shared internal context model.

No layer is allowed to directly access another layer’s internal implementation.

This ensures:

- low coupling  
- high cohesion  
- predictable data flow  
- safe extensibility  

---

## 🧩 High-Level Flow

Project Path
↓
Directory Scanner
↓
Language Detection
↓
Language Parser (AST)
↓
Unified Context Model
↓
Writer Engine
↓
Generated Output


---

## 📂 Project Structure

go-doc-agent/
│
├── cmd/
│ └── commentingo/
│ └── main.go
│
├── internal/
│ ├── agent/
│ │ └── orchestrator.go
│ │
│ ├── io/
│ │ ├── scanner.go
│ │ └── file.go
│ │
│ ├── language/
│ │ ├── detect.go
│ │ └── registry.go
│ │
│ ├── context/
│ │ ├── project.go
│ │ ├── file.go
│ │ ├── function.go
│ │ └── symbol.go
│ │
│ ├── parser/
│ │ └── interface.go
│ │
│ ├── writer/
│ │ └── interface.go
│ │
│ └── language/
│ └── go/
│ └── parser/
│ └── ast_parser.go
│
└── README.md


---

## 🔍 Core Components

---

### 1️⃣ Directory Scanner (`internal/io`)

Responsible for:

- walking the project directory
- identifying readable source files
- ignoring unsupported formats
- loading file content safely

This layer has **zero knowledge of language syntax**.

---

### 2️⃣ Language Detection (`internal/language`)

Each file is analyzed to determine:

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

Each language owns:

its AST logic

syntax rules

semantic extraction

The core system never interacts with AST directly.

4️⃣ Go AST Parser

The Go implementation uses:

go/parser

go/ast

go/token

Responsibilities:

extract functions

capture receivers

identify exported vs private symbols

preserve comments

map source positions

The output is normalized into the unified context model.

5️⃣ Unified Context Model (internal/context)

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

Go AST

Python AST

Tree-sitter

LLM output

Only the context matters.

6️⃣ Writer Layer (internal/writer)

Writers consume the context model and generate output:

GoDoc comments

Markdown documentation

README files

future formats (HTML, JSON, etc.)

Writers never parse code.

They only translate structured context into output.

🤖 AI Integration (Optional)

AI is treated as:

a comment generator

not a parser

not a decision engine

The system remains functional without AI.

This prevents:

vendor lock-in

non-deterministic builds

dependency on external services

AI is a plugin — not a foundation.

🔒 Why AST-Based Parsing

Regex-based documentation tools fail because they:

break on formatting

cannot understand scope

cannot identify receivers

misinterpret nested logic

AST provides:

syntactic correctness

structural certainty

safe refactoring support

future-proof parsing

🧪 Error Handling Strategy

Errors are isolated per file:

one malformed file does not break the project

parsing failures are reported, not fatal

partial results are allowed

This mirrors real-world CI behavior.

🚀 Extending the System

To add a new language:

Implement a parser adapter

Map AST data into the context model

Register the language

Reuse all existing writers

No refactoring required.

🧭 Summary

go-doc-agent is not a script.

It is a documentation engine designed with:

strict separation of concerns

predictable data flow

extensibility as a first-class feature

production-oriented architecture

The system prioritizes clarity, safety and long-term maintainability.

Documentation should not be written.

It should be derived.