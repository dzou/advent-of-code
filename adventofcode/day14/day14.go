package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Assignment struct {
	key   int64
	value int64
}

type OperationSet struct {
	clearMask   int64
	setMasks    []int64
	assignments []Assignment
}

func main() {
	file, _ := os.Open("adventofcode/day14/day14.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var opSetList []OperationSet
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "mask") {
			clearMask, setMasks := loadMask(line[7:])
			opSet := OperationSet{clearMask, setMasks, []Assignment{}}
			opSetList = append(opSetList, opSet)

		} else if strings.HasPrefix(line, "mem") {
			var key, val int64
			fmt.Sscanf(line, "mem[%d] = %d", &key, &val)
			currOp := &opSetList[len(opSetList)-1]
			currOp.assignments = append(currOp.assignments, Assignment{key, val})
			fmt.Println(currOp.assignments)
		}
	}

	fmt.Println(opSetList)

	answer := solve(opSetList)
	fmt.Println(answer)

}

func solve(opslist []OperationSet) int64 {
	memory := make(map[int64]int64)

	for _, op := range opslist {
		for _, assignment := range op.assignments {
			idxStart := assignment.key & op.clearMask
			for _, m := range op.setMasks {
				idx := idxStart | m
				memory[idx] = assignment.value
			}
		}
	}

	sum := int64(0)
	for _, v := range memory {
		sum += v
	}

	fmt.Println(memory)
	return sum
}

func loadMask(mask string) (int64, []int64) {
	clearMask := int64(0)
	setMasks := genMask(mask, 0)

	for _, r := range mask {
		clearMask <<= 1

		if r == 'X' {
			clearMask |= 0
		} else if r == '0' {
			clearMask |= 1
		}
	}

	return clearMask, setMasks
}

func genMask(input string, acc int64) []int64 {
	if len(input) == 0 {
		return []int64{acc}
	}

	var results []int64

	acc <<= 1
	if input[0] == 'X' {
		results = append(results, genMask(input[1:], acc|1)...)
	} else if input[0] == '1' {
		acc |= 1
	}

	results = append(results, genMask(input[1:], acc)...)
	return results
}
