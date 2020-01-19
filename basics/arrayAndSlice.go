package main

import (
	"fmt"
	"sort"
)

func main() {

	num := []int{1, 2, 9, 10, 3, 4, 5, 6, 7, 8}

	num = append(num, 11)

	fmt.Println("Unsorted NUmber: ", num)

	sort.Ints(num)

	fmt.Println("Sorted NUmber: ", num)

	//slicing an array
	numSlice := num[:5]
	fmt.Println("array before", num)
	for i := range numSlice {
		//original array will also be modified
		numSlice[i] = numSlice[i] * 3
	}
	fmt.Println("array after: ", num)

	fmt.Println("Sliced Array: ", numSlice)

}
