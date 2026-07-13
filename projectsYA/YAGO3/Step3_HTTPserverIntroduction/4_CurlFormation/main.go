package main

import (
	"fmt"
	"strings"
)

func MakeCurlCommand(method, url, headers, body string) string {
	//тут нам нужно будет поработать с заголовками (headers), т.к. их может быть n - колич. и их нужно привести к такому виду -H '...'

	//и тут еще нужно учитвать пустые параметры, если они пустые (""), то нужно делать проверку и формировать итог строку без них (т.е. если нет body то не должно быть --data)

	//если через fmt.Sprintf, то нужно для каждого условия формировать отдельную строку, этот вариант запушу на гит, и сделаем через strings.Builder
	if method == "GET" {

	}

	fmtHeaders := formatHeaders(headers)
	strCurl := fmt.Sprintf("curl -X %s %s --data '%s' %s",method, fmtHeaders, body, url)

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
  "",
  "",
)
	fmt.Println(res)
}