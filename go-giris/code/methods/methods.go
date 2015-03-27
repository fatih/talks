package main

import "fmt"

// START OMIT
type Point struct {
	X int
	Y int
}

func (p Point) Add(a int) int {
	return p.X + p.Y + a
}

func main() {
	p := Point{3, 4}
	fmt.Println(p.Add(5))
}

// END OMIT
