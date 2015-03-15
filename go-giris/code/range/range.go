package main

import "fmt"

// START OMIT
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
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
		fmt.Println(i)
	}
}

// END OMIT
