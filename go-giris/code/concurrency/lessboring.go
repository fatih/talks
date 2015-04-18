// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	boring("boring!") // HL
}

// START OMIT
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

// STOP OMIT
