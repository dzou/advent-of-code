package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Hexagon struct {
	x int
	y int
	z int
}

func main() {
	file, _ := os.Open("adventofcode/day24/day24.txt")
	defer file.Close()

	tracker := make(map[Hexagon]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		hexagon := parseHexagon(line)
		fmt.Println(hexagon)

		if tracker[hexagon] {
			delete(tracker, hexagon)
		} else {
			tracker[hexagon] = true
		}
	}

	fmt.Println(len(tracker))
	solve(tracker)
}

func solve(tiles map[Hexagon]bool) int {
	count := 0

	for count < 100 {
		result := make(map[Hexagon]bool)
		for tile := range tiles {
			num := countFlippedNeighbors(tile, tiles)
			if num >= 1 && num <= 2 {
				result[tile] = true
			}

			for _, neighbor := range getNeighbors(tile) {
				num = countFlippedNeighbors(neighbor, tiles)
				if num == 2 {
					result[neighbor] = true
				}
			}
		}

		count++
		tiles = result
		fmt.Println("%d: %d", count, len(result))
	}

	return len(tiles)
}

func countFlippedNeighbors(center Hexagon, tiles map[Hexagon]bool) int {
	count := 0

	for _, neighbor := range getNeighbors(center) {
		if tiles[neighbor] {
			count++
		}
	}
	return count
}

func getNeighbors(hexagon Hexagon) []Hexagon {
	deltas := []Hexagon{
		Hexagon{0, -1, 1},
		Hexagon{0, 1, -1},
		Hexagon{-1, 0, 1},
		Hexagon{1, 0, -1},
		Hexagon{-1, 1, 0},
		Hexagon{1, -1, 0},
	}

	var result []Hexagon

	for _, d := range deltas {
		result = append(result, Hexagon{hexagon.x + d.x, hexagon.y + d.y, hexagon.z + d.z})
	}
	return result
}

func parseHexagon(input string) Hexagon {
	x := 0
	y := 0
	z := 0

	for len(input) > 0 {
		if strings.HasPrefix(input, "e") {
			x++
			y--
			input = input[1:]
		} else if strings.HasPrefix(input, "w") {
			x--
			y++
			input = input[1:]
		} else {
			prefix := input[0:2]
			if prefix == "ne" {
				x++
				z--
			} else if prefix == "se" {
				y--
				z++
			} else if prefix == "nw" {
				y++
				z--
			} else if prefix == "sw" {
				x--
				z++
			}
			input = input[2:]
		}
	}

	return Hexagon{x, y, z}
}
