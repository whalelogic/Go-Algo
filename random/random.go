// Package random provides functions to generate random PINs.
package random

import (
	"fmt"
	"math/rand"
	"time"


)


// RandomPIN generates a random 4-digit PIN as an integer (0-9999)
func RandomPIN() int {
	rand.NewSource(time.Now().UnixNano())
	pin := rand.Intn(10000) 
	return pin
}


// RandomPINString is same as RandomPIN but returns a string with leading zeros
func RandomPINString() string {
	pin := randIntn(10000)
	return fmt.Sprintf("%04d", pin) // 4-digits and leading zeros
}


// randIntn generates a random integer in [0, n)
func randIntn(n int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(n)
}
