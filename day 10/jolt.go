package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	adapters := readInput()
	sort.Ints(adapters)
	partOne(adapters)
	partTwo(adapters)
}

func partOne(adapters []int) {
	numDifferences := make(map[int]int)
	jolt := 0
	for _, val := range adapters {
		diff := val - jolt
		if _, exists := numDifferences[diff]; !exists {
			numDifferences[diff] = 0
		}
		numDifferences[diff] += 1
		jolt = val
	}
	jolt += 3
	if _, exists := numDifferences[3]; !exists {
		numDifferences[3] = 0
	}
	numDifferences[3] += 1
	fmt.Println(numDifferences[1] * numDifferences[3])
}

func partTwo(adapters []int) {
	islands := toIslands(adapters)
	fmt.Println(islands)
	numCombinationsForLength := []int{0, 0, 0, 2, 4, 7}
	result := 1
	for _, island := range islands {
		if len(island) > 2 {
			result *= numCombinationsForLength[len(island)]
		}
	}
	fmt.Println(result)
}

// Divide input into 'islands' of numbers, with diff of 3 on either side. The edges can't be removed, but the numbers
// in-between can and only have diff 1 (as part 1 hinted).
func toIslands(adapters []int) (result [][]int) {
	lastIndex := -1
	jolt := 0
	for i, val := range adapters {
		diff := val - jolt
		if diff == 3 {
			if lastIndex == -1 {
				result = append(result, append([]int{0}, adapters[:i]...))
			} else {
				result = append(result, adapters[lastIndex:i])
			}
			lastIndex = i
		}
		jolt = val
	}
	result = append(result, adapters[lastIndex:])
	return result
}

func readInput() (r []int) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, toInt(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return r
}

func toInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}
