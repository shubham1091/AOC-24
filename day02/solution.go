package day02

import (
	"AOC-24/utils"
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// day02/solution.go
func Run() {
	data, err := utils.FetchInput(2)
	if err != nil {
		fmt.Println("Error fetching input:", err)
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	safeCount, safeCountWithDampener := processSafeData(input)

	fmt.Println("Part 1 - Number of safe reports:", safeCount)
	fmt.Println("Part 2 - Number of safe reports with dampener:", safeCountWithDampener)
}

func isSafeWithDampener(numbers []int) bool {
	// Try removing each number one at a time
	for i := 0; i < len(numbers); i++ {
		// Create a new slice without the current number
		temp := make([]int, 0, len(numbers)-1)
		temp = append(temp, numbers[:i]...)
		temp = append(temp, numbers[i+1:]...)

		if isSafe(temp) {
			return true
		}
	}
	return false
}

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

// day02/solution.go
func processSafeData(input []string) (int, int) {
	var safeCount, safeCountWithDampener int
	for _, line := range input {
		numbers := strings.Fields(line)

		var data []int
		for _, number := range numbers {
			n, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				continue
			}
			data = append(data, n)
		}

		if isSafe(data) {
			safeCount++
			safeCountWithDampener++
		} else if isSafeWithDampener(data) {
			safeCountWithDampener++
		}
	}

	return safeCount, safeCountWithDampener
}

// Helper function to get absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
