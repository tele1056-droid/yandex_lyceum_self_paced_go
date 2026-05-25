package main

import (
	"fmt"
	"slices"
	//"sort"

)

type CompanyInterface interface{
    AddWorkerInfo(name string, position string, salary uint, experience uint) error
    SortWorkers() ([]string, error) 
}

type Worker struct {
	Name string
	Position string
	Salary uint
	Experience uint
}

type Company struct {
	Workers []Worker
}



//тут нужно сделать метод который добавляет инф. о сотруднике и сохраняет, т.е в методе мы делаем экземпляр типа и потом добавляем (сохраняем) его в слайс структур в хранилище так сказать
func (c *Company) AddWorkerInfo(name string, position string, salary uint, experience uint) error {
	//и нужно сделать проверки
	if name == "" {
		return fmt.Errorf("net name")
	}

	//проверяем валид должность
	switch position {
	case "директор", "зам. директора", "начальник цеха", "мастер", "рабочий":
		//тут если должность верная, то просто выйдет из свитч, если фалс везде, то ошибка
	default:
		return fmt.Errorf("invalid position")
	}

	// if salary < 0 {
	// 	return fmt.Errorf("negative oplata")
	// }

	//создаем экземпляр по шаблону Worker
	worker := Worker{Name: name, Position: position, Salary: salary, Experience: experience}
	
	//и добавляем в хранилище Company
	c.Workers = append(c.Workers, worker)
	return nil
}

//добавляем функцию для расчёта дохода за всё время
func (w Worker) TotalIncome() uint {
	return w.Salary * w.Experience
}

//создаем мапу с должностями и каждой долж. присваимваем вес, чтобы в дальнейшем сортировать по должностям, тк по длине не получится, по алфавиту тоже
var roleWeight = map [string]int {
	"директор": 5,
	"зам. директора": 4,
	"начальник цеха": 3,
	"мастер": 2,
	"рабочий": 1,

}

//функцию для получения веса
func getRoleWeight(role string) int {
	w, ok := roleWeight[role] // тут получаем в w - вес для ключа (долж.) и в ok: true - если ключ есть в мапе, false если нет
	if ok == true {
		return w //если ключ есть в мапе, то возвращаем вес
	}
	return 0
}

//реализуем метод сортировки воркеров
func (c *Company) SortWorkers() ([]string, error) {
	//тут первым делом нужно отсортировать сотрудников, я так понимаю нужно именно сортировать структуры внутри слайса, по доходу за время работы на предпр. (это наверное произведение Salsry * Experience (и опыт нужно привести к float64))
	//и сортировка по убыванию

	slices.SortFunc(c.Workers, func(a, b Worker) int {

		//проверяем равны ли доходы, если равны то будем сортировать по должностям
		if a.TotalIncome() != b.TotalIncome() {
			//этот свитч для сорт. по убыванию, и наоборот для возраст.
			switch {
		case a.TotalIncome() > b.TotalIncome():
			return -1
		case a.TotalIncome() < b.TotalIncome():
			return 1
		default:
			return 0
		}
		}

		//тут сортировка для равных доходов (по убыванию), если доходы равны
		switch {
		case getRoleWeight(a.Position) > getRoleWeight(b.Position):
			return -1
		case getRoleWeight(a.Position) < getRoleWeight(b.Position):
			return 1
		default:
			return 0
		}
	})

	//тут будем составлять слайс стринговый для возврата, с фиксированым слайсом?
	result := make([]string, len(c.Workers))
	for i, w := range c.Workers {
		result[i] = fmt.Sprintf("%s — %d — %s", w.Name, w.TotalIncome(), w.Position)
	}
	return result, nil
}

func main() {

}
