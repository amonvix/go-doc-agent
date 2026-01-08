# languages/python_lang.py
from core.language_registry import LanguageRegistry


class PythonLanguage: ...


LanguageRegistry.register("python", PythonLanguage)
