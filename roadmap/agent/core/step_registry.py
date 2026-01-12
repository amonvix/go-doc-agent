class StepRegistry:
    _steps = {}

    @classmethod
    def register(cls, name, step_cls):
        cls._steps[name] = step_cls

    @classmethod
    def get(cls, name):
        if name not in cls._steps:
            raise ValueError(f"Step '{name}' not registered")
        return cls._steps[name]
