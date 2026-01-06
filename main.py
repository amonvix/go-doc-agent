import os

from agent.commenter import GoCodeCommenter
from agent.readme_generator import ReadmeGenerator
from agent.scanner import find_go_files
from agent.writer import write_commented_file


def main():
    raw_path = input("Enter the path to your Go project: ").strip()
    project_path = os.path.expanduser(raw_path)

    go_files = find_go_files(project_path)

    if not go_files:
        print("No Go files found.")
        return

    commenter = GoCodeCommenter(prompt_path="prompts/go_comment_prompt.txt")
    readme_generator = ReadmeGenerator(template_path="templates/readme_template.md")

    print(f"\nFound {len(go_files)} Go files. Processing...\n")

    for go_file in go_files:
        print(f"🧠 Commenting: {go_file}")

        code = go_file.read_text(encoding="utf-8")
        print("📤 Sending code to AI...")

        commented_code = commenter.comment_code(code)
        print("📥 Received response from AI")

        # 1) Overwrite the original Go file (no new main files)
        write_commented_file(go_file, commented_code)
        print("✅ Updated:", go_file)

        # 2) Generate README.md in the same folder using the commented code
        updated_code = go_file.read_text(encoding="utf-8")

        sections = readme_generator.generate_sections(updated_code)
        readme_content = readme_generator.build_readme(sections)
        readme_generator.write_readme(go_file.parent, readme_content)

        print("📄 README generated in:", go_file.parent, "\n")

    print("🎉 All files processed successfully.")


if __name__ == "__main__":
    main()
