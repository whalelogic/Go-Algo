package main

import (
	"fmt"
	"iter"
	"slices"
)

func printAll(seq iter.Seq2[int, int]) {
    for i, v := range seq {
        fmt.Printf("[%d] = %d\n", i, v)
    }
}


func timer(start int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i >= 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := slices.Backward(a[:])
	for i, v := range b {
		fmt.Printf("Index: %v  Value: %v\n", i, v)
	}

	newArr := [9]int{11,22,33,44,55,66,77,88,99}
	sorted := slices.Concat(a[:], newArr[:])
	fmt.Println(sorted)
	printAll(slices.Backward(sorted))

	seq := slices.All([]int{1, 2, 3})

	var firstVal int
	seq(func(i, v int) bool {
    	firstVal = v
		fmt.Println("First value:", firstVal)
    	return false
	})

	var sum int
	for _, v := range newArr {
		sum += v
	}
	fmt.Println("Sum of newArr integers: ", sum)
}
