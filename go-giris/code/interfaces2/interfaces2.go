package main

import (
	"fmt"
	"strconv"
)

type Adder interface {
	Add(int) int
}

type Point struct {
	X, Y int
}

func (p Point) Add(a int) int {
	return p.X + p.Y + a
}

func (p Point) String() string {
	return fmt.Sprintf("X: %d Y: %d Sum: %d", p.X, p.Y, p.Add(0))
}

type MyInt int

func (m MyInt) Add(a int) int {
	return int(m) + a
}

func (m MyInt) String() string {
	return "*** " + strconv.Itoa(int(m)) + " ***"
}

// START OMIT
func main() {
	p := Point{3, 4}
	m := MyInt(3)

	fmt.Println(p, "\n", m)
}

// END OMIT
