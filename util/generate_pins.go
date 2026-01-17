// Package util create a file containing all 4-digit PIN numbers from 0000 to 9999, one per line.
package util

import (
	"fmt"
)


// GeneratePINS Generates all 4-digit PINs from 0000 to 9999
func GeneratePINS() []string {	
	pins := []string{}
	for i := 0; i <= 9999; i++ {
		pin := fmt.Sprintf("%04d", i)
		pins = append(pins, pin)
	}
	return pins
}


//
//
// // Main function testing GeneratePINs and writing to a file
//
// func main() {
//
// 	path := "/home/whaler/github/goalgo/data/pins.txt"
// 	f, err := os.Create(path)
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 	}
// 	defer f.Close()
//
// 	w := bufio.NewWriter(f)
// 	var byteSum int
//
// 	fmt.Println(GeneratePINS()[0])       // Print first PIN
// 	fmt.Println(GeneratePINS()[len(GeneratePINS())-1]) // Print last PIN
// 	fmt.Println(len(GeneratePINS()))          // Print total number of PINs
// 	pins := GeneratePINS()
// 	fmt.Println("Total PINs generated: ", len(pins))
// 	fmt.Println(pins)
// 	for _, pin := range pins {
// 		b, e := fmt.Fprintf(w, "%s\n", pin)
// 		if e != nil {
// 			fmt.Println("error: ", e)
// 		}
// 		byteSum += b
// 	}
// 	fmt.Println("Total bytes written: ", byteSum)
//
// 	// Alternative way without storing all PINs in memory
// 	//
// 	// 	b, e := fmt.Fprintf(w, "%04d\n", i)
// 	// 	byteSum += b
// 	// fmt.Println("Total bytes written: ", byteSum)
// 	//
// 	w.Flush()
// 	fmt.Println("Done!")
//
// }
