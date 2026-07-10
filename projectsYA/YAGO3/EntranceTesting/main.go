package main

import (
    "context"
    "sync"
	"fmt"
	"time"
)

func ParallelMapCtx(ctx context.Context, inputs []int, fn func(int) int, workers int) ([]int, error) {
    if workers <= 0 {
        workers = 1
    }

    results := make([]int, len(inputs))
    errCh := make(chan error, 1)
    var wg sync.WaitGroup

    // Канал задач: каждый элемент — это индекс и значение
    tasks := make(chan struct {
        idx int
        val int
    }, len(inputs))

    // Заполняем канал задачами
    for i, v := range inputs {
        tasks <- struct {
            idx int
            val int
        }{idx: i, val: v}
    }
    close(tasks)

    // Запускаем воркеров
    for w := 0; w < workers; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()

            for task := range tasks {
                // Проверяем отмену контекста
                select {
                case <-ctx.Done():
                    select {
                    case errCh <- ctx.Err():
                    default:
                    }
                    return
                default:
                }

                // Обрабатываем задачу
                results[task.idx] = fn(task.val)
            }
        }()
    }

    // Ждём завершения всех воркеров
    wg.Wait()

    // Проверяем, была ли ошибка
    select {
    case err := <-errCh:
        return nil, err
    default:
        return results, nil
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fn := func(x int) int {
        time.Sleep(100 * time.Millisecond) // имитация работы
        return x * 2
    }

    results, err := ParallelMapCtx(ctx, inputs, fn, 3)
    if err != nil {
        fmt.Println("Ошибка:", err) // context deadline exceeded
        return
    }

    fmt.Println(results) // [2, 4, 6, 8, 10, 12, 14, 16, 18, 20]
}

/*
Как работает:
Шаг	Что происходит
1	Создаётся канал tasks с задачами (индекс + значение)
2	Запускаются workers горутин
3	Каждая горутина берёт задачу из канала
4	Проверяется контекст (отмена/таймаут)
5	Выполняется fn(task.val)
6	Результат сохраняется по индексу task.idx
7	После завершения всех воркеров возвращаются результаты
*/

/*
ВАЖНЫЕ НЮАНСЫ:
Что								Почему
tasks с буфером len(inputs)		Чтобы не блокировать отправку
errCh с буфером 1				Чтобы не блокировать отправку ошибки
select с default				Неблокирующая проверка контекста
wg.Wait() перед чтением ошибки	Гарантирует, что все горутины завершены
*/