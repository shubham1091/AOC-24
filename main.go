package main

import (
	"AOC-24/day01"
	"AOC-24/day02"
	"AOC-24/day03"
	"AOC-24/day04"
	"AOC-24/day05"
	"AOC-24/utils"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go [day]")
	}

	utils.LoadEnv()

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid day: %v", err)
	}

	// Check if the day is between 1 and 25
	if day < 1 || day > 25 {
		log.Fatalf("Day must be between 1 and 25, got %d", day)
	}

	switch day {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	case 3:
		day03.Run()
	case 4:
		day04.Run()
	case 5:
		day05.Run()
	default:
		log.Fatalf("Day %d not implemented", day)
	}
}
