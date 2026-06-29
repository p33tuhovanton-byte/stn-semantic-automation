use actix_web::{post, web, App, HttpServer, Responder, HttpResponse};
use serde_json::Value;
use std::fs;
use std::path::Path;

#[post("/execute/{naimenovanie}")]
async fn execute_semantic_node(
    naimenovanie: web::Path<String>,
    payload: web::Json<Value>,
) -> impl Responder {
    let name = naimenovanie.into_inner().to_lowercase();
    
    // 1. Считываем внешнее изолированное описание
    let schema_path = format!("../schemas/{}.json", name);
    if !Path::new(&schema_path).exists() {
        return HttpResponse::NotFound().body(format!("Описание сути '{}' не найдено", name));
    }

    let schema_content = fs::read_to_string(&schema_path).unwrap();
    let schema_json: Value = serde_json::from_str(&schema_content).unwrap();

    // 2. Автоматическая валидация по международному стандарту JSON Schema
    match jsonschema::JSONSchema::compile(&schema_json) {
        Ok(compiled_schema) => {
            if let Err(errors) = compiled_schema.validate(&payload) {
                let error_messages: Vec<String> = errors.map(|e| e.to_string()).collect();
                return HttpResponse::BadRequest().body(format!(
                    "Ошибка валидации контекста (Кто/Зачем): {}",
                    error_messages.join(", ")
                ));
            }
        }
        Err(_) => return HttpResponse::InternalServerError().body("Ошибка компиляции схемы"),
    }

    // 3. Финальный автоматизм исполнения сути
    HttpResponse::Ok().json(serde_json::json!({
        "status": "Успех",
        "message": format!("Наименование успешно обработано на языке Rust."),
        "details": {
            "выявленная_суть": "Контекст полностью соответствует внешнему описанию",
            "полученные_данные": payload.into_inner()
        }
    }))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    println!("🚀 Rust-сервис семантического автоматизма запущен на порту :8020");
    HttpServer::new(|| {
        App::new().service(execute_semantic_node)
    })
    .bind(("127.0.0.1", 8020))?
    .run()
    .await
}
