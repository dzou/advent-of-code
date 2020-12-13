package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("adventofcode/day5/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var boardingPasses []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.ReplaceAll(line, "F", "0")
		line = strings.ReplaceAll(line, "B", "1")
		line = strings.ReplaceAll(line, "R", "1")
		line = strings.ReplaceAll(line, "L", "0")

		num, _ := strconv.ParseInt(line, 2, 0)
		boardingPasses = append(boardingPasses, int(num))
	}

	answer := solve(boardingPasses)
	fmt.Println(answer)

}

func solve(passes []int) int {
	sort.Ints(passes)
	fmt.Println(passes)

	for i := 0; i < len(passes)-1; i++ {
		if passes[i+1]-passes[i] > 1 {
			fmt.Println(passes[i])
			fmt.Println(passes[i+1])
			break
		}
	}

	return 0
}
