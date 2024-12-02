package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}
func main() {
	data, err := fetchInput()
	if err != nil {
		fmt.Println("Error fetching input:", err)
		return
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))

	var safeCount, safeCountWithDampener int
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)

		var data []int
		for _, number := range numbers {
			n, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				return
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

	fmt.Println("Part 1 - Number of safe reports:", safeCount)
	fmt.Println("Part 2 - Number of safe reports with dampener:", safeCountWithDampener)
}

func fetchInput() ([]byte, error) {
	sessionCookie := os.Getenv("AOC_SESSION")
	if sessionCookie == "" {
		return nil, fmt.Errorf("AOC_SESSION environment variable not set")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://adventofcode.com/2024/day/2/input", nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
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

// Helper function to get absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
