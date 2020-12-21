package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	leftRules  []int
	rightRules []int
	prefix     string
}

func main() {
	file, _ := os.Open("adventofcode/day19/day19_rules.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rules = make(map[int]Rule)

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, ":")

		id, _ := strconv.Atoi(tokens[0])

		if strings.Contains(tokens[1], "\"") {
			prefix := strings.TrimSpace(strings.ReplaceAll(tokens[1], "\"", ""))
			rules[id] = Rule{[]int{}, []int{}, prefix}
		} else {
			tokens = strings.Split(strings.TrimSpace(tokens[1]), "|")

			var leftTokens []int
			var rightTokens []int

			for i, chunk := range tokens {
				chunkTokens := strings.Split(strings.TrimSpace(chunk), " ")

				var curr *[]int
				if i == 0 {
					curr = &leftTokens
				} else {
					curr = &rightTokens
				}

				for _, token := range chunkTokens {
					num, _ := strconv.Atoi(token)
					*curr = append(*curr, num)
				}
			}

			rules[id] = Rule{leftTokens, rightTokens, ""}
		}
	}

	file, _ = os.Open("adventofcode/day19/day19.txt")
	defer file.Close()

	count := 0

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		answer := validateStart(line, rules)
		if answer {
			count++
		}
	}

	fmt.Println(count)
}

func validateStart(input string, rules map[int]Rule) bool {
	var aList []int
	var bList []int

	result := false

	for a := 0; a < 5; a++ {
		aList = append(aList, 42)
		for b := 0; b < 5; b++ {
			bList = append(append([]int{42}, bList...), 31)

			tmpRule := Rule{append(aList, bList...), []int{}, ""}

			tmpResult, remainder := validate(input, tmpRule, rules)

			result = result || (tmpResult && len(remainder) == 0)
		}
		bList = nil
	}

	return result
}

func validate(input string, currRule Rule, rules map[int]Rule) (bool, string) {
	if len(input) == 0 {
		return false, ""
	}

	if len(currRule.leftRules) == 0 {
		return input[0:1] == currRule.prefix, input[1:]
	}

	var result bool
	var remainder = input

	for _, otherRuleId := range currRule.leftRules {
		result, remainder = validate(remainder, rules[otherRuleId], rules)

		if !result {
			break
		}
	}

	if !result && len(currRule.rightRules) > 0 {
		remainder = input
		for _, otherRuleId := range currRule.rightRules {
			result, remainder = validate(remainder, rules[otherRuleId], rules)

			if !result {
				break
			}
		}
	}

	return result, remainder
}
