package main

import (
	"fmt"
)

func Filter[T any](arr []T, predicate func(T) bool) []T {
	var FilterArr []T
	for _ ,elem := range arr {
		if predicate(elem) {
			FilterArr = append(FilterArr, elem)
		}
	}
	return FilterArr
}


func main() {

}