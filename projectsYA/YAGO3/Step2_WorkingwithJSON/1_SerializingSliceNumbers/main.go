package main

import (
	"encoding/json"
	"fmt"
)
//тут нужно слайс целых чисел сериализировать в формат JSON
//*сериализация - это кодирование и запись значений лбюбых типов данных в JSON. Отвечает за это - JSON Encoder
//или *Сериализация в JSON - это преобразование данных из формата Go в JSON. *Десериализация из JSON в Go работает в обратную сторону.



func SerializeIntSlice(nums []int) ([]byte, error) {
	//тут преобразуем данные Go(слайс интов) в данные JSON
	jsonBytes, err := json.Marshal(nums)
	if err != nil {
		panic(err) //тут вызываем панику без ретерна
	}
	return jsonBytes, nil
}

func main() {
	testNums := []int{1, 2, 3, 4}

	result, _ := SerializeIntSlice(testNums)
	fmt.Println(string(result))
}