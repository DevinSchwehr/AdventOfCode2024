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

	partOneResult := partOne("input.txt")
	partTwoResult := partTwo("input.txt")

	fmt.Printf("Result for Part One is: %v \n", partOneResult)
	fmt.Printf("Result for Part Two is: %v \n", partTwoResult)
}

func getInputsFromLine(s string) (int, int) {
	inputs := strings.Fields(s)
	if len(inputs) != 2 {
		log.Fatal("Improper input line")
	}
	firstNum, err1 := strconv.Atoi(inputs[0])
	secondNum, err2 := strconv.Atoi(inputs[1])

	if err1 != nil || err2 != nil {
		log.Fatal("error parsing integers: %v, %t", err1, err2)
	}

	return firstNum, secondNum

}

func partOne(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var firstList []int
	var secondList []int

	for scanner.Scan() {

		firstNum, secondNum := getInputsFromLine(scanner.Text())

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

func partTwo(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := make(map[int]int)

	var keyList []int
	var occurenceList []int

	for scanner.Scan() {
		key, number := getInputsFromLine(scanner.Text())
		m[key] = 0
		keyList = append(keyList, key)
		occurenceList = append(occurenceList, number)
	}
	for _, value := range occurenceList {
		_, exists := m[value]
		if exists {
			m[value]++
		}
	}
	result := 0
	for _, key := range keyList {
		result += key * m[key]
	}
	return result
}
