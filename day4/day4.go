package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Execute() {
	os.Chdir("day4")

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
	var crossword [][]string

	for scanner.Scan() {
		line := scanner.Text()
		lineChars := strings.Split(line, "")
		crossword = append(crossword, lineChars)
	}

	matches := 0
	for lineIndex, line := range crossword {
		for charIndex, letter := range line {
			if letter == "X" {
				matches += isXMAS(searchCrossword("", charIndex, lineIndex, "horizontal-forward", crossword))
				matches += isXMAS(searchCrossword("", charIndex, lineIndex, "horizontal-backward", crossword))
				matches += isXMAS(searchCrossword("", charIndex, lineIndex, "vertical-up", crossword))
				matches += isXMAS(searchCrossword("", charIndex, lineIndex, "vertical-down", crossword))
				matches += isXMAS(searchCrossword("", charIndex, lineIndex, "diagonal-right-up", crossword))
				matches += isXMAS(searchCrossword("", charIndex, lineIndex, "diagonal-right-down", crossword))
				matches += isXMAS(searchCrossword("", charIndex, lineIndex, "diagonal-left-up", crossword))
				matches += isXMAS(searchCrossword("", charIndex, lineIndex, "diagonal-left-down", crossword))
			}
		}
	}

	return matches
}

func isXMAS(word string) int {
	if word == "XMAS" {
		return 1
	}
	return 0
}

func searchCrossword(word string, charIndex int, lineIndex int, direction string, crossword [][]string) string {
	//Base cases
	if lineIndex < 0 || lineIndex >= len(crossword) {
		return word
	}
	if charIndex < 0 || charIndex >= len(crossword[lineIndex]) {
		return word
	}

	word += crossword[lineIndex][charIndex]
	if word == "XMAS" {
		return word
	} else if len(word) == 4 {
		return word
	}

	switch direction {
	case "horizontal-forward":
		charIndex++
		word = searchCrossword(word, charIndex, lineIndex, direction, crossword)
	case "horizontal-backward":
		charIndex--
		word = searchCrossword(word, charIndex, lineIndex, direction, crossword)
	case "vertical-up":
		lineIndex--
		word = searchCrossword(word, charIndex, lineIndex, direction, crossword)
	case "vertical-down":
		lineIndex++
		word = searchCrossword(word, charIndex, lineIndex, direction, crossword)
	case "diagonal-right-up":
		charIndex++
		lineIndex--
		word = searchCrossword(word, charIndex, lineIndex, direction, crossword)
	case "diagonal-right-down":
		charIndex++
		lineIndex++
		word = searchCrossword(word, charIndex, lineIndex, direction, crossword)
	case "diagonal-left-up":
		charIndex--
		lineIndex--
		word = searchCrossword(word, charIndex, lineIndex, direction, crossword)
	case "diagonal-left-down":
		charIndex--
		lineIndex++
		word = searchCrossword(word, charIndex, lineIndex, direction, crossword)
	}
	return word

}

func partTwo(filename string) int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var crossword [][]string

	for scanner.Scan() {
		line := scanner.Text()
		lineChars := strings.Split(line, "")
		crossword = append(crossword, lineChars)
	}

	matches := 0
	for lineIndex, line := range crossword {
		for charIndex, letter := range line {
			if letter == "A" {
				leftToRightUp, leftToRightDown := getDiagonals(charIndex, lineIndex, crossword)
				if containsMAS(leftToRightDown) && containsMAS(leftToRightUp) {
					matches++
				}

			}
		}
	}

	return matches
}

func containsMAS(word string) bool {
	if strings.Contains(word, "A") &&
		strings.Contains(word, "M") &&
		strings.Contains(word, "S") &&
		len(word) == 3 {
		return true
	}
	return false
}

func getDiagonals(charIndex int, lineIndex int, crossword [][]string) (string, string) {
	lineBelow := lineIndex + 1
	lineAbove := lineIndex - 1
	charForward := charIndex + 1
	charBackward := charIndex - 1

	leftToRightUp := "A"
	leftToRightDown := "A"

	if lineBelow < len(crossword) && charBackward >= 0 {
		leftToRightUp += crossword[lineBelow][charBackward]
	}
	if lineAbove >= 0 && charForward < len(crossword[lineAbove]) {
		leftToRightUp += crossword[lineAbove][charForward]
	}

	if lineAbove >= 0 && charBackward >= 0 {
		leftToRightDown += crossword[lineAbove][charBackward]
	}
	if lineBelow < len(crossword) && charForward < len(crossword[lineBelow]) {
		leftToRightDown += crossword[lineBelow][charForward]
	}

	return leftToRightUp, leftToRightDown
}
