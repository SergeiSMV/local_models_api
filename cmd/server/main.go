package main

import (
	// Пакет для кодирования данных в JSON-формат.
	"encoding/json"
	// Логирование служебных сообщений и ошибок.
	"log"
	// Базовый HTTP-сервер из стандартной библиотеки Go.
	"net/http"
	// Работа со временем (здесь используется для timeout).
	"time"

	// Chi — легковесный роутер (маршрутизация HTTP-запросов).
	"github.com/go-chi/chi/v5"
	// Набор готовых middleware для логов, recovery, timeout и др.
	"github.com/go-chi/chi/v5/middleware"
)

// main — точка входа приложения.
func main() {
	// Создаем роутер, который будет обрабатывать входящие HTTP-запросы.
	r := chi.NewRouter()
	// Логирует каждый запрос (метод, путь, статус, время выполнения).
	r.Use(middleware.Logger)
	// Перехватывает panic внутри хендлеров, чтобы сервер не падал целиком.
	r.Use(middleware.Recoverer)
	// Ограничивает максимальное время обработки одного запроса.
	r.Use(middleware.Timeout(6 * time.Minute))

	// Регистрируем GET endpoint для проверки "живости" сервиса.
	r.Get("/health", handleHealth)

	// Сообщаем в лог, что сервер запускается на порту 8080.
	log.Println("server starting on :8080")
	// Запускаем HTTP-сервер. При фатальной ошибке завершаем программу.
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

// handleHealth возвращает простой JSON-ответ со статусом сервиса.
func handleHealth(w http.ResponseWriter, r *http.Request) {
	// Указываем, что ответ будет в формате JSON.
	w.Header().Set("Content-Type", "application/json")
	// Отправляем клиенту JSON: {"status":"ok"}.
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
