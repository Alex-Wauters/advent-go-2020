package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// byr,iyr,eyr,hgt,hcl,ecl,pid,cid string
type passport map[string]string

func (p passport) isValid() bool {
	_, hasCid := p["cid"]
	if hasCid {
		return len(p) == 8
	}
	return len(p) == 7
}

func main() {
	passports := readInput()
	fmt.Println(partOne(passports))
}

func partOne(passports []passport) (r int) {
	for _, p := range passports {
		if p.isValid() {
			r++
		}
	}
	return r
}

func readInput() (r []passport) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	isAdding := false
	p := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isAdding = false
		} else {
			if !isAdding {
				p = make(map[string]string)
				r = append(r, p)
				isAdding = true
			}
			addFields(p, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return r
}

func addFields(p passport, s string) {
	fields := strings.Split(s, " ")
	for _, f := range fields {
		values := strings.Split(f, ":")
		p[values[0]] = values[1]
	}
}
