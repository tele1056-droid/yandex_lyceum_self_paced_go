package main

import (
	"fmt"
	"strings"
)

func MakeCurlCommand(method, url, headers, body string) string {
	//тут нам нужно будет поработать с заголовками (headers), т.к. их может быть n - колич. и их нужно привести к такому виду -H '...'

	//и тут еще нужно учитвать пустые параметры, если они пустые (""), то нужно делать проверку и формировать итог строку без них (т.е. если нет body то не должно быть --data)

	//попробуем через вариант strings.Builder - это тип из стандартной библиотеки, который позволяет эффективно собирать строки из частей без лишних аллокаций памяти
	var sb strings.Builder // это буфер в ктором будем формировать итог. строку, последовательно командой за командой будет собираться строка
	//*strings.Builder — это как конвейер: ты добавляешь части одну за другой, а в конце получаешь готовую строку.
	sb.WriteString("curl")
	//Каждый WriteString добавляет в конец!!!
	//тут отрабатываем все методы кроме GET, если GET то ничего в итог строку не вставляем
	if method != "GET" {
		sb.WriteString(" -X ")
		sb.WriteString(method)
		sb.WriteString(" ")

	} else {
		sb.WriteString(" ")
	}

	fmtHeaders := formatHeaders(headers) //тут формируем заголовки и подсатвляем в буфер
	if headers != "" {
		
		sb.WriteString(fmtHeaders)
		sb.WriteString(" ")
	}

	if body != "" {
		sb.WriteString("--data '")
		sb.WriteString(body)
		sb.WriteString("' ")
		//тут обвернуть body в одинарные кавычки лучше так
	}

	
	
	sb.WriteString(url)

	
	strCurl := sb.String() //формируем итог. строку

	return strCurl
}

//для форм. заголовков сделаем отдельную функцию
func formatHeaders(headers string) string {
	//сплитим исходник с заголовками по разделителю "\n", получаем слайс заголовков
	lines := strings.Split(strings.TrimSpace(headers), "\n")
	//создаем пустой слайс стринг, в который будем аппендить уже форматированые заголовки (-H '...')
	parts := make([]string, 0, len(lines))
	for _, line := range lines {
		//делаем проверку на пустые заголовки
		if line == "" {
			continue
		}
		parts = append(parts, fmt.Sprintf("-H '%s'", line))
	}
	return strings.Join(parts, " ") //и соединяем через пробел
}

//			ОСНОВНЫЕ МЕТОДЫ ДЛЯ strings.Builder
/*
Метод					Что делает				Пример
WriteString(s string)	Добавляет строку		sb.WriteString("Hello")
WriteByte(b byte)		Добавляет байт			sb.WriteByte(' ')
WriteRune(r rune)		Добавляет символ		sb.WriteRune('😊')
Write(p []byte)			Добавляет байты			sb.Write([]byte("abc"))
String()				Возвращает собранную строку	result := sb.String()
Len()					Длина текущей строки	sb.Len()
Reset()					Очищает буфер			sb.Reset()
Grow(n int)				Зарезервировать память	sb.Grow(1024)
*/

func main() {
// 	res := MakeCurlCommand(
//   "POST",
//   "https://example.com/api/users",
//   "Content-Type: application/json\nAuthorization: Bearer abc123\n",
//   `{"name":"John Doe","email":"johndoe@example.com","password":"123456"}`,
// )

res := MakeCurlCommand(
  "POST",
  "https://example.com/api/users",
  "Content-Type: application/json\n",
  `{"name":"John Doe","email":"johndoe@example.com","password":"123456"}`,
)
	fmt.Println(res)
}