package main

import "fmt"

func main() {
	defer fmt.Println("ÖYG!")
	defer fmt.Println("Dünya")

	fmt.Println("Merhaba")
}
