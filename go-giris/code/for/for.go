package main

import "fmt"

// START OMIT
func main() {
	sum := 0
	for i := 0; i < 4; i++ {
		sum += i
	}

	sum2 := 1
	for sum2 < 6 {
		sum2 += sum2
	}

	sum3 := 0
	for {
		sum3++
		if sum3 == 10 {
			break
		}
	}

	fmt.Printf("%d %d %d\n", sum, sum2, sum3)
}

// END OMIT
