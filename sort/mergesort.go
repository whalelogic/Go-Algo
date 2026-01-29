package sort

import (
	"fmt"
)


// Merge Sort Algorithm

// list is for reference. You would typically pass your own list to the MergeSort function.
var list = []int{64, 34, 25, 12, 22, 11, 90}


func Merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// MergeSort sorts a list of integers using the merge sort algorithm.
// This works by dividing the list into halves, sorting each half recursively,
// and then merging the sorted halves back together.

func MergeSort(list []int) []int {
	fmt.Println("Original list:", list)
	if len(list) <= 1 {
		return list
	}

	// divide the list
	mid := len(list) / 2
	// make 2 halves
	left := list[:mid]
	right := list[mid:]
	// sort each half
	left = MergeSort(left)
	right = MergeSort(right)
	// merge sorted halves
	return Merge(left, right)
}





	

