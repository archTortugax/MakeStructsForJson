import json

def textFromJsonFile(path: str) -> str:
    with open(path, "r", encoding="utf-8") as f:
        return f.read()

def dataFromJson(text: str):
    return json.loads(text)
