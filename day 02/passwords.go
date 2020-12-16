package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type password struct {
	min      int
	max      int
	letter   string
	password string
}

func (p password) isValid() bool {
	count := strings.Count(p.password, p.letter)
	return p.min <= count && count <= p.max
}

func (p password) isValid2() bool {
	first, second := string(p.password[p.min-1]) == p.letter, string(p.password[p.max-1]) == p.letter
	return (first || second) && first != second
}

func main() {
	passwords := readInput()
	fmt.Println(partOne(passwords))
	fmt.Println(partTwo(passwords))
}

func partOne(passwords []password) (r int) {
	for _, p := range passwords {
		if p.isValid() {
			r++
		}
	}
	return
}

func partTwo(passwords []password) (r int) {
	for _, p := range passwords {
		if p.isValid2() {
			r++
		}
	}
	return
}

func readInput() (r []password) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, toPassword(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return r
}

func toPassword(text string) password {
	re := regexp.MustCompile(`[: -]`)
	fields := re.Split(text, -1)
	return password{min: toInt(fields[0]), max: toInt(fields[1]), letter: fields[2],
		password: fields[4]}
}

func toInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}
