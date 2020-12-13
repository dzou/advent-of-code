package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	name string
	val  int
}

func main() {
	file, err := os.Open("adventofcode/day8/day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var operations []Operation

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		opName := tokens[0]
		opVal, _ := strconv.Atoi(tokens[1])

		operations = append(operations, Operation{opName, opVal})
	}

	for i := range operations {
		op := &operations[i]
		if op.name == "jmp" {
			op.name = "nop"
		} else if op.name == "nop" {
			op.name = "jmp"
		}

		fmt.Println(operations)
		result := solve(operations)
		if result {
			break
		}

		if op.name == "jmp" {
			op.name = "nop"
		} else if op.name == "nop" {
			op.name = "jmp"
		}
	}
}

func solve(opList []Operation) bool {
	acc := 0
	visited := make(map[int]bool)

	for i := 0; i < len(opList); i++ {
		op := opList[i]
		if visited[i] {
			return false
		}

		visited[i] = true
		if op.name == "acc" {
			acc += op.val
		} else if op.name == "jmp" {
			i += op.val - 1
		}
	}

	fmt.Println(acc)
	return true
}
