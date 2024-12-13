package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Execute() {

	os.Chdir("day10")

	partOneResult := partOne("input.txt")
	partTwoResult := partTwo("input.txt")

	fmt.Printf("Result for Part One is: %v \n", partOneResult)
	fmt.Printf("Result for Part Two is: %v \n", partTwoResult)
}

type Coordinate struct {
	x int
	y int
}

func newCoordinate(x int, y int) *Coordinate {
	coord := new(Coordinate)
	coord.x = x
	coord.y = y
	return coord
}

func partOne(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var topoMap [][]int
	var coordinates []*Coordinate

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStringArray := strings.Split(scanner.Text(), "")
		var lineIntArray []int
		for index, value := range lineStringArray {
			valueNum, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal("error parsing int from string")
			}
			if valueNum == 0 {
				coordinates = append(coordinates, newCoordinate(index, len(topoMap)))
			}
			lineIntArray = append(lineIntArray, valueNum)
		}
		topoMap = append(topoMap, lineIntArray)
	}

	result := 0
	for _, coord := range coordinates {
		var topCoords []Coordinate
		result += trailSearch(coord.x, coord.y, topoMap, &topCoords)
	}
	return result
}

func partTwo(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var topoMap [][]int
	var coordinates []*Coordinate

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStringArray := strings.Split(scanner.Text(), "")
		var lineIntArray []int
		for index, value := range lineStringArray {
			valueNum, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal("error parsing int from string")
			}
			if valueNum == 0 {
				coordinates = append(coordinates, newCoordinate(index, len(topoMap)))
			}
			lineIntArray = append(lineIntArray, valueNum)
		}
		topoMap = append(topoMap, lineIntArray)
	}

	result := 0
	for _, coord := range coordinates {
		result += trailSearchPartTwo(coord.x, coord.y, topoMap)
	}
	return result
}

func trailSearch(currX int, currY int, topoMap [][]int, topCoords *[]Coordinate) int {
	if topoMap[currY][currX] == 9 {
		for _, value := range *topCoords {
			if value.x == currX && value.y == currY {
				return 0
			}
		}
		*topCoords = append(*topCoords, *newCoordinate(currX, currY))
		return 1
	}
	result := 0
	//Check up
	if currY-1 >= 0 {
		if topoMap[currY-1][currX] == topoMap[currY][currX]+1 {
			result += trailSearch(currX, currY-1, topoMap, topCoords)
		}
	}
	//Check down
	if currY+1 < len(topoMap) {
		if topoMap[currY+1][currX] == topoMap[currY][currX]+1 {
			result += trailSearch(currX, currY+1, topoMap, topCoords)
		}
	}
	//Check right
	if currX+1 < len(topoMap[currY]) {
		if topoMap[currY][currX+1] == topoMap[currY][currX]+1 {
			result += trailSearch(currX+1, currY, topoMap, topCoords)
		}
	}
	//Check left
	if currX-1 >= 0 {
		if topoMap[currY][currX-1] == topoMap[currY][currX]+1 {
			result += trailSearch(currX-1, currY, topoMap, topCoords)
		}
	}
	return result
}

func trailSearchPartTwo(currX int, currY int, topoMap [][]int) int {
	if topoMap[currY][currX] == 9 {
		return 1
	}
	result := 0
	//Check up
	if currY-1 >= 0 {
		if topoMap[currY-1][currX] == topoMap[currY][currX]+1 {
			result += trailSearchPartTwo(currX, currY-1, topoMap)
		}
	}
	//Check down
	if currY+1 < len(topoMap) {
		if topoMap[currY+1][currX] == topoMap[currY][currX]+1 {
			result += trailSearchPartTwo(currX, currY+1, topoMap)
		}
	}
	//Check right
	if currX+1 < len(topoMap[currY]) {
		if topoMap[currY][currX+1] == topoMap[currY][currX]+1 {
			result += trailSearchPartTwo(currX+1, currY, topoMap)
		}
	}
	//Check left
	if currX-1 >= 0 {
		if topoMap[currY][currX-1] == topoMap[currY][currX]+1 {
			result += trailSearchPartTwo(currX-1, currY, topoMap)
		}
	}
	return result
}
