package main

import "time"
import "fmt"

// START OMIT
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "bir"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "iki"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// END OMIT
