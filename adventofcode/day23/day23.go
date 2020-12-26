package main

import "fmt"

type Node struct {
	value int
	prev  *Node
	next  *Node
}

func main() {
	cupsInput := "219347865"

	var root *Node
	var prev *Node

	cache := make(map[int]*Node)

	for _, r := range cupsInput {
		num := r - '0'

		if root == nil {
			root = &Node{int(num), nil, nil}
			prev = root
		} else {
			tmp := Node{int(num), prev, nil}
			prev.next = &tmp
			prev = &tmp
		}
		cache[int(num)] = prev
	}

	for i := 10; i <= 1000000; i++ {
		tmp := Node{i, prev, nil}
		prev.next = &tmp
		prev = &tmp
		cache[i] = prev
	}

	prev.next = root
	root.prev = prev

	fmt.Println(getSize(root))
	simulate(root, cache)

	printCircle(root)
}

func simulate(root *Node, cache map[int]*Node) {
	count := 0
	currNode := root
	size := getSize(root)

	for count < 10000000 {
		var chunkNums = map[int]bool{0: true}

		chunkStart := currNode.next
		chunkEnd := currNode
		for i := 0; i < 3; i++ {
			chunkEnd = chunkEnd.next
			chunkNums[chunkEnd.value] = true
		}

		currNode.next = chunkEnd.next
		chunkEnd.next.prev = currNode

		for i := 1; i < size+1; i++ {
			searchVal := (currNode.value - i + (size + 1)) % (size + 1)
			if chunkNums[searchVal] {
				continue
			}

			foundNode := cache[searchVal]

			tmp := foundNode.next
			foundNode.next = chunkStart
			chunkStart.prev = foundNode

			chunkEnd.next = tmp
			tmp.prev = chunkEnd
			break
		}
		currNode = currNode.next

		count++
	}
}

func findNum(root *Node, num int) *Node {
	count := 0
	if root.value == num {
		return root
	}

	tmp := root.next
	for tmp != root {
		count += 1
		if tmp.value == num {
			fmt.Println(count)
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
