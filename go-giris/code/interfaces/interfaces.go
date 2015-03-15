package main

import "fmt"

type Adder interface {
	Add(int) int
}

type Point struct {
	X, Y int
}

func (p Point) Add(a int) int {
	return p.X + p.Y + a
}

type MyInt int

func (m MyInt) Add(a int) int {
	return int(m) + a
}

// START OMIT
func addMagicNumber(a Adder) int {
	return a.Add(5)
}

func main() {
	p := Point{3, 4}
	m := MyInt(3)

	fmt.Printf("Point: %d\n", addMagicNumber(p))
	fmt.Printf("MyInt: %d\n", addMagicNumber(m))
}

// END OMIT
