package main

import (
	//"bufio"
	"fmt"
	//"strings"
	"os"
)

func ModifyFile(filename string, pos int, val string) {
	//нам нужно получить дескриптор файла и при этои дописать данные которые уже есть, для этого исп. os.OpenFile()
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err) // обрабатываем ошибку
	}

	//не забываем закрывать файл
	defer file.Close()

	//НО, если исп. O_APPEND - то всегда будет запись в конец файла, этот режим отключает возможность смещения .Seek. Нужно открывать без O_APPEND
	// тут указываем именно два флага и O_APPEND и O_WRONLY, т.к.:
	/*
	O_WRONLY — говорит: «я буду только записывать» (не читать)
	O_APPEND — говорит: «записывай в конец файла»
	Без O_WRONLY файл не откроется для записи.
	*/
	// теперь у нас в file содержиться дескриптор и при записи через file исходные данные останутся в файле
	

	//теперь делаем смещение .Seek
	offset := int64(pos) // переводим int в int64
	file.Seek(offset, 0)

	//и записываем данные в файл с позиции смещения
	file.WriteString(val)
}

func main() {
	//тут составляем путь к файлу, который будем передавать
	//testFile := "../io_ReaderWriter/toRead.txt" // это относительно папки запуска программы, т.е. тут запуск - WorkingWithFiles\ReadContent.go и нам надо ".." - поднятся на уровень выше в папку YAGO2, потом зайти в папку io_ReaderWriter и там файл лежит testRead.txt
	fileWrite := "../io_ReaderWriter/toWrite.txt"

	//testFile1 := "E:/freeborm/projectsYA/YAGO2/io_ReaderWriter/testRead.txt" // или абсолютный путь, работает независимо от того, откуда запущена программа

	ModifyFile(fileWrite, 58, "POE1")
}