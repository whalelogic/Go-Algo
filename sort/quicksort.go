// Package sort implements the sorting algorithms for integers.
package sort

import (
	"fmt"
)


// Quick Sort Algorithm

// QuickSort sorts a list of integers using the quick sort algorithm.
// Similar to Merge Sort, Quick Sort is a divide-and-conquer algorithm.
// It uses recursion to sort sub-lists created by partitioning the main list around a pivot element.

func QuickSort(list []int) []int {
	fmt.Println("Original list:", list)
	if len(list) <= 1 {
		return list
	}

	pivot := list[len(list)/2]
	left := []int{}
	right := []int{}
	middle := []int{}

	for _, value := range list {
		if value < pivot {
			left = append(left, value)
		} else if value == pivot {
			middle = append(middle, value)
		} else {
			right = append(right, value)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, middle...), right...)
}


