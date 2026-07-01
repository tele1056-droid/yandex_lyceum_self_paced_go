package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	Data string // тело ответа. Приводим к стринг string(body)
	StatusCode int // код ответа. Статус код берётся из resp.StatusCode
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	//создаём контекст с таймаутом из аргумента timeout и передаем его в метод WithContext запроса. Если запрос не будет выполнен за это время, контекст отменится и возвратится ошибка context.DeadlineExceeded.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	//делаем запрос и читаем ответ
	//cоздаём новый HTTP-запрос с помощью функции NewRequest()
	req, err := http.NewRequest(http.MethodGet, url, nil) // первый аргумент можно задавать и так http.MethodGet, и так "GET"
	if err != nil {
		return nil, err
	}

	//создаём новый клиент и отправляем запрос с помощью функции Do()
	client := &http.Client{}
	//отправляем запрос с контекстом, и именно тут получим ошибку context.DeadlineExceeded если запрос не будет выполнен за время таймаута
	resp, errDeadEx := client.Do(req.WithContext(ctx))
	if errDeadEx != nil {
		return nil, errDeadEx
	}

	//когда отправляем HTTP-запрос, сервер возвращает ответ в виде потока байтов. Чтобы прочесть его, нам нужен resp.Body — объект типа io.ReadCloser.

	defer resp.Body.Close() //и закрываем resp.Body, ра

	//читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//возвращаем указатель на структуру
	response := APIResponse{
		Data: string(body),
		StatusCode: resp.StatusCode}
	
	//и возвращаем либо так
	return &response, nil
	
	/*
	либо так
	return &Response{
        StatusCode: resp.StatusCode,
        Data:       string(body),
    }, nil
	*/

}

/*
Читать поток байтов напрямую может быть неудобно, поэтому рекомендуем функцию io.ReadAll. Она читает все данные из потока и возвращает их в виде байтового массива.
*/

/*
Важно
- resp.Body нужно закрывать (defer resp.Body.Close())
*/
