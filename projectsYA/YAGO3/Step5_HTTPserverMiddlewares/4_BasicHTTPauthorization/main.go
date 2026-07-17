package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	//"sync"
)

/*
curl -u admin:12345 localhost:8080 - автоматически кодирует логин и пароль в Base64 и отправляет в заголовке Authorization: Basic ...

Что делает -u:
curl -u admin:12345 localhost:8080
curl автоматически:
1. Собирает строку admin:12345
2. Кодирует в Base64: YWRtaW46MTIzNDU=
3. Добавляет заголовок: Authorization: Basic YWRtaW46MTIzNDU=

Сравнение:
Команда								Что отправляет
curl -u admin:12345 localhost:8080	Authorization: Basic YWRtaW46MTIzNDU=
curl -H "Authorization: Basic YWRtaW46MTIzNDU=" localhost:8080						То же самое
*/

/*
ВАЖНО:
- curl -u не шифрует пароль, а только кодирует в Base64
- Base64 — это не шифрование, его легко декодировать
- Всегда используй HTTPS, чтобы пароль не передавался в открытом виде
*/
//------------------------------------------------------

/*
ЧТО НУЖНО СДЕЛАТЬ:
1. Пишем middleware-функцию Authorization(http.HandlerFunc):
	- этот слой будет проверять запрос на наличие заголовка Authorization с логином и паролем в формате Basic Auth, т.е. проверять отсутствует ли заголовок, и логин или пароль пустые, или формат неверный
*/

//          == MIDDLEWARE ==

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*
		мы делаем валидацию запроса в middleware, и если валидация не проходит то мы «отклоняем запрос на уровне middleware». Т.е. решаем что делать с запросом: пропустить дальше или вернуть ответ прямо из middleware.

		Когда запрос приходит, middleware может:
		1. Пропустить -> next.ServeHTTP(w, r), т.е. отправить запрос дальше по цепочке до конечного обработчика (хендлера)
		2. Отклонить → отправить ответ и не вызывать next. Т.е. в проверке return (выходим, не вызывая next)
		*/

		// 1. Проверяем наличие заголовка
		//тут в случае невалид запроса (заголовок Authorization отсутствует, или логин или пароль пустые, или формат неверный), то возвращаем статус (401) и устанавливаем заголовок
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`) //устанавливает заголовок ответа
			w.WriteHeader(http.StatusUnauthorized) //устанавливает статус-код
			fmt.Fprintf(w, "Unauthorized\n")
			return
		}

		//2. Проверяем, что это Basic Auth.
		//тут проверяем формат
		if !strings.HasPrefix(authHeader, "Basic ") {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized\n")
			return 
		}

		// 3. Декодируем Base64
		encoded := authHeader[6:] // убираем "Basic "
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprint(w, "Unauthorized\n")
            return
		}

		// 4. Разделяем логин и пароль
		parts := strings.SplitN(string(decoded), ":", 2)
		if len(parts) != 2 {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprint(w, "Unauthorized\n")
            return
		}

		username := parts[0]
		password := parts[1]

		// 5. Проверяем, что логин и пароль не пустые
		if username == "" || password == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprint(w, "Unauthorized\n")
            return
		}

		// 6. Если валидация проходит, то пропускаем запрос дальше, по ситуации, либо в некст мидлваре или в конечный обработчик
		// + тут еще можно сохранить username в контексте, то есть передаем значения между обработчиками HTTP, можно использовать метод context.WithValue(). н добавляет значения в контекст запроса и получает их в других обработчиках запросов.

		ctx := context.WithValue(r.Context(), "username", username)
		//r = r.WithContext(ctx)

		next.ServeHTTP(w, r.WithContext(ctx))
		//про передачу в контексте значений: Go-2 -> Step5 -> Middleware и context

		/*
		СОВЕТ: Тип ключа
		// ❌ Неправильно (строки могут конфликтовать)
		ctx := context.WithValue(r.Context(), "username", username)

		// ✅ Правильно (используй свой тип)
		type contextKey string
		const usernameKey contextKey = "username"

		ctx := context.WithValue(r.Context(), usernameKey, username)
		*/

	})
}

//					== HANDLER == 
func answerHandler(w http.ResponseWriter, r *http.Request) {
	//в мидлваре мы в контексте передали значение username
	//тут получаем из контекста значени
	username := r.Context().Value("username")
	if username == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
	}

	fmt.Fprintf(w, "Welcome, %s!", username)

}

/* ЗАПРОС curl
Сравнение флагов
Флаг				Что показывает
curl (без флагов)	Только тело ответа
curl -i				Заголовки + тело
curl -v				Запрос + заголовки + тело
curl -I				Только заголовки
*/

func main() {
	mux := http.NewServeMux()

	handlerHello := Authorization(http.HandlerFunc(answerHandler))

	mux.Handle("/answer/", handlerHello)

	fmt.Println("Server started on :8080")
    http.ListenAndServe(":8080", mux)
}


/*
Части запроса:
Часть		Что это				Пример
Метод		Действие			GET, POST, PUT, DELETE
Путь		Ресурс				/hello, /users/123
Версия		HTTP-версия			HTTP/1.1, HTTP/2
Заголовки	Метаданные			Host, Content-Type, Authorization
Тело		Данные (POST, PUT)	{"name": "Alice"}

// Заголовки
    host := r.Header.Get("Host")
    userAgent := r.Header.Get("User-Agent")
*/

/*
	Заголовки (Header)
Заголовок		Что содержит
Host			Домен и порт
User-Agent		Клиент (браузер, curl)
Content-Type	Тип данных (JSON, HTML)
Authorization	Токен авторизации
Accept			Формат ответа, который ожидает клиен
*/

