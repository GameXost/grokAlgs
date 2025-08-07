package selection

func findSmallest(array []int) int {
	smallest := array[0]
	smallestInd := 0
	for i := 1; i < len(array); i++ {
		if array[i] < smallest {
			smallest = array[i]
			smallestInd = i
		}
	}
	return smallestInd
}

// not in-place sort
func SelectionSort(array []int) []int {
	copiedArray := make([]int, len(array))
	copy(copiedArray, array)
	sortedArray := make([]int, len(array))
	for ind := 0; ind < len(array); ind++ {
		smallestInd := findSmallest(copiedArray) // index smallest element in modified array
		smallestEl := copiedArray[smallestInd]
		sortedArray[ind] = smallestEl
		copiedArray = append(copiedArray[:smallestInd], copiedArray[smallestInd+1:]...)
	}
	return sortedArray
}

//in-place sort
func SelectionSortInPLace(array []int) {
	for i := 0; i < len(array)-1; i++ {
		minInd := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minInd] {
				minInd = j
			}
		}
		array[i], array[minInd] = array[minInd], array[i]
	}
}
