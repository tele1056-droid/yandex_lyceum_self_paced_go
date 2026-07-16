package main

import (
	"fmt"
	//"unicode"
	"net/http"
	//"log"
)

/*основные ошибки с автотестом:
1. Всегда используй http.HandlerFunc для передачи функции в middleware.
Т.е. мой основной хендлер (helloHandler) в автотесте передается как объект с типом http.Handler, и вызывается в цепочке  middleware. А у меня хендлер оформлен как функция

2. 
*/


//основной хендлер, и тут определяем именно функцию-обработчик
 func HelloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	//обрабатываем запрос и возвращаем строку клиенту
	fmt.Fprintf(w, "hello %s", name)
}

// 2. Создаём переменную типа http.Handler, а здесь к helloHandler делаем тип http.Handler
//var HelloHandler http.Handler = http.HandlerFunc(helloHandler)

//тут проверка мидлваре что name (только буквы)
func Sanitize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		name := query.Get("name")

		//проверяем только на англ. буквы (если есть цифры, не англ. буквы, другие символы)
		if !CheckOnlyASCII(name) {
			fmt.Fprint(w, "hello dirty hacker")
			return 
		}

		next.ServeHTTP(w, r) //не забываем передать управление след. обработчику

	})
}

//тут отдельный мидлаваре для параметр name отсутствует или пустой
func SetDefaultName(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		name := query.Get("name")
		if name == "" {
			// Может быть и отсутствие, и пустота
		// Проверяем, есть ли ключ
		if _, ok := query["name"]; !ok { //если "отсутствует", то в ok=false
			//заходим сюда, значит ключа нет
			fmt.Fprint(w, "hello stranger") //это по условию нужно вернуть
			return
			}
			//тут ключ есть, но пустой
			fmt.Fprintf(w, "hello stranger")
			return
		}

		next.ServeHTTP(w, r) //не забываем передать управление след. обработчику
	})
}



func CheckOnlyASCII(s string) bool {
	countASCII := 0
	for _, letter := range s {
		//тут нужно смотреть символы в этих диапазонах 65–90, 97–122
		if (letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z'){
			countASCII++
			continue
		}
	}
	if countASCII == len(s) {
		return true
	} else {
		return false
	}

	/* эту проверку можно реализовать эффективнее
	if s == "" {
        return false // пустая строка → невалидна
    }
    for _, letter := range s {
        if !((letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z')) {
            return false
        }
    }
    return true
	*/
}

func main() {
	/*
	mux := http.NewServeMux()
	hello := http.HandlerFunc(HelloHandler)

	mux.Handle("/hello", SetDefaultName(Sanitize(hello)))
	//тут Порядок выполнения: Sanitize → SetDefaultName → helloHandler

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	} // --handler = LoggingMiddleware(mux)
	*/

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)

	// Добавляем middleware для записи запросов
	handler := SetDefaultName(Sanitize(mux))

	http.ListenAndServe(":8080", handler)
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