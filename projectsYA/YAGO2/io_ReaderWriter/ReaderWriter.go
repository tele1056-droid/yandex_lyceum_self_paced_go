package main

import (
	//"errors"
	"fmt"
	"io"
	"os"
)

func WriteString(s string, w io.Writer) error {
	strBytes := []byte(s)
	n, err := w.Write(strBytes) //тут в w я передавал созданный файл, т.е нужен такой объект который у которого есть метод Write
	if err != nil {
		return err // возвращаем реальную ошибку, а не свою
	}

	//тут проверяем колич. запсанных байт, чтобы убедиться что все данные записаны
	if n != len(strBytes) {
		return fmt.Errorf("записано %d из %d байт", n, len(strBytes))
	}
	return nil
}

func main() {

	//тут мы создаем "объект" (файл test.txt) для передачи в аргумент func WriteString, мы создаем файл, обрабатываем ошибки, потом через defer закрываем файл, и в конце для перемен. err присваиваем новое значение (это WriteString вернет ошибку) и так же обрабатываем 
	test := "Hello, !"
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err) 
	}
	defer file.Close()

	if err = WriteString(test,file); err != nil {
		fmt.Println(err)
	}
}