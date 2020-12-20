package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("adventofcode/day20/day20.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// line := scanner.Text()

	}

}

func solve(input string) int {
	//result := 0
	//operand := '+'
	//
	//for _, r := range input {
	//  if r >= '0' && r <= '9' {
	//    num, _ := strconv.Atoi(string(r))
	//    if operand == '+' {
	//
	//    }
	//
	//  }
	//}
	//

	return 0
}
