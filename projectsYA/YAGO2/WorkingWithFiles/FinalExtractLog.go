package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)


func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	//открываем файл / закрываем файл
	file, err := os.Open(inputFileName)
	if err != nil {
		return nil, err //возвращаем пустой слайс и ошибку
	}

	defer file.Close()

	//начинаем читать через сканер и обрабатывать строки
	var lines []string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		//убираем пробелы и пустые строки в начале и в конце строки
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		//сплитим строки по разделителю " "
		parts := strings.Split(line, " ")

		//сделаем проверку на колич. элементов в parts, если вот такой формат логов 12.12.2022 info, то должно быть не меньше двух элементов
		if len(parts) < 2 {
			continue
		}

		//ПРОВЕРКА ДАТЫ обрабатываем дату вот с таким layout = "02.01.2006"
		timeStamp, errTime := time.Parse("02.01.2006", parts[0]) //тут time.Parse возвращает дату, либо ошибку, если формат не совпадает или дата невозможная, если ошибку то пропускаем эту строку
		if errTime != nil {
			continue
		}

		//ПРОВЕРЯЕМ ДИАПАЗОН входит ли дата (timeStamp) в диапазон который передали в функцию
		//тут будем использовать !timeStamp.Before(start) && !timeStamp.After(end), т.е. проверяет что timeStamp находится в диапазоне [start, end] (включая границы)
		/*
		t.Before(start)		t раньше, чем start
		!t.Before(start)	t НЕ раньше start → t >= start
		t.After(end)		t позже, чем end
		!t.After(end)		t НЕ позже end → t <= end
		&&					Оба условия должны быть true
		*/
		/*
		НАПРИМЕР:
		start = 10.12.2022
		end   = 15.12.2022

		t = 12.12.2022
		├── !t.Before(start) → true (12 >= 10)
		├── !t.After(end)    → true (12 <= 15)
		└── true && true     → true ✅ в диапазоне

		t = 08.12.2022
		├── !t.Before(start) → false (8 < 10)
		└── false && ...     → false ❌ вне диапазона

		t = 20.12.2022
		├── !t.After(end)    → false (20 > 15)
		└── ... && false     → false ❌ вне диапазона
		*/

		if !timeStamp.Before(start) && !timeStamp.After(end) {
			lines = append(lines, line)
		}
	}

	//проверяем наш собранный lines
	if lines == nil {
		return nil, errors.New("Log was not found")
	}

	return lines, nil
}

func main() {
	testFile := "../io_ReaderWriter/toRead.txt"

	layout := "02.01.2006"
	start, _ := time.Parse(layout, "13.12.2022")
	end, _ := time.Parse(layout, "15.12.2022")

	result, err := ExtractLog(testFile, start, end)
	fmt.Println(result)
	fmt.Println(err)
}