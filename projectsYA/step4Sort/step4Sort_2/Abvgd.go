package main

import (
	"fmt"
	//"slices"
	"sort"
)

func SortNames(names []string) {
	//тут примитивный тип данных string отсортирует пакет sort
	sort.Strings(names)
	fmt.Println(names)
}

func main() {
	names := []string{"Есения", "Аксинья", "Варвара", "Арина"}
	SortNames(names)
}