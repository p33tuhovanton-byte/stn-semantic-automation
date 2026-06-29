library(plumber)
library(jsonvalidate)
library(jsonlite)

# Инициализируем API
api <- pr()

# Определяем семантический узел обработки наименования
api %>% pr_post("/execute/<naimenovanie>", function(naimenovanie, req, res) {
  
  # Извлекаем сырое тело запроса (payload)
  payload_raw <- req$postBody
  
  # 1. Считываем внешнее изолированное описание
  schema_path <- paste0("../schemas/", tolower(naimenovanie), ".json")
  
  if (!file.exists(schema_path)) {
    res$status <- 404
    return(list(error = paste0("Описание сути '", naimenovanie, "' не найдено.")))
  }
  
  # 2. Автоматическая валидация по международному стандарту JSON Schema
  # Пакет jsonvalidate проверяет структуру (Кто и С какой целью) автоматически
  validator <- jsonvalidate::json_validator(schema_path, engine = "ajv")
  validation_result <- validator(payload_raw, verbose = TRUE)
  
  if (!validation_result) {
    res$status <- 400
    # Извлекаем ошибки валидации структуры человека или цели
    errors <- attr(validation_result, "errors")
    return(list(
      error = "Ошибка автоматической валидации контекста (Кто/Зачем).",
      details = errors
    ))
  }
  
  # Если валидация пройдена, парсим данные для успешного ответа
  payload_data <- jsonlite::fromJSON(payload_raw)
  
  # 3. Финальный автоматизм исполнения сути
  res$status <- 200
  return(list(
    status = "Успех",
    message = paste0("Наименование '", naimenovanie, "' успешно обработано на языке R."),
    details = list(
      vyyavlennaya_sut = "Контекст полностью соответствует внешнему описанию",
      dannie = payload_data
    )
  ))
})

# Запускаем R-сервис семантического автоматизма
print("🚀 R-сервис семантического автоматизма запущен на порту :8030")
api %>% pr_run(host = "127.0.0.1", port = 8030)
