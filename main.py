from agent.scanner import find_go_files


def main():
    project_path = input("Enter the path to your Go project: ").strip()

    go_files = find_go_files(project_path)

    print(f"Found {len(go_files)} Go files:")
    for file in go_files:
        print(f" - {file}")


if __name__ == "__main__":
    main()
