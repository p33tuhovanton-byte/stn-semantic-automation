import json
import os
from fastapi import FastAPI, HTTPException
import jsonschema

app = FastAPI(title="Python STN Service", version="1.0.0")

def load_json_schema(name: str) -> dict:
    schema_path = f"../schemas/{name.lower()}.json"
    if not os.path.exists(schema_path):
        raise HTTPException(status_code=404, detail=f"Описание сути '{name}' не найдено.")
    with open(schema_path, "r", encoding="utf-8") as file:
        return json.load(file)

@app.post("/execute/{naimenovanie}")
def execute_semantic_node(naimenovanie: str, payload: dict):
    # 1. Считываем внешнее описание
    schema = load_json_schema(naimenovanie)
    
    # 2. Автоматическая валидация по стандарту JSON Schema (Кто и С какой целью)
    try:
        jsonschema.validate(instance=payload, schema=schema)
    except jsonschema.exceptions.ValidationError as err:
        raise HTTPException(status_code=400, detail=f"Ошибка валидации контекста: {err.message}")
        
    # 3. Финальный автоматизм исполнения сути
    return {
        "status": "Успех",
        "message": f"Наименование '{naimenovanie}' успешно обработано на языке Python.",
        "details": {
            "выявленная_суть": "Контекст полностью соответствует внешнему описанию",
            "данные": payload
        }
    }
