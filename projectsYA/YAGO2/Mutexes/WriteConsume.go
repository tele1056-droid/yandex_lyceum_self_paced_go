package main

import (
	"fmt"
	"sync"
)

// Используйте мьютекс для синхронизации доступа к буферу
var (
	Buf []int
	mutex sync.Mutex
)

//функция записывает данные в буфер Buf []int
func Write(num int) {
	mutex.Lock()
	Buf = append(Buf, num)
	mutex.Unlock()
}

//функция которая будет забирать первое значение из буфера и возвращать его
func Consume() int {
	mutex.Lock()
	defer mutex.Unlock()

	//проверяем на пустой буфер
	if len(Buf) == 0 {
		return 0
	}

	data := Buf[0] //забираем первое значение
	Buf = Buf[1:] //удаляем первый элемент

	return data
}

func main() {
	nums := []int{1, 2, 3, 65, 98}

	for _, element := range nums {
		Write(element)
	}

	for _, expected := range nums {
		actual := Consume()
		if actual != expected {
			fmt.Println("error")
		}
		fmt.Println(actual)
	}
}