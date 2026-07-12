package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseHTTPStatus(statusLine string) (code int, reason string) {
	//будем применять strings.Fields - разделитель по пробелам
	parts := strings.Fields(statusLine) // тут мы получим слайс из стринг (части, разделенные пробелом)
	codes, err := strconv.Atoi(parts[1]) //дополнительно нужно преобразовать в число из стринги
	if err != nil {
		return
	}

	//тут получаем текстовое пояснение (reason phrase) из статуса HTTP-ответа, со 2 индекса (включительно) и до конца слайса - достаем текст. пояснение
	reasons := parts[2:]
	//и нужно еще соединить элементы, исп. strings.Join
	resultReason := strings.Join(reasons, " ")

	return codes, resultReason
}

func main() {
	code, reas := ParseHTTPStatus("HTTP/1.1 418 I'm a teapot")
	fmt.Println(code)
	fmt.Println(reas)
}