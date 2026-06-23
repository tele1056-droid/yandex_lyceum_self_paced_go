package main

import (
	"bytes"
	"fmt"
	"io"
)

//тут по заданию нужно найти в данных первое вхождение байт seq., т.е. первого вхождения (это найти первый раз, когда встречается последовательность) последовательности (seq - Sequence)
func Contains(r io.Reader, seq []byte) (bool, error) {
	//делаем чтение в буффер
	data := make([]byte, 1024)
	_, err := r.Read(data)
	if err != nil {
		return false, err
	}

	//ищем первое совпадение, исп. метод .Index который возвращает индекс первого экземпляра или -1 если совпадений не найдено
	pos := bytes.Index(data, seq)

	//проверяем найдена ли последовательность
	if pos >= 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func main() {
	testInput := bytes.NewReader([]byte("Hello, world!"))
	seq := []byte("Hel")

	fmt.Println(Contains(testInput, seq))
}