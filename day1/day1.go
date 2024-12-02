package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Execute() {

	os.Chdir("day1")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	partOneResult := partOne(file)

	fmt.Printf("Result for Part One is: %v \n", partOneResult)
}

func partOne(file *os.File) int {
	scanner := bufio.NewScanner(file)

	var firstList []int
	var secondList []int

	for scanner.Scan() {
		inputs := strings.Fields(scanner.Text())
		if len(inputs) != 2 {
			log.Fatal("Improper input line")
		}
		firstNum, err1 := strconv.Atoi(inputs[0])
		secondNum, err2 := strconv.Atoi(inputs[1])

		if err1 != nil || err2 != nil {
			log.Fatal("error parsing integers: %v, %t", err1, err2)
		}

		firstList = append(firstList, firstNum)
		secondList = append(secondList, secondNum)

	}

	sort.Ints(firstList)
	sort.Ints(secondList)

	var differences int = 0

	var index int = 0
	for index < len(firstList) {
		differences += int(math.Abs(float64(firstList[index] - secondList[index])))
		index++
	}
	return differences
}
