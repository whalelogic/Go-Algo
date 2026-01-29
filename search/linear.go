package search

import "fmt"


func LinearSearch(arr []int, target int) int {
	// Super slow linear search implementation
	for i, v := range arr {
		if v == target {
			fmt.Println("Found target at index: ", i)
			return i
		}
	}
	return -1
}
