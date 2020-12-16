package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := readInput()
	for i, _ := range numbers {
		for k := i + 1; k < len(numbers); k++ {
			if numbers[i]+numbers[k] == 2020 {
				fmt.Println(numbers[i] * numbers[k])
				return
			}
		}
	}
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
	i, _ = strconv.Atoi(s)
	return
}
