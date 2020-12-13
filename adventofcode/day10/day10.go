package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("adventofcode/day10/day10.txt")
	defer file.Close()

	numbers := []int{0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}

	sort.Ints(numbers)
	numbers = append(numbers, numbers[len(numbers)-1]+3)

	cache := make([]int, len(numbers))
	for i := 0; i < len(cache); i++ {
		cache[i] = -1
	}

	answer := count(0, numbers, cache)
	fmt.Println(answer)
}

func solve(adapters []int) (int, int) {
	oneDiff := 0
	threeDiff := 0

	for i := 0; i < len(adapters)-1; i++ {
		diff := adapters[i+1] - adapters[i]

		if diff == 1 {
			oneDiff++
		} else if diff == 3 {
			threeDiff++
		}
	}

	return oneDiff, threeDiff
}

func count(idx int, adapters []int, cache []int) int {
	if idx == len(adapters)-1 {
		return 1
	}

	if cache[idx] != -1 {
		return cache[idx]
	}

	sum := 0
	currAdapter := adapters[idx]
	for i := idx + 1; i < len(adapters); i++ {
		if adapters[i]-currAdapter <= 3 {
			sum += count(i, adapters, cache)
		}
	}

	cache[idx] = sum
	return sum
}
