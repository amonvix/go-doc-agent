from agent.core.language_registry import LanguageRegistry


class PythonLanguage:
    name = "python"

    def get_extension(self) -> str:
        return ".py"

    def format(self, code: str) -> str:
        return code.strip()

    def generate(self, context) -> str:
        request = context.request

        code = f"""
def main():
    print("Request: {request}")


if __name__ == "__main__":
    main()
"""
        return code


# auto-registro
LanguageRegistry.register("python", PythonLanguage)
