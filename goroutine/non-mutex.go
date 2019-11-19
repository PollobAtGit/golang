package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	m   map[string]int
	mux sync.Mutex
}

func (sc *safeCounter) IncUnsafely(n int) {
	sc.m["x"] = n
}

// seems like there isn't any difference between locking and unlocking
func (sc *safeCounter) Inc(n int) {
	sc.mux.Lock()

	sc.m["x"] = n

	sc.mux.Unlock()
}

func (sc *safeCounter) value(k string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()

	return sc.m["x"]
}

func main() {

	sc := safeCounter{m: make(map[string]int)}

	for i := 0; i < 20; i++ {
		time.Sleep(time.Second)
		go sc.IncUnsafely(i)
	}

	time.Sleep(time.Second)

	fmt.Println(sc.m["x"])
}
