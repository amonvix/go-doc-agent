# Pipeline used

This file represents the full execution pipeline of the go-doc-agent.
Each step answers exactly one question and has a single responsibility.

INPUT
↓
scanner
↓
builder
↓
parser
↓
extractor
↓
Semantic Analizer
↓
generator
↓
writer
↓
OUTPUT