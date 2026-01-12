from core.step_registry import StepRegistry


class GenerateCodeStep:
    def execute(self, context):
        print(f"⚙️ Generating code for language: {context.language}")
        context.code = f"# Code generated for {context.language}"


StepRegistry.register("generate_code", GenerateCodeStep)
