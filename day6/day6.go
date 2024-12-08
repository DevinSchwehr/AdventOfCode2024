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

	fmt.Printf("Result for Part One is: %d \n", partOne("input.txt"))
	fmt.Printf("Result for Part Two is: %d \n", partTwo("input.txt"))
}

func partOne(filename string) int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0

	var grid [][]string

	guardX, guardY := 0, 0
	direction := "up"

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

	exitedX := func() bool { return guardX <= 0 || guardX >= len(grid[0])-1 }
	exitedY := func() bool { return guardY <= 0 || guardY >= len(grid)-1 }

	for !exitedX() && !exitedY() {
		nextX, nextY :=
		switch direction {
		case "up":
			checkNext := grid[guardY-1][guardX]
			if checkNext == "#" {
				direction = "right"
			} else {
				guardY--
				if grid[guardY][guardX] != "X" {
					grid[guardY][guardX] = "X"
					result++
				}
			}
		case "right":
			checkNext := grid[guardY][guardX+1]
			if checkNext == "#" {
				direction = "down"
			} else {
				guardX++
				if grid[guardY][guardX] != "X" {
					grid[guardY][guardX] = "X"
					result++
				}
			}
		case "down":
			checkNext := grid[guardY+1][guardX]
			if checkNext == "#" {
				direction = "left"
			} else {
				guardY++
				if grid[guardY][guardX] != "X" {
					grid[guardY][guardX] = "X"
					result++
				}
			}
		case "left":
			checkNext := grid[guardY][guardX-1]
			if checkNext == "#" {
				direction = "up"
			} else {
				guardX--
				if grid[guardY][guardX] != "X" {
					grid[guardY][guardX] = "X"
					result++
				}
			}
		}
	}

	return result
}

type ObstacleCoordinates struct {
	x int
	y int
}

type Guard struct {
	x         int
	y         int
	direction string
}

func (g Guard) getNext() (int, int) {
	switch g.direction {
	case "up":
		return g.x, g.y - 1
	case "left":
		return g.x - 1, g.y
	case "right":
		return g.x + 1, g.y
	case "down":
		return g.x, g.y + 1
	}
	return g.x, g.y
}

func (g Guard) changeDirection() {
	switch g.direction {
	case "up":
		g.direction = "right"
	case "right":
		g.direction = "down"
	case "down":
		g.direction = "left"
	case "left":
		g.direction = "up"
	}
}

func newGuard(x int, y int) *Guard {
	guard := new(Guard)
	guard.x = x
	guard.y = y
	guard.direction = "up"
	return guard
}

func newBlockCoordinates(x int, y int) *ObstacleCoordinates {
	coords := new(ObstacleCoordinates)
	coords.x = x
	coords.y = y
	return coords
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

// func PartTwoV2(filename string) int {

// 	grid, guardX, guardY := createGrid(filename)

// 	guard := newGuard(guardX, guardY)
// 	var usedBlockLocations []ObstacleCoordinates
// 	exitedX := func() bool { return guard.x <= 0 || guard.x >= len(grid[0])-1 }
// 	exitedY := func() bool { return guard.y <= 0 || guard.y >= len(grid)-1 }

// 	resetGrid := func() {
// 		direction = "up"
// 		grid, guardX, guardY = createGrid(filename)
// 	}

// 	for !exitedX() && !exitedY() {
// 		nextX, nextY := guard.getNext()
// 		if grid[nextY][nextX] == "#" {
// 			grid[guard.y][guard.x] = "+"
// 			guard.changeDirection()
// 		}
// 	}

// 	return len(usedBlockLocations)
// }

func partTwo(filename string) int {

	grid, guardX, guardY := createGrid(filename)
	direction := "up"

	var usedBlockLocations []ObstacleCoordinates
	exitedX := func() bool { return guardX <= 0 || guardX >= len(grid[0])-1 }
	exitedY := func() bool { return guardY <= 0 || guardY >= len(grid)-1 }

	resetGrid := func() {
		direction = "up"
		grid, guardX, guardY = createGrid(filename)
	}

	for !exitedX() && !exitedY() {
		switch direction {
		case "up":
			checkNext := grid[guardY-1][guardX]
			//Got back to loop location, Note Location and reset
			if checkNext == "#" {
				//Regular logic below
				grid[guardY][guardX] = "+"
				direction = "right"
			} else if checkNext == "O" {
				for _, coord := range usedBlockLocations {
					if coord.x == guardX && coord.y == guardY-1 {
						resetGrid()
						continue
					}
				}
				usedBlockLocations = append(usedBlockLocations, *newBlockCoordinates(guardX, guardY-1))
				resetGrid()
				continue
			} else if obstacleAhead(guardX, guardY, direction, grid) &&
				grid[guardY][guardX+1] == "-" &&
				grid[guardY][guardX] != "+" {
				//Check if loopable location
				locationUsed := false
				for _, coords := range usedBlockLocations {
					if coords.x == guardX && coords.y == guardY-1 {
						locationUsed = true
						break
					}
				}
				if !locationUsed {
					grid[guardY-1][guardX] = "O"
					grid[guardY][guardX] = "+"
					direction = "right"
				} else {
					guardY--
					grid[guardY][guardX] = "|"
				}
			} else {
				guardY--
				grid[guardY][guardX] = "|"
			}
		case "right":
			checkNext := grid[guardY][guardX+1]
			//Got back to loop location, Note Location and reset
			if checkNext == "#" {
				//Regular logic below
				grid[guardY][guardX] = "+"
				direction = "down"
			} else if checkNext == "O" {
				for _, coord := range usedBlockLocations {
					if coord.x == guardX+1 && coord.y == guardY {
						resetGrid()
						continue
					}
				}
				usedBlockLocations = append(usedBlockLocations, *newBlockCoordinates(guardX+1, guardY))
				resetGrid()
				continue
			} else if obstacleAhead(guardX, guardY, direction, grid) &&
				grid[guardY+1][guardX] == "|" &&
				grid[guardY][guardX] != "+" {
				//Check if loopable location
				locationUsed := false
				for _, coords := range usedBlockLocations {
					if coords.x == guardX+1 && coords.y == guardY {
						locationUsed = true
						break
					}
				}
				if !locationUsed {
					grid[guardY][guardX+1] = "O"
					grid[guardY][guardX] = "+"
					direction = "down"
				} else {
					guardX++
					grid[guardY][guardX] = "-"
				}
			} else {
				guardX++
				grid[guardY][guardX] = "-"
			}
		case "down":
			checkNext := grid[guardY+1][guardX]
			//Got back to loop location, Note Location and reset
			if checkNext == "#" {
				//Regular logic below
				grid[guardY][guardX] = "+"
				direction = "left"
			} else if checkNext == "O" {
				for _, coord := range usedBlockLocations {
					if coord.x == guardX && coord.y+1 == guardY {
						resetGrid()
						continue
					}
				}
				usedBlockLocations = append(usedBlockLocations, *newBlockCoordinates(guardX, guardY+1))
				resetGrid()
				continue
			} else if obstacleAhead(guardX, guardY, direction, grid) &&
				grid[guardY][guardX-1] == "-" &&
				grid[guardY][guardX] != "+" {
				//Check if loopable location
				locationUsed := false
				for _, coords := range usedBlockLocations {
					if coords.x == guardX-1 && coords.y == guardY {
						locationUsed = true
						break
					}
				}
				if !locationUsed {
					grid[guardY+1][guardX] = "O"
					grid[guardY][guardX] = "+"
					direction = "left"
				} else {
					guardY++
					grid[guardY][guardX] = "|"
				}
			} else {
				guardY++
				grid[guardY][guardX] = "|"
			}
		case "left":
			checkNext := grid[guardY][guardX-1]
			//Got back to loop location, Note Location and reset
			if checkNext == "#" {
				//Regular logic below
				grid[guardY][guardX] = "+"
				direction = "up"
			} else if checkNext == "O" {
				for _, coord := range usedBlockLocations {
					if coord.x == guardX-1 && coord.y == guardY {
						resetGrid()
						continue
					}
				}
				usedBlockLocations = append(usedBlockLocations, *newBlockCoordinates(guardX-1, guardY))
				resetGrid()
				continue
			} else if obstacleAhead(guardX, guardY, direction, grid) &&
				grid[guardY-1][guardX] == "|" &&
				grid[guardY][guardX] != "+" {
				//Check if loopable location
				locationUsed := false
				for _, coords := range usedBlockLocations {
					if coords.x == guardX-1 && coords.y == guardY {
						locationUsed = true
						break
					}
				}
				if !locationUsed {
					grid[guardY][guardX-1] = "O"
					grid[guardY][guardX] = "+"
					direction = "up"
				} else {
					guardX--
					grid[guardY][guardX] = "-"
				}
			} else {
				guardX--
				grid[guardY][guardX] = "-"
			}
		}
	}

	return len(usedBlockLocations)
}

// func checkSpot(x int, y int, xStep int, yStep int, stepString string, grid [][]string) {
// 	nextStep := grid[y+yStep][x+xStep]
// 	if nextStep == "#" {
// 		stepString
// 	}
// }

func obstacleAhead(currX int, currY int, direction string, grid [][]string) bool {
	switch direction {
	case "up":
		for currY >= 0 {
			if grid[currY][currX] == "#" {
				return true
			}
			currY--
		}
		return false
	case "right":
		for currX < len(grid[currY]) {
			if grid[currY][currX] == "#" {
				return true
			}
			currX++
		}
		return false
	case "down":
		for currY < len(grid) {
			if grid[currY][currX] == "#" {
				return true
			}
			currY++
		}
		return false
	case "left":
		for currX >= 0 {
			if grid[currY][currX] == "#" {
				return true
			}
			currX--
		}
		return false
	default:
		return false
	}
}
