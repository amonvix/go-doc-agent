from pathlib import Path

from openai import OpenAI


class ReadmeGenerator:
    def __init__(self, template_path: str):
        self.template_path = Path(template_path)
        self.client = OpenAI()

        if not self.template_path.exists():
            raise FileNotFoundError(f"README template not found: {template_path}")

        self.template = self.template_path.read_text(encoding="utf-8")

    def generate_sections(self, go_code: str) -> dict:
        prompt = f"""
You are a senior Go engineer and technical writer.

Your task is to analyze the following Go code and generate content to fill a predefined README template.

IMPORTANT RULES:
- You MUST NOT create headings, titles, or section names.
- You MUST NOT repeat the topic name inside the text.
- You MUST NOT write phrases like "This example demonstrates" or similar.
- You MUST NOT explain what the code does line by line.
- You MUST NOT use markdown headings.
- You MUST NOT include code blocks.
- You MUST write only the raw content for each field.

You will generate exactly 4 sections:

1. TOPIC_NAME  
A short, technical, descriptive title (no punctuation, no markdown, no extra words).

2. CONCEPTS  
A bullet list of key technical concepts covered (each line starting with "- ").

3. PURPOSE  
A concise technical explanation of the intent of the code. No tutorial tone. No didactic language.

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

    def build_readme(self, sections: dict) -> str:
        readme = self.template
        readme = readme.replace("{{topic_name}}", sections["topic_name"])
        readme = readme.replace("{{concepts_list}}", sections["concepts_list"])
        readme = readme.replace("{{purpose_text}}", sections["purpose_text"])
        readme = readme.replace("{{notes_text}}", sections["notes_text"])
        return readme

    def write_readme(self, folder: Path, content: str) -> None:
        readme_path = folder / "README.md"
        readme_path.write_text(content, encoding="utf-8")
