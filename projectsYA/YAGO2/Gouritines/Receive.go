package main

import (
	"fmt"
)

func Receive(ch chan int) int {
	val := <-ch //получаем значение из канала
	return val
}

func main() {
	ch := make(chan int) // создаем канал

	go func() { 
		ch <- 452 // запускаем горутину, которая отправляет значение в канал
	}()

	result := Receive(ch)

	fmt.Println(result)

}