package main

import (
	"fmt"
	"github.com/whalelogic/goalgo/util"
	"github.com/whalelogic/goalgo/search"
	"github.com/whalelogic/goalgo/sort"
	"slices"
	"strings"
)


func main() {
	str := "This is my Algorithms library"
	target := "Go"

	// Insert "Go" to complete the sentence.

	words := slices.Contains(strings.Fields(str), target)
	remaining := util.RemainingSpace([]any{str, target})

	fmt.Printf("Does the string contain '%s'? %v\n", target, words)
	fmt.Printf("Remaining space in the slice: %d\n", remaining)


	arr := []int{1, 2, 3, 4, 5}
	targetInt := 3

	t := search.LinearSearch(arr, targetInt)
	if t != -1 {
		fmt.Printf("Found target %d at index: %d\n", targetInt, t)
	} else {
		fmt.Printf("Target %d not found in the array.\n", targetInt)
	}


	newArr := []int{38, 27, 43, 3, 9, 82, 10}

	sortedNewArr := sort.MergeSort(newArr)
	
	fmt.Println("Sorted array:", sortedNewArr)


}
