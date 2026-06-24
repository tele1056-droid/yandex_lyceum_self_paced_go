package main

import (
	//"bufio"
	"fmt"
	//"io"
	"os"
	//"strings"
)

func ReadContent(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}

func main() {
	//тут составляем пусть к файлу, который будем передавать
	testFile := "../io_ReaderWriter/testRead.txt" // это относительно папки запуска программы, т.е. тут запуск - WorkingWithFiles\ReadContent.go и нам надо ".." - поднятся на уровень выше в папку YAGO2, потом зайти в папку io_ReaderWriter и там файл лежит testRead.txt


	//testFile1 := "E:/freeborm/projectsYA/YAGO2/io_ReaderWriter/testRead.txt" // или абсолютный путь, работает независимо от того, откуда запущена программа

	fmt.Println(ReadContent(testFile))
}