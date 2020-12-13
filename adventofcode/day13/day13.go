package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("adventofcode/day13/day13.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	departTime, _ := strconv.Atoi(scanner.Text())

	var schedule []int
	scanner.Scan()
	scheduleTokens := strings.Split(scanner.Text(), ",")
	for _, time := range scheduleTokens {
		if time != "x" {
			num, _ := strconv.Atoi(time)
			schedule = append(schedule, num)
		} else {
			schedule = append(schedule, -1)
		}
	}

	scheduleMap := make(map[int]int64)
	for idx, time := range schedule {
		if time != -1 {
			//if idx > time {
			//  idx = idx % time
			//}
			scheduleMap[idx] = int64(time)
		}
	}

	fmt.Println(departTime)
	fmt.Println(scheduleMap)

	answer := solve(scheduleMap)
	fmt.Println(answer)
}

func solve(scheduleMap map[int]int64) int64 {
	increment := int64(1)
	time := int64(0)
	solved := make(map[int]bool)

	checks := 0
	for len(solved) < len(scheduleMap) {
		time += increment

		for desiredTime, busTime := range scheduleMap {
			if !solved[desiredTime] && (busTime-(time%busTime))%busTime == int64(desiredTime)%busTime {
				solved[desiredTime] = true
				increment *= busTime
				fmt.Printf("Solved: %d, %d at %d\n", desiredTime, busTime, time)
			}
		}

		checks++
		if checks%1000000 == 0 {
			fmt.Printf("checked %d\n", checks)
			fmt.Println(time)
		}
	}

	return time
}

func validate(currTime int, scheduleMap map[int]int) bool {
	return true
}

func getEarliestTime(departTime int, schedule []int) (int, int) {
	earliestWait := 999999999
	earliestTime := -1

	for _, time := range schedule {
		timeWaited := time - (departTime % time)

		if timeWaited < earliestWait {
			earliestWait = timeWaited
			earliestTime = time
		}
	}

	return earliestTime, earliestWait
}
