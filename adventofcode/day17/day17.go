package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
	z int
	w int
}

func (self Point) add(other Point) Point {
	return Point{self.x + other.x, self.y + other.y, self.z + other.z, self.w + other.w}
}

var DELTAS = initPoints()

func initPoints() []Point {
	var arr []Point
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					arr = append(arr, Point{x, y, z, w})
				}
			}
		}
	}

	return arr
}

func main() {
	file, _ := os.Open("adventofcode/day17/day17.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	activeCells := make(map[Point]bool)
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		for x, r := range line {
			if r == '#' {
				p := Point{x, y, 0, 0}
				activeCells[p] = true
			}
		}
		y++
	}

	answer := solve(activeCells)
	fmt.Println(answer)
}

func solve(activeCells map[Point]bool) int {
	for iter := 0; iter < 6; iter++ {
		nextCells := make(map[Point]bool)

		for k := range activeCells {
			neighbors := getActiveNeighbors(k, activeCells)
			if neighbors == 2 || neighbors == 3 {
				nextCells[k] = true
			}

			for _, delta := range DELTAS {
				neighbor := k.add(delta)

				if !activeCells[neighbor] {
					count := getActiveNeighbors(neighbor, activeCells)
					if count == 3 {
						nextCells[neighbor] = true
					}
				}
			}
		}
		activeCells = nextCells
	}

	return len(activeCells)
}

func getActiveNeighbors(point Point, activeCells map[Point]bool) int {
	sum := 0

	for _, delta := range DELTAS {
		neighbor := point.add(delta)
		if _, ok := activeCells[neighbor]; ok {
			sum += 1
		}
	}

	return sum
}
