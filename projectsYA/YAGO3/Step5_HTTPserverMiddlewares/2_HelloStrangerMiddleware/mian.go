/*основные ошибки с автотестом:
1. Всегда используй http.HandlerFunc для передачи функции в middleware.
Т.е. мой основной хендлер (helloHandler) в автотесте передается как объект с типом http.Handler, и вызывается в цепочке  middleware. А у меня хендлер оформлен как функция

					РЕШЕНИЕ ПРОБЛЕМЫ АВТОТЕСТА
поменял у Sanitize и SetDefaultName:
(next http.Handler) http.Handler -> (next http.HandlerFunc) http.HandlerFunc и автотест прошел. Т.е. замени Sanitize и SetDefaultName на работу с http.HandlerFunc напрямую
*/


package main

import (
    "fmt"
    "net/http"
)

// ====== ВСПОМОГАТЕЛЬНАЯ ФУНКЦИЯ ======

// CheckOnlyASCII проверяет, что строка состоит только из английских букв
func CheckOnlyASCII(s string) bool {
    if s == "" {
        return false
    }
    for _, r := range s {
        if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
            return false
        }
    }
    return true
}

// ====== MIDDLEWARE ======

// Sanitize — проверяет, что name состоит только из английских букв
func Sanitize(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")

        // Если имя пустое — пропускаем (дальше SetDefaultName)
        if name == "" {
            next.ServeHTTP(w, r)
            return
        }

        // Если имя невалидно — возвращаем "hello dirty hacker"
        if !CheckOnlyASCII(name) {
            fmt.Fprint(w, "hello dirty hacker")
            return
        }

        // Всё ок — передаём дальше
        next.ServeHTTP(w, r)
    })
}

// SetDefaultName — подставляет "stranger", если name отсутствует или пустой
func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        query := r.URL.Query()
        name := query.Get("name")

        // Если name пустой или отсутствует
        if name == "" {
            // Проверяем, есть ли ключ "name" в запросе
            if _, ok := query["name"]; !ok || name == "" {
                fmt.Fprint(w, "hello stranger")
                return
            }
        }

        // Передаём дальше (имя есть и валидно)
        next.ServeHTTP(w, r)
    })
}

// ====== ОСНОВНОЙ ОБРАБОТЧИК ======

// HelloHandler — возвращает "hello <name>"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    fmt.Fprintf(w, "hello %s", name)
}

// ====== ЗАПУСК СЕРВЕРА ======

func main() {
    mux := http.NewServeMux()

    // Оборачиваем обработчик в middleware
    // Порядок: Sanitize → SetDefaultName → HelloHandler
    handler := SetDefaultName(Sanitize(http.HandlerFunc(HelloHandler)))

    mux.Handle("/hello", handler)

    fmt.Println("Server started on :8080")
    http.ListenAndServe(":8080", mux)
}

//Порядок выполнения middleware определяется тем, как ты оборачиваешь их, а не тем, как они расположены в коде.
/*
Можно выстраивать цепочку из middleware последовательно
Запрос → Middleware 1 → Middleware 2 → Middleware 3 → Хендлер
           │               │               │
       логирование    авторизация    сжатие ответа


*/

/*				ПРИМЕР ЦЕПОЧКИ middleware
// Middleware 1: логирование
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        slog.Info("Запрос", "method", r.Method, "path", r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

// Middleware 2: авторизация
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return // ❌ цепочка прерывается
        }
        next.ServeHTTP(w, r)
    })
}

// Middleware 3: замер времени
func TimerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        slog.Info("Время выполнения", "duration", time.Since(start))
    })
}

					СБОРКА ЦЕПОЧКИ
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", homeHandler)

    // Цепочка: Auth → Timer → Logging → mux
    handler := AuthMiddleware(
        TimerMiddleware(
            LoggingMiddleware(mux),
        ),
    )

    http.ListenAndServe(":8080", handler)
}
*/