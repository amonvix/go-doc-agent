import yaml
from core.context import AgentContext
from core.step_registry import StepRegistry


class AgentEngine:
    def __init__(self, pipelines_path):
        with open(pipelines_path, "r") as f:
            self.pipelines = yaml.safe_load(f)

    def run(self, pipeline_name, request, language, api):
        if pipeline_name not in self.pipelines:
            raise ValueError(f"Pipeline '{pipeline_name}' not found")

        context = AgentContext()
        context.request = request
        context.language = language
        context.api = api

        steps = self.pipelines[pipeline_name]["steps"]

        for step_name in steps:
            step_cls = StepRegistry.get(step_name)
            step_instance = step_cls()
            step_instance.execute(context)

        return context
