package main

import (
	"fmt"
	"sync"
)

var GEncounter = 0
var mu sync.Mutex

func executer(waitGroup *sync.WaitGroup) {
	mu.Lock()
	GEncounter += 1
	mu.Unlock()
	waitGroup.Done()
}

func main() {
	var varWaitGroup sync.WaitGroup
	for i := 0; i < 1000; i++ {
		varWaitGroup.Add(1)
		go executer(&varWaitGroup)
	}

	varWaitGroup.Wait()
	fmt.Println("Value of encounter:", GEncounter)
}
