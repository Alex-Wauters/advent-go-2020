package main

import (
	"container/list"
	"fmt"
)

func main() {
	playGame(2020)
	playGame(30000000)
}

func playGame(max int) {
	seen := make(map[int]*list.List)
	starting := []int{1, 20, 11, 6, 12, 0}
	lastSpoken := 0
	for i := 1; i <= max; i++ {
		if i <= len(starting) {
			lastSpoken = starting[i-1]
		} else {
			lastSpoken = getNextValue(seen, lastSpoken)
		}
		addNumber(seen, lastSpoken, i)
	}
	fmt.Println(lastSpoken)
}

func addNumber(seen map[int]*list.List, value, iteration int) {
	prevValues, exists := seen[value]
	if !exists {
		list := list.New()
		list.PushFront(iteration)
		seen[value] = list
	} else {
		prevValues.PushFront(iteration)
		if prevValues.Len() > 2 {
			prevValues.Remove(prevValues.Back())
		}
	}
}

func getNextValue(seen map[int]*list.List, lastSpoken int) int {
	prevValues, _ := seen[lastSpoken]
	if prevValues.Len() == 1 {
		return 0
	}
	return prevValues.Front().Value.(int) - prevValues.Back().Value.(int)
}
