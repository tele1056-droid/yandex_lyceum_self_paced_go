package main

import (
	"fmt"
	"context"
	"sync"
)

//создами структуру чтобы создать безопасный слайс, + устанавливаем приватные (неэкспортируемые) поля, чтобы не давать пользователю доступ к полям структуры
// и создаем эту структуру чтобы в дальнейшем в коде использовать мьютексы (Lock и Unlock) для ограничения доступа к слайсу только одной горутины в один момент времени
type SafeSlice struct {
  results []int
  mx      *sync.Mutex
}

func New() *SafeSlice {
	return  &SafeSlice{
		mx: &sync.Mutex{},
		results: []int{},
	}
}

//добавляем эелемент к салйсу
func (s *SafeSlice) Append(item int) {
	// вызван Lock, поэтому только одна горутина за раз может получить доступ к слайсу
	s.mx.Lock()
	defer s.mx.Unlock()

	s.results = append(s.results, item)

}

func ParallelMapCtx(ctx context.Context, inputs []int, fn func(int) int, workers int) ([]int, error) {
	safeSlice := New()
	// создаём экземпляр WaitGroup
	wg := &sync.WaitGroup{}

	//работа горутин
	wg.Add(workers)
	for elem := range inputs {
		go func() {
			defer wg.Done()
			item := fn(elem)
			safeSlice.Append(item)
		}()
	}

	//ждём выполнения всех горутин группы
  	wg.Wait()

}