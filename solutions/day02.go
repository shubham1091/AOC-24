package solutions

import (
	"AOC-24/utils"
	"fmt"
	"strconv"
	"strings"
)

// Day02 represents the solution for Day 02
type Day02 struct{}

// Solve executes both parts of the day's problem
func (d *Day02) Solve() (interface{}, error) {
	// Fetch input data
	data, err := utils.FetchInput(2)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch input: %w", err)
	}

	// Parse input
	input := utils.ParseInput(data)

	// Solve both parts
	partOneResult, err := d.PartOne(input)
	if err != nil {
		return nil, fmt.Errorf("failed to solve part one: %w", err)
	}

	partTwoResult, err := d.PartTwo(input)
	if err != nil {
		return nil, fmt.Errorf("failed to solve part two: %w", err)
	}

	// Return combined results
	return map[string]interface{}{
		"Part One": partOneResult,
		"Part Two": partTwoResult,
	}, nil
}

// PartOne calculates the number of safe reports
func (d *Day02) PartOne(input []string) (interface{}, error) {
	safeCount := 0
	for _, line := range input {
		numbers, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse line: %w", err)
		}
		if isSafe(numbers) {
			safeCount++
		}
	}
	return safeCount, nil
}

// PartTwo calculates the number of safe reports including dampener logic
func (d *Day02) PartTwo(input []string) (interface{}, error) {
	safeCountWithDampener := 0
	for _, line := range input {
		numbers, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse line: %w", err)
		}
		if isSafe(numbers) || isSafeWithDampener(numbers) {
			safeCountWithDampener++
		}
	}
	return safeCountWithDampener, nil
}

// parseLine converts a line of numbers into a slice of integers
func parseLine(line string) ([]int, error) {
	fields := strings.Fields(line)
	numbers := make([]int, len(fields))

	for i, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, fmt.Errorf("invalid number '%s': %w", field, err)
		}
		numbers[i] = num
	}
	return numbers, nil
}

// isSafe checks if a sequence of numbers is "safe"
func isSafe(numbers []int) bool {

	// Check first difference to determine if we should check for increasing or decreasing
	diff := numbers[1] - numbers[0]

	if diff == 0 {
		return false // Adjacent numbers cannot be equal
	}

	isIncreasing := diff > 0

	// Check all adjacent pairs
	for i := 1; i < len(numbers); i++ {
		currentDiff := numbers[i] - numbers[i-1]

		// Check if difference is between 1 and 3 (inclusive)
		if abs(currentDiff) < 1 || abs(currentDiff) > 3 {
			return false
		}

		// For increasing sequence, all differences must be positive
		if isIncreasing && currentDiff <= 0 {
			return false
		}
		// For decreasing sequence, all differences must be negative
		if !isIncreasing && currentDiff >= 0 {
			return false
		}
	}

	return true
}

// isSafeWithDampener checks if a sequence is "safe" after removing one number
func isSafeWithDampener(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		temp := make([]int, 0, len(numbers)-1)
		temp = append(temp, numbers[:i]...)
		temp = append(temp, numbers[i+1:]...)
		if isSafe(temp) {
			return true
		}
	}
	return false
}
