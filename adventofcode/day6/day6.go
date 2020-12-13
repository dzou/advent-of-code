package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("adventofcode/day6/day6.txt")

	sum := 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\r\n\r\n") {
		sum += getCount(s)
	}

	fmt.Println(sum)
}

func getCount(input string) int {
	var sets []map[rune]bool

	for _, s := range strings.Split(strings.TrimSpace(input), "\r\n") {
		set := map[rune]bool{}

		for _, char := range s {
			set[char] = true
		}

		sets = append(sets, set)
	}

	masterSet := sets[0]

	for _, set := range sets {
		for k := range masterSet {
			if !set[k] {
				masterSet[k] = false
			}
		}
	}

	count := 0
	for _, v := range masterSet {
		if v {
			count++
		}
	}

	return count
}
