package main

import "fmt"

// START OMIT
func main() {
	a := true       // bool
	b := "Istanbul" // string
	c := -13        // int (int8, int16, int32, int64)
	d := 2015       // uint (uint8, uint16, uint32, uint64)
	e := 3.14       // float32 (float64)

	fmt.Printf("%v %v %v %v %v\n", a, b, c, d, e)

	f := [4]int{1, 2, 3, 4}           // array
	g := []int{2, 3, 5, 7, 11, 13}    // slice
	h := map[string]int{"Ankara": 06} // map
	j := struct {                     // struct
		id     int
		server string
	}{id: 123, server: "Stargate"}

	fmt.Printf("%v %v %v %v\n", f, g, h, j)

	// ve daha nicesi ...
	// functions
	// channels
}

// END OMIT
