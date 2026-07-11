package main

import (
	"log/slog"
	//"time"
)

func LogHTTPRequest(logger *slog.Logger, method, path string, status int, durationMs int64) {
	logger.Info("http request",
	"Method", method,
	"Path", path,
	"Status", status,
	"duration_ms", durationMs, //тут была проблема - Ключи в логах должны совпадать с тегами json:, ключ нужно было писать такой `json:"duration_ms"`.
	// И регистр тоже важен, в данном случае остальные ключи начинаются с заглавной буквы, а в json с прописной, но slog (стандартный логгер) автоматически приводит ключи к нижнему регистру при выводе в в JSON.

	//тут для логгирования durationMs - нужно преобразовать тип durationMs int64 к time.Duration и тогда в логгах будет 150ms.
	// но тут нужно только число (без ms)
)
}

/*
	АВТОТЕТСТ
type httpLogEntry struct {
	Level      string `json:"level"`
	Msg        string `json:"msg"`
	Method     string `json:"method"`
	Path       string `json:"path"`
	Status     int    `json:"status"`
	DurationMs int64  `json:"duration_ms"`
}
*/

func main() {
	logger := slog.Default()
	tests := struct {
		method string
		path string
		status int
		duration int64

	}{
		method: "GET",
		path: "/api/user",
		status: 200,
		duration: 123,
	}

	LogHTTPRequest(logger, tests.method, tests.path, tests.status, tests.duration)
}