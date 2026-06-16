package main

import (
	"fmt"
)

type WeatherCondition int

const (
    Clear WeatherCondition = iota // Ключевое слово iota присваивает каждой константе числовое значение по порядку (0, 1, 2, 3 и т.д.) // 0
    Rain // 1
    HeavyRain // 2
    Snow // 3

	/*
	то есть тут у всех констан тип WeatherCondition, значение каждой константе присваивается автоматически, т.е. iota - это счетчик строк внутри блока const (начиная с 0 и увелич на 1 на каждой новой строке).
	А присваимваем значение константам чтобы, когда мы будем передавать состояние в функцию, например WeatherData{Condition: HeavyRain}, в GetWeatherMultiplier(weather WeatherData), то мы будем сравнивать weather == с константой, например, HeavyRain, то по сути мы передаем число 2, и в дальнейшем будем увеличивать коэф. на соответ. значение.

	По сути мы числам даём имена, 1 - это Rain, например.

	w := WeatherData{
    Condition: HeavyRain, // HeavyRain = 2
    WindSpeed: 15,
	}
	И момент когда мы полю Condition присваиваем значение и :
	1. w.Condition имеет тип WeatherCondition
	2. HeavyRain тоже имеет тип WeatherCondition
	3. HeavyRain = 2 (по iota)
	4. При присваивании w.Condition = HeavyRain → в w.Condition записывается число 2

	Визуализация:
	HeavyRain = 2
   		↓
	w.Condition = 2
   		↓
	if w.Condition == HeavyRain
	if 2 == 2 → true

	*/
)

type WeatherData struct {
    Condition WeatherCondition
    WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	multiplier := 1.0 // тут определяем базовый коэф.

	//и через switch case будем увелич базовый коэф. относительно состояния
	switch weather.Condition {
	case Clear:
		//тут не изменяем коэф. базовый
	case Rain:
		multiplier += 0.125
	case HeavyRain:
		multiplier += 0.2
	case Snow:
		multiplier += 0.15
	}

	// и для скорости ветра
	if weather.WindSpeed > 15 {
		multiplier += 0.1
	}

	return multiplier
}

func main() {
	test := WeatherData{Condition: Snow, WindSpeed: 15}

	fmt.Println(GetWeatherMultiplier(test))
}

