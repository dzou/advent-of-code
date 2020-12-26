package main

import "fmt"

func main() {
	cardPub := 6270530
	doorPub := 14540258

	cardLoop := findLoopSize(cardPub)
	doorLoop := findLoopSize(doorPub)

	fmt.Println(cardLoop)
	fmt.Println(doorLoop)

	result := transform(cardPub, doorLoop)
	fmt.Println(result)
}

func transform(subject int, iters int) int {
	result := 1
	for i := 0; i < iters; i++ {
		result = (result * subject) % 20201227
	}
	return result
}

func findLoopSize(pubkey int) int {
	sub := 1
	for i := 1; i < 100000000000; i++ {
		sub = (sub * 7) % 20201227

		if sub == pubkey {
			return i
		}
	}

	return -1
}
