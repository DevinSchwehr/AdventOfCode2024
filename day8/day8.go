package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Execute() {
	os.Chdir("day8")

	fmt.Printf("Result for Part One is: %d \n", partOne("input.txt"))
	fmt.Printf("Result for Part Two is: %d \n", partTwo("input.txt"))
}

type Antenna struct {
	x         int
	y         int
	frequency string
}

func newAntenna(x int, y int, frequency string) *Antenna {
	result := new(Antenna)
	result.x = x
	result.y = y
	result.frequency = frequency
	return result
}

type Antinode struct {
	x int
	y int
}

func newAntinode(x int, y int) *Antinode {
	result := new(Antinode)
	result.x = x
	result.y = y
	return result
}

func partOne(filename string) int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]string
	var antennaLocations []*Antenna
	var antinodes []*Antinode

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		grid = append(grid, line)
		for index, char := range line {
			if char != "." {
				antenna := newAntenna(index, len(grid)-1, char)
				antennaLocations = append(antennaLocations, antenna)
			}
		}
	}

	for index, currentAntenna := range antennaLocations {
		for secondIndex, secondAntenna := range antennaLocations {
			//Look for two different antenna of the same frequency
			if currentAntenna.frequency == secondAntenna.frequency && index != secondIndex {
				xdiff := currentAntenna.x - secondAntenna.x
				ydiff := currentAntenna.y - secondAntenna.y
				closeX, closeY := (currentAntenna.x + xdiff), (currentAntenna.y + ydiff)
				farX, farY := (currentAntenna.x + (xdiff * -2)), (currentAntenna.y + (ydiff * -2))
				//Check close can be added
				if inBounds(closeX, closeY, grid) && doesNotContainAntinode(closeX, closeY, antinodes) {
					antinodes = append(antinodes, newAntinode(closeX, closeY))
				}
				//Check if far can be added
				if inBounds(farX, farY, grid) && doesNotContainAntinode(farX, farY, antinodes) {
					antinodes = append(antinodes, newAntinode(farX, farY))
				}
			}
		}
	}

	return len(antinodes)
}

func inBounds(x int, y int, grid [][]string) bool {
	if y >= 0 && y <= len(grid)-1 {
		if x >= 0 && x <= len(grid[y])-1 {
			return true
		}
	}
	return false
}

func doesNotContainAntinode(x int, y int, antinodes []*Antinode) bool {
	for _, antinode := range antinodes {
		if antinode.x == x && antinode.y == y {
			return false
		}
	}
	return true
}

func partTwo(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]string
	var antennaLocations []*Antenna
	var antinodes []*Antinode

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		grid = append(grid, line)
		for index, char := range line {
			if char != "." {
				antenna := newAntenna(index, len(grid)-1, char)
				antennaLocations = append(antennaLocations, antenna)
			}
		}
	}

	for index, currentAntenna := range antennaLocations {
		for secondIndex, secondAntenna := range antennaLocations {
			//Look for two different antenna of the same frequency
			if currentAntenna.frequency == secondAntenna.frequency && index != secondIndex {
				xdiff := currentAntenna.x - secondAntenna.x
				ydiff := currentAntenna.y - secondAntenna.y
				currX, currY := currentAntenna.x, currentAntenna.y
				//Work in 'increment' direction
				for inBounds(currX, currY, grid) {
					if doesNotContainAntinode(currX, currY, antinodes) {
						antinodes = append(antinodes, newAntinode(currX, currY))
					}
					currX += xdiff
					currY += ydiff
				}
				currX, currY = currentAntenna.x, currentAntenna.y
				//work in 'decrement' direction
				for inBounds(currX, currY, grid) {
					if doesNotContainAntinode(currX, currY, antinodes) {
						antinodes = append(antinodes, newAntinode(currX, currY))
					}
					currX -= xdiff
					currY -= ydiff
				}
			}
		}
	}

	return len(antinodes)
}
