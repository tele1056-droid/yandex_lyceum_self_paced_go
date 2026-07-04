package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func Compare(name1, name2 string) (string, error) {
	//нужно сформировать запрос по адресу localhost:8082/mark?name=<имя студента>
	// net/url (правильный способ) и используем его т.к. могут быть имена с пробелами, но при этом способе будет создаваться валид URL
	base := "http://localhost:8082/mark" // тут обяз. http://
	// nameFrst := name1
	// nameScnd := name2

	//создаем параметры
	params1 := url.Values{}
	params1.Set("name", name1)
	url1 := base + "?" + params1.Encode()
	
	params2 := url.Values{}
	params2.Set("name", name2)
	url2 := base + "?" + params2.Encode()

	//сделать запрос с контекстом, ну чтобы отменить запрос через время, если не отвечает сервер
	//создаем запрос с контекстом
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	//создаем клиент
	client := &http.Client{}

	//запрос req (request)
	req1, err := http.NewRequest(http.MethodGet, url1, nil)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	//отправляем запрос resp (response)
	resp1, err := client.Do(req1.WithContext(ctx))
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	//закрываем resp
	defer resp1.Body.Close()

	//тут делаем проверку на статус, http.StatusOK - это константа из пакета net/http там int 200
	if resp1.StatusCode != http.StatusOK {
		return "", errors.New(resp1.Status)
	}

	//читаем тело ответа
	body1, err := io.ReadAll(resp1.Body)
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return "", nil
	}
	

	//если статутс 200, то сохраним содержимое body1(а там слайс байт) в перемен. приведеную к int. Т.к. по HTTP всё передаётся как текс (поток байт), то нам нужно этот поток привести к стринг и потом к инт
	comp_body1, err := strconv.Atoi(string(body1))
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	//выше мы сохранили значение запроса для первого name1, осталось сделать все тоже самое для name2, и потом сделать саму работу
	req2, err := http.NewRequest(http.MethodGet, url2, nil)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	resp2, err := client.Do(req2.WithContext(ctx))
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != http.StatusOK {
		return "", errors.New(resp2.Status)
	}

	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return "", nil
	}

	comp_body2, err := strconv.Atoi(string(body2))
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	//делаем работу
	switch {
	case comp_body1 > comp_body2:
		return ">", nil
	case comp_body1 < comp_body2:
		return "<", nil
	default:
		return "=", nil
	}
}

func main() {

	
}

/*
Код ответа		resp.StatusCode
Тело ответа		io.ReadAll(resp.Body)
Заголовки		resp.Header.Get("...")
Статус текстом	resp.Status
*/

//тут конечно много дублирования, надо по сути запрос и обработку в отдельную функцию