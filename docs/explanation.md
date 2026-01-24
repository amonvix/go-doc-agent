# Go explanation

INPUT

↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓
**SCANNER**
The scanner is responsible for receiving the path or file and loading it into memory.
It converts filesystem data into a program structure:

type SourceFile struct {
Path     string
Name     string
Content  string
}

After running through all directories in the path, the final return is:
"[]SourceFile" 

🚫 What the scanner do not do

- ❌ do not understand sintax
- ❌ do not know what a function is
- ❌ do not know what a package is
- ❌ do not know AST -> is the structured representation of source code.
- ❌ do not change anything
- ❌ do not create comment

It only answers one question:
“These are the files that exist.”.

Analogy:
 The scanner is like someone who enters an archive, retrieves all the relevant folders, and lays everything out on the table.

↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓
**BUILDER**
The builder coordinates the entire pipeline.
It does not analyze code.
It does not understand AST.
It does not extract meaning.

It only orchestrates.

Builder responsibilities:

- receives files from the scanner

- iterates through them

- detects the language

- selects the correct parser

- calls the parser

- calls the extractor

- aggregates all results into a Project model


Example flow:

lang := language.Detect(file.Path)
	and get the response:

	//"go"

Then:
astTree := parser.Parse(file)

and finally:
functions := extractor.ExtractFunctions(astTree)

The extractor reads the AST and identifies:

#functions
#receivers
#parameters
#return values
#visibility

The builder then aggregates:
project.Functions = append(project.Functions, functions...)

The builder does not extract.
It only collects knowledge.

🧠 Important concept

project.Functions is not an analyzer.

It is the project knowledge repository.

It represents everything the system knows about the source code.

projectFunctions is the final result of the pipeline of the builder. It is the repository of the project. (source code)

↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓

**PARSER**
The parser converts source code into an AST.

	It answers only one question: 

	#“Is this code syntactically valid?”

It does not interpret meaning.

↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓

**EXTRACTOR**
The extractor reads the AST and converts it into structured data.

	It answers:

	“What structures exist?”

Functions, structs, interfaces, receivers, visibility.

↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓

**SEMANTIC ANALYZER** -> The semantic analyzer is the layer that translates code structure into human meaning. It doesn’t use AI.

It uses:

naming rules
patterns
conventions
context

Example rules:

Get* → returns something

Set* → modifies something

Parse* → converts input

Load* → reads from IO

Save* → writes to IO

New* → constructor

Is* / Has* → boolean

This is the core intelligence of the engine.

↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓

⚙️ **GENERATOR**

The generator decides:

	#“What should be documented?”

It selects templates and documentation strategies.

It does not write files.

↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓

✍️ **WRITER**

The writer takes generated text and saves it to disk.

It does not think.
It only executes the creation of the final file. 

↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓

OUTPUT

🧠 To finish the explanation of the app

The extractor tells us what exists.
The semantic analyzer tells us what it means.
The generator decides what to document.
The writer applies the result.



