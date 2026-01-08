# go-doc-agent 🧠⚙️

**go-doc-agent** is an AI-powered agent that automatically documents Go source code using professional GoDoc-style comments and generates clean README files.

## 🚀 What it does

- Scans a Go project directory
- Reads `.go` source files
- Uses AI to generate clean, professional comments in English
- Injects comments directly into the code
- (Coming soon) Generates README.md based on code structure and comments

## 🎯 Why this exists

Documentation is often neglected, inconsistent, or outdated.  
`go-doc-agent` was created to automate documentation in a clean, professional, and scalable way.

This project is part of a real-world automation initiative focused on AI-assisted software engineering.

## 🧠 How it works (high level)

1. Directory scanner locates Go files
2. Each file is sent to the AI engine with a technical documentation prompt
3. The AI returns the commented version
4. The agent writes the result back to disk

## 🛠 Tech Stack

- Python 3.10+
- OpenAI API (or local LLM in future)
- Go (target language)

## 📌 Roadmap

- [x] Go file scanner
- [ ] AI code commenter
- [ ] README generator
- [ ] Directory watcher (auto mode)
- [ ] CLI interface

## ⚠️ Disclaimer

This tool modifies source files. Always use version control.

## ⚙️ How to run

# python3 main.py --path ~/.project_location --readme

# python3 main.py --path ~/.project_location --no-readme

# python3 main.py --path ~/.project_location --dry-run

---

Built with engineering discipline and a slight obsession for automation.
