package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type group struct {
	answers []string
}

func (g group) count() int {
	letters := make(map[int32]bool)
	for _, a := range g.answers {
		for _, c := range a {
			letters[c] = true
		}
	}
	return len(letters)
}

func main() {
	groups := readInput()
	fmt.Println(partOne(groups))
}

func partOne(groups []group) (c int) {
	for _, g := range groups {
		c += g.count()
	}
	return c
}

func readInput() (r []group) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	isAdding := false
	var currentGroup group
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isAdding = false
			r = append(r, currentGroup)
		} else {
			if !isAdding {
				isAdding = true
				currentGroup = group{}
			}
			currentGroup.answers = append(currentGroup.answers, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return r
}
