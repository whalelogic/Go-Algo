package main

import (
	"fmt"
	"github.com/whalelogic/goalgo/util"
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
}
