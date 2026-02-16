package util

import "fmt"

func RemainingSpace(a []any) int {
	r := cap(a) - len(a)
	fmt.Printf("Remaining space: %d\n", r)
	return r
}
