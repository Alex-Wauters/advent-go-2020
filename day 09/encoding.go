package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := readInput()
	invalidNumber := partOne(numbers)
	fmt.Println("Part one:", invalidNumber)
	fmt.Println("Part two", partTwo(numbers, invalidNumber))
}

func hasSum(list *list.List, sum int) bool {
	for a := list.Front(); a != nil; a = a.Next() {
		for b := a.Next(); b != nil; b = b.Next() {
			if a.Value.(int)+b.Value.(int) == sum {
				return true
			}
		}
	}
	return false
}

func partOne(numbers []int) int {
	last25 := list.New()
	for i, val := range numbers {
		if i >= 25 {
			if !hasSum(last25, val) {
				return val
			}
		}
		if last25.Len() >= 25 {
			last25.Remove(last25.Back())
		}
		last25.PushFront(val)
	}
	panic("All entries had a valid sum")
}

func partTwo(numbers []int, target int) int {
mainloop:
	for i, val := range numbers {
		sum := val
		smallest, largest := val, val
		for k := i + 1; k < len(numbers); k++ {
			sum += numbers[k]
			if sum > target {
				continue mainloop
			}
			if numbers[k] < smallest {
				smallest = numbers[k]
			}
			if numbers[k] > largest {
				largest = numbers[k]
			}
			if sum == target {
				return smallest + largest
			}
		}
	}
	panic("Could not find a combination of contiguous values that sum up to the invalid value")
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
