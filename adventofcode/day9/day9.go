package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const GOAL = 22406676

func main() {
	file, err := os.Open("adventofcode/day9/day9.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numberList []int
	for scanner.Scan() {
		line := scanner.Text()

		num, _ := strconv.Atoi(line)
		numberList = append(numberList, num)
	}

	subList := findSum(numberList)
	sort.Ints(subList)
	fmt.Println(subList[0] + subList[len(subList)-1])
}

func solve(numberList []int) int {

	for i := range numberList {
		slice := numberList[i : i+26]

		if !validSlice(slice) {
			return numberList[i+25]
		}
	}

	return -1
}

func findSum(numberList []int) []int {
	i := 0
	j := 0
	sum := 0

	for sum != GOAL {
		if sum < GOAL {
			sum += numberList[j]
			j++
		} else if sum > GOAL {
			sum -= numberList[i]
			i++
		}
	}

	return numberList[i:j]
}

func validSlice(numberSlice []int) bool {
	val := numberSlice[25]
	for i := 0; i < 25; i++ {
		for j := i + 1; j < 25; j++ {
			if numberSlice[i]+numberSlice[j] == val {
				return true
			}
		}
	}

	return false
}
