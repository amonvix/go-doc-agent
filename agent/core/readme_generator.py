import os
from pathlib import Path
from typing import Dict

from openai import OpenAI


class ReadmeGenerator:
    def __init__(self, template_path: str):
        self.client = OpenAI(api_key=os.getenv("OPENAI_API_KEY"))
        self.template_path = Path(template_path)
        self.template = self._load_template()

    def _load_template(self) -> str:
        if not self.template_path.exists():
            raise FileNotFoundError(f"README template not found: {self.template_path}")
        return self.template_path.read_text(encoding="utf-8")

    def generate_sections(self, go_code: str) -> Dict[str, str]:
        prompt = f"""
You are a senior Go engineer and technical writer.

Your task is to analyze the following Go code and generate content to fill a predefined README template.

IMPORTANT RULES (STRICT - MUST FOLLOW):
- You MUST NOT use the word "example" anywhere.
- You MUST NOT use teaching, tutorial, or explanatory tone.
- You MUST NOT start sentences with "This example", "This code", "This program", or similar.
- You MUST write in a neutral, professional, production documentation tone.
- You MUST NOT explain concepts for learning purposes.
- You MUST NOT describe what the code does step by step.
- You MUST NOT use markdown headings.
- You MUST NOT include code blocks.
- You MUST NOT repeat the topic name inside the text.
- PURPOSE must start directly with the technical intent, not with introductory phrases.
- Output MUST follow the exact format. If you break the format, the output is invalid.

If you violate any rule, regenerate the output internally before returning.

You will generate exactly 4 sections:

1. TOPIC_NAME  
A short, technical, properly capitalized title using title case. Use correct technical formatting (e.g. "8-bit", "Go").

2. CONCEPTS  
A bullet list of key technical concepts covered (each line starting with "- ").

3. PURPOSE  
Describe the technical intent of the code in an impersonal, documentation-style tone. Do not use teaching language. Do not use "illustrate", "demonstrate", or similar verbs.

4. NOTES  
Technical notes, caveats, or best practices relevant to this code. Bullet points. No generic statements.

Output format MUST be exactly:

TOPIC_NAME:
<text>

CONCEPTS:
- item 1
- item 2

PURPOSE:
<text>

NOTES:
- item 1
- item 2

Do not add anything outside this structure.

Go code:
{go_code}
"""

        response = self.client.responses.create(
            model="gpt-4.1-mini",
            input=prompt,
            temperature=0.3,
        )

        text = response.output_text.strip()

        sections = {
            "topic_name": "",
            "concepts_list": "",
            "purpose_text": "",
            "notes_text": "",
        }

        current_key = None
        buffer = []

        for line in text.splitlines():
            line = line.strip()

            if line.lower().startswith("topic"):
                current_key = "topic_name"
                buffer = []
                continue

            if line.lower().startswith("concept"):
                current_key = "concepts_list"
                buffer = []
                continue

            if line.lower().startswith("purpose"):
                current_key = "purpose_text"
                buffer = []
                continue

            if line.lower().startswith("notes"):
                current_key = "notes_text"
                buffer = []
                continue

            if current_key:
                buffer.append(line)
                sections[current_key] = "\n".join(buffer)

        return sections

    def build_readme(self, sections: Dict[str, str]) -> str:
        readme = self.template
        readme = readme.replace("{{topic_name}}", sections["topic_name"])
        readme = readme.replace("{{concepts_list}}", sections["concepts_list"])
        readme = readme.replace("{{purpose_text}}", sections["purpose_text"])
        readme = readme.replace("{{notes_text}}", sections["notes_text"])
        return readme

    def write_readme(self, folder: Path, content: str) -> None:
        readme_path = folder / "README.md"
        readme_path.write_text(content, encoding="utf-8")
