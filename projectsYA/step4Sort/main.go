package main

import (
	"fmt"
	//"sort"
	"slices"
)

func SortNums(nums []uint) {
	// тут []uint - слайс целых неотрицательных целых чисел от 0
	slices.Sort(nums)
	fmt.Println(nums)
}

func main() {
	uiInts := []uint{20, 65, 0, 9, 5, 15, 2, 34}

	SortNums(uiInts)
}