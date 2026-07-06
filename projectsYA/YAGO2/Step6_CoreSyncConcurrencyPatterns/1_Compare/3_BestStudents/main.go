package main

import (
	"context"
	//"math"
	//"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
    "sort"
    "strings"
)

type result struct {
    name  string
    score int
    err   error
}

func  BestStudents(names []string) (string, error) {
	results := make(chan result, len(names))
    var wg sync.WaitGroup

    // Запускаем запросы параллельно
    for _, name := range names {
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

    // Собираем результаты, мапа с именами и их оценками
    scores := make(map[string]int)
    for res := range results {
        if res.err != nil {
            return "", fmt.Errorf("ошибка для %s: %w", res.name, res.err)
        }
        scores[res.name] = res.score
    }

    // Делаем работу, ищем среднюю успеваемость
    //т.к. мапу нельзя напрямую сортировать, то будет доставать в отдельный слайс значения (оценки) из мапы и этот слайс сортируем
	value := make([]int, 0, len(scores)) //создаем слайс
	for _, v := range scores {
		value = append(value, v)
	}
	// и ищем среднее
	sum := 0
	for _, v := range value {
		sum += v // находим сумму слайса value
	}
	avgStud := float64(sum) / float64(len(value)) // тут считаем среднее, главное привести слагаемые к float64

    //формируем список имен выше среднего
    namesAvg := make([]string, 0, len(scores))
    for key, value := range scores {
        if float64(value) > avgStud {
            namesAvg = append(namesAvg, key)
        } 
    }

    //сортируем по алф.
    sort.Strings(namesAvg)

    itogReturn := strings.Join(namesAvg, ", ")

	return itogReturn, nil
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