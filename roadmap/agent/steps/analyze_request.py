from core.step_registry import StepRegistry


class AnalyzeRequestStep:
    def execute(self, context):
        print(f"🧠 Analyzing request: {context.request}")
        context.metadata["analysis"] = "Request analyzed"


StepRegistry.register("analyze_request", AnalyzeRequestStep)
