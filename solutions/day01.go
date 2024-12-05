package solutions

import (
	"AOC-24/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Day01 represents the solution for Day 01
type Day01 struct{}

// Solve implements the main solution entry point
func (d *Day01) Solve() (interface{}, error) {
	// Fetch and parse input data
	data, err := utils.FetchInput(1)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch input: %w", err)
	}

	input := utils.ParseInput(data)
	parsedLeft, parsedRight, parseErr := parseInput(input)
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing input: %w", parseErr)
	}

	// Calculate results for both parts
	totalDistance := calculateTotalDistance(parsedLeft, parsedRight)
	similarityScore := calculateSimilarityScore(parsedLeft, parsedRight)

	return map[string]int{
		"TotalDistance":   totalDistance,
		"SimilarityScore": similarityScore,
	}, nil
}

// PartOne solves the first part of the problem
func (d *Day01) PartOne(input []string) (interface{}, error) {
	leftList, rightList, err := parseInput(input)
	if err != nil {
		return nil, fmt.Errorf("error parsing input for PartOne: %w", err)
	}
	return calculateTotalDistance(leftList, rightList), nil
}

// PartTwo solves the second part of the problem
func (d *Day01) PartTwo(input []string) (interface{}, error) {
	leftList, rightList, err := parseInput(input)
	if err != nil {
		return nil, fmt.Errorf("error parsing input for PartTwo: %w", err)
	}
	return calculateSimilarityScore(leftList, rightList), nil
}

// parseInput parses the input lines into two integer slices
func parseInput(input []string) ([]int, []int, error) {
	var leftList, rightList []int

	for _, line := range input {
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		left, err1 := strconv.Atoi(numbers[0])
		right, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("error parsing numbers: %v, %v", err1, err2)
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}
	return leftList, rightList, nil
}

// calculateTotalDistance computes the total Manhattan distance between two sorted lists
func calculateTotalDistance(leftList, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += abs(leftList[i] - rightList[i])
	}
	return totalDistance
}

// calculateSimilarityScore computes the weighted similarity score between two lists
func calculateSimilarityScore(leftList, rightList []int) int {
	sort.Ints(rightList)

	// Build frequency map for the right list
	freq := make(map[int]int)
	for _, num := range rightList {
		freq[num]++
	}

	// Calculate similarity score
	similarityScore := 0
	for _, num := range leftList {
		similarityScore += num * freq[num]
	}
	return similarityScore
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
