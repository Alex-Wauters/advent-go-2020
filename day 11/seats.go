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
	occupied(lines, nextChar1)
	occupied(lines, nextChar2)
}

func occupied(lines []string, nextChar func([]string, int, int) rune) {
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

func nextChar1(lines []string, row, col int) rune {
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

func isOccupiedInDir(lines []string, row, col, deltaRow, deltaCol int) bool {
	if (deltaCol == 0 && deltaRow == 0) || (col+deltaCol) < 0 || (col+deltaCol) >= len(lines[row]) || (row+deltaRow) < 0 || (row+deltaRow) >= len(lines) {
		return false
	}
	if lines[row+deltaRow][col+deltaCol] == '#' {
		return true
	}
	if lines[row+deltaRow][col+deltaCol] == 'L' {
		return false
	}
	return isOccupiedInDir(lines, row+deltaRow, col+deltaCol, deltaRow, deltaCol)
}

func nextChar2(lines []string, row, col int) rune {
	currChar := lines[row][col]
	if currChar == '.' {
		return '.'
	}
	numAdjacent := 0
	for _, deltaRow := range []int{-1, 0, 1} {
		for _, deltaCol := range []int{-1, 0, 1} {
			if isOccupiedInDir(lines, row, col, deltaRow, deltaCol) {
				numAdjacent++
			}
		}
	}
	if currChar == 'L' && numAdjacent == 0 {
		return '#'
	}
	if currChar == '#' && numAdjacent >= 5 {
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
