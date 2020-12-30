package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	cube := []string{"...#..#.", "..##.##.", "..#.....", "....#...", "#.##...#", "####..##", "...##.#.", "#.#.#..."}
	conway(cube, false)
	conway(cube, true)
}

func key(x, y, z int) string {
	return fmt.Sprintf("%v,%v,%v", x, y, z)
}

func key4(x, y, z, w int) string {
	return fmt.Sprintf("%v,%v,%v,%v", x, y, z, w)
}

func toCoords(key string) (x, y, z int) {
	fields := strings.Split(key, ",")
	return toInt(fields[0]), toInt(fields[1]), toInt(fields[2])
}

func toCoords4(key string) (x, y, z, w int) {
	fields := strings.Split(key, ",")
	return toInt(fields[0]), toInt(fields[1]), toInt(fields[2]), toInt(fields[3])
}

func neighbors(space map[string]bool, k string, is4dim bool) (numActive int, coordsInactive []string) {
	if is4dim {
		return neighbors4(space, k)
	} else {
		return neighbors3(space, k)
	}
}

func neighbors3(space map[string]bool, k string) (numActive int, coordsInactive []string) {
	x, y, z := toCoords(k)
	for _, deltaX := range []int{-1, 0, 1} {
		for _, deltaY := range []int{-1, 0, 1} {
			for _, deltaZ := range []int{-1, 0, 1} {
				if deltaX == 0 && deltaY == 0 && deltaZ == 0 {
					continue
				}
				if isActive, _ := space[key(x+deltaX, y+deltaY, z+deltaZ)]; isActive {
					numActive++
				} else {
					coordsInactive = append(coordsInactive, key(x+deltaX, y+deltaY, z+deltaZ))
				}
			}
		}
	}
	return
}

func neighbors4(space map[string]bool, k string) (numActive int, coordsInactive []string) {
	x, y, z, w := toCoords4(k)
	for _, deltaX := range []int{-1, 0, 1} {
		for _, deltaY := range []int{-1, 0, 1} {
			for _, deltaZ := range []int{-1, 0, 1} {
				for _, deltaW := range []int{-1, 0, 1} {
					if deltaX == 0 && deltaY == 0 && deltaZ == 0 && deltaW == 0 {
						continue
					}
					if isActive, _ := space[key4(x+deltaX, y+deltaY, z+deltaZ, w+deltaW)]; isActive {
						numActive++
					} else {
						coordsInactive = append(coordsInactive, key4(x+deltaX, y+deltaY, z+deltaZ, w+deltaW))
					}
				}
			}
		}
	}
	return
}

func changeNode(space map[string]bool, key string, next map[string]bool, inactiveNeighbors []string, is4dim bool) []string {
	isActive, _ := space[key]
	if isActive {
		numActive, inactive := neighbors(space, key, is4dim)
		if numActive == 2 || numActive == 3 {
			next[key] = true
		} else {
			next[key] = false
		}
		inactiveNeighbors = append(inactiveNeighbors, inactive...)
	} else {
		numActive, _ := neighbors(space, key, is4dim)
		if numActive == 3 {
			next[key] = true
		}
	}
	return inactiveNeighbors
}

func conway(cube []string, is4dim bool) {
	space := make(map[string]bool)
	for y, line := range cube {
		for x, val := range line {
			if val == '#' {
				if is4dim {
					space[key4(x, y, 0, 0)] = true
				} else {
					space[key(x, y, 0)] = true
				}
			}
		}
	}
	for i := 1; i <= 6; i++ {
		next := make(map[string]bool)
		// Keep a list of inactive nodes next to active nodes, to check if they need to spring to active
		inactiveNeighbors := make([]string, 0)
		for k := range space {
			inactiveNeighbors = changeNode(space, k, next, inactiveNeighbors, is4dim)
		}
		for _, n := range inactiveNeighbors {
			changeNode(space, n, next, inactiveNeighbors, is4dim)
		}
		space = next
	}
	sum := 0
	for _, v := range space {
		if v {
			sum++
		}
	}
	fmt.Println(sum)
}

func toInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}
