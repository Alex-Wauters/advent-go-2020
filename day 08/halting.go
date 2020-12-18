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

func turnaround(o op) op {
	if o.action == "nop" {
		return op{action: "jmp", delta: o.delta}
	}
	if o.action == "jmp" {
		return op{action: "nop", delta: o.delta}
	}
	return o
}

func main() {
	ops := readInput()
	fmt.Println(partOne(ops))
	fmt.Println(partTwo(ops))
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

func partTwo(ops []op) int {
	for i, operation := range ops {
		replaced := append(make([]op, 0), ops[:i]...)
		replaced = append(replaced, turnaround(operation))
		replaced = append(replaced, ops[i+1:]...)
		executes, acc := executes(replaced)
		if executes {
			return acc
		}
	}
	return 0
}

func executes(ops []op) (bool, int) {
	acc := 0
	executed := make(map[int]bool)
	index := 0
	for {
		if index >= len(ops) {
			return true, acc
		}
		if _, isExecuted := executed[index]; isExecuted {
			return false, acc
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
