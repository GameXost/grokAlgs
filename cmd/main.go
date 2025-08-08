package main

import (
	"fmt"
	rec "grokAlgs/recursion"
)

func main() {
	// array := []int{1, 3, 2, 5, 4, 2, 2, 4, 5}
	// // fmt.Println(array)
	// sortedArray := selection.SelectionSort(array)
	// fmt.Println(sortedArray)
	// // fmt.Println(binary.BinarySearch(sortedArray, 2))
	// selection.SelectionSortInPLace(array)
	// fmt.Println(array)

	// rec.CountdownReversed(10)
	// fmt.Println(rec.Sum([]int{1, 2, 3, 4}))
	fmt.Println(rec.Greatest([]int{1, 2, 3, 4, 5, 6, 7, 8}))
}
