package main

import "fmt"

// START OMIT

var a = func() {
	fmt.Println("Function assigned to a variable")
}

func b(x func()) {
	fmt.Println("Private (unexported) function")
	x()
}

func Exported() {
	fmt.Println("Public Exported function")
}

func main() {
	a()
	b(a)
	Exported()

	a = func() {
		fmt.Println("Override function variable")
	}
	a()
}

// END OMIT
