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

type passport map[string]string

func (p passport) isValid() bool {
	_, hasCid := p["cid"]
	if hasCid {
		return len(p) == 8
	}
	return len(p) == 7
}
func (p passport) isValid2() bool {
	return p.isValidYear("byr", 1920, 2002) && p.isValidYear("iyr", 2010, 2020) && p.isValidYear("eyr", 2020, 2030) && p.isValidHeight() && p.isValidEyeColor() && p.isValidHairColor() && p.isValidPid()
}

func (p passport) isValidYear(key string, min, max int) bool {
	field, hasField := p[key]
	if !hasField {
		return false
	}
	i, err := strconv.Atoi(field)
	if err != nil {
		return false
	}
	return min <= i && i <= max
}

func (p passport) isValidEyeColor() bool {
	color, hasColor := p["ecl"]
	if !hasColor {
		return false
	}
	for _, s := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if s == color {
			return true
		}
	}
	return false
}

func (p passport) isValidPid() bool {
	pid, hasPid := p["pid"]
	if !hasPid {
		return false
	}
	re := regexp.MustCompile(`^[0-9]{9}$`)
	return re.MatchString(pid)
}

func (p passport) isValidHairColor() bool {
	hcl, hasHcl := p["hcl"]
	if !hasHcl {
		return false
	}
	re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	return re.MatchString(hcl)
}

func (p passport) isValidHeight() bool {
	height, hasHeight := p["hgt"]
	if !hasHeight {
		return false
	}
	if strings.HasSuffix(height, "cm") {
		i, err := strconv.Atoi(height[:len(height)-2])
		if err != nil {
			return false
		}
		return 150 <= i && i <= 193
	}
	if strings.HasSuffix(height, "in") {
		i, err := strconv.Atoi(height[:len(height)-2])
		if err != nil {
			return false
		}
		return 59 <= i && i <= 76
	}
	return false
}

func main() {
	passports := readInput()
	fmt.Println(partOne(passports))
	fmt.Println(partTwo(passports))
}

func partOne(passports []passport) (r int) {
	for _, p := range passports {
		if p.isValid() {
			r++
		}
	}
	return r
}

func partTwo(passports []passport) (r int) {
	for _, p := range passports {
		if p.isValid2() {
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
