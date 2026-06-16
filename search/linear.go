// Package search implements a simple linear search algorithm.
package search

import "fmt"


func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			fmt.Println("Found target at index: ", i)
			return i
		}
	}
	return -1
}


func LinearSearchString(arr []string, target string) int {
	for i, v := range arr {
		if v == target {
			fmt.Println("Found target at index: ", i)
			return i
		}
	}
	return -1
}


