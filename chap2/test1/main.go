package main

import (
	"fmt"
	"time"
)

func execute(message string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(message, ":", i)
	}
}

func main() {
	fmt.Println("Go concurrency example")

	go execute("Go execute 1")
	go execute("Go execute 2")
	go execute("Go execute 3")

	time.Sleep(time.Second * 10)

	fmt.Println("Example completed.")
}
