package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type op struct {
	action string
	delta  int
}

func main() {
	ops := readInput()
	fmt.Println(partOne(ops))
}

func partOne(ops []op) (acc int) {
	executed := make(map[int]bool)
	index := 0
	for {
		if _, isExecuted := executed[index]; isExecuted {
			return acc
		}
		if ops[index].action == "acc" {
			acc += ops[index].delta
		}
		executed[index] = true
		if ops[index].action == "jmp" {
			index += ops[index].delta
		} else {
			index += 1
		}
	}
}

func readInput() (r []op) {
	re := regexp.MustCompile(`^(?P<Op>.+) (?P<Delta>.+)$`)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := re.FindStringSubmatch(line)
		r = append(r, op{
			action: fields[1],
			delta:  toInt(fields[2]),
		})
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
