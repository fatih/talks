package main

import "fmt"

// START OMIT

var globalVar = func() {
	fmt.Println("Function assigned to a variable")
}

func unexported(x func()) {
	fmt.Println("Private (unexported) function")
	x()
}

func Exported() {
	fmt.Println("Public Exported function")
}

func main() {
	globalVar()
	unexported(globalVar)
	Exported()

	globalVar = func() {
		fmt.Println("Override function variable")
	}
	globalVar()
}

// END OMIT
