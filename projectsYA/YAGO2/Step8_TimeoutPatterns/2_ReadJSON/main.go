package main

import (
	"context"
	//"fmt"
	"os"
	//"time"
)

//тут параметр (result chan<- []byte) - это канал в который можно только отправлять(записывать) данные типа слайс байт []byte

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	//тут контекст передает вызывающий код (у себя уже создал контекст с таймаутом) , мы его просто используем

	//тут проверяем контекст сразу
    select {
    case <-ctx.Done():
        close(result)
        return // завершаем сразу
    default:
    }

	//читаем из json файла, будем исп. os.ReadFile (открывает и читает и закрывает файл)
	 // создаем канал в который будем записывать ошибки из горутины
	go func() {
		data, err := os.ReadFile(path)
		if err != nil {
			close(result)
			return
		}

		 // Проверяем контекст ПОСЛЕ чтения файла
        select {
        case <-ctx.Done():
            close(result)
            return
        default:
        }
        
		//проверяем контекст перед отправкой
		select {
		case result <- data: //пишем в канал
		case <-ctx.Done():
            close(result)
			return
		}
	}()

	// Просто ждём завершения горутины
    <-ctx.Done()
	
}

 
/*

Запись		Что можно делать
chan int	Читать и писать
chan<- int	Только писать (отправлять)
<-chan int	Только читать (получать)

*/

/*
Где и когда ставить select:
Где select			Зачем
Внутри горутины		Чтобы реагировать на сигналы (отмена, тайм-аут) во время работы
После горутины		Чтобы ждать результат или тайм-аут

*Внутри горутины — для проверки. После горутины — для ожидания.

		ВНУТРИ ГОРУТИНЫ
go func() {
    data, err := os.ReadFile(path)
    if err != nil {
        errCh <- err
        return
    }
    select {
    case result <- data:
        // ✅ успешно отправлено
    case <-ctx.Done():
        // ❌ контекст отменён — не отправляем
        return
    }
}()

		ПОСЛЕ ГОРУТИНЫ
select {
case data := <-result:
    // ✅ результат получен
case err := <-errCh:
    // ❌ ошибка
case <-ctx.Done():
    // ❌ тайм-аут
}
*/