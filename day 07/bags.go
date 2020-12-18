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

type content struct {
	bag string
	num int
}
type rule struct {
	bag     string
	content []content
}

var rules map[string]rule

func main() {
	readInput()
	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func canCarryGoldBag(bag string) bool {
	if bag == "shiny gold" {
		return true
	}
	rule, hasRule := rules[bag]
	if !hasRule {
		panic("Could not find rule for " + bag)
	}
	if len(rule.content) == 0 {
		return false
	}
	for _, c := range rule.content {
		if canCarryGoldBag(c.bag) {
			return true
		}
	}
	return false
}

func partOne() (c int) {
	for _, r := range rules {
		if r.bag != "shiny gold" && canCarryGoldBag(r.bag) {
			c++
		}
	}
	return c
}

func partTwo() (c int) {
	return totalBags("shiny gold") - 1
}

func totalBags(bag string) (count int) {
	rule := rules[bag]
	count = 1
	for _, c := range rule.content {
		count += c.num * totalBags(c.bag)
	}
	return count
}

func readInput() {
	rules = make(map[string]rule)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rule := toRule(scanner.Text())
		rules[rule.bag] = rule
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func toRule(s string) (r rule) {
	re := regexp.MustCompile(`^(?P<Bag>.+) bags contain (?P<Content>.+)\.$`)
	if !re.MatchString(s) {
		panic("Did not recognize " + s)
	}
	res := re.FindStringSubmatch(s)
	r.bag = res[1]
	r.content = make([]content, 0)
	if res[2] == "no other bags" {
		return r
	}
	cre := regexp.MustCompile(`^(?P<Count>\d+) (?P<Bag>.+) (bags|bag)$`)
	contents := strings.Split(res[2], ", ")
	for _, c := range contents {
		contentFields := cre.FindStringSubmatch(c)
		r.content = append(r.content, content{
			bag: contentFields[2],
			num: toInt(contentFields[1]),
		})
	}
	return r

}

func toInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}
