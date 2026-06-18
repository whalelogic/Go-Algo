package main

import (
	"fmt"
)


func main() {
	m := make(map[rune]int)
	fmt.Println("Map initialized with size:", len(m))

	m['a'] = 00001
	m['b'] = 00010
	m['c'] = 00011
	m['d'] = 00100
	m['🚀'] = 333

	fmt.Printf("Value for 'a': %o\n", m['a'])
	fmt.Printf("Value for a in binary: %b\n", m['a'])
	fmt.Printf("Value for e: %U\n", m['🚀'])

	for _, v := range m {
		fmt.Println("values of map: ", v)
	}

}
