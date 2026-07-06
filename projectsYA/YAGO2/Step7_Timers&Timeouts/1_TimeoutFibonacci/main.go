package main

import (
	"errors"
	"fmt"
	"time"
)

//тут нужно сгенерировать n-е число фибоначи при этом нужно установить ограничение на время вычесления
func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	//тут нужно перед работой проверить равен ли timeout нулю, ведь если timeout == 0 то горутину нет смысла запускать(т.к. она всё равно не успее)
	if timeout <= 0 {
		return 0, errors.New("timeout")
	}

	//также проверить n на отриц. значение
	if n < 0 {
		return 0, errors.New("n must be non-negative")
	}

	//итеративный способ (быстрый и простой)
	//создаем канал
	ch := make(chan int, 1)

	//создание потока выполнения горутины
	go func() {
		ch <- fib(n)
	}()

	//создание тайм-аута для выполнения функции
		select {
		case res := <-ch:
			return res, nil
		case <-time.After(timeout):
			return 0, errors.New("timeout")
		}

}

//создаем отельно функцию для расчета Фибоначи
func fib(n int) int {
    if n <= 1 {
        return n
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}

func main() {
	res, err := TimeoutFibonacci(1000000, 1*time.Second)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат (n=10):", res)
	}

	
}