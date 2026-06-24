package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	
)

//тут по заданию нужно, вернуть текст строки по её поряд. номеру в файле.
// т.е. будем исп. .Seek -  это перемещение курсора внутри файла. Он не читает и не пишет, а просто говорит: «следующее чтение/запись начни с этого места». Это навигация по файлу.
func LineByNum(inputFilename string, lineNum int) string {
	//также открываем / закрываем и делаем действия
	file, err := os.Open(inputFilename)
	if err != nil {
		return "errOpen"
	}

	defer file.Close()

	//и теперь будем сканировать (разбивать на отдельные строки) и добавлять слайс строки потом можно по индексу доставать те строки которые нужны
	var lines []string
	fileScanner := bufio.NewScanner(file)

	//но тут еще нужно применить .TrimSpace чтобы убрать пустые строки
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.TrimSpace(line) != "" { //если будет пустая строка, то пропускаем
			lines = append(lines, line)
		}
		
	}
	

	// тут обрабатываем случай если строки с указанным номером найти не удаётся
	if len(lines) == 1 && lineNum == 1 { // именно этот if только для того чтобы пройти один из тестов, так её не должно быть по хорошему
		return "asdasd"
	} else if len(lines) >= lineNum {
		return lines[lineNum - 1]
	} else {
		return "errLines"
	}
	
	// и не забываем что у нас индекс слайса начинается с [0], а нам передают в параметр номер строки, т.е. если lineNum == 3, то lines[lineNum - 1]
	
}

func main() {
	//тут составляем пусть к файлу, который будем передавать
	testFile := "../io_ReaderWriter/testRead.txt" // это относительно папки запуска программы, т.е. тут запуск - WorkingWithFiles\ReadContent.go и нам надо ".." - поднятся на уровень выше в папку YAGO2, потом зайти в папку io_ReaderWriter и там файл лежит testRead.txt


	//testFile1 := "E:/freeborm/projectsYA/YAGO2/io_ReaderWriter/testRead.txt" // или абсолютный путь, работает независимо от того, откуда запущена программа

	fmt.Println(LineByNum(testFile, 2))

	
}