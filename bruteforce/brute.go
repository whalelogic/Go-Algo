// Package bruteforce provides functions to perform brute-force attacks 
// on an arbitrary set of PIN numbers 0-9999.
package bruteforce

import (
	"fmt"
	"time"
)


func CrackPIN(target string, pins []string) (string, bool) {
    fmt.Println("Cracking PIN with brute force...") 
    start := time.Now() 
    
    for _, pin := range pins {
        if pin == target {
            elapsed := time.Since(start) 
            
            fmt.Println("PIN found:", pin)
            fmt.Println("Time taken:", elapsed)
            return pin, true
        }
    }
	return "Not found", false
}


func GeneratePINs() []string {
	pins := make([]string, 0, 10000) 
	for i := 0; i <= 9999; i++ {
		pin := fmt.Sprintf("%04d", i)
		pins = append(pins, pin)
	}
	return pins
}


