package sort

import (
	"fmt"
)


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

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}


func MergeSort(list []int) []int {
	fmt.Println("Original list:", list)
	if len(list) <= 1 {
		return list
	}

	// divide the list
	mid := len(list) / 2
	fmt.Println("mid: ", mid)
	// make 2 halves
	left := list[:mid]
	right := list[mid:]
	// sort each half
	left = MergeSort(left)
	right = MergeSort(right)
	// merge sorted halves
	return Merge(left, right)
}





	

