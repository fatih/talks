package main

import "fmt"

// START OMIT
var sayı = 13

const aciklama = "Örnek bir Go uygulamasi"

func hesapla(x, y, z int) (int, bool) {
	a := x + y + z
	b := 20 < a
	return a, b
}

func main() {
	fmt.Printf("%s\n", aciklama)

	var a int
	b := 8

	toplam, buyuk := hesapla(a, b, sayı)
	fmt.Println(toplam, buyuk)
}

// END OMIT
