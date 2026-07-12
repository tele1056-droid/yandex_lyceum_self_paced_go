package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"io"
)
type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}

func DecodeStudentFromReader(r io.Reader) (Student, error) {
	// Создаём Decoder для чтения JSON из буфера
	decoder := json.NewDecoder(r)

	// Создаём переменную для хранения декодированных данных
	var student Student

	// Читаем JSON из буфера и записываем в переменную person
	err := decoder.Decode(&student)
	if err != nil {
		return Student{}, err //тут при ошибке возвращаем пустую структуру
	}

	return student, nil
}

func main() {
	// КАК СОЗДАТЬ ОБЪЕКТЫ io.Reader
	//1. Из строки
	jsonStr := `{"name": "Alice", "grade": 25}`
    r := strings.NewReader(jsonStr) //  io.Reader из строки

	student, err := DecodeStudentFromReader(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(student)

	//2. Из файла
	file, _ := os.Open("student.json")
    defer file.Close()
    r2 := file

	//3. Из байтов
	data := []byte(`{"name": "Alice", "age": 25}`)
    r3 := bytes.NewReader(data)
}