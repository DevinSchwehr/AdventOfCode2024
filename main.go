package main

import (
	"flag"
	"log"

	"github.com/DevinSchwehr/AdventOfCode2024/day1"
	"github.com/DevinSchwehr/AdventOfCode2024/day10"
	"github.com/DevinSchwehr/AdventOfCode2024/day2"
	"github.com/DevinSchwehr/AdventOfCode2024/day3"
	"github.com/DevinSchwehr/AdventOfCode2024/day4"
	"github.com/DevinSchwehr/AdventOfCode2024/day5"
	"github.com/DevinSchwehr/AdventOfCode2024/day6"
	"github.com/DevinSchwehr/AdventOfCode2024/day7"
	"github.com/DevinSchwehr/AdventOfCode2024/day8"
	"github.com/DevinSchwehr/AdventOfCode2024/day9"
)

func main() {
	var day int
	flag.IntVar(&day, "day", 1, "day number")

	flag.Parse()

	if day < 1 || day > 25 {
		log.Fatalf("invalid day value, must be 1 through 25 but got %v", day)
	}

	switch day {
	case 1:
		day1.Execute()
	case 2:
		day2.Execute()
	case 3:
		day3.Execute()
	case 4:
		day4.Execute()
	case 5:
		day5.Execute()
	case 6:
		day6.Execute()
	case 7:
		day7.Execute()
	case 8:
		day8.Execute()
	case 9:
		day9.Execute()
	case 10:
		day10.Execute()
	}
}
