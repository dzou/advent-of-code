package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strconv"
)

type Move struct {
	code   string
	amount int
}

var DIRECTIONS = map[string]image.Point{
	"N": {0, 1},
	"E": {1, 0},
	"S": {0, -1},
	"W": {-1, 0},
}

var DEGREES = map[int]string{
	0:   "N",
	90:  "E",
	180: "S",
	270: "W",
}

func main() {
	file, _ := os.Open("adventofcode/day12/day12.txt")
	defer file.Close()

	var moveList []Move

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		code := string(line[0])

		num, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		moveList = append(moveList, Move{code, num})
	}

	newLoc := solve(moveList)
	fmt.Println(newLoc)
	fmt.Println(abs(newLoc.X) + abs(newLoc.Y))
}

func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}

func rotate(point image.Point, degrees int) image.Point {
	if degrees == 90 {
		return image.Point{point.Y, -point.X}
	} else if degrees == 180 {
		return image.Point{-point.X, -point.Y}
	} else if degrees == 270 {
		return image.Point{-point.Y, point.X}
	} else {
		return point
	}
}

func solve(moveList []Move) image.Point {
	myPos := image.Point{0, 0}
	wayPoint := image.Point{10, 1}

	for _, move := range moveList {
		if move.code == "R" {
			wayPoint = rotate(wayPoint, move.amount)
		} else if move.code == "L" {
			wayPoint = rotate(wayPoint, 360-move.amount)
		} else if move.code == "F" {
			delta := wayPoint.Mul(move.amount)
			myPos = myPos.Add(delta)
		} else {
			delta := DIRECTIONS[move.code].Mul(move.amount)
			wayPoint = wayPoint.Add(delta)
		}
	}

	return myPos
}
