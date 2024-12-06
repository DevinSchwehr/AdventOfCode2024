package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Execute() {
	os.Chdir("day5")

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

	pageOrderMap := make(map[int][]int)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			orderings := strings.Split(line, "|")
			keyPage, err1 := strconv.Atoi(orderings[0])
			valuePage, err2 := strconv.Atoi(orderings[1])
			if err1 != nil || err2 != nil {
				log.Fatal("Error getting page orders")
			}
			pages := pageOrderMap[keyPage]
			pages = append(pages, valuePage)
			pageOrderMap[keyPage] = pages
		}
		if strings.Contains(line, ",") {
			result += getMiddle(line, pageOrderMap)
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

	pageOrderMap := make(map[int][]int)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			orderings := strings.Split(line, "|")
			keyPage, err1 := strconv.Atoi(orderings[0])
			valuePage, err2 := strconv.Atoi(orderings[1])
			if err1 != nil || err2 != nil {
				log.Fatal("Error getting page orders")
			}
			pages := pageOrderMap[keyPage]
			pages = append(pages, valuePage)
			pageOrderMap[keyPage] = pages
		}
		if strings.Contains(line, ",") {
			pages := strings.Split(line, ",")
			var pageNums []int
			for _, page := range pages {
				num, err := strconv.Atoi(page)
				if err != nil {
					log.Fatal("couldn't parse integer from string")
				}
				pageNums = append(pageNums, num)
			}
			var incorrect bool
			incorrect, pageNums = getMiddlePartTwo(pageNums, pageOrderMap)
			if incorrect {
				for incorrect {
					incorrect, pageNums = getMiddlePartTwo(pageNums, pageOrderMap)
				}
				result += pageNums[len(pageNums)/2]
			}
		}
	}

	return result
}

func getMiddle(line string, orderings map[int][]int) int {
	pages := strings.Split(line, ",")
	var pageNums []int
	for _, page := range pages {
		num, err := strconv.Atoi(page)
		if err != nil {
			log.Fatal("couldn't parse integer from string")
		}
		pageNums = append(pageNums, num)
	}

	for index, page := range pageNums {
		for _, value := range orderings[page] {
			indexes := getPageIndexes(value, pageNums)
			for _, pageIndex := range indexes {
				if pageIndex < index {
					return 0
				}
			}
		}
	}

	return pageNums[len(pageNums)/2]
}

func getMiddlePartTwo(pageNums []int, orderings map[int][]int) (bool, []int) {
	incorrect := false
	for index, page := range pageNums {
		var incorrectIndexes []int
		//Go through the values for the key
		for _, value := range orderings[page] {
			indexes := getPageIndexes(value, pageNums)
			//Go through the indexes for page
			for _, pageIndex := range indexes {
				if pageIndex < index {
					incorrect = true
					incorrectIndexes = append(incorrectIndexes, pageIndex)
				}
			}
		}
		if incorrect {
			//For any invalid indexes, move that element to be in front of page. Reverse order for proper insertion
			slices.Reverse(incorrectIndexes)
			for _, badIndex := range incorrectIndexes {
				editNums := make([]int, len(pageNums))
				copy(editNums, pageNums)
				var reOrder []int
				elemBefore := editNums[:badIndex]
				elemBetween := editNums[badIndex+1 : index]
				elemAfter := editNums[index+1:]
				reOrder = append(reOrder, elemBefore...)
				reOrder = append(reOrder, elemBetween...)
				reOrder = append(reOrder, pageNums[index])
				reOrder = append(reOrder, pageNums[badIndex])
				reOrder = append(reOrder, elemAfter...)
				pageNums = reOrder
				break
			}
		}
	}

	return incorrect, pageNums
}

func getPageIndexes(target int, array []int) []int {
	var indexes []int
	for index, value := range array {
		if value == target {
			indexes = append(indexes, index)
		}
	}
	return indexes
}
