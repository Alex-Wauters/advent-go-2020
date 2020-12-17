package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	passes := readInput()
	fmt.Println(partOne(passes))
	fmt.Println(partTwo(passes))
}

func partOne(passes []string) int {
	if seatId("FBFBBFFRLR") != 357 {
		panic("Test case failed for 357")
	}
	highest := 0
	for _, s := range passes {
		seat := seatId(s)
		if seat > highest {
			highest = seat
		}
	}
	return highest
}

func partTwo(passes []string) int {
	var seats []int
	for _, s := range passes {
		seats = append(seats, seatId(s))
	}
	sort.Ints(seats)
	for i := 1; i < len(seats)-1; i++ {
		if seats[i+1] != seats[i]+1 {
			return seats[i] + 1
		}
	}
	panic("Could not find seat")
}

func seatId(s string) int {
	return row(s[0:7])*8 + col(s[7:])
}

func row(s string) int {
	lower, higher := 0.0, 127.0
	for _, c := range s {
		diff := math.Ceil((higher - lower) / 2)
		if c == 'F' {
			higher = higher - diff
		} else {
			lower = lower + diff
		}
	}
	if lower != higher {
		fmt.Errorf("lower: %v, higher: %v", lower, higher)
	}
	return int(lower)
}

func col(s string) int {
	lower, higher := 0.0, 7.0
	for _, c := range s {
		diff := math.Ceil((higher - lower) / 2)
		if c == 'L' {
			higher = higher - diff
		} else {
			lower = lower + diff
		}
	}
	if lower != higher {
		fmt.Errorf("lower: %v, higher: %v", lower, higher)
	}
	return int(lower)
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
