package main

import (
	"fmt"
	"sync"
)

// START OMIT
var mu sync.Mutex

func main() {
	deadlock(4) // daha büyük bir sayı yazalım

	mu.Lock()
	fmt.Println("finished")
	mu.Unlock()
}

func deadlock(a int) {
	mu.Lock()
	if 4 < a {
		return
	}
	mu.Unlock()
}

// END OMIT
