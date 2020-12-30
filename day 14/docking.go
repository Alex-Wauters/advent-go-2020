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
	instructions := readInput()
	partOne(instructions)
	partTwo(instructions)
}

func partOne(instructions []instruction) {
	memory := make(map[int]int64)
	mask := ""
	for _, inst := range instructions {
		if inst.action == "mask" {
			mask = inst.mask
		} else {
			memory[inst.destination] = applyMask(mask, inst.value)
		}
	}
	sum := int64(0)
	for _, val := range memory {
		sum += val
	}
	fmt.Println(sum)
}

func partTwo(instructions []instruction) {
	memory := make(map[int64]int)
	mask := ""
	for _, inst := range instructions {
		if inst.action == "mask" {
			mask = inst.mask
		} else {
			destinations := applyMask2(mask, inst.destination)
			for _, d := range destinations {
				memory[d] = inst.value
			}
		}
	}
	sum := 0
	for _, val := range memory {
		sum += val
	}
	fmt.Println(sum)
}

type instruction struct {
	action             string
	mask               string
	destination, value int
}

func readInput() (r []instruction) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, toInstruction(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return r
}

func applyMask(mask string, value int) int64 {
	binary := fmt.Sprint(strconv.FormatInt(int64(value), 2))
	resultString := ""
	for i := 0; i < len(mask); i++ {
		maskChar := string(mask[len(mask)-1-i])
		numberChar := "0"
		if len(binary)-1-i >= 0 {
			numberChar = string(binary[len(binary)-1-i])
		}
		if maskChar == "X" {
			resultString = numberChar + resultString
		} else {
			resultString = maskChar + resultString
		}
	}
	result, err := strconv.ParseInt(resultString, 2, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func applyMask2(mask string, value int) (r []int64) {
	binary := fmt.Sprint(strconv.FormatInt(int64(value), 2))
	strings := []string{""}
	for i := 0; i < len(mask); i++ {
		maskChar := string(mask[len(mask)-1-i])
		for sIndex, s := range strings {
			numberChar := "0"
			if len(binary)-1-i >= 0 {
				numberChar = string(binary[len(binary)-1-i])
			}
			if maskChar == "X" {
				strings[sIndex] = "0" + s
				strings = append(strings, "1"+s)
			} else if maskChar == "0" {
				strings[sIndex] = numberChar + s
			} else if maskChar == "1" {
				strings[sIndex] = maskChar + s
			}
		}
	}
	for _, s := range strings {
		number, _ := strconv.ParseInt(s, 2, 64)
		r = append(r, number)
	}
	return r
}

func toInstruction(s string) instruction {
	memRe := regexp.MustCompile(`^mem\[(?P<dest>\d+)] = (?P<val>\d+)$`)
	if strings.HasPrefix(s, "mask") {
		fields := strings.Split(s, " ")
		return instruction{action: "mask", mask: fields[len(fields)-1]}
	} else if strings.HasPrefix(s, "mem") {
		if !memRe.MatchString(s) {
			panic("Did not recognize " + s)
		}
		res := memRe.FindStringSubmatch(s)
		return instruction{action: "assign", destination: toInt(res[1]), value: toInt(res[2])}
	} else {
		panic(s)
	}
}

func toInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}
