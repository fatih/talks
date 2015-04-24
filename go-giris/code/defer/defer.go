package main

import "fmt"

func main() {
	defer fmt.Println("ve Ay")
	defer fmt.Println("Dünya")

	fmt.Println("Merhaba")
}
