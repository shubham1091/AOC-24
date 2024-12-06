package main

import (
	"AOC-24/solutions"
	"AOC-24/utils"
	"log"
	"os"
	"strconv"
)

func main() {
	// Ensure a day is provided as a command-line argument
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go [day]")
	}

	// Load environment variables
	utils.LoadEnv()

	// Parse the day argument to an integer
	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 || day > 25 {
		log.Fatalf("Invalid day: %v. Please provide a day between 1 and 25.", err)
	}

	// Map of day numbers to their corresponding DaySolver implementations
	daySolvers := map[int]utils.DaySolver{
		1: &solutions.Day01{},
		2: &solutions.Day02{},
		3: &solutions.Day03{},
		4: &solutions.Day04{},
		5: &solutions.Day05{},
		6: &solutions.Day06{},
	}

	// Check if the solver for the provided day exists
	daySolver, exists := daySolvers[day]
	if !exists {
		log.Fatalf("Day %d not implemented", day)
	}

	// Solve the problem using the corresponding solver
	result, err := daySolver.Solve()
	if err != nil {
		log.Fatalf("Error solving day %d: %v", day, err)
	}

	// Print the result
	log.Printf("Solution for Day %d:\n%v", day, result)
}
