package main

import (
	"encoding/json"
	"fmt"
)

//			ЧТЕНИЕ JSON
/*
1. Структуру (обычно):
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

var user User
json.Unmarshal(data, &user)
---------------------------------------
2. Карту (map[string]interface{})
var data map[string]interface{}
json.Unmarshal(jsonData, &data)

name := data["name"].(string)
-----------------------------------------
3. Слайс ([]interface{})
var users []User
json.Unmarshal(jsonData, &users)

var data []interface{}
json.Unmarshal(jsonData, &data)
--------------------------------------------
4. any (любой тип)
var data any
json.Unmarshal(jsonData, &data)
// data может быть map, slice, string, number...
-----------------------------------------------
*/

/*
Действие				Название
JSON → Go-структура		Десериализация / Парсинг / Декодирование
Go-структура → JSON		Сериализация / Маршалинг / Кодирование
*/

//тут нужно принять строку (т.е. текст вида {"key": "value"}) JSON и парсить её в структуру
func DeserializeStringMap(data string) (map[string]string, error){
	//т.е. тут нужно применять json.Unmarshal (десериализация, из JSON в Go)

	//создаем мапу в котору будем парсить JSON
	var result map[string]string
	//парсим
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		//panic(err) //panic - это критическая ошибка, которая останавливает выполнение программы.
		return nil, err
	}
	return result, nil	
}

func main() {
	//jsonStr := `{"name": "John", "Gender": "male"}`
	//jsonStr1 := `{"a":"1","b":"2"}`
	//jsonStr2 := `{}`
	jsonStr3 := `{bad json}`

	fmt.Println(DeserializeStringMap(jsonStr3))
}


//			Разница между Unmarshal/Marshal и Decoder/Encoder
/*
Что					Когда использовать
json.Unmarshal		Уже есть []byte (строка)
json.Marshal		Нужно получить []byte (строку)
json.NewDecoder		Читаешь JSON из io.Reader (файл, сеть)
json.NewEncoder		Записываешь JSON в io.Writer (файл, сеть)
*/

/*
Операция	Вход		Выход
Unmarshal	[]byte		Структура
Marshal		Структура	[]byte
Decoder		io.Reader	Структура
Encoder		Структура	io.Writer
*/

//			ПРИМЕРЫ
/*
Unmarshal — из строки в структуру:
jsonStr := `{"name": "Alice", "age": 25}`
var user User
err := json.Unmarshal([]byte(jsonStr), &user)
-------------------------------------------------------
Marshal — из структуры в строку:
user := User{Name: "Alice", Age: 25}
data, err := json.Marshal(user) // data — []byte
jsonStr := string(data)
---------------------------------------------------------
Decoder — из файла/сети в структуру:
file, _ := os.Open("user.json")
defer file.Close()

var user User
err := json.NewDecoder(file).Decode(&user)
-----------------------------------------------------------
Encoder — из структуры в файл/сеть:
file, _ := os.Create("user.json")
defer file.Close()

user := User{Name: "Alice", Age: 25}
err := json.NewEncoder(file).Encode(user)
*/

//					ПРАВИЛО ВЫБОРА
/*
Ситуация								Использовать
Есть строка JSON (string или []byte)	Unmarshal / Marshal
Читаю из файла							Decoder
Пишу в файл								Encoder
Читаю из HTTP-ответа					Decoder
Пишу HTTP-ответ							Encoder
*/


