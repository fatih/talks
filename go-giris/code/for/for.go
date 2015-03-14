package main

import "fmt"

// START OMIT
func main() {
	a := 0
	for i := 0; i < 4; i++ {
		a += i
	}

	b := 1
	for b < 6 {
		b += b
	}

	c := 0
	for {
		c++
		if c == 10 {
			break
		}
	}

	fmt.Printf("%d %d %d\n", a, b, c)
}

// END OMIT
