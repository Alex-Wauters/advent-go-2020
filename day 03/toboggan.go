package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type row struct {
	pattern string
}

func (r row) isTree(i int) bool {
	return string(r.pattern[i%len(r.pattern)]) == "#"
}

func main() {
	rows := readInput()
	fmt.Println(partOne(rows))
	fmt.Println(partTwo(rows))
}

func trees(rows []row, addX, addY int) (r int) {
	for x, y := 0, 0; y < len(rows); x, y = x+addX, y+addY {
		if rows[y].isTree(x) {
			r++
		}
	}
	return
}

func partOne(rows []row) (r int) {
	return trees(rows, 3, 1)
}

func partTwo(rows []row) (r int) {
	return trees(rows, 1, 1) * trees(rows, 3, 1) * trees(rows, 5, 1) * trees(rows, 7, 1) * trees(rows, 1, 2)
}

func readInput() (r []row) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, row{scanner.Text()})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return r
}
