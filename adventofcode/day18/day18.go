package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("adventofcode/day18/day18.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var results []int

	for scanner.Scan() {
		line := scanner.Text()
		answer, _ := solve(line)
		results = append(results, answer)
	}

	sum := 0
	for _, n := range results {
		sum += n
	}

	fmt.Println(results)
	fmt.Println(sum)
}

func solve(input string) (int, int) {
	result := 0
	operand := '+'

	idx := 0
	for idx < len(input) {
		r := input[idx]

		if r >= '0' && r <= '9' {
			num, _ := strconv.Atoi(string(r))
			if operand == '+' {
				result += num
			} else if operand == '*' {
				result *= num
			}

		} else if r == '(' {
			tmp, tmpIdx := solve(input[idx+1:])

			if operand == '+' {
				result += tmp
			} else if operand == '*' {
				result *= tmp
			}
			idx += tmpIdx + 1
		} else if r == ')' {
			break
		} else if r == '+' || r == '*' {
			operand = rune(r)
			if r == '*' {
				tmp, tmpIdx := solve(input[idx+1:])
				result *= tmp
				idx += tmpIdx
			}
		}

		idx++
	}

	return result, idx
}
func solve2(input string) (int, int) {
	result := 0
	operand := '+'

	idx := 0
	for idx < len(input) {
		r := input[idx]

		if r >= '0' && r <= '9' {
			num, _ := strconv.Atoi(string(r))
			if operand == '+' {
				result += num
			} else if operand == '*' {
				result *= num
			}

		} else if r == '(' {
			tmp, tmpIdx := solve(input[idx+1:])

			if operand == '+' {
				result += tmp
			} else if operand == '*' {
				result *= tmp
			}
			idx += tmpIdx + 1
		} else if r == ')' {
			break
		} else if r == '+' || r == '*' {
			operand = rune(r)
		}
		idx++
	}

	return result, idx
}
