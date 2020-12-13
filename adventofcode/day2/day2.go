package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var problems []PasswordProblem
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		numberTokens := strings.Split(tokens[0], "-")

		lowerBound, _ := strconv.Atoi(string(numberTokens[0]))
		upperBound, _ := strconv.Atoi(string(numberTokens[1]))
		character := tokens[1][0]
		password := tokens[2]

		problem := PasswordProblem{
			lowerBound,
			upperBound,
			string(character),
			string(password),
		}

		problems = append(problems, problem)
	}

	answer := solve(problems)
	fmt.Println(answer)
}

func solve(inputs []PasswordProblem) int {
	count := 0

	for _, problem := range inputs {

		pass := problem.password

		lowerHit := problem.lowerBound-1 < len(problem.password) && string(pass[problem.lowerBound-1]) == problem.letter
		higherHit := problem.upperBound-1 < len(problem.password) && string(pass[problem.upperBound-1]) == problem.letter

		if (lowerHit || higherHit) && !(lowerHit && higherHit) {
			count++
		}
	}

	return count
}

type PasswordProblem struct {
	lowerBound int
	upperBound int
	letter     string
	password   string
}
