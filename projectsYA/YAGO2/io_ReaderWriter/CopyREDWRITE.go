package main

import (
	"fmt"
	"os"
	"io"
	"bytes"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	//переводим n uint в int для удобства испльзования
	num := int(n)
	//первым делом работаем с io.Reader, читаем данные в буфер, обрабатываем ошибки, 

	buffer := make([]byte, 1024)
	total := 0

	//тут мы в цикле каждую итерацию читаем в буфер по 1024 байт
	for {
		bytesRead, err := r.Read(buffer)
		if err != nil {
			return err
		}

		if bytesRead == 0 { // если bytesRead ровно нулю, то значит данных для чтения больше нет, и выходим из цикла, это случай когда прочитаных байт меньше чем сколько всего нужно
			return fmt.Errorf("не удалось прочитать данные")
			//break
		} 
		
		//
		if total + bytesRead > num {
			// тут мы смотрим, если total - сколько байт прочит. всего, и bytesRead - сколько байт прочитано на итерации,  больше n - то есть сколько всего нужно скопировать, то записываем (исп. .Write) недостоющую часть и выходим из цикла
			need := num - total
			_, errWrite2 := w.Write(buffer[:need])
			//обрабатываем ошибку
			if errWrite2 != nil {
				return errWrite2
			}
			break
		}

		//тут, записываем весь прочитанный кусок, т.е. это случай если total+bytesRead < n
		_, errWrite := w.Write(buffer[:bytesRead])
		//обрабатываем ошибку w.Write
		if errWrite != nil {
			return errWrite
		}
		total += num

		//и проверяем сколько байт прочитали и сколько всего нужно
		if total == num {
			break
		}
	}

	
	return nil
}

/*
Т.е. в этой задаче нужно написать функцию, в которой мы читаем объект через метод Read, метод возвращает колич. прочитаных байт, параметр n uint - это сколько просят скопировать. Если в источнике больше или ровно n байт, то копируешь ровно n (тоесть записываю .Write), если в источнике меньше n байт, то копируем всё что есть.
Визуализация:
Источник: [A B C D E F G H] (8 байт)

Случай 1: n = 5
→ копируем [A B C D E] (5 байт) ✅

Случай 2: n = 10
→ копируем [A B C D E F G H] (8 байт) ✅
→ данных меньше, чем просили, копируем всё
------------------------------------------------------------

«Количество байт, доступных для чтения» — это сколько данных ещё осталось в источнике. По сути это размер источника (файл, строка, слайс байт, сеть, ), и не у всех источников есть размер. Файл и строка - есть. Сеть и io.Reader - нет, нужно читать до конца.

Для файла:
file, _ := os.Open("test.txt")

Для строки:
r := strings.NewReader("hello world")

Для слайса байт:
data := []byte("hello")
r := bytes.NewReader(data)

Для сети (net.Conn) - нельзя:

Для io.Reader (общий случай) — нельзя:
*/

func main() {
	/*
	//тут мы пишем конструкцию для чтения файла - это один объект
	fileRead, err := os.Open("toRead.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return // этим ретерном завершаем программу
	}

	defer fileRead.Close() // именно тут закрываем файл, сразу после открытия и проверки ошибки, если в конце то есть риск не закрыть файл

	//тут пишем конструкцию для записи файла - это второй объект(в данном случае я читал и записывал .txt файлы)
	fileWrite, err := os.Create("toWrite.txt")
	if err != nil {
		fmt.Println(err)
	}

	//вызываем функцию Copy
	if errFunc := Copy(fileRead, fileWrite, 100); errFunc != nil {
		fmt.Println(errFunc)
	}
		*/
	
	
	
						//test "Copy all bytes"
	testInput := bytes.NewReader([]byte("Hello, World!")) // тут мы создаем строку как объект для реалзиации io.Reader
	//т.е.: берет строку (string) -> переводим её в слайс байтов -> применяем bytes.NewReader для создания объекта

	fileWrite, err := os.Create("toWrite.txt")
    if err != nil {
        fmt.Println("Ошибка создания файла:", err)
        return
    }
    defer fileWrite.Close()

	if errFunc := Copy(testInput, fileWrite, 5); errFunc != nil {
        fmt.Println("Ошибка копирования:", errFunc)
	}
		



}