package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	arrival := 1003240
	busses := "19,x,x,x,x,x,x,x,x,41,x,x,x,37,x,x,x,x,x,787,x,x,x,x,x,x,x,x,x,x,x,x,13,x,x,x,x,x,x,x,x,x,23,x,x,x,x,x,29,x,571,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,17"
	earliestDepart := 99999999
	earliestBus := 0
	for _, bus := range strings.Split(busses, ",") {
		if bus != "x" {
			busId := toInt(bus)
			time := busId
			for ; time < arrival; time = time + busId {
			}
			if time < earliestDepart {
				earliestDepart = time
				earliestBus = busId
			}
		}
	}
	fmt.Println((earliestDepart - arrival) * earliestBus)
}

type schedule struct {
	bus    int
	offset int
}

func partTwo() {
	busses := "19,x,x,x,x,x,x,x,x,41,x,x,x,37,x,x,x,x,x,787,x,x,x,x,x,x,x,x,x,x,x,x,13,x,x,x,x,x,x,x,x,x,23,x,x,x,x,x,29,x,571,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,17"
	var schedules []schedule
	for i, bus := range strings.Split(busses, ",") {
		if bus != "x" {
			schedules = append(schedules, schedule{toInt(bus), i})
		}
	}
	largestIncrement := 1
	for i := 0; ; {
		allLeaving := true
		for _, s := range schedules {
			if (i+s.offset)%(s.bus) != 0 {
				allLeaving = false
				break
			} else if s.bus > largestIncrement {
				largestIncrement = s.bus
			}
		}
		if allLeaving {
			fmt.Println("All busses depart at ", i)
			return
		}
		i += largestIncrement
	}
}

func toInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}
