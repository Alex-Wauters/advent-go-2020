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

func main() {
	expressions := readInput()
	partOne(expressions)
	partTwo(expressions)
}

func partOne(expressions []string) {
	sum := 0
	for _, e := range expressions {
		sum += solve(e)
	}
	fmt.Println(sum)
}

func partTwo(expressions []string) {
	sum := 0
	for _, e := range expressions {
		sum += solve2(e)
	}
	fmt.Println(sum)
}

func solve(expression string) int {
	subexpressionRe := regexp.MustCompile(`(\([+*\d\s]+\))`)
	reduced := expression
	for subexpressionRe.MatchString(reduced) {
		subexpr := subexpressionRe.FindStringSubmatch(reduced)[1]
		solution := calculate(subexpr)
		reduced = strings.ReplaceAll(reduced, subexpr, strconv.Itoa(solution))
	}
	return calculate(reduced)
}

func solve2(expression string) int {
	subexpressionRe := regexp.MustCompile(`(\([+*\d\s]+\))`)
	plusRe := regexp.MustCompile(`(\d+ \+ \d+)`)
	reduced := expression
	for subexpressionRe.MatchString(reduced) {
		originalSubExpression := subexpressionRe.FindStringSubmatch(reduced)[1]
		reducedPlusOperations := originalSubExpression
		for plusRe.MatchString(reducedPlusOperations) {
			plusExpr := plusRe.FindStringSubmatch(reducedPlusOperations)[1]
			solution := calculate(plusExpr)
			reducedPlusOperations = strings.Replace(reducedPlusOperations, plusExpr, strconv.Itoa(solution), 1)
		}
		solution := calculate(reducedPlusOperations)
		reduced = strings.Replace(reduced, originalSubExpression, strconv.Itoa(solution), 1)
	}
	for plusRe.MatchString(reduced) {
		plusExpr := plusRe.FindStringSubmatch(reduced)[1]
		solution := calculate(plusExpr)
		reduced = strings.Replace(reduced, plusExpr, strconv.Itoa(solution), 1)
	}
	return calculate(reduced)
}

// Calculate parts without subexpressions. It may contain begin and end parenthesis
func calculate(s string) int {
	lastOperand := " "
	var result int
	fields := strings.Split(strings.ReplaceAll(strings.ReplaceAll(s, "(", ""), ")", ""), " ")
	for _, c := range fields {
		if c == "+" || c == "*" {
			lastOperand = c
		} else {
			number := toInt(c)
			switch lastOperand {
			case " ":
				result = number
			case "+":
				result += number
			case "*":
				result *= number
			}
		}
	}
	return result
}

func toInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
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
	return
}
