package main

import (
	"fmt"
	"os"
	"time"
	"strconv"
	"github.com/whalelogic/goalgo/util"
	"github.com/whalelogic/goalgo/random"
	"github.com/whalelogic/goalgo/bruteforce"
	"github.com/whalelogic/goalgo/search"
)


func main() {
	t := time.Now()
	f := t.Format("15:04:05 01-02-2006")
	fmt.Println("Time and Date: ", f)
	fmt.Println(os.Getwd())
	// Use GeneratePINs to create the same dataset in memory
	pins := util.GeneratePINS()
	fmt.Println("Total PINs generated: ", len(pins))

	pinsIntArray := make([]int, len(pins))
	for i, pin := range pins {
		pinInt, err := strconv.Atoi(pin)
		if err != nil {
			fmt.Println("Error converting PIN to int: ", err)
			continue
		}
		pinsIntArray[i] = pinInt
	}



	linSearch := search.LinearSearch(pinsIntArray, 1234)
	fmt.Println("Linear Search Result: ", linSearch)


	// Writing to memory is always faster than reading from a disk

	targetPIN := random.RandomPINString()
	fmt.Println("Target PIN to crack: ", targetPIN)

	// bruteforce CrackPIN is also linear search but over strings
	bruteforce.CrackPIN(targetPIN, pins)

}
