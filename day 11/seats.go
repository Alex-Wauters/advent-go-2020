package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := readInput()
	partOne(lines)
}

func partOne(lines []string) {
	previous := lines
	for true {
		changed := make([]string, 0)
		for row, line := range previous {
			var sb strings.Builder
			for col := range line {
				sb.WriteRune(nextChar(previous, row, col))
			}
			changed = append(changed, sb.String())
		}
		if strings.Join(previous, ",") == strings.Join(changed, ",") {
			occupied := 0
			for _, line := range changed {
				occupied += strings.Count(line, "#")
			}
			fmt.Println("Occupied seats", occupied)
			return
		}
		previous = changed
	}
}

func nextChar(lines []string, row, col int) rune {
	currChar := lines[row][col]
	if currChar == '.' {
		return '.'
	}
	numAdjacent := 0
	for _, deltaRow := range []int{-1, 0, 1} {
		for _, deltaCol := range []int{-1, 0, 1} {
			if !(deltaCol == 0 && deltaRow == 0) && (col+deltaCol) >= 0 && (col+deltaCol) < len(lines[row]) && (row+deltaRow) >= 0 && (row+deltaRow) < len(lines) {
				if lines[row+deltaRow][col+deltaCol] == '#' {
					numAdjacent++
				}
			}
		}
	}
	if currChar == 'L' && numAdjacent == 0 {
		return '#'
	}
	if currChar == '#' && numAdjacent >= 4 {
		return 'L'
	}
	return rune(currChar)
}

func readInput() (r []string) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return r
}
