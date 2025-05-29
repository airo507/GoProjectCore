package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PathValueOrError(w http.ResponseWriter, r *http.Request, name string) (string, bool) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: Если ты пишешь REST API, то тебе не нужно вызывать FormValue.
	// Лучше используй Unmarshal тела запроса на структуру Request.
	value := r.FormValue(name)

	// TODO: Логировать параметры из запроса это не безопасно. В эту функцию, например,
	// могут передать name = "access-token", а это уже чувствительная информация.
	// Логировать запросы можно, но стоит задумываться над логированием паролей, токенов и тп.
	//
	// TODO: Тут не используешь отдельный логгер. Лучше настроить один логгер и
	// использовать везде. Так у тебя в логах будет куча разных по формату сообщений.
	fmt.Println(value)
	if value == "" {
		// TODO: Функция делает слишком много.
		// - Достает параметр.
		// - Пишет ответ.
		// А что если параметр опциональный? Запись ответа лучше оставить на сторону вызывающего.
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(DefaultResponse{
			Code:    InvalidRequest,
			Message: fmt.Sprintf("invalid path parameter '%s'", name),
		})
		return "", false
	}

	return value, true
}
