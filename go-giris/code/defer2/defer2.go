package main

import (
	"fmt"
	"sync"
)

// START OMIT
var mu sync.Mutex

func main() {
	deadlock(4)

	mu.Lock()
	fmt.Println("finished")
	mu.Unlock()
}

func deadlock(a int) {
	mu.Lock()
	if 5 < a {
		return
	}
	mu.Unlock()
}

// END OMIT
