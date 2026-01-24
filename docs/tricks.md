# Shorter is better than verboses
___________________________________________________
|     Question                 |  Responsible     |
| -----------------------------|------------------|
| where are the files?         | Scanner          |
| Who coordinates everything?  | Builder          |
| What language is that?       | Detector         |
| Is it a valid code?          | Parser           |
| What structures exist?       | Extractor        |
| what this realy means?       | Semantic Analyzer|
| what do we have to document? | Generator     	  |
| Where do I save this?        | Writer           |
---------------------------------------------------

One rule that no one tells you and is the core of the Semantic Analizer
If function name starts with:
- Get     → returns something
- Set     → modifies something
- Parse   → reads and converts input
- Load    → reads from IO
- Save    → writes to IO
- New     → constructor
- Is/Has  → boolean

Explaning AST
Abstract Syntax Tree (AST) is a tree representation of the source code structure.
It represents how the code is written, not what the code does.

Each node corresponds to a syntactic element, such as:
- function declarations (*ast.FuncDecl)
- literals (*ast.BasicLit)
- identifiers (*ast.Ident)
- statements and expressions

Why this is important?
Because the extractor walks through the AST to identify structures such as functions, receivers, parameters and visibility.

Filosofy part:
Official documentation is written to be correct.
Good documentation is written to be understood.

and our goal is to masters the second...