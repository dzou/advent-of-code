package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("adventofcode/day22/day22.txt")

	tokens := strings.Split(string(input), "\r\n\r\n")

	var decks [][]int

	for i := 0; i < len(tokens); i++ {
		currTokens := strings.Split(tokens[i], "\r\n")
		var result []int
		for _, token := range currTokens[1:] {
			num, _ := strconv.Atoi(token)
			result = append(result, num)
		}
		decks = append(decks, result)
	}

	answer := solve(decks[0], decks[1], make(map[string]bool))
	fmt.Println(answer)
}

func thumbprint(a []int, b []int) string {
	result := ""
	for _, num := range a {
		result += strconv.Itoa(num) + ","
	}
	result += "--"
	for _, num := range b {
		result += strconv.Itoa(num) + ","
	}
	return result
}

func solve(aList []int, bList []int, visited map[string]bool) int {
	wtf := false

	for len(aList) > 0 && len(bList) > 0 {
		ta := thumbprint(aList, bList)
		if visited[ta] {
			wtf = true
			break
		}
		visited[ta] = true

		rv := -1
		if aList[0] < len(aList) && bList[0] < len(bList) {
			rv = solve(
				append([]int{}, aList[1:aList[0]+1]...),
				append([]int{}, bList[1:bList[0]+1]...),
				make(map[string]bool))
		}

		if rv == 0 || (rv == -1 && aList[0] > bList[0]) {
			aList = append(aList[1:], aList[0], bList[0])
			bList = append([]int{}, bList[1:]...)
		} else if rv == 1 || (rv == -1 && bList[0] > aList[0]) {
			bList = append(bList[1:], bList[0], aList[0])
			aList = append([]int{}, aList[1:]...)
		}
	}

	var theList []int
	var retVal int
	if wtf || len(aList) > 0 {
		theList = aList
		retVal = 0
	} else {
		theList = bList
		retVal = 1
	}

	fmt.Println(theList)

	sum := 0
	for i, card := range theList {
		sum += (len(theList) - i) * card
	}
	fmt.Printf("Sum: %d\n", sum)

	return retVal
}
