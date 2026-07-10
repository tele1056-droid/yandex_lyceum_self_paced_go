package main

import (
	"fmt"
	//"log"
	"os"
)

func WriteToLogFile(message string, fileName string) error {
	//открываем файл
	file, err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close() // Закрываем файл после выхода из main

	// Конфигурируем логгер, чтобы он выводил лог в файл
	// log.SetOutput(file)
	// log.Println(message)

	//записываем данные в файл
	_, err = file.WriteString(message)
	if err != nil {
		return err
	}

	return nil
}

/*
Использовать абсолютный путь: (Работает только на твоём компьютере)
path := `E:\freeborm\projectsYA\YAGO3\ab.txt`, тут если слеш влево направлен то нужно исп. " `` " - обратные кавычки

Относительный путь от корня проекта: (на любом компьютере)
path := "../../YAGO3/ab.txt" - т.е. на две папки вверх
*/


/*
1. log.txt": строка с именем файла, который вы хотите открыть или создать
2. os.O_CREATE|os.O_WRONLY|os.O_APPEND: комбинация флагов, которые определяют, как будет открыт файл:
 - os.O_CREATE: указывает, что файл должен быть создан, если его нет. Если файл существует, этот флаг ничего не делает
 - os.O_WRONLY: файл будет открыт только для записи (write-only). Вы не сможете читать из этого файла внутри программы
 - os.O_APPEND: данные будут добавляться в конец файла, а не перезаписывать его содержимое
3. 0644: восьмеричное число, которое даёт права доступа к файлу. 0644 означает, что файл будет доступен для чтения и записи владельцу файла, а остальным — только для чтения
*/

func main() {
	message := "My first logOOOOOO1"
	filename := `E:\freeborm\projectsYA\YAGO3\ab.txt`

	err := WriteToLogFile(message, filename)
	if err != nil {
		fmt.Println(err)
	}

	
}