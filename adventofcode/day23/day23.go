package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func main() {
	cupsInput := "389125467"

	var root *Node
	var prev *Node

	for _, r := range cupsInput {
		num := r - '0'

		if root == nil {
			root = &Node{int(num), nil}
			prev = root
		} else {
			tmp := Node{int(num), nil}
			prev.next = &tmp
			prev = &tmp
		}
	}

	//for i := 10; i <= 1000000; i++ {
	//  tmp := Node{i, nil}
	//  prev.next = &tmp
	//  prev = &tmp
	//}
	prev.next = root

	fmt.Println(getSize(root))
	simulate(root)

	printCircle(root)
}

func simulate(root *Node) {
	count := 0
	currNode := root
	size := getSize(root)

	for count < 100 {
		chunkStart := currNode.next
		chunkEnd := currNode
		for i := 0; i < 3; i++ {
			chunkEnd = chunkEnd.next
		}
		currNode.next = chunkEnd.next

		for i := 1; i < size+1; i++ {
			searchVal := (currNode.value - i + (size + 1)) % (size + 1)

			if foundNode := findNum(currNode, searchVal); foundNode != nil {
				tmp := foundNode.next
				foundNode.next = chunkStart
				chunkEnd.next = tmp
				break
			}
		}
		currNode = currNode.next

		printCircle(root)
		count++
	}
}

func findNum(root *Node, num int) *Node {
	if root.value == num {
		return root
	}

	tmp := root.next
	for tmp != root {
		if tmp.value == num {
			return tmp
		}
		tmp = tmp.next
	}

	return nil
}

func getSize(root *Node) int {
	count := 0
	tmp := root

	for tmp != root || count == 0 {
		count++
		tmp = tmp.next
	}
	return count
}

func printCircle(root *Node) {
	count := 0

	for root.value != 1 {
		root = root.next
	}

	tmp := root
	for val := -1; val != root.value; val = tmp.value {
		fmt.Printf("%d,", tmp.value)
		tmp = tmp.next
		count++

		if count > 100 {
			break
		}
	}
	fmt.Println()
}
