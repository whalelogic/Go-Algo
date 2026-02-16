package main

import (
	"fmt"
)

type Iterator[T any] struct {
	collection []T
	index      int
}

func NewIterator[T any](collection []T) *Iterator[T] {
	return &Iterator[T]{collection: collection, index: 0}
}

func (it *Iterator[T]) HasNext() bool {
	return it.index < len(it.collection)
}

func (it *Iterator[T]) Next() T {
	if !it.HasNext() {
		panic("No more elements in the collection")
	}
	element := it.collection[it.index]
	it.index++
	return element
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	iterator := NewIterator(numbers)

	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}

}
