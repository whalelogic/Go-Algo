package main

import (
	"fmt"
	"os"
	"bufio"
	"time"
	"github.com/whalelogic/goalgo/util"
	"github.com/whalelogic/goalgo/random"
	"github.com/whalelogic/goalgo/bruteforce"
)


func main() {
	fmt.Println(time.Now())
	fmt.Println(os.Getwd())
	var b = bufio.NewReader(os.Stdin)
	// Use GeneratePINs to create the same dataset in memory
	pins := util.GeneratePINS()
	fmt.Println("Total PINs generated: ", len(pins))
	fmt.Println("Press Enter to continue...")
	b.ReadString('\n')



	// Writing to memory is always faster than reading from a disk

	targetPIN := random.RandomPINString()
	fmt.Println("Target PIN to crack: ", targetPIN)

	bruteforce.CrackPIN(targetPIN, pins)

}
