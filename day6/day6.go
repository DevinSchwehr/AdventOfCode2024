package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Execute() {
	os.Chdir("day6")

	partOneResult, partOneGuard := partOne("input.txt")

	fmt.Printf("Result for Part One is: %d \n", partOneResult)
	fmt.Printf("Result for Part Two is: %d \n", partTwo("input.txt", partOneGuard))
}

func createGrid(filename string) ([][]string, int, int) {
	var grid [][]string
	guardX, guardY := 0, 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineArray := strings.Split(scanner.Text(), "")
		hasStart := false
		for index, char := range lineArray {
			if char == "^" {
				guardX = index
				hasStart = true
			}
		}
		grid = append(grid, lineArray)
		if hasStart {
			guardY = len(grid) - 1
		}
	}
	return grid, guardX, guardY
}

func partOne(filename string) (int, *Guard) {

	grid, guardX, guardY := createGrid(filename)
	result := 0
	guard := newGuard(guardX, guardY)

	for {
		nextX, nextY := guard.getNext()
		if (nextX < 0 || nextX >= len(grid[guard.y])) || (nextY < 0 || nextY >= len(grid)) {
			break
		}
		if grid[nextY][nextX] == "#" {
			guard.changeDirection()
		} else {
			guard.moveNext()
			if grid[guard.y][guard.x] != "X" {
				guard.addToPath(guard.x, guard.y, guard.direction)
				grid[guard.y][guard.x] = "X"
				result++
			}
		}
	}

	return result, guard
}

func partTwo(filename string, guard *Guard) int {

	grid, guardX, guardY := createGrid(filename)
	guard.set(guardX, guardY, "up")
	validObstacles := 0

	//Go along part one guard's traveled path and try to put obstacle on each spot
	for _, step := range guard.traveledPath {
		//Don't try to put obstacle on starting position
		if step.x == guard.x && step.y == guard.y {
			continue
		}
		grid[step.y][step.x] = "#"
		copyGuard := newGuard(guard.x, guard.y)
		copyGuard.direction = guard.direction
		if checkLooped(copyGuard, grid) {
			validObstacles++
		}
		grid[step.y][step.x] = "."
	}

	return validObstacles
}

func checkLooped(guard *Guard, grid [][]string) bool {
	var traveledPath []visitedCell
	for {
		nextX, nextY := guard.getNext()
		if (nextX < 0 || nextX >= len(grid[guard.y])) || (nextY < 0 || nextY >= len(grid)) {
			break
		}
		if grid[nextY][nextX] == "#" {
			guard.changeDirection()
			continue
		}

		//If on already traveled path
		for _, value := range traveledPath {
			if value.x == guard.x && value.y == guard.y {
				if !value.checkDirections(guard.direction) {
					value.directions = append(value.directions, guard.direction)
				} else {
					return true
				}
			}
		}
		traveledPath = append(traveledPath, *newPosAndDirection(guard.x, guard.y, guard.direction))
		guard.moveNext()
	}
	return false
}
