package main

import (
	 "encoding/json"
	 "fmt"
	 "io/ioutil"
	 "net/http"
	 "os"
	 "strings"


   "://github.com"
)

// Структура для автоматического ответа
type Response struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}

func handleSemanticNode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Извлекаем наименование из URL пути
	parts := strings.Split(r.URL.Path, "/")
	naimenovanie := parts[len(parts)-1]

	// 1. Считываем внешнее изолированное описание
	schemaPath := fmt.Sprintf("../schemas/%s.json", strings.ToLower(naimenovanie))
	if _, err := os.Stat(schemaPath); os.IsNotExist(err) {
		http.Error(w, fmt.Sprintf("Описание сути '%s' не найдено", naimenovanie), http.StatusNotFound)
		return
	}

	// Читаем тело запроса (payload)
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	
	// 2. Автоматическая валидация по международному стандарту JSON Schema
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
	documentLoader := gojsonschema.NewBytesLoader(bodyBytes)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !result.Valid() {
		var errors []string
		for _, desc := range result.Errors() {
			errors = append(errors, desc.String())
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Ошибка валидации контекста (Кто/Зачем): %s", strings.Join(errors, ", "))
		return
	}

	// Парсим данные для успешного ответа
	var payload map[string]interface{}
	json.Unmarshal(bodyBytes, &payload)

	// 3. Финальный автоматизм исполнения сути
	w.Header().Set("Content-Type", "application/json")
	resp := Response{
		Status:  "Успех",
		Message: fmt.Sprintf("Наименование '%s' успешно обработано на языке Go.", naimenovanie),
		Details: map[string]interface{}{
			"выявленная_суть": "Контекст полностью соответствует внешнему описанию",
			"данные":          payload,
		},
	}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/execute/", handleSemanticNode)
	fmt.Println("🚀 Go-сервис семантического автоматизма запущен на порту :8010")
	http.ListenAndServe(":8010", nil)
}
