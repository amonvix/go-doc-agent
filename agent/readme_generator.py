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

Based on the following Go code, generate:
1. A concise topic name.
2. A list of key concepts covered (bullet points).
3. A short purpose explanation.
4. Technical notes, caveats, or best practices.

Guidelines:
- Be concise and technical.
- Avoid tutorial tone.
- Do NOT include code blocks.
- Output plain text only.

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
