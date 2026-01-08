from pathlib import Path

EXCLUDED_DIRS = {"vendor", "node_modules", ".git", "__pycache__"}


def find_go_files(base_path: str) -> list[Path]:
    base = Path(base_path)
    go_files = []

    for path in base.rglob("*.go"):
        if any(part in EXCLUDED_DIRS for part in path.parts):
            continue
        go_files.append(path)

    return go_files
