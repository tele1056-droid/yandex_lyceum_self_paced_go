package main

import (
	"fmt"
	//"io"
	//"os"
	"strings"
)

type UpperWriter struct {
    UpperString string
}

func (uw *UpperWriter) Write(p []byte) (n int, err error) {
	//тут переводи слайс байт в string и переводим в верх. регистр
	upper := strings.ToUpper(string(p))

	// тут мы "записываем" в поле структуры, т.е. присваиваем значение полю (поля это переменные - присваимаем значение)
	uw.UpperString = upper

	//и возвращем значения
	return len(p), nil

}

func main() {

	//тут мы создаем w типа io.Writer (интерфейс) и хранит адрес структуры, и w используем только как Write, т.е. через w.UpperString - нельзя обратиться к полям структуры
	/*
	 var w io.Writer = &UpperWriter{}
	result, _ := w.Write([]byte("hey, hello"))
	*/

	//а тут у uw будет тип *UpperWriter (указатель на структуру), и можно вызывать любые методы и обращаться к полям
	uw := &UpperWriter{}
	uw.Write([]byte("hey, hello"))

	//fmt.Println(result)
	fmt.Println(uw.UpperString)
}
