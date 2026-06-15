package main

import (
	"fmt"
	"slices"
)

//interface в этом задании я думаю не надо создавать т.к. в программе используется одна структура и

// наши объекты (игроки) в виде структур, это шаблон
type Player struct {
    Name      string
    Goals     int
    Misses    int
    Assists   int
    Rating    float64
}

//тут будем хранить данные о каждом игроке, создаем стурктуру в которой будем хранить игроков (т.е. слайс структур)
type PlayerStorage struct {
	Players []Player
}

//метод для добавления игроков в хранилище, и тут метод принимает ЗНАЧЕНИЕ
func (ps *PlayerStorage) AddPlayer(p Player) {
    ps.Players = append(ps.Players, p)
}

//метод для расчета рейтинга, мы оформляем его отдельно, и будем вызывать в конструкторе для расчета поля Rating    float64
func (p *Player) calculateRating() float64 {
    rating := 0.0
    if p.Misses == 0 {
        rating = float64(p.Goals) + float64(p.Assists) / 2
    } else {
        rating = (float64(p.Goals) + float64(p.Assists) / 2) / float64(p.Misses) 
    }
	return rating
}

//конструктор для создания игрока и расчет рейтинга, и тут конструктор по условию задачи должен возвращать ЗНАЧЕНИЕ (не указатель)
func NewPlayer(name string, goals int, misses int, assists int) Player {
    p := Player{
        Name: name,
        Goals: goals,
        Misses: misses,
        Assists: assists,
    }
    p.Rating = p.calculateRating()
    return p
}

//функции для сортировок
func goalsSort(players []Player) []Player {
    //создаем копию слайса структур, чтобы не изменять исходный слайс
    sorted := make([]Player, len(players))
    copy(sorted, players)

    //сортируем по голам, по убыванию + если колич. голов одинак. то сорт. по именам
    slices.SortFunc(sorted, func(a, b Player) int {
            if a.Goals != b.Goals {
                    switch {
            case a.Goals < b.Goals:
                return 1
            case a.Goals > b.Goals:
                return -1
            default:
                return 0
            }
        }

        switch {
        case a.Name > b.Name:
            return 1
        case a.Name < b.Name:
            return -1
        default:
            return 0
        }
            
            
    })
    
    return sorted
}

func ratingSort(players []Player) []Player {
    sorted := make([]Player, len(players))
    copy(sorted, players)

    slices.SortFunc(sorted, func(a, b Player) int {
        if a.Rating != b.Rating {
            switch {
            case a.Rating < b.Rating:
                return 1
            case a.Rating > b.Rating:
                return -1
            default:
                return 0
            }
        }

        //тут сортировка если рейтинги равны, то переставляем по именам
        switch {
        case a.Name > b.Name:
            return 1
        case a.Name < b.Name:
            return -1
        default:
            return 0
        }
    })

    return sorted
}

//сортировка Убыванию отношения голов к промахам
func effectiveRatio (p Player) int {
    if p.Misses == 0 {
        return 1e9 // большое число (условная "бесконечность")
    }
    return p.Goals / p.Misses
}

func gmSort(players []Player) []Player {
    sorted := make([]Player, len(players))
    copy(sorted, players)

    //так, в этой сортировке нам нужно учесть момент когда колич. промахов равно нулю (т.к. делить на ноль нельзя), т.е мы считаем что игрок не промахивался и он должнен стоят выше всех и мы считаем такое отношение очень большим числом

    //для этого сделаем отедльную функуцию выше

    //+ нужно при равных G/M сделать сорт. по имени

    slices.SortFunc(sorted, func(a, b Player) int {
        ra := effectiveRatio(a)
        rb := effectiveRatio(b)
        
        // сорт. по убыв.
        if ra != rb {
                switch {
            case ra < rb:
                return 1
            case ra > rb:
                return -1
            default:
                return 0
            }
        }

        switch {
        case a.Name > b.Name:
            return 1
        case a.Name < b.Name:
            return -1
        default:
            return 0
        }
        
    })

    return sorted
}


func main() {
	// fmt.Println("Hello World")
    // p := NewPlayer("Jane", 20, 5, 12)
    // p1 := NewPlayer("Casey", 32, 2, 24)
    // fmt.Println(p)

    // storage := &PlayerStorage{
    //     Players: []Player{},
    // }

    // storage.AddPlayer(p)
    // storage.AddPlayer(p1)


    // fmt.Println(storage.Players)
 //--------------------------------------------------------
    players := []Player{
        {"Charlie", 10, 5, 6, 1.6}, //10/5 2
        {"Blice", 12, 2, 4, 4.0}, // 12/2 6
        {"Aob", 5, 1, 3, 4.0}, // 5/1 5
        {"Player4", 10, 0, 4, 1.2},
        {"Ioby", 4, 1, 3, 3.5}, // 4/1 4
    }

    // {"Name", Goals, Misses, Assists, Rating}
    

    sorted := goalsSort(players)
    sorted2 := ratingSort(players)
    sorted3 := gmSort(players)
    

    fmt.Println(sorted)
    fmt.Println(sorted2)
    fmt.Println(sorted3)
}