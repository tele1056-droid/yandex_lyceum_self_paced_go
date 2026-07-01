package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
)

/*
Пакет context — это стандартный пакет Go для управления временем жизни и отменой операций
Что даёт пакет context
Отмена долгих операций (таймауты, ручная отмена)
Передача данных через цепочку вызовов (например, userID, traceID)
Синхронизация горутин (сигнал «стоп»)

select case - работает с каналами, Ожидание событий на каналах:
	1. Проверяет все case (каналы)
	2. Если один готов — выполняет его
	3. Если несколько готовы — выбирает случайный
	4. Если ни один не готов:
		- Без default → блокируется (ждёт)
		- С default → выполняет default сразу

	проверяй контекст там, где есть смысл остановиться. В начале функции — чтобы не начинать. В цикле — чтобы не продолжать. Перед вызовом БД/HTTP — чтобы не ждать.
*/

// функция которая должна найти первое вхождение байт seq в данных, доступных через Reader r.
func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	//тут сделаем начальную проверку контекста на отмену
	if ctx.Err() != nil {
		return false, ctx.Err()
	}
	//тут для чтения будем использовать io.ReadAll - читает все данные из потока и возвращает слайс байтов
	data, err := io.ReadAll(r)
	if err != nil {
		return false, err
	}

	//ищем первое вхождение seq
	pos := bytes.Index(data, seq)
	//проверяем найдена ли последовательность
	if pos >= 0 {
		return true, nil
	}

	//проверяем после работы
	 if ctx.Err() != nil {
        return false, ctx.Err()
    }
	
	//тут ретерн если последовательность не найдена
	return false, nil
}
/*
Если последовательность найдена, верните true, nil, иначе false, nil. В случае ошибки — false и саму ошибку. В случае отмены контекста — false и причину отмены.
*/

func main() {
	
}

/*
			ВАРИАНТЫ ПРОВЕРОК КОНТЕКСТА НА ОТМЕНУ:

	1. Проверка в начале + долгая работа: Если работа быстрая и не требует прерывания

	func process(ctx context.Context) error {
    // Проверяем в начале
    if ctx.Err() != nil {
        return ctx.Err()
    }

    // Делаем работу
    fmt.Println("Начинаем работу...")
    time.Sleep(3 * time.Second)

    // Проверяем после работы (опционально)
    if ctx.Err() != nil {
        return ctx.Err()
    }

    fmt.Println("Работа завершена")
    return nil
	}

	2. Цикл с работой внутри: Если работа состоит из повторяющихся шагов

	func process(ctx context.Context) error {
    for i := 0; i < 10; i++ {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            // делаем часть работы
            fmt.Println("Шаг", i)
            time.Sleep(500 * time.Millisecond)
        }
    }
    return nil
}

	3. select для ожидания (не бесконечный цикл): Если работа — просто ожидание

	func process(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    case <-time.After(5 * time.Second):
        // работа завершена
        fmt.Println("Работа завершена")
        return nil
    }
}

*/