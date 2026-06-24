package main

import (
	"bufio"
	"fmt"
	//"io"
	"os"
	"strings"
)

//тут по заданию функция принмает на вход путь к файлу и возвращает его содержимое, то есть нам нужно прочитать его и буффер перевести в стринг и вернуть.
// т.к. мы не знаем заранее какой файл будет передаваться, то тут будем читать по байтам, т.е. в цикле (for) читаем по опред. колич. байт
func ReadContent(filename string) string {
	//первым делом на надо открыть файл (который передадут в випути стринг формата)
	file, err := os.Open(filename) //тут получили объект типа *os.File который можно читать
	if err != nil {
		return ""
	}
	//и закрываем файл используя defer, и если мы записываем, то исп. вот такой defer:
	// defer func() {
  	// 	errClose := f.Close() 
		// if errClose != nil { //т.к только тут можно обработать ошибку

		// }
	// }()
	defer file.Close()

	//теперь читаем, методом "считывать файл построчно"
	fileScanner := bufio.NewScanner(file) // создаем сканер

	var builder strings.Builder // создаем билдер - в него будем собирать строки со сканера
	for fileScanner.Scan() {
		builder.WriteString(fileScanner.Text())
		builder.WriteString("\n")
	}

	//обрабатываем ошибку сканера
	if err := fileScanner.Err(); err != nil {
		return ""
	}

	//возвращаем итог (билдер)
	return builder.String()
}

func main() {
	//тут составляем пусть к файлу, который будем передавать
	testFile := "../io_ReaderWriter/testRead.txt" // это относительно папки запуска программы, т.е. тут запуск - WorkingWithFiles\ReadContent.go и нам надо ".." - поднятся на уровень выше в папку YAGO2, потом зайти в папку io_ReaderWriter и там файл лежит testRead.txt


	//testFile1 := "E:/freeborm/projectsYA/YAGO2/io_ReaderWriter/testRead.txt" // или абсолютный путь, работает независимо от того, откуда запущена программа

	fmt.Println(ReadContent(testFile))
}