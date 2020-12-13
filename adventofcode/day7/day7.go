package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BagNode struct {
	bagName string
	count   int
}

var exp = regexp.MustCompile(``)

func main() {
	file, err := os.Open("adventofcode/day7/day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bagGraph := map[string][]BagNode{}

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, "contain")
		fromBag := getFromKey(strings.TrimSpace(tokens[0]))
		toBags := getToBags(strings.TrimSpace(tokens[1]))

		bagGraph[fromBag] = toBags
	}

	answer := solve(bagGraph)
	fmt.Println(answer)
}

func solve(graph map[string][]BagNode) int {
	//values := map[string]bool{}
	//
	//for _, v := range graph {
	//  for _, bagNode := range v {
	//    values[bagNode.bagName] = true
	//  }
	//}
	//
	//var roots []string
	//for k, _ := range graph {
	//  if !values[k] {
	//    fmt.Println("found root: " + k)
	//    roots = append(roots, k)
	//  }
	//}
	//
	//bagHolders := map[string]bool{"shiny gold": true}
	//
	//for _, k := range roots {
	//  countNodes(k, graph, bagHolders)
	//}

	// return len(bagHolders) - 1
	return sumNodes("shiny gold", graph) - 1
}

func sumNodes(root string, graph map[string][]BagNode) int {
	count := 1

	for _, node := range graph[root] {
		count += node.count * sumNodes(node.bagName, graph)
	}

	return count
}
func countNodes(root string, graph map[string][]BagNode, visited map[string]bool) int {
	if visited[root] {
		return 1
	}

	count := 0

	for _, node := range graph[root] {
		tmp := countNodes(node.bagName, graph, visited)
		if tmp > 0 {
			visited[root] = true
		}
		count += tmp
	}

	return count
}

func getFromKey(input string) string {
	return strings.TrimSuffix(input, " bags")
}

func getToBags(input string) []BagNode {
	var result []BagNode
	if input == "no other bags" {
		return result
	}

	tokens := strings.Split(input, ",")

	for _, token := range tokens {
		token = strings.TrimSpace(token)
		num, _ := strconv.Atoi(string(token[0]))
		bagType := token[2 : strings.Index(token, "bag")-1]
		result = append(result, BagNode{bagType, num})
	}

	return result
}
