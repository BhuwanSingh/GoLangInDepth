package main

import "fmt"

func shareChannel(varChannel chan int) {
	for i := 0; i < 7; i++ {
		varChannel <- i
	}
	close(varChannel)
}

func main() {
	varIntChannel := make(chan int)
	go shareChannel(varIntChannel)
	for i := 0; i < 7; i++ {
		fmt.Println(<-varIntChannel)
	}

}
