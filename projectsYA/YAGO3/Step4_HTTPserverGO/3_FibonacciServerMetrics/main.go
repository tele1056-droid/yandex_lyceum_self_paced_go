package main

import (
	"fmt"
	"net/http"
)

//тут по заданию нужно добавтиь хендлер /metrics

//Добавить хендлер = сказать серверу: «на этот путь отвечай этой функцией». Т.е. «Добавить хендлер» — значит зарегистрировать обработчик для определённого пути (URL)

var count int = 0

func fib(n int) int {
    if n <= 1 {
        return n
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}


func fibHandler(w http.ResponseWriter, r *http.Request) {
    // Увеличиваем счётчик при каждом запросе, и тут сначала подсчитываем фибоначи, а потом увелич счетчик
	result := fib(count)
    count++
    
    fmt.Fprint(w, result)
}

func metrickHandler(w http.ResponseWriter, r *http.Request) {
	countReq := +count
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", countReq)
}



func main() {
    http.HandleFunc("/", fibHandler) //вот эта строчка - и есть хендлер для "/", то есть регистрируем обработчик для "/"

	//добавляем хендлер /metrics который возвращает число запросов
	http.HandleFunc("/metrics", metrickHandler)

	//и хендлеры мы размещаем до запуска сервера сначала все http.HandleFunc, в конце http.ListenAndServe

    //fmt.Println("Server started on :8080")
    http.ListenAndServe(":8080", nil)
}


/*
Хендлер (handler) — это функция, которая отвечает на HTTP-запросы по определённому пути.

Когда пользователь заходит на http://localhost:8080/hello, сервер должен знать, какая функция должна обработать этот запрос. Регистрация этой функции и есть добавление хендлера.
*/