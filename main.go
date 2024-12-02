package main

import (
	"flag"
	"log"

	"github.com/DevinSchwehr/AdventOfCode2024/day1"
)

func main() {
	var day int
	flag.IntVar(&day, "day", 1, "day number")

	flag.Parse()

	if day < 1 || day > 25 {
		log.Fatalf("invalid day value, must be 1 through 25 but got %v", day)
	}

	day1.Execute()
}
