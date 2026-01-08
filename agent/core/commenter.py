import os
from pathlib import Path

from openai import OpenAI


class GoCodeCommenter:
    def __init__(self, prompt_path: str):
        self.client = OpenAI(api_key=os.getenv("OPENAI_API_KEY"))
        self.prompt_template = self._load_prompt(prompt_path)

    def _load_prompt(self, prompt_path: str) -> str:
        path = Path(prompt_path)
        if not path.exists():
            raise FileNotFoundError(f"Prompt file not found: {prompt_path}")
        return path.read_text(encoding="utf-8")

    def comment_code(self, code: str) -> str:
        prompt = self.prompt_template + "\n\n" + code

        response = self.client.responses.create(
            model="gpt-4.1-mini",
            input=prompt,
            temperature=0.2,
        )

        content = response.output_text

        if not content:
            raise ValueError("AI returned empty content. Cannot generate comments.")

        return content.strip()
