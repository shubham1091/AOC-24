package day03

import (
	"AOC-24/utils"
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

func Run() {
	// Fetch the input data
	data, err := utils.FetchInput(3)
	if err != nil {
		fmt.Println("Error fetching input:", err)
		return
	}

	// Initialize scanner
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// Calculate and sum the results from valid mul instructions
	solutionOne := PartOne(input)
	fmt.Println("Part 1 - Sum of all mul results:", solutionOne)

	// Calculate and sum the results from valid mul instructions, considering do() and don't()
	solutionTwo := PartTwo(input)
	fmt.Println("Part 2 - Sum of all enabled mul results:", solutionTwo)
}

func PartOne(input []string) int {
	// Regex to match mul(x,y) pattern
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	var totalSum int

	// Iterate over each line in the input
	for _, line := range input {
		// Find all mul() matches in the current line
		matches := mulRegex.FindAllStringSubmatch(line, -1)

		// Sum the results of the mul operations
		for _, match := range matches {
			// Convert the numbers to integers and calculate their product
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			totalSum += a * b
		}
	}

	return totalSum
}

func PartTwo(input []string) int {
	commandRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	var total int
	enabled := true

	// Process each line in the input
	for _, line := range input {
		matches := commandRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			switch match[0] {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if enabled {
					a, _ := strconv.Atoi(match[1])
					b, _ := strconv.Atoi(match[2])
					total += a * b
				}
			}
		}
	}
	return total
}
