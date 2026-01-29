// Package sort defines some typical sorting algorithms.
package sort

// Bubble Sort Algorithm

import (
	"fmt"
)

// list is for reference. You would typically pass your own list to the BubbleSort function.
var list = []int{64, 34, 25, 12, 22, 11, 90}


// BubbleSort sorts a list of integers using the bubble sort algorithm.
// This works by repeatedly stepping through the list, comparing elements with 
// elements + 1, and swapping them if they are in the wrong order.
// The pass through the list is repeated until the list is sorted.
func BubbleSort(list []int) []int {
	fmt.Println("Original list:", list)

	l := len(list)

	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}


