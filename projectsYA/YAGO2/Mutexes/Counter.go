package main

import (
	"fmt"
	"sync"
)

//написать потокобезопасный счётчик
type Counter struct {
    value int
    mu sync.RWMutex
}

type Count interface{
    Increment() // увеличение счётчика на единицу
    GetValue() int // получение текущего значения
}

func (c *Counter) Increment() {
	c.mu.Lock() //блокируем запись — НИКТО не может читать и писать
	c.value += 1
	c.mu.Unlock()
}

func (c *Counter) GetValue() int {
	//тут блокируем чтение только в момент записи другими горутинами, блокируем чтение — другие тоже могут читать
	c.mu.RLock()
	// теперь в эту секцию могут зайти несколько горутин
	data := c.value
	c.mu.RUnlock()

	return data
}

func main() {

	tets := Counter{value: 1}

	tets.Increment()

	fmt.Println(tets.GetValue())

	//fmt.Println(tets.value)
}


