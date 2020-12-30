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

func (r rule) isValid(v int) bool {
	return (r.r1start <= v && v <= r.r1end) || (r.r2start <= v && v <= r.r2end)
}

type rule struct {
	field                          string
	r1start, r1end, r2start, r2end int
	found                          bool
}

func main() {
	rules, mine, others := readInput()
	valid := partOne(rules, others)
	partTwo(rules, mine, append(valid, mine))
}

func partOne(r []rule, tickets [][]int) (valid [][]int) {
	sum := 0
	for _, t := range tickets {
		isValidTicket := true
		for _, f := range t {
			isValidField := false
			for _, r := range r {
				if r.isValid(f) {
					isValidField = true
					break
				}
			}
			if !isValidField {
				sum += f
				isValidTicket = false
			}
		}
		if isValidTicket {
			valid = append(valid, t)
		}
	}
	fmt.Println(sum)
	return valid
}

func filterPossibleRules(rules []rule, value int) (result []rule) {
	for _, r := range rules {
		if r.isValid(value) {
			result = append(result, r)
		}
	}
	return result
}

func partTwo(r []rule, mine []int, tickets [][]int) {
	possibleRules := make(map[int][]rule)
	for _, ticket := range tickets {
		for i, field := range ticket {
			alreadyFiltered, exists := possibleRules[i]
			if !exists {
				possibleRules[i] = filterPossibleRules(r, field)
			} else {
				possibleRules[i] = filterPossibleRules(alreadyFiltered, field)
			}
		}
	}
	for rulesFound := 0; rulesFound < len(r); {
		for _, rules := range possibleRules {
			if len(rules) == 1 && !rules[0].found {
				for otherIndex, others := range possibleRules {
					if len(others) > 1 {
						for i, o := range others {
							if o.field == rules[0].field {
								possibleRules[otherIndex] = append(others[:i], others[i+1:]...)
								break
							}
						}
					}
				}
				rules[0].found = true
				rulesFound++
			}
		}
	}
	result := 1
	for i, rule := range possibleRules {
		if strings.HasPrefix(rule[0].field, "departure") {
			result *= mine[i]
		}
	}
	fmt.Println(result)
}

func readInput() (r []rule, mine []int, others [][]int) {
	ruleRe := regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)
	ticketRe := regexp.MustCompile(`^(\d+,?)+`)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if ruleRe.MatchString(line) {
			fields := ruleRe.FindStringSubmatch(line)
			r = append(r, rule{fields[1], toInt(fields[2]), toInt(fields[3]), toInt(fields[4]), toInt(fields[5]), false})
		} else if ticketRe.MatchString(line) {
			fields := strings.Split(line, ",")
			ticket := make([]int, 0)
			for _, f := range fields {
				ticket = append(ticket, toInt(f))
			}
			if len(mine) == 0 {
				mine = ticket
			} else {
				others = append(others, ticket)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func toInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}
