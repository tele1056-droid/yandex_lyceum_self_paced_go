package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	//"strings"
)

/*
Структура json.Encoder
JSON Encoder в Go — это структура, которая позволяет кодировать данные в JSON и записывать их в writer — например, в файл или буфер.
*/

type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}

func EncodeStudentsToWriter(w io.Writer, students []Student) error {
	// Создаём буфер для записи JSON-данных
	//var buf bytes.Buffer Тут не нужно создавать буфер (объект io.Writer), т.к. уже передают в функцию

	// Создаём Encoder для записи JSON-данных в буфер
	encoder := json.NewEncoder(w)

	// Записываем JSON-данные в буфер с помощью метода Encode() Encoder
	err := encoder.Encode(students)
	if err != nil {
		return err
	}
	return nil
}

//				ВАРИАНТЫ СОЗДАНИЯ ОБЪЕКТА io.Writer
/*
Способ 1: bytes.Buffer (в память):
import "bytes"
var buf bytes.Buffer // io.Writer
--------------------------------------
Способ 2: Файл:
import "os"
file, err := os.Create("output.txt")
    if err != nil {
        return
    }
    defer file.Close()
	//// file — это io.Writer
---------------------------------------
Способ 3: os.Stdout (консоль):
import "os"
// os.Stdout — это io.Writer
WriteToWriter(os.Stdout) // выведет в консоль - это как пример
--------------------------------------------------
Способ 4: strings.Builder (для строк):
import "strings"
var builder strings.Builder // io.Writer
*/

func main() {
	var buf bytes.Buffer 
	students := []Student{
        {Name: "Alice", Grade: 7},
        {Name: "Bob", Grade: 8},
        {Name: "Charlie", Grade: 9},
    }

	EncodeStudentsToWriter(&buf, students) //тут важно передавать &buf - указатель
	fmt.Println(buf.String()) // и вот так слайс байт приводить к стринге

}
