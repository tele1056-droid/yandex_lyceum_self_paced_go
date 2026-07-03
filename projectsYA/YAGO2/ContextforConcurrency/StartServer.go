package main

import (
	"time"
	"net/http"
)


package main

import (
    "io"
    "net/http"
    "time"
)

func StartServer(maxTimeout time.Duration) {
    // 1. Создаём обработчик, который делает запрос к localhost:8081
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Делаем запрос к другому серверу
        resp, err := http.Get("http://localhost:8081/provideData")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        // Читаем данные
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Отправляем данные клиенту
        w.Write(body)
    })

    // 2. Оборачиваем обработчик в TimeoutHandler
    // maxTimeout — время ожидания ответа от localhost:8081
    timeoutHandler := http.TimeoutHandler(
        handler,
        maxTimeout,
        "Service Unavailable", // сообщение при таймауте
    )

    // 3. Регистрируем обработчик на /readSource
    http.HandleFunc("/readSource", timeoutHandler.ServeHTTP)

    // 4. Запускаем сервер
    http.ListenAndServe("localhost:8080", nil)
}
/*
По заданию:
Входящий запрос (от клиента)	r (*http.Request)
Исходящий запрос (к другому серверу)	http.Client
Ответ клиенту	w (http.ResponseWriter)
*/

/* http.HandleFunc - Регистрирует обработчик для URL
1. http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request):
	http.HandleFunc("/", ...) - это главная страница
	func(w http.ResponseWriter, r *http.Request) - что делать, когда приходит запрос на это URL:
		- w - Ответ - туда пишем , что вернуть клиенту. Запись ответа (текст, JSON, HTML)
		- r - Запрос - откуда читаем, что прислал клиент. Данные запроса (метод, URL, заголовки, тело). r	*http.Request	Запрос — читаем данные, которые прислал клиент

	http.HandleFunc — это способ сказать: «когда кто-то зайдёт на этот URL, выполни эту функцию».

2. запуск сервера:
	http.ListenAndServe(":8080", nil)
	http.ListenAndServe принимает строку с адресом в формате: "хост:порт". Либо так ":8080", так "localhost:8080"
	http.ListenAndServe ждёт только адрес, без http://. А в 
	браузере уже добавляешь http://

3. второй аргумент: nil во втором аргументе — это «используй стандартный обработчик». В 90% случаев так и делают.

4. func(w http.ResponseWriter, r *http.Request) - Это функция-обработчик, которая вызывается, когда приходит HTTP-запрос.

*/