package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Constraint struct {
	title  string
	ranges []int
}

func main() {
	input, _ := ioutil.ReadFile("adventofcode/day16/day16.txt")
	sections := strings.Split(strings.TrimSpace(string(input)), "\r\n\r\n")

	scanner := bufio.NewScanner(strings.NewReader(sections[0]))
	var constraints []Constraint
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "-") {
			var title string
			ranges := make([]int, 4)

			fmt.Sscanf(
				line, "%s %d-%d or %d-%d",
				&title,
				&ranges[0], &ranges[1], &ranges[2], &ranges[3])

			constraint := Constraint{title, ranges}
			constraints = append(constraints, constraint)
		}

	}
	fmt.Println(constraints)

	otherTicketTokens := strings.Split(sections[2], "\r\n")

	var tickets [][]int

	for i := 1; i < len(otherTicketTokens); i++ {
		numberTokens := strings.Split(otherTicketTokens[i], ",")
		ticket := make([]int, len(numberTokens))
		for x, ns := range numberTokens {
			ticket[x], _ = strconv.Atoi(ns)
		}
		tickets = append(tickets, ticket)
	}
	fmt.Println(tickets)

	answer := solve(tickets, constraints)
	for _, x := range answer {
		fmt.Println(x)
	}

	constraintOrder := make([]Constraint, len(constraints))

	for i := 0; i < len(answer); i++ {
		var constraint Constraint

		for x, constraintList := range answer {
			if len(constraintList) == 1 {
				constraint = constraintList[0]
				constraintOrder[x] = constraint
				break
			}
		}

		for x := 0; x < len(answer); x++ {
			idx := -1
			for tmp := 0; tmp < len(answer[x]); tmp++ {
				if answer[x][tmp].title == constraint.title {
					idx = tmp
					break
				}
			}

			if idx > -1 {
				answer[x] = append(answer[x][:idx], answer[x][idx+1:]...)
			}
		}
	}

	fmt.Println(constraintOrder)

	myTicket := []int{139, 67, 71, 59, 149, 89, 101, 83, 107, 103, 79, 157, 151, 113, 61, 109, 73, 97, 137, 53}

	product := 1

	for i := 0; i < len(constraintOrder); i++ {
		constraint := constraintOrder[i]
		if strings.HasPrefix(constraint.title, "departure") {
			product *= myTicket[i]
		}
	}

	fmt.Println(product)

}

func (self Constraint) Matches(ticketValue int) bool {
	if ticketValue >= self.ranges[0] && ticketValue <= self.ranges[1] {
		return true
	}
	if ticketValue >= self.ranges[2] && ticketValue <= self.ranges[3] {
		return true
	}
	return false
}

func solve(tickets [][]int, constraints []Constraint) [][]Constraint {
	ticketSchema := make([][]Constraint, len(constraints))
	for i := 0; i < len(constraints); i++ {
		ticketSchema[i] = append(ticketSchema[i], constraints...)
	}

	for _, ticket := range tickets {
		if !isValidTicket(ticket, constraints) {
			continue
		}

		for i, ticketValue := range ticket {
			var validConstraints []Constraint
			for _, constraint := range ticketSchema[i] {
				if constraint.Matches(ticketValue) {
					validConstraints = append(validConstraints, constraint)
				}
			}
			ticketSchema[i] = validConstraints
		}
	}

	return ticketSchema
}

func isValidTicket(ticket []int, constraints []Constraint) bool {
	for _, num := range ticket {
		if !isValid(num, constraints) {
			return false
		}
	}
	return true
}

func isValid(number int, constraints []Constraint) bool {
	for _, constraint := range constraints {
		if number >= constraint.ranges[0] && number <= constraint.ranges[1] {
			return true
		}

		if number >= constraint.ranges[2] && number <= constraint.ranges[3] {
			return true
		}
	}

	return false
}
