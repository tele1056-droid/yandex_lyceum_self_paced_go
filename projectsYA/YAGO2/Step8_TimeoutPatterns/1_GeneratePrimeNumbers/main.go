package main

import (
  //"fmt"
  "time"
  //"sync"
)

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	//var wg sync.WaitGroup
	// создаем поток выполнениея горутины (основной работы - генрации простых чисел)
	go func() {
		for i := 2; i < N; i++ {
			if isPrime(i) {
				prime_nums <-i
			}
		}
		close(prime_nums) // именно тут закрываем канал, т.к. одна горутина пишет в канал, как все записалось, то закрываем
	}()
	
	/*
	Как выбрать как закрывать канал close(ch) или sync.WaitGroup
	1. Если одна горутина отправляет данные, то close(ch) внутри этой горутины;
	2, Если несколько горутин отправляют в один канал, то sync.WaitGroup + отдельная горутина для закрытия
	*отдельая горутина для закрытия:
	// Закрываем, когда все отправители завершены
		go func() {
        wg.Wait()
        close(ch) 
    }()
	*/

	//создаем тайм-аут, который не даёт функции выполнятся дольше заданого время
	timeout := time.AfterFunc(100 * time.Millisecond, func() {
		stop <- struct{}{} // тут передаем анонимную пустую структуру в канал для структур, как сигнал в канале
	})
	close(stop)

	timeout.Stop() // отмена функции тайм-аута
}

//тут создаем функцию для провреки на простое число
func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main() {

}