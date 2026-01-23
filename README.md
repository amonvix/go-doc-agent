# go-doc-agent 🧠⚙️

**go-doc-agent** is a documentation engine designed to automatically analyze source code using AST and generate structured, consistent technical documentation.

Built for engineers who care about **clarity, scale, and systems that survive real-world complexity**.

---

## 🚀 What is this?

A modular, language-oriented documentation engine that:

- Parses source code using Abstract Syntax Trees (AST)
- Extracts structural and semantic context
- Generates professional documentation automatically
- Keeps documentation aligned with the actual codebase

No manual comments.  
No outdated README files.  
No guesswork.

---

## 🎯 Why this exists

Large systems rarely fail because of bad code.

They fail because:

- Documentation becomes outdated  
- Knowledge stays locked inside people  
- Codebases scale faster than understanding  

**go-doc-agent** was created to close that gap.

It treats documentation as a **derivative of the source code**, not a separate artifact that inevitably rots over time.

---

## 🧩 High-level architecture

Source Code
↓
Language Parser (AST)
↓
Unified Context Model
↓
Writer Engine
↓
Comments / Docs / README

---

## ⚙️ Core design principles

- **AST-based parsing** — no regex, no heuristics  
- **Language adapters** — each language owns its parser logic  
- **Unified context model** — writers are language-agnostic  
- **Strict separation of concerns**  
- **Extensible by design**  
- **Production-first mindset**  

AI is a feature — **not a dependency**.

The system works with or without external LLMs.

---

## 🧠 What makes it different

Most documentation tools are:

- language-locked  
- template-driven  
- fragile at scale  

**go-doc-agent** is built as an engine, not a script.

It is designed to support:

- multi-language expansion  
- multiple output formats  
- different documentation strategies  
- local or cloud-based AI models  

Without rewriting the core.

---

## 🧪 Real-world use case

> A company with hundreds of microservices and no standardized documentation can automatically generate consistent comments and README files directly from the source code.

The documentation becomes reproducible, auditable, and scalable.

---

## 🛠 Tech Stack

- Go (core engine)
- Native AST parsing
- Modular internal architecture
- CLI-oriented execution model
- Optional AI integration

---

## 📌 Roadmap

- [x] Project directory scanner  
- [x] Language detection layer  
- [x] Unified context model  
- [x] Go AST parser  
- [ ] Comment writer engine  
- [ ] README generator  
- [ ] CLI interface  
- [ ] Multi-language support  
- [ ] Local LLM integration  

---

## ⚠️ Important

This tool can modify source files.

Always use version control before execution.

---

## 🧭 Philosophy

This project follows a simple belief:

> Systems should explain themselves.

Documentation should not depend on memory, discipline, or hero developers.

It should be **generated from truth — the code itself**.

---

Built with engineering discipline  
and a strong dislike for systems that collapse under pressure.
