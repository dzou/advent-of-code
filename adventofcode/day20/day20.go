package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var MONSTER = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
}

var DELTAS = []Point{
	Point{-1, 0},
	Point{1, 0},
	Point{0, -1},
	Point{0, 1},
}

type Point struct {
	x int
	y int
}

type Tile struct {
	id   int
	grid [][]rune
}

func main() {
	input, _ := ioutil.ReadFile("adventofcode/day20/day20.txt")
	tileTokens := strings.Split(strings.TrimSpace(string(input)), "\r\n\r\n")

	var tiles []Tile

	for _, tileToken := range tileTokens {
		lines := strings.Split(strings.TrimSpace(tileToken), "\r\n")

		id, _ := strconv.Atoi(lines[0][5 : len(lines[0])-1])

		grid := make([][]rune, len(lines)-1)
		for i := 1; i < len(lines); i++ {
			grid[i-1] = make([]rune, len(lines[i]))
			for j := 0; j < len(lines[i]); j++ {
				grid[i-1][j] = rune(lines[i][j])
			}
		}

		tiles = append(tiles, Tile{id, grid})
	}

	answer := solve(tiles)

	corner := len(answer) - 1
	fmt.Println(answer[0][0].id * answer[corner][0].id * answer[0][corner].id * answer[corner][corner].id)

	image := removeBorders(answer)
	prettyPrint(image)

	var monsterMask [][]rune
	for _, line := range MONSTER {
		monsterMask = append(monsterMask, []rune(line))
	}
	prettyPrint(monsterMask)

	for i := 0; i < 4; i++ {
		answer := countMonsters(rotateTimes90(image, i), monsterMask)
		fmt.Println(answer)
	}

	image = flip(image)
	for i := 0; i < 4; i++ {
		answer := countMonsters(rotateTimes90(image, i), monsterMask)
		fmt.Println(answer)
	}
}

func countMonsters(grid [][]rune, monsterMask [][]rune) int {
	hashCount := 0
	monsterCount := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' {
				hashCount++
			}

			if matchAndReplace(grid, monsterMask, i, j) {
				monsterCount++
			}
		}
	}

	return hashCount - 15*monsterCount
}

func matchAndReplace(grid [][]rune, monsterMask [][]rune, i int, j int) bool {
	height := len(monsterMask)
	width := len(monsterMask[0])

	if i+height > len(grid) || j+width > len(grid[0]) {
		return false
	}

	for mi := 0; mi < len(monsterMask); mi++ {
		for mj := 0; mj < len(monsterMask[mi]); mj++ {
			di := i + mi
			dj := j + mj

			if monsterMask[mi][mj] == '#' && grid[di][dj] != '#' {
				return false
			}
		}
	}

	return true
}

func removeBorders(answer [][]Tile) [][]rune {
	for _, row := range answer {
		for ti, tile := range row {
			for i := 0; i < len(tile.grid); i++ {
				row[ti].grid[i] = tile.grid[i][1 : len(tile.grid[i])-1]
			}
			row[ti].grid = tile.grid[1 : len(tile.grid)-1]
		}
	}

	gridSize := len(answer[0][0].grid)
	totalSize := len(answer) * len(answer[0][0].grid)

	result := make([][]rune, totalSize)
	for i := 0; i < totalSize; i++ {
		result[i] = make([]rune, totalSize)
	}

	for i := 0; i < totalSize; i++ {
		for j := 0; j < totalSize; j++ {
			tileRow := i / gridSize
			tileCol := j / gridSize

			gridRow := i % gridSize
			gridCol := j % gridSize

			result[i][j] = answer[tileRow][tileCol].grid[gridRow][gridCol]
		}
	}

	return result
}

func solve(tiles []Tile) [][]Tile {
	size := int(math.Sqrt(float64(len(tiles))))

	placements := make([][]Tile, size)
	for i := 0; i < size; i++ {
		placements[i] = make([]Tile, size)
	}

	usedTiles := make(map[int]bool)

	return solveHelper(placements, tiles, usedTiles)
}

func solveHelper(placements [][]Tile, tiles []Tile, usedTiles map[int]bool) [][]Tile {
	if len(usedTiles) == len(tiles) {
		return placements
	}

	var i = len(usedTiles) / len(placements)
	var j = len(usedTiles) % len(placements)

	for _, tile := range tiles {
		if usedTiles[tile.id] {
			continue
		}

		for _, orientation := range getTileOrientations(tile) {
			placements[i][j] = orientation
			usedTiles[orientation.id] = true

			if isValidPlacement(i, j, placements) {
				result := solveHelper(placements, tiles, usedTiles)
				if result != nil {
					return result
				}
			}

			delete(usedTiles, orientation.id)
			placements[i][j].id = 0
		}
	}

	return nil
}

func isValidPlacement(i int, j int, placements [][]Tile) bool {
	neighbors := getNeighbors(i, j, placements)

	for dir, neighbor := range neighbors {
		if neighbor.id == 0 {
			continue
		}

		var mine [][]rune
		var other [][]rune

		if dir == "N" {
			mine = placements[i][j].grid
			other = rotateTimes90(neighbor.grid, 2)
		} else if dir == "S" {
			mine = rotateTimes90(placements[i][j].grid, 2)
			other = rotateTimes90(neighbor.grid, 0)
		} else if dir == "E" {
			mine = rotateTimes90(placements[i][j].grid, 1)
			other = rotateTimes90(neighbor.grid, 3)
		} else if dir == "W" {
			mine = rotateTimes90(placements[i][j].grid, 3)
			other = rotateTimes90(neighbor.grid, 1)
		}

		if !isFit(mine, other) {
			return false
		}
	}

	return true
}

func isFit(myGrid [][]rune, otherGrid [][]rune) bool {
	for i := 0; i < len(myGrid); i++ {
		if myGrid[0][i] != otherGrid[0][len(myGrid)-i-1] {
			return false
		}
	}

	return true
}

func getNeighbors(i int, j int, placements [][]Tile) map[string]Tile {
	result := make(map[string]Tile)

	for _, delta := range DELTAS {
		nextX := i + delta.x
		nextY := j + delta.y

		if nextX >= 0 && nextX < len(placements) && nextY >= 0 && nextY < len(placements) {
			if delta.x == -1 {
				result["N"] = placements[nextX][nextY]
			} else if delta.x == 1 {
				result["S"] = placements[nextX][nextY]
			} else if delta.y == -1 {
				result["W"] = placements[nextX][nextY]
			} else if delta.y == 1 {
				result["E"] = placements[nextX][nextY]
			}
		}
	}

	return result
}

func getTileOrientations(root Tile) []Tile {
	var result []Tile

	curr := root.grid
	for i := 0; i < 4; i++ {
		result = append(result, Tile{root.id, rotateTimes90(curr, i)})
	}

	curr = flip(curr)
	for i := 0; i < 4; i++ {
		result = append(result, Tile{root.id, rotateTimes90(curr, i)})
	}

	return result
}

func flip(grid [][]rune) [][]rune {
	next := make([][]rune, len(grid))

	for i := 0; i < len(grid); i++ {
		next[i] = make([]rune, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			next[i][j] = grid[len(grid[i])-1-i][j]
		}
	}
	return next
}

func rotateTimes90(grid [][]rune, times int) [][]rune {
	result := grid
	for i := 0; i < times; i++ {
		result = rotate90(result)
	}
	return result
}

func rotate90(grid [][]rune) [][]rune {
	next := make([][]rune, len(grid))

	for i := 0; i < len(grid); i++ {
		next[i] = make([]rune, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			next[i][j] = grid[j][len(grid[i])-i-1]
		}
	}
	return next
}

func prettyPrint(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
