# core/language_registry.py


class LanguageRegistry:
    _languages = {}

    @classmethod
    def register(cls, name, handler_cls):
        cls._languages[name] = handler_cls

    @classmethod
    def get(cls, name):
        if name not in cls._languages:
            raise ValueError(f"Language '{name}' not registered")
        return cls._languages[name]
