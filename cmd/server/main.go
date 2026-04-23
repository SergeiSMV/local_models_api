package main

import (
	// encoding/json нужен для сериализации Go-структур/данных в JSON.
	"encoding/json"
	// log используется для вывода служебных сообщений и ошибок в консоль.
	"log"
	// net/http — стандартный пакет для HTTP-сервера и маршрутизации запросов.
	"net/http"
)

// main — точка входа приложения.
func main() {
	// Создаем маршрутизатор (mux), который направляет запросы в нужные обработчики.
	mux := http.NewServeMux()
	// Регистрируем endpoint проверки состояния сервиса.
	// Формат "GET /health" означает: обрабатываем только GET-запросы на путь /health.
	mux.HandleFunc("GET /health", handleHealth)

	// Логируем, что сервер стартует на порту 8080.
	log.Println("server starting on :8080")
	// Запускаем HTTP-сервер и передаем ему маршрутизатор.
	// Если сервер завершился с ошибкой — останавливаем программу.
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

// handleHealth отвечает на запрос /health и возвращает JSON со статусом сервиса.
func handleHealth(w http.ResponseWriter, r *http.Request) {
	// Сообщаем клиенту, что в ответе будет JSON.
	w.Header().Set("Content-Type", "application/json")
	// Отправляем простой JSON-объект: {"status":"ok"}.
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
