package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Execute() {

	os.Chdir("day9")

	partOneResult := partOne("input.txt")
	fmt.Printf("Result for Part One is: %v \n", partOneResult)
	partTwoResult := partTwo("input.txt")
	fmt.Printf("Result for Part Two is: %v \n", partTwoResult)
}

type Node struct {
	value     string
	blockSize int
	prev      *Node
	next      *Node
}

func newNode(v string, n *Node, p *Node, blockSize int) *Node {
	result := new(Node)
	result.value = v
	result.prev = p
	result.next = n
	result.blockSize = blockSize
	return result
}

func partOne(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	blocks := strings.Split(scanner.Text(), "")
	var head *Node
	// var tail *Node
	var currNode *Node
	currId := 0
	for index, blockString := range blocks {
		blockSize, err := strconv.Atoi(blockString)
		if err != nil {
			log.Fatal("error parsing int from string")
		}
		if index == 0 {
			//Initial Block
			head = newNode(strconv.Itoa(currId), nil, nil, 1)
			currNode = head
			currNode = advanceLinkedList(currNode, strconv.Itoa(currId), blockSize-1)
			currId++
		} else if index%2 == 0 {
			//File block
			currNode = advanceLinkedList(currNode, strconv.Itoa(currId), blockSize)
			currId++
		} else {
			//free block
			currNode = advanceLinkedList(currNode, ".", blockSize)
		}
	}
	//CurrNode should be at the tail
	for currNode != head {
		if currNode.value != "." {
			//Find first free node
			freeNode := head
			noFreeSpace := false
			for freeNode.value != "." {
				freeNode = freeNode.next
				// //escape case
				if freeNode == currNode {
					noFreeSpace = true
					break
				}
			}
			if noFreeSpace {
				break
			}
			//Swap node positions
			tempNode := *currNode
			currNode.next = freeNode.next
			currNode.prev = freeNode.prev
			if freeNode.prev != nil {
				freeNode.prev.next = currNode
			}
			if freeNode.next != nil {
				freeNode.next.prev = currNode
			}
			freeNode.prev = tempNode.prev
			freeNode.next = tempNode.next
			if tempNode.prev != nil {
				tempNode.prev.next = freeNode
			}
			if tempNode.next != nil {
				tempNode.next.prev = freeNode
			}
			currNode = freeNode
		} else {
			currNode = currNode.prev
		}
	}

	result := 0
	curr := head
	index := 0
	for curr.value != "." {
		number, err := strconv.Atoi(curr.value)
		if err != nil {
			log.Fatal("Error parsing int from string")
		}
		result += (number * index)
		curr = curr.next
		index++
	}
	// printLinkedList(head)

	return result
}

func printLinkedList(head *Node) {
	printString := ""
	currNode := head
	for currNode != nil {
		printIndex := 0
		for printIndex < currNode.blockSize {
			printString += currNode.value
			printIndex++
		}
		currNode = currNode.next
	}
	log.Println(printString)
}

func advanceLinkedList(currNode *Node, value string, blockSize int) *Node {
	currIncrement := 0
	for currIncrement < blockSize {
		new := newNode(value, nil, currNode, 1)
		currNode.next = new
		currNode = new
		currIncrement++
	}
	return currNode
}

func partTwo(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	blocks := strings.Split(scanner.Text(), "")
	var head *Node
	// var tail *Node
	var currNode *Node
	currId := 0
	for index, blockString := range blocks {
		blockSize, err := strconv.Atoi(blockString)
		if blockSize == 0 {
			continue
		}
		if err != nil {
			log.Fatal("error parsing int from string")
		}
		if index == 0 {
			//Initial Block
			head = newNode(strconv.Itoa(currId), nil, nil, blockSize)
			currNode = head
			currId++
		} else if index%2 == 0 {
			//File block
			new := newNode(strconv.Itoa(currId), nil, currNode, blockSize)
			currNode.next = new
			currNode = new
			currId++
		} else {
			//free block
			new := newNode(".", nil, currNode, blockSize)
			currNode.next = new
			currNode = new
		}
	}
	//CurrNode should be at the tail
	for currNode != head {
		if currNode.next != nil {
			log.Printf("Node Next value: %s \n", currNode.next.value)
		}
		if currNode.prev != nil {
			log.Printf("Node Prev value: %s \n", currNode.prev.value)
		}
		if currNode.value != "." {
			//Find first free block
			freeNode := head
			noFreeSpace := false
			for {
				if freeNode.value == "." && freeNode.blockSize >= currNode.blockSize {
					break
				}
				freeNode = freeNode.next
				// could not find enough free space
				if freeNode == currNode {
					noFreeSpace = true
					break
				}
			}
			if noFreeSpace {
				currNode = currNode.prev
				continue
			}
			//found free block that has enough size for new block
			//Swap node positions
			tempNode := *currNode
			difference := freeNode.blockSize - currNode.blockSize
			//if there is remaining free space
			if difference > 0 {
				remainderNode := newNode(".", freeNode.next, currNode, difference)
				currNode.next = remainderNode
				currNode.prev = freeNode.prev
				if freeNode.prev != nil {
					freeNode.prev.next = currNode
				}
				if freeNode.next != nil {
					freeNode.next.prev = remainderNode
				}

			} else {
				currNode.next = freeNode.next
				currNode.prev = freeNode.prev
				if freeNode.prev != nil {
					freeNode.prev.next = currNode
				}
				if freeNode.next != nil {
					freeNode.next.prev = currNode
				}
			}

			freeNode.prev = tempNode.prev
			freeNode.next = tempNode.next
			freeNode.blockSize = currNode.blockSize
			if tempNode.prev != nil {
				tempNode.prev.next = freeNode
			}
			if tempNode.next != nil {
				tempNode.next.prev = freeNode
			}
			currNode = freeNode
		} else {
			currNode = currNode.prev
		}
	}

	printLinkedList(head)
	result := 0
	curr := head
	index := 0
	for curr != nil {
		if curr.value != "." {
			number, err := strconv.Atoi(curr.value)
			if err != nil {
				log.Fatal("Error parsing int from string")
			}
			blockIndex := 0
			for blockIndex < curr.blockSize {
				result += (number * (blockIndex + index))
				blockIndex++
			}
			index += curr.blockSize
		} else {
			index += curr.blockSize
		}
		curr = curr.next
	}

	return result
}
