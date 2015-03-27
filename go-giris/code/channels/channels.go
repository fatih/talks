package main

import (
	"fmt"
	"time"
)

// START OMIT

func main() {
	messages := make(chan string)

	fmt.Println("ping")

	go func() {
		time.Sleep(2 * time.Second)
		messages <- "pong"
	}()

	fmt.Println("bekliyorum ...")

	msg := <-messages
	fmt.Println("geldi:", msg)
}

// END OMIT
