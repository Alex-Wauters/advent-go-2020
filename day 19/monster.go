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

type rule struct {
	letter   string
	subrules [][]int
}

func main() {
	rules, messages := readInput()
	allCombos := findAllCombinations(rules, 0)
	sumPartOne := 0
	for _, message := range messages {
		if contains(allCombos, message) {
			sumPartOne++
		}
	}
	fmt.Println("Part 1: ", sumPartOne)
	all42, all31 := findAllCombinations(rules, 42), findAllCombinations(rules, 31)
	sumPartTwo := 0
	for _, message := range messages {
		if isValidP2(all42, all31, message, 0, 0) {
			sumPartTwo++
		}
	}
	fmt.Println("Part two", sumPartTwo)
}

// Original rule in part 1 is [0: 8 11]
// In the changes in p2, The first part (8) repeated as: (42+ 11)
// Second part (11) is repeated as 42{n+} 31{n+} with n >= 1
// So a message must be compliant with 42+ 42{n+} 31{n+}.
// Check this by removing all possible prefixes (42) and suffixes (31) from the word and comparing the count.
func isValidP2(all42, all31 []string, word string, prefixCount, suffixCount int) bool {
	if word == "" && prefixCount > 0 && suffixCount > 0 && prefixCount > suffixCount {
		return true
	}
	if prefixCount > 0 {
		for _, w31 := range all31 {
			if strings.HasSuffix(word, w31) {
				replaced := word[:len(word)-len(w31)]
				if newSuffixCount := suffixCount + 1; isValidP2(all42, all31, replaced, prefixCount, newSuffixCount) {
					return true
				}
			}
		}
	}
	for _, w42 := range all42 {
		if strings.HasPrefix(word, w42) {
			replaced := strings.Replace(word, w42, "", 1)
			if newPrefixCount := prefixCount + 1; isValidP2(all42, all31, replaced, newPrefixCount, suffixCount) {
				return true
			}
		}
	}
	return false
}

func contains(messages []string, s string) bool {
	for _, m := range messages {
		if m == s {
			return true
		}
	}
	return false
}

func findAllCombinations(rules map[int]rule, ruleNumber int) (result []string) {
	r := rules[ruleNumber]
	if r.letter != "" {
		return append(result, r.letter)
	}
	for _, subrule := range r.subrules {
		newcombos := make([]string, 0)
		if len(subrule) == 1 {
			newcombos = findAllCombinations(rules, subrule[0])
		} else if len(subrule) == 2 {
			newcombosp1 := findAllCombinations(rules, subrule[0])
			newcombosp2 := findAllCombinations(rules, subrule[1])
			for _, p1 := range newcombosp1 {
				for _, p2 := range newcombosp2 {
					newcombos = append(newcombos, p1+p2)
				}
			}
		}
		result = append(result, newcombos...)
	}
	return result
}

func readInput() (rules map[int]rule, messages []string) {
	rules = make(map[int]rule)
	outputRe := regexp.MustCompile(`^([ab]+)$`)
	charRuleRe := regexp.MustCompile(`^(\d+): "([ab]+)"$`)
	singleRuleRe := regexp.MustCompile(`^(\d+): (\d+)$`)
	multiRuleRe := regexp.MustCompile(`^(\d+): (\d+) (\d+)$`)
	eitherSingleRuleRe := regexp.MustCompile(`^(\d+): (\d+) \| (\d+)$`)
	eitherMultiRuleRe := regexp.MustCompile(`^(\d+): (\d+) (\d+) \| (\d+) (\d+)$`)

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if outputRe.MatchString(line) {
			messages = append(messages, line)
		} else if charRuleRe.MatchString(line) {
			fields := charRuleRe.FindStringSubmatch(line)
			rules[toInt(fields[1])] = rule{letter: fields[2]}
		} else if singleRuleRe.MatchString(line) {
			fields := singleRuleRe.FindStringSubmatch(line)
			rules[toInt(fields[1])] = rule{subrules: [][]int{{toInt(fields[2])}}}
		} else if multiRuleRe.MatchString(line) {
			fields := multiRuleRe.FindStringSubmatch(line)
			rules[toInt(fields[1])] = rule{subrules: [][]int{{toInt(fields[2]), toInt(fields[3])}}}
		} else if eitherSingleRuleRe.MatchString(line) {
			fields := eitherSingleRuleRe.FindStringSubmatch(line)
			rules[toInt(fields[1])] = rule{subrules: [][]int{{toInt(fields[2])}, {toInt(fields[3])}}}
		} else if eitherMultiRuleRe.MatchString(line) {
			fields := eitherMultiRuleRe.FindStringSubmatch(line)
			rules[toInt(fields[1])] = rule{subrules: [][]int{{toInt(fields[2]), toInt(fields[3])}, {toInt(fields[4]), toInt(fields[5])}}}
		} else {
			if line != "" {
				panic("Could not parse " + line)
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
