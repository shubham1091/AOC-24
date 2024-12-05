package solutions

import (
	"AOC-24/utils"
	"fmt"
	"regexp"
	"strconv"
)

// Day03 represents the solution for Day 03
type Day03 struct{}

// Solve executes both parts of the day's problem
func (d *Day03) Solve() (interface{}, error) {
	// Fetch input data
	data, err := utils.FetchInput(3)
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

// PartOne sums all results from valid `mul(x,y)` instructions
func (d *Day03) PartOne(input []string) (interface{}, error) {
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	totalSum := 0

	for _, line := range input {
		matches := mulRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, err1 := strconv.Atoi(match[1])
			b, err2 := strconv.Atoi(match[2])
			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("invalid numbers in line: %s", line)
			}
			totalSum += a * b
		}
	}
	return totalSum, nil
}

// PartTwo sums the results from `mul(x,y)` instructions, considering `do()` and `don't()`
func (d *Day03) PartTwo(input []string) (interface{}, error) {
	commandRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	totalSum := 0
	enabled := true

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
					a, err1 := strconv.Atoi(match[1])
					b, err2 := strconv.Atoi(match[2])
					if err1 != nil || err2 != nil {
						return nil, fmt.Errorf("invalid numbers in line: %s", line)
					}
					totalSum += a * b
				}
			}
		}
	}
	return totalSum, nil
}
