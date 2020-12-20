package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("adventofcode/day15/day15.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	numbersText := strings.Split(scanner.Text(), ",")

	var numbers []int
	for _, nString := range numbersText {
		n, _ := strconv.Atoi(nString)
		numbers = append(numbers, n)
	}

	fmt.Println(numbers)

	answer := solve(numbers, 30000000-1)
	fmt.Println(answer)
}

func solve(seed []int, goalIdx int) int {
	prevMap := make(map[int]int)
	for i := 0; i < len(seed)-1; i++ {
		prevMap[seed[i]] = i
	}

	currIdx := len(seed) - 1
	lastNum := seed[len(seed)-1]

	for currIdx < goalIdx {

		next := 0
		if val, ok := prevMap[lastNum]; ok {
			next = currIdx - val
		}

		prevMap[lastNum] = currIdx
		lastNum = next
		currIdx++
	}

	return lastNum
}
