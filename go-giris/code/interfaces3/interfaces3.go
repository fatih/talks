package main

import "fmt"

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

type Adder interface {
	Add(int) int
}

func addMagicNumber(a Adder) int {
	return a.Add(5)
}

// START OMIT
func main() {
	p := Point{3, 4}
	m := MyInt(3)

	fmt.Println(p, "\n", m)
}

// END OMIT
