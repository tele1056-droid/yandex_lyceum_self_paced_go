package main

import (
	"fmt"
	"net/http"
	"unicode"
)

//будем писать веб-сервер через обработчик http.HandleFunc
//отдельно пишем хендлер
func helloStrangerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Запрос получен!") //тоже будет выводится в терминале с сервером при отправке запроса (через curl)

	//первым делом нужно из параметра запроса достать имя для возврата приветствия, для этого воспользуемся паттерном проверки на "пустой" или "отсутствует"

	//здесь получаем значения параметра name, например из URL вида "http://example.com/?name=John&age=25"

	//проверяем параметр на "пустой" или "отсутствует", это разные вещи и нужно и то и то проверять
	query := r.URL.Query()

	name := query.Get("name")
	if name == "" {
		// Может быть и отсутствие, и пустота
		// Проверяем, есть ли ключ
		if _, ok := query["name"]; !ok { //если "отсутствует", то в ok=false
			//заходим сюда, значит ключа нет
			fmt.Fprintf(w, "hello stranger") //это по условию нужно вернуть
			return
		}

		//тут ключ есть, но пустой
		fmt.Fprintf(w, "hello stranger")
	}

	//проверяем если не только английские буквы
	if !CheckOnlyASCII(name) {
		//сюда мы попадаем если CheckOnlyASCII вернул false - т.е. не все символы ASCII
		fmt.Fprintf(w, "hello dirty hacker")
		return
	}

	//дефолт ответ с именем из параметра запроса
	fmt.Fprintf(w, name)
	
}

//эта функция проверяет символы строки, если все они в ASCII - true, иначе false
func CheckOnlyASCII(s string) bool {
	for _, letter := range s {
		if letter > unicode.MaxASCII {
			return false
		}
	}
	return true
}

/*
Запрос			Get("name")	query["name"]
/?name=John		"John"		["John"], ok=true
/?name=			""			[""], ok=true
/				""			nil, ok=false
*/

/*
Два состояния параметра
Состояние		Что значит						Как проверить
Отсутствует		Ключа нет в запросе				val, ok := r.URL.Query()["key"]
Пустой			Ключ есть, но значение пустое	val := r.URL.Query().Get("key") → ""
*/

/*
					ПРИМЕРЫ:
1. Параметр отсутствует:
// Запрос: http://example.com/
name := r.URL.Query().Get("name") // "" (пустая строка)

2. Параметр есть, но пустой
// Запрос: http://example.com/?name=
name := r.URL.Query().Get("name") // "" (пустая строка)

	НО, ПРОБЛЕМА: Get() не отличает отсутствие от пустоты
*/

func main() {
	http.HandleFunc("/", helloStrangerHandler)

	fmt.Println("Server started on :8080") // этот принт будет выводится, в терминале сервера, когда будет запускаться сервер (go run main.go)
	http.ListenAndServe(":8080", nil)
}

/*
Основные компоненты
Компонент			Что делает
http.HandleFunc		Регистрирует обработчик для пути
http.ResponseWriter	Пишет ответ клиенту
*http.Request		Содержит данные запроса
http.ListenAndServe	Запускает сервер
*/

//Простое правило: используйте http.HandleFunc для большинства случаев — это проще и короче. Переходите на http.Handler, когда вам нужно хранить состояние или внедрять зависимости.