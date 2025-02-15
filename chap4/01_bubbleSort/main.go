package main

import "fmt"

type BubbleSortUtil struct {
}

func NewBubbleSortUtil() *BubbleSortUtil {
	return &BubbleSortUtil{}
}

func (bubbleSortUtil *BubbleSortUtil) sort(integers []int) {
	isSwapped := true
	for isSwapped {
		isSwapped = false
		for i := 1; i < 11; i++ {
			if integers[i-1] > integers[i] {
				temp := integers[i]
				integers[i] = integers[i-1]
				integers[i-1] = temp
				isSwapped = true
			}
		}
	}
	fmt.Println(integers)
}

func main() {
	integers := []int{
		31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10,
	}
	bubbleSortUtil := NewBubbleSortUtil()
	bubbleSortUtil.sort(integers)
}
