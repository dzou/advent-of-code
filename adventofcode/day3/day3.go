package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	answer :=
		solve(lines, 1, 1) * solve(lines, 3, 1) * solve(lines, 5, 1) * solve(lines, 7, 1) * solve(lines, 1, 2)
	fmt.Println(answer)
}

func solve(lines []string, right int, down int) int {
	width := len(lines[0])

	count := 0
	for i := down; i < len(lines); i += down {
		if lines[i][((i/down)*right)%width] == '#' {
			count++
		}
	}

	return count
}
