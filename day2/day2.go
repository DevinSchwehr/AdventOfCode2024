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
		if isSafe {
			result++
		}
	}

	return result
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
		removedIndex := -1

		fmt.Println(levelArray)

		firstValid, firstRemoved := checkFirstIndex(levelArray)
		lastValid, lastRemoved := checkLastIndex(levelArray)
		if firstValid && firstRemoved {
			removedIndex = 0
		}
		if lastValid && lastRemoved {
			removedIndex = len(levelArray) - 1
		}
		//if the first or last index is not removable, report is not valid
		if !firstValid || !lastValid {
			isSafe = false
		} else {
			for index, value := range levelArray {
				//Middle Indexes
				if index > 0 && index < len(levelArray)-1 && index != removedIndex {
					prevIndex := index - 1
					if removedIndex == prevIndex {
						if removedIndex == 0 {
							increasing := value < levelArray[index+1]
							decreasing := value > levelArray[index+1]
							validDifference := math.Abs(float64(value-levelArray[index+1])) <= 3
							isSafe = (increasing || decreasing) && validDifference
						} else {
							prevIndex = removedIndex - 1
							increasing := value > levelArray[prevIndex] && value < levelArray[index+1]
							decreasing := value < levelArray[prevIndex] && value > levelArray[index+1]
							validDifference := (math.Abs(float64(value-levelArray[prevIndex])) <= 3) &&
								(math.Abs(float64(value-levelArray[index+1])) <= 3)
							isSafe = (increasing || decreasing) && validDifference
						}
					} else if removedIndex == index+1 {
						increasing := value < levelArray[index-1]
						decreasing := value > levelArray[index-1]
						validDifference := math.Abs(float64(value-levelArray[index-1])) <= 3
						isSafe = (increasing || decreasing) && validDifference
					} else {
						increasing := value > levelArray[prevIndex] && value < levelArray[index+1]
						decreasing := value < levelArray[prevIndex] && value > levelArray[index+1]
						validDifference := (math.Abs(float64(value-levelArray[prevIndex])) <= 3) &&
							(math.Abs(float64(value-levelArray[index+1])) <= 3)
						isSafe = (increasing || decreasing) && validDifference
					}
					if !isSafe {
						//Able to remove level?
						if removedIndex == -1 {
							removable :=
								(levelArray[index-1] < levelArray[index+1] || levelArray[index-1] > levelArray[index+1]) &&
									(math.Abs(float64(levelArray[index-1]-levelArray[index+1])) <= 3)
							if removable {
								isSafe = true
								removedIndex = index
							} else {
								//Check if next element is breaking
								nextRemovable := checkSkipNextValidity(index, levelArray)
								if nextRemovable {
									isSafe = true
									removedIndex = index + 1
								} else {
									break
								}
							}
						} else if removedIndex == index-1 {
							//Check if removed index could be shifted by one
							removable :=
								(levelArray[index-1] < levelArray[index+1] || levelArray[index-1] > levelArray[index+1]) &&
									(math.Abs(float64(levelArray[index-1]-levelArray[index+1])) <= 3)
							if removable {
								isSafe = true
								removedIndex = index
							}
						} else {
							break
						}
					}
				}
			}
		}

		fmt.Println(isSafe)
		if isSafe {
			result++
		}
	}

	return result
}

func checkSkipNextValidity(index int, levelArray []int) bool {
	if index+2 <= len(levelArray)-1 {
		value := levelArray[index]
		increasing := value > levelArray[index-1] && value < levelArray[index+2]
		decreasing := value < levelArray[index-1] && value > levelArray[index+2]
		validDifference := (math.Abs(float64(value-levelArray[index-1])) <= 3) &&
			(math.Abs(float64(value-levelArray[index+2])) <= 3)
		return (increasing || decreasing) && validDifference
	}
	return false
}

func checkLastIndex(levelArray []int) (bool, bool) {
	removed := false
	index := len(levelArray) - 1
	value := levelArray[index]
	increasing := value < levelArray[index-1]
	decreasing := value > levelArray[index-1]
	difference := math.Abs(float64(value - levelArray[index-1]))
	isSafe := (increasing || decreasing) && difference <= 3
	if !isSafe {
		increasing = levelArray[index-2] < levelArray[index-1]
		decreasing = levelArray[index-2] > levelArray[index-1]
		difference = math.Abs(float64(levelArray[index-2] - levelArray[index-1]))
		removable := (increasing || decreasing) && difference <= 3
		if removable {
			isSafe = true
			removed = removable
		}
	}
	return isSafe, removed
}

func checkFirstIndex(levelArray []int) (bool, bool) {
	value := levelArray[0]
	removed := false
	increasing := value < levelArray[1]
	decreasing := value > levelArray[1]
	difference := math.Abs(float64(value - levelArray[1]))
	isSafe := (increasing || decreasing) && difference <= 3
	if !isSafe {
		increasing = levelArray[1] < levelArray[2]
		decreasing = levelArray[1] > levelArray[2]
		difference = math.Abs(float64(levelArray[1] - levelArray[2]))
		removable := (increasing || decreasing) && difference <= 3
		if removable {
			isSafe = true
			removed = removable
		}
	}
	return isSafe, removed
}
