# adapters/openai_adapter.py
from core.registry import AdapterRegistry


class OpenAIAdapter: ...


AdapterRegistry.register("openai", OpenAIAdapter)
