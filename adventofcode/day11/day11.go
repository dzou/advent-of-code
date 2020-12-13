package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction struct {
	dx int
	dy int
}

var ALL_DIRECTIONS = []Direction{
	Direction{0, 1},
	Direction{0, -1},
	Direction{1, 0},
	Direction{-1, 0},
	Direction{-1, 1},
	Direction{-1, -1},
	Direction{1, -1},
	Direction{1, 1},
}

func main() {
	file, _ := os.Open("adventofcode/day11/day11.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune

	for scanner.Scan() {
		var acc []rune
		line := scanner.Text()
		for _, r := range line {
			acc = append(acc, r)
		}

		grid = append(grid, acc)
	}

	answer := simulate(grid)
	fmt.Println(answer)
}

func simulate(grid [][]rune) int {
	changed := true

	for changed {
		changed = false
		workingGrid := createEmptyCopy(grid)

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				filledNeighbors := countFilledSeenNeighbors(i, j, grid)
				if grid[i][j] == 'L' && filledNeighbors == 0 {
					workingGrid[i][j] = '#'
					changed = true
				} else if grid[i][j] == '#' && filledNeighbors >= 5 {
					workingGrid[i][j] = 'L'
					changed = true
				}
			}
		}
		prettyPrint(workingGrid)
		grid = workingGrid
	}

	sum := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '#' {
				sum++
			}
		}
	}

	return sum
}

func createEmptyCopy(grid [][]rune) [][]rune {
	b := make([][]rune, len(grid))
	for i := range b {
		b[i] = make([]rune, len(grid[0]))
		for j := range b[i] {
			b[i][j] = grid[i][j]
		}
	}
	return b
}

func countFilledNeighbors(row int, col int, grid [][]rune) int {
	count := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if inBounds(i, j, grid) && grid[i][j] == '#' && (i != row || j != col) {
				count++
			}
		}
	}
	return count
}

func countFilledSeenNeighbors(row int, col int, grid [][]rune) int {
	count := 0

	for _, dir := range ALL_DIRECTIONS {
		currRow := row + dir.dx
		currCol := col + dir.dy

		for inBounds(currRow, currCol, grid) && grid[currRow][currCol] == '.' {
			currRow += dir.dx
			currCol += dir.dy
		}

		if inBounds(currRow, currCol, grid) && grid[currRow][currCol] == '#' {
			count++
		}
	}

	return count
}

func inBounds(row int, col int, grid [][]rune) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}

func prettyPrint(grid [][]rune) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
