package main

import "fmt"

func BinarySearch(arr []int, item int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else if guess < item {
			low = mid + 1
		}
	}
	return -1
}
func main() {
	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1))
}
