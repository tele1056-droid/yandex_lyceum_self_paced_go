package main

import (
	"io"
	"net/http"
	"sync"
	"time"
	"context"
	"fmt"
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
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	results := make([]*APIResponse, len(urls)) // создаем слайс APIResponse размером с колич. urls

	// запускаем цикл который будет запускать горутину с каждым url для создания запроса и получения данных, но тут нужно еще сделать concurrently для горутин - это wg := sync.WaitGroup{}, wg.Add(2) и defer wg.Done()
	var wg sync.WaitGroup //объявляем wg один раз, до цикла, чтобы wg.Wait() видел wg
	for i, url := range urls {
		wg.Add(1)
		go func(idx int, u string) {
			defer wg.Done()
			//тут нужно запускать запрос и обработку желательно отдельной функцие, но пока так
			resp := &APIResponse{URL: u} //тут создаем "заготовку" структуры с заполн. полем URL и в resp сохраняем адрес этой заготвки и дальше по этому адресу заполним остальные поля
			req, err := http.NewRequest(http.MethodGet, u, nil)
			if err != nil {
				 //TODO: как тут ошибку вернуть в объект APIResponse
				 //тут в ошибках, в заготовку структуры в поле Err записываем ошибку и потом эту заготовку (с заполн. полями URL и Err) записываем в слайс заготовок
				resp.Err = err
				results[idx] = resp
				return
			}

			//отправляем запрос с контекстом
			res, err := client.Do(req.WithContext(ctx))
			if err != nil {
				resp.Err = err
				results[idx] = resp
				return 
			}
			defer res.Body.Close()

			//читаем тело ответа
			body, err := io.ReadAll(res.Body)
			if err != nil {
				 //TODO: err
				resp.Err = err 
				results[idx] = resp
				return
			}

			//собираем данные для структуры
			resp.Data = string(body)
			resp.StatusCode = res.StatusCode
			results[idx] = resp

		}(i ,url)
	}
	wg.Wait() // тут Wait() должен быть?
	return results // и тут надо ретернить?
}


func main() {
    // 1. Создаём контекст (можно Background, но лучше с таймаутом)
    ctx := context.Background()

    // 2. Список URL для запросов
    urls := []string{
        "https://httpbin.org/get",
        "https://httpbin.org/get?test=1",
        "https://httpbin.org/status/404",
        "https://httpbin.org/delay/1",
    }

    // 3. Вызываем функцию с таймаутом 5 секунд
    results := FetchAPI(ctx, urls, 5*time.Second)

    // 4. Выводим результаты
    for i, res := range results {
        fmt.Printf("--- Запрос %d ---\n", i+1)
        fmt.Printf("URL: %s\n", res.URL)
        
        if res.Err != nil {
            fmt.Printf("Ошибка: %v\n", res.Err)
            continue
        }
        
        fmt.Printf("Статус: %d\n", res.StatusCode)
        fmt.Printf("Длина данных: %d байт\n", len(res.Data))
        fmt.Println()
    }
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