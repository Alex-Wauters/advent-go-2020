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
}

func partOne(rows []row) (r int) {
	for x, y := 0, 0; y < len(rows); x, y = x+3, y+1 {
		if rows[y].isTree(x) {
			r++
		}
	}
	return
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
