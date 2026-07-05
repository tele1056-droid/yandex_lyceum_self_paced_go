package main

import (
	"context"
	//"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"sync"
)

type result struct {
    name  string
    score int
    err   error
}

func Compare(name1, name2 string) (string, error) {
    results := make(chan result, 2)
    var wg sync.WaitGroup

    // Запускаем запросы параллельно
    for _, name := range []string{name1, name2} {
        wg.Add(1)
        go func(n string) {
            defer wg.Done()
            score, err := fetchScore(n)
            results <- result{name: n, score: score, err: err}
        }(name)
    }

    // Закрываем канал после завершения всех горутин
    go func() {
        wg.Wait()
        close(results)
    }()

    // Собираем результаты
    scores := make(map[string]int)
    for res := range results {
        if res.err != nil {
            return "", fmt.Errorf("ошибка для %s: %w", res.name, res.err)
        }
        scores[res.name] = res.score
    }

    // Сравниваем
    score1 := scores[name1]
    score2 := scores[name2]

    switch {
    case score1 > score2:
        return ">", nil
    case score1 < score2:
        return "<", nil
    default:
        return "=", nil
    }
}

// fetchScore — функция для получения оценки
func fetchScore(name string) (int, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    base := "http://localhost:8082/mark"
    params := url.Values{}
    params.Set("name", name)
    fullURL := base + "?" + params.Encode()

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
    if err != nil {
        return 0, err
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return 0, fmt.Errorf("статус %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return 0, err
    }

    return strconv.Atoi(string(body))
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

//				ПАТТЕРН ДЛЯ ПОДОБНЫХ ЗАДАЧ
/*
func Compare(...) {
    results := make(chan result, 2)
    var wg sync.WaitGroup

    // 1. Запускаем горутины
    for _, name := range []string{name1, name2} {
        wg.Add(1)
        go func(n string) {
            defer wg.Done()
            // ... делаем запрос ...
            results <- result{name: n, score: score, err: err}
        }(name)
    }

    // 2. Запускаем горутину-«закрывалку»
    go func() {
        wg.Wait()
        close(results)
    }()

    // 3. Читаем результаты
    for res := range results {
        // обрабатываем
    }
}
*/