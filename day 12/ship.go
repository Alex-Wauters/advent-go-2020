package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	lines := readInput()
	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	x, y := 0, 0
	dir := 90 // in degrees
	for _, line := range lines {
		val := toInt(line[1:])
		switch line[0] {
		case 'N':
			y += val
		case 'E':
			x += val
		case 'S':
			y -= val
		case 'W':
			x -= val
		case 'F':
			if dir == 0 {
				y += val
			} else if dir == 90 {
				x += val
			} else if dir == 180 {
				y -= val
			} else if dir == 270 {
				x -= val
			} else {
				panic(dir)
			}
		case 'R':
			dir = newDir(dir, val)
		case 'L':
			dir = newDir(dir, -val)
		}
	}
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func newDir(dir, delta int) int {
	newDir := dir + delta
	if newDir >= 360 {
		newDir = newDir % 360
	}
	if newDir < 0 {
		newDir = 360 + (newDir % 360)
	}
	return newDir
}

func rotateWaypoint(wpX, wpY, degrees int) (int, int) {
	switch degrees {
	case 90:
		return wpY, -wpX
	case 180:
		return -wpX, -wpY
	case 270:
		return -wpY, wpX
	case 360:
		return wpX, wpY
	}
	panic(degrees)
}

func partTwo(lines []string) {
	x, y := 0, 0
	wpX, wpY := 10, 1
	for _, line := range lines {
		val := toInt(line[1:])
		switch line[0] {
		case 'N':
			wpY += val
		case 'E':
			wpX += val
		case 'S':
			wpY -= val
		case 'W':
			wpX -= val
		case 'F':
			x += val * wpX
			y += val * wpY
		case 'R':
			wpX, wpY = rotateWaypoint(wpX, wpY, val)
		case 'L':
			wpX, wpY = rotateWaypoint(wpX, wpY, 360-val)
		}
	}
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
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

func toInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}
