package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func CopyFilePart(inputFilename string, outFileName string, startpos int) error {

	//открываем файл + закрываем
	file, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	
	defer file.Close()

					//СМЕЩЕНИЕ
	//читаем в буфер с определенной позиции, применяем .Seek
	// мы двигаем (курсор) в файле - смещаем позицию, как только применили f.Seek(offset, 0), то след. чтение / запись будет начинаться с этой позиции
	offset := int64(startpos)
	_, errSeek := file.Seek(offset, 0) // смещаемся от начала файла
	if errSeek != nil {
		return errSeek
	}
	
	/*					//SCANER
	// и делаем чтение попробуем через сканер
	var data []byte // cюда сразу будем аппендить строки в виде байт со сканера

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {	
		data = append(data, fileScanner.Text()...) //тут "..." - это оператор распаковки (разворачивания) слайса. И для string можно сразу применять "...", а .Text() возвращает строку
		data = append(data, '\n')
	}
	*/

					//BUILDER
	//тут применяем билдер, т.к. исходные данные файла имеют пробелы и перенос строк, и когда мы через сканер читаем файл будет на последней строке \n, а к билдеру можно применить .TrimSpace и убрать лишние пробелы и переносы строк
	var builder strings.Builder
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		builder.WriteString(fileScanner.Text())
		builder.WriteString("\n")
	}

	result := strings.TrimSpace(builder.String())
	data := []byte(result)


	// и делаем запись в файл через os.WriteFile()
	errWrite := os.WriteFile(outFileName, data, 0666)
	if errWrite != nil {
		return errWrite
	}
	return nil
}

func main() {
	//тут составляем путь к файлу, который будем передавать
	testFile := "../io_ReaderWriter/toRead.txt" // это относительно папки запуска программы, т.е. тут запуск - WorkingWithFiles\ReadContent.go и нам надо ".." - поднятся на уровень выше в папку YAGO2, потом зайти в папку io_ReaderWriter и там файл лежит testRead.txt
	fileWrite := "../io_ReaderWriter/toWrite.txt"

	//testFile1 := "E:/freeborm/projectsYA/YAGO2/io_ReaderWriter/testRead.txt" // или абсолютный путь, работает независимо от того, откуда запущена программа

	fmt.Println(CopyFilePart(testFile,fileWrite, 8))
}