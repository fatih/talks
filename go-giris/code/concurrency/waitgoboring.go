// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Bekliyorum...")
	go boring("boring!")

	time.Sleep(2 * time.Second)

	fmt.Println("Canım sıkıldı, ben çıkıyorum.")
}

// STOP OMIT

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
