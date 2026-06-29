package main

import (
	"fmt"
	"sync"
)

//Напишите потокобезопасную очередь
type Queue interface {
    Enqueue(element interface{}) // положить элемент в очередь
    Dequeue() interface{} // забрать первый элемент из очереди
}

type ConcurrentQueue struct {
    queue []interface{}
    mutex sync.Mutex
}

//тут ставим в очередь
func (c *ConcurrentQueue) Enqueue(element interface{}) {
	c.mutex.Lock()
	c.queue = append(c.queue, element)
	c.mutex.Unlock()
}

//тут выводим из очереди
func (c *ConcurrentQueue) Dequeue() interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	//проверяем на пустую очередь
	if len(c.queue) == 0 {
		return nil
	}

	data := c.queue[0] // забираем первый элемент
	c.queue = c.queue[1:] //это способ удалить первый элемент слайса (сдвигаем слайс)
 	return data
}


func main() {
	
}