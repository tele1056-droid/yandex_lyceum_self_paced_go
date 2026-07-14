package main

import (
	"fmt"
	"net/http"
	"log/slog"
	"log"
)

// Middleware (Промежуточный слой) — это «прослойка» между запросом и ответом

//в данном случае нам нужен middleware для записи запросов в лог
// это будет func LoggingMiddleware()

//и еще нужно использовать http.ServeMux — механизм маршрутизации. Это диспетчер, который получает запрос и направляет его к нужному обработчику в зависимости от пути. mux := http.NewServeMux()
//без mux , используется глобальный DefaultServeMux
/*
С nil	http.ListenAndServe(":8080", nil)	Используется глобальный DefaultServeMux
С mux	http.ListenAndServe(":8080", mux)	Используется свой маршрутизатор
*/

/*
Почему использовать свой mux:

Причина			Объяснение
Контроль		Не мешаешь другим пакетам
Безопасность	Не конфликтуешь с глобальным
Организация		Легче группировать пути
Middleware		Проще оборачивать
*/

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//логируем инф. о запросе, используем slog, и нужно передавать пары ключ-занчение
		slog.Info("incoming request",
		"method", r.Method,
		"path", r.URL.Path)

		//сигнатура slog.Info("сообщение", "ключ1", "значение1", "ключ2", "значение2")

		//Передаём управление следующему обработчику, ну основному хендлеру
		next.ServeHTTP(w, r) // это mux
	})
}
/*
LoggingMiddleware возвращает новый http.Handler, который внутри себя:

 - Сначала выполняет свой код (логирование)

 - Потом вызывает next.ServeHTTP(w, r) — это и есть mux

 - mux вызывает homeHandler

Middleware оборачивает роутер (mux), а роутер вызывает хендлер.
*/

//наш основной обработчик (хендлер)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	//обрабатываем запрос и возвращаем строку клиенту
	fmt.Fprint(w, "Hello, middleware!")
}


func main() {
	mux := http.NewServeMux()
	//--mux.HandleFunc("/hello", helloHandler)

	//создаём обработчик для маршрута "/hello"
	hello := http.HandlerFunc(helloHandler)

	// применяем logging middleware (Logger) к обработчику "/hello"
	mux.Handle("/hello", Logger(hello))

	//--Оборачиваем mux в Middleware
	//--handler := Logger(mux)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	} // --handler = LoggingMiddleware(mux)
}


/*
Клиент → Запрос → Middleware → Обработчик → Ответ
                     │            │
                     │            │
            логирование      основная работа
*/

/*
Запрос → LoggingMiddleware
    │
    ├── log.Println("Запрос начат")
    │
    ├── next.ServeHTTP(w, r) → основной хендлер
    │                               │
    │                               ├── fmt.Fprintf(w, "Hello!")
    │                               │
    │                               └── return
    │
    ├── log.Println("Запрос завершён")
    │
    └── Ответ клиенту
*/


//slog.Info("msg", "key", "value")

/*
r *http.Request содержит всю информацию о запросе

Что достать			Код
Метод				r.Method
Путь				r.URL.Path
Полный URL			r.URL.String()
Параметры (query)	r.URL.Query().Get("key")
Заголовки			r.Header.Get("Content-Type")
Тело				io.ReadAll(r.Body)
IP клиента			r.RemoteAddr
Контекст			r.Context()
*/