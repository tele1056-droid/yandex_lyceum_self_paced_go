package main

import (
	"fmt"
)

/*
// этот вариант дженерика не будет работать, т.к. оператор "+" поддерживают не все типы (например, struct не поддерживает), и омпилятор не позволяет использовать операции, которые не гарантированы для всех типов
func Sum[T any](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}
*/	

//поэтому пишем своё ограничение, используя свой интрефейс
type MyConstraint interface {
	int | float64 | string
}

func Sum[T MyConstraint](nums []T) T {
	var total T
	for _, n := range nums {
        total += n
    }
    return total
}

func main() {
	intSum := Sum([]int{1, 2, 3})
	if intSum != 6 {
		fmt.Printf("Expected Sum([]int{1, 2, 3}) to be 6, but got %d", intSum)
	}

	floatSum := Sum([]float64{1.5, 2.5, 3.0})

	fmt.Println(intSum)
	fmt.Println(floatSum)
}

