package main

import "fmt"

// START OMIT
var sayı = 13

const aciklama = "Örnek bir Go uygulamasi"

func hesapla(x, y, z int) (int, bool) {
	var a int = x + y + z
	b := 20 < a
	return a, b
}

func main() {
	fmt.Printf("%s\n", aciklama)

	var empty int
	number := 8

	toplam, buyuk := hesapla(empty, number, sayı)
	fmt.Println(toplam, buyuk)
}

// END OMIT
