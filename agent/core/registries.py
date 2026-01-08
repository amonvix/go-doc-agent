class AdapterRegistry:
    _adapters = {}

    @classmethod
    def register(cls, name, adapter_cls):
        cls._adapters[name] = adapter_cls

    @classmethod
    def get(cls, name):
        if name not in cls._adapters:
            raise ValueError(f"Adapter '{name}' not registered")
        return cls._adapters[name]


class LanguageRegistry:
    _languages = {}

    @classmethod
    def register(cls, name, language_cls):
        cls._languages[name] = language_cls

    @classmethod
    def get(cls, name):
        if name not in cls._languages:
            raise ValueError(f"Language '{name}' not registered")
        return cls._languages[name]
