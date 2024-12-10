package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Execute() {
	os.Chdir("day7")

	fmt.Printf("Result for Part One is: %d \n", partOne("input.txt"))
	fmt.Printf("Result for Part Two is: %d \n", partTwo("input.txt"))
}

type Node struct {
	value    int
	children []*Node
}

func createNode(value int) *Node {
	node := new(Node)
	node.value = value
	return node
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
		line := scanner.Text()
		splitString := strings.Split(line, ": ")
		total, err := strconv.Atoi(splitString[0])
		if err != nil {
			log.Fatal("Error parsing int from string")
		}
		numberStrings := strings.Split(splitString[1], " ")
		var numbers []int
		for _, value := range numberStrings {
			number, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal("Error parsing int from string")
			}
			numbers = append(numbers, number)
		}

		if canEqual(total, numbers, true) {
			result += total
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
		line := scanner.Text()
		splitString := strings.Split(line, ": ")
		total, err := strconv.Atoi(splitString[0])
		if err != nil {
			log.Fatal("Error parsing int from string")
		}
		numberStrings := strings.Split(splitString[1], " ")
		var numbers []int
		for _, value := range numberStrings {
			number, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal("Error parsing int from string")
			}
			numbers = append(numbers, number)
		}

		if canEqual(total, numbers, false) {
			result += total
		}
	}

	return result
}

func canEqual(sum int, numbers []int, isPartOne bool) bool {
	//Base cases
	if len(numbers) == 1 {
		return numbers[0] == sum
	}
	if len(numbers) == 2 {
		return numbers[0]+numbers[1] == sum ||
			numbers[0]*numbers[1] == sum ||
			concatNumbers(numbers[0], numbers[1]) == sum
	}

	baseNode := createNode(numbers[0])
	if isPartOne {
		propagateTree(sum, 1, baseNode, numbers)
	} else {
		propagateTreePartTwo(sum, 1, baseNode, numbers)
	}

	return findTarget(sum, baseNode, 1, numbers)
}

func propagateTree(target int, index int, node *Node, numbers []int) {
	sumResult := node.value + numbers[index]
	multResult := node.value * numbers[index]

	//At end of numbers
	if index == len(numbers)-1 {
		children := []*Node{
			createNode(sumResult),
			createNode(multResult),
		}
		node.children = children
		return
	}
	index++
	addNode := createNode(sumResult)
	if addNode.value <= target {
		propagateTree(target, index, addNode, numbers)
	}
	multNode := createNode(multResult)
	if multNode.value <= target {
		propagateTree(target, index, multNode, numbers)
	}
	children := []*Node{addNode, multNode}
	node.children = children

}

func propagateTreePartTwo(target int, index int, node *Node, numbers []int) {
	sumResult := node.value + numbers[index]
	multResult := node.value * numbers[index]
	concatResult := concatNumbers(node.value, numbers[index])

	//At end of numbers
	if index == len(numbers)-1 {
		children := []*Node{
			createNode(sumResult),
			createNode(multResult),
			createNode(concatResult),
		}
		node.children = children
		return
	}
	index++
	addNode := createNode(sumResult)
	if addNode.value <= target {
		propagateTreePartTwo(target, index, addNode, numbers)
	}
	multNode := createNode(multResult)
	if multNode.value <= target {
		propagateTreePartTwo(target, index, multNode, numbers)
	}
	concatNode := createNode(concatResult)
	if concatNode.value <= target {
		propagateTreePartTwo(target, index, concatNode, numbers)
	}
	children := []*Node{addNode, multNode, concatNode}
	node.children = children
}

func findTarget(target int, node *Node, index int, numbers []int) bool {
	if index == len(numbers)-1 {
		for _, child := range node.children {
			if child.value == target {
				return true
			}
		}
		return false
	}
	for _, value := range node.children {
		if len(value.children) > 0 {
			nextLevel := index + 1
			childValid := findTarget(target, value, nextLevel, numbers)
			if childValid {
				return true
			}
		}
	}
	return false
}

func concatNumbers(first int, second int) int {
	firstString := strconv.Itoa(first)
	secondString := strconv.Itoa(second)
	firstString += secondString
	concatResult, err := strconv.Atoi(firstString)
	if err != nil {
		log.Fatal("error parsing int from string")
	}
	return concatResult
}
