package main

import (
	"fmt"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	name := query.Get("msg")

	if name == "" {
		// Может быть и отсутствие, и пустота
		// Проверяем, есть ли ключ
		if _, ok := query["msg"]; !ok { //если "отсутствует", то в ok=false
			//заходим сюда, значит ключа нет
			fmt.Fprintf(w, "empty") //это по условию нужно вернуть
			return
		}

		//тут ключ есть, но пустой
		fmt.Fprintf(w, "empty")
		return // тут ретерн тоже нужен после отправки ответа обработчика
	}

	fmt.Fprint(w, name)
}

func main () {
	http.HandleFunc("/echo", echoHandler)

	http.ListenAndServe(":8080", nil)
}