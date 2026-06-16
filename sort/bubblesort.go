// Package sort defines some typical sorting algorithms.
package sort

// Bubble Sort Algorithm

import (
	"fmt"
)

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
