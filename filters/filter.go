package main

import (
	"fmt"
)

func FilterInPlace(s []string, filter func(string) bool) {
	j := 0
	for _, v := range s {
		if filter(v) {
			s[j] = v
			j++
		}
	}
	s = s[:j]
	r := cap(s) - len(s)
	// Clear the remaining elements to allow garbage collection
	for i := len(s); i < cap(s); i++ {
		s = append(s, "")
	}
	fmt.Printf("Filtered slice: %v\n", s)
	fmt.Println("Remaining space: ", r)
}


func main() {
	strs := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Printf("Original slice: %v\n", strs)

	FilterInPlace(strs, func(s string) bool {
		return len(s) > 5
	})
}


