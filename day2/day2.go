package day2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Execute() {
	os.Chdir("day2")

	result := partOne("input.txt")

	fmt.Printf("Result for Part One is: %d \n", result)

	result = partTwo("input.txt")
	fmt.Printf("Result for Part Two is: %d \n", result)
}

func partOne(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		levels := strings.Fields(scanner.Text())
		var levelArray []int
		for _, value := range levels {
			levelNumber, err := strconv.Atoi(value)
			if err != nil {
				log.Panic("unable to parse integer")
			}
			levelArray = append(levelArray, levelNumber)
		}

		if checkLevelSafe(levelArray) {
			result++
		}
	}

	return result
}

func checkLevelSafe(levelArray []int) bool {
	isSafe := true

	for index, value := range levelArray {
		//Base cases
		if index == 0 || index == len(levelArray)-1 {
			continue
		}
		//check is increasing or decreasing
		increasing := levelArray[index-1] < value && value < levelArray[index+1]
		decreasing := levelArray[index-1] > value && value > levelArray[index+1]
		prevDifference := math.Abs(float64(levelArray[index-1] - value))
		nextDifference := math.Abs(float64(levelArray[index+1] - value))

		isSafe = (increasing || decreasing) && (prevDifference <= 3 && nextDifference <= 3)

		if !isSafe {
			break
		}

	}
	return isSafe
}

func partTwo(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		levels := strings.Fields(scanner.Text())
		var levelArray []int
		for _, value := range levels {
			levelNumber, err := strconv.Atoi(value)
			if err != nil {
				log.Panic("unable to parse integer")
			}
			levelArray = append(levelArray, levelNumber)
		}

		isSafe := true

		if len(levelArray) == 1 {
			result++
			continue
		}

		if len(levelArray) == 2 {
			increasing := levelArray[0] < levelArray[1]
			decreasing := levelArray[0] > levelArray[1]
			validDifference := math.Abs(float64(levelArray[0]-levelArray[1])) <= 3
			if (increasing || decreasing) && validDifference {
				result++
				continue
			}
		}

		if !checkLevelSafe(levelArray) {
			validWithRemove := false
			for index, _ := range levelArray {
				level := make([]int, len(levelArray))
				copy(level, levelArray)
				var slice []int
				switch index {
				case 0:
					slice = level[1:]
				case len(level) - 1:
					slice = level[:index]
				default:
					first := level[:index]
					second := level[index+1:]
					slice = append(first, second...)
				}
				if checkLevelSafe(slice) {
					validWithRemove = true
					break
				}
			}
			isSafe = validWithRemove
		}

		if isSafe {
			result++
		}
	}

	return result
}
