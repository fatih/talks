package main

import (
	"fmt"
	"time"
)

// placeholder so we have the time package imported, so it doesn't break
// when we uncomment the time.Sleep statement inside the fibonacci
// function
var b = time.Second

// START OMIT
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		// slow producer
		// time.Sleep(time.Millisecond * 300)
		x, y = y, x+y
	}
	close(c)
}

func main() {
	bufSize := 10
	c := make(chan int, bufSize)

	// collect numbers
	go fibonacci(bufSize, c)

	for i := range c {
		// slow consumer
		// time.Sleep(time.Millisecond * 300)
		fmt.Println(i)
	}

	select {
	case v1 := <-chan1:
	case chan2 <- v2:
	default:
	}
}

// END OMIT
