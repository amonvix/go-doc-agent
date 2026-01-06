from pathlib import Path


def write_commented_file(original_path: Path, commented_code: str) -> None:
    original_path.write_text(commented_code, encoding="utf-8")
