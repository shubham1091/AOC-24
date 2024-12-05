package solutions

import (
	"AOC-24/utils"
	"fmt"
	"strconv"
	"strings"
)

type Day05 struct{}

func (d *Day05) Solve() (interface{}, error) {
	// Fetch input data
	data, err := utils.FetchInput(5)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch input: %w", err)
	}
	grid := utils.ParseInput(data)

	// Solve both parts
	partOneResult, err := d.PartOne(grid)
	if err != nil {
		return nil, fmt.Errorf("failed to solve part one: %w", err)
	}

	partTwoResult, err := d.PartTwo(grid)
	if err != nil {
		return nil, fmt.Errorf("failed to solve part two: %w", err)
	}

	// Return combined results
	return map[string]interface{}{
		"Part One": partOneResult,
		"Part Two": partTwoResult,
	}, nil
}

// ProcessData separates the parsing logic for better organization
func ProcessData(input []string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	var updates [][]int
	parsingRules := true

	for _, line := range input {
		if line == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			if x, y, ok := ParseRule(line); ok {
				rules[x] = append(rules[x], y)
			}
		} else {
			if update := ParseUpdate(line); len(update) > 0 {
				updates = append(updates, update)
			}
		}
	}
	return rules, updates
}

func ParseRule(line string) (int, int, bool) {
	parts := strings.Split(line, "|")
	if len(parts) != 2 {
		return 0, 0, false
	}
	x, errx := strconv.Atoi(strings.TrimSpace(parts[0]))
	y, erry := strconv.Atoi(strings.TrimSpace(parts[1]))
	if errx != nil || erry != nil {
		return 0, 0, false
	}
	return x, y, true
}

func ParseUpdate(line string) []int {
	nums := strings.Split(line, ",")
	update := make([]int, 0, len(nums))
	for _, num := range nums {
		if n, err := strconv.Atoi(strings.TrimSpace(num)); err == nil {
			update = append(update, n)
		}
	}
	return update
}

func (d *Day05) PartOne(input []string) (interface{}, error) {
	rules, updates := ProcessData(input)
	sum := 0
	for _, update := range updates {
		if isValidOrder(update, rules) {
			sum += update[len(update)/2]
		}
	}
	return sum, nil
}

func (d *Day05) PartTwo(input []string) (interface{}, error) {
	rules, updates := ProcessData(input)
	sum := 0
	for _, update := range updates {
		if !isValidOrder(update, rules) {
			sortedUpdate := correctOrder(update, rules)
			sum += sortedUpdate[len(sortedUpdate)/2]
		}
	}
	return sum, nil
}

func correctOrder(update []int, rules map[int][]int) []int {
	graph := make(map[int][]int, len(update))
	inDegree := make(map[int]int, len(update))

	// Initialize maps with capacity
	for _, page := range update {
		inDegree[page] = 0
		graph[page] = make([]int, 0, len(update))
	}

	// Build graph
	for page, deps := range rules {
		if _, exists := inDegree[page]; exists {
			for _, dep := range deps {
				if _, exists := inDegree[dep]; exists {
					graph[page] = append(graph[page], dep)
					inDegree[dep]++
				}
			}
		}
	}

	// Perform topological sort
	result := make([]int, 0, len(update))
	queue := make([]int, 0, len(update))

	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	for len(queue) > 0 {
		page := queue[0]
		queue = queue[1:]
		result = append(result, page)

		for _, neighbor := range graph[page] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

func isValidOrder(update []int, rules map[int][]int) bool {
	positions := make(map[int]int, len(update))
	for i, page := range update {
		positions[page] = i
	}

	for page, dependencies := range rules {
		if pos1, exists1 := positions[page]; exists1 {
			for _, mustComeAfter := range dependencies {
				if pos2, exists2 := positions[mustComeAfter]; exists2 && pos2 <= pos1 {
					return false
				}
			}
		}
	}
	return true
}
