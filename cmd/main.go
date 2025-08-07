package main

import (
	"fmt"
	"grokAlgs/sorts/selection"
)

func main() {
	array := []int{1, 3, 2, 5, 4, 2, 2, 4, 5}
	// fmt.Println(array)
	sortedArray := selection.SelectionSort(array)
	fmt.Println(sortedArray)
	// fmt.Println(binary.BinarySearch(sortedArray, 2))
	selection.SelectionSortInPLace(array)
	fmt.Println(array)
}
