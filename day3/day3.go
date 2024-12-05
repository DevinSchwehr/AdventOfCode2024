package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Execute() {
	os.Chdir("day3")

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

	regex, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		log.Fatal("error creating regex")
	}

	var matches []string

	for scanner.Scan() {
		line := scanner.Text()
		matches = append(matches, regex.FindAllString(line, -1)...)
	}

	result := 0

	for _, value := range matches {
		digitRegex, err := regexp.Compile(`\d{1,3}`)
		if err != nil {
			log.Fatal("error creating regex")
		}
		numbers := digitRegex.FindAllString(value, 2)

		firstDigit, err1 := strconv.Atoi(numbers[0])
		secondDigit, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			log.Fatal("error parsing numbers")
		}

		result += (firstDigit * secondDigit)
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

	regex, err := regexp.Compile(`(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\))`)
	if err != nil {
		log.Fatal("error creating regex")
	}

	var matches []string

	for scanner.Scan() {
		line := scanner.Text()
		matches = append(matches, regex.FindAllString(line, -1)...)
	}

	result := 0

	accepting := true
	for _, value := range matches {
		if value == "do()" {
			accepting = true
			continue
		} else if value == "don't()" {
			accepting = false
			continue
		}
		if accepting {
			digitRegex, err := regexp.Compile(`\d{1,3}`)
			if err != nil {
				log.Fatal("error creating regex")
			}
			numbers := digitRegex.FindAllString(value, 2)

			firstDigit, err1 := strconv.Atoi(numbers[0])
			secondDigit, err2 := strconv.Atoi(numbers[1])
			if err1 != nil || err2 != nil {
				log.Fatal("error parsing numbers")
			}
			if accepting {
				result += (firstDigit * secondDigit)
			}
		}

	}

	return result

}
