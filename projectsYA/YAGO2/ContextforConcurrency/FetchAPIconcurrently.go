package main

import (
	"io"
	"net/http"
	"sync"
	"time"
	"context"
)
var client = &http.Client{} // создаем клиента глобально для отправки запроса, т.к. не нужно создавать нового клиента для каждого запроса

type APIResponse struct {
	URL string  // запрошенный URL
	Data string // тело ответа
	StatusCode int // код ответа
	Err error // ошибка, если возникла

}



//тут мы отрабатываем Concurrently - т.е. несколько задач выполняются одновременно (параллельно или с переключением контекста)
// и это достигается через горутины. concurrently = go (горутины) + sync.WaitGroup - это стандартный способ делать несколько вещей одновременно. То есть задачи выполняются одновременно, а не по очереди

//В программировании: несколько HTTP-запросов, вычислений или операций ввода-вывода выполняются одновременно.

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	//создаем контекст с таймаутом, чтобы ограничить время запроса и отмены ожидания свыше timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	results := make([]*APIResponse, len(urls)) // созаем слайс APIResponse размером с колич. urls

	// запускаем цикл который будет запускать горутину с каждым url для создания запроса и получения данных, но тут нужно еще сделать concurrently для горутин - это wg := sync.WaitGroup{}, wg.Add(2) и defer wg.Done()
	wg := sync.WaitGroup{} //объявляем wg один раз, до цикла, чтобы wg.Wait() видел wg
	for i, url := range urls {
		
		wg.Add(1)
		go func(idx int, u string) {
			defer wg.Done()
			//тут нужно запускать запрос и обработку желательно отдельной функцие, но пока так
			req, err := http.NewRequest(http.MethodGet, u, nil)
			if err != nil {
				return //TODO: как тут ошибку вернуть в объект APIResponse
			}
			resp, errDeadEx := client.Do(req.WithContext(ctx))
			if errDeadEx != nil {
				return //TODO: errDeadEx
			}
			defer resp.Body.Close()

			//читаем тело ответа
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return //TODO: err
			}

			//собираем данные для структуры
			response := APIResponse{
				URL: u,
				Data: string(body),
				StatusCode: resp.StatusCode,
				Err: ,// как сюда положить ошибки?
			}

		}(url)
	}
	wg.Wait() // тут Wait() должен быть?
	return // и тут надо ретернить?
}



func main() {

}

/*
2. sync.WaitGroup
------------------
var wg sync.WaitGroup
Это счетчик горутин:
Метод		Что делает
wg.Add(1)	Увеличивает счётчик на 1
wg.Done()	Уменьшает счётчик на 1 (вызывается в горутине)
wg.Wait()	Блокирует выполнение, пока счётчик не станет 0

defer wg.Done() - Гарантирует, что Done() вызовется даже при ошибке.

Последовательно:
[запрос 1] ──► [запрос 2] ──► [запрос 3]  ← 6 секунд

Concurrently:
[запрос 1] ──┐
[запрос 2] ──┼──► ждём все ──► 3 секунды
[запрос 3] ──┘

3. Горутины:
----------------
for _, url := range urls {
    wg.Add(1)
    go func(u string) {
        defer wg.Done()
        fetch(u)
    }(url)
}
	Это вызов анонимной функции с передачей аргумента
	go - запустить как горутину
	func(u string) - объявляем функцию, принимающую string
	(url) - Передаём значение url в функцию (в параметр u)

	for _, url := range urls {
    // url = "https://httpbin.org/delay/1"
    go func(u string) {
        // внутри горутины u = "https://httpbin.org/delay/1"
        fetch(u)
    	}(url) // ← значение url передаётся в параметр u
	}
	
	ВОТ ТАК НЕПРАВИЛЬНО ИСПОЛЬЗОВАТЬ (ЗАМЫКАНИЕ)
	for _, url := range urls {
    go func() {
        fetch(url) // ❌ все горутины возьмут последний url
    }()
}
*/