package main

import "fmt"

// START OMIT

func main() {
	messages := make(chan string, 1)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}

// END OMIT
