package main

import (
	"fmt"
	"time"
)

// START OMIT
func f(from string, wait time.Duration) {
	time.Sleep(wait)
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct", 0)

	f("goroutine", time.Second) // go ekleyelim

	func(msg string) { // go ekleyelim
		fmt.Println(msg)
	}("going")

	// time.Sleep(time.Second * 2)
}

// END OMIT
