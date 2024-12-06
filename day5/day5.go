package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func getPageIndexes(target int, array []int) []int {
	var indexes []int
	for index, value := range array {
		if value == target {
			indexes = append(indexes, index)
		}
	}
	return indexes
}

func partTwo(filename string) int {

	// file, err := os.Open(filename)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)

	return 0
}
