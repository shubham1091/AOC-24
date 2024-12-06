package solutions

import (
	"AOC-24/utils"
	"fmt"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

// State represents the robot's current state
type State struct {
	Row, Col  int
	Direction Direction
}

type Day06 struct{}

type Movement struct {
	dx, dy int
}

type GuardSimulation struct {
	grid       [][]rune
	guardX     int
	guardY     int
	guardDir   int
	directions []Movement
}

func (d *Day06) Solve() (interface{}, error) {
	data, err := utils.FetchInput(6)
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

func (d Day06) PartOne(input []string) (interface{}, error) {
	directions := []Movement{
		{dx: -1, dy: 0}, // Up
		{dx: 0, dy: 1},  // Right
		{dx: 1, dy: 0},  // Down
		{dx: 0, dy: -1}, // Left
	}

	grid, guardX, guardY, guardDir := parseLabMap(input)

	simulation := GuardSimulation{
		grid:       grid,
		guardX:     guardX,
		guardY:     guardY,
		guardDir:   guardDir,
		directions: directions,
	}

	visited := simulation.simulateMovement()

	return len(visited), nil
}

func (d Day06) PartTwo(input []string) (interface{}, error) {
	directions := []Movement{
		{dx: -1, dy: 0}, // Up
		{dx: 0, dy: 1},  // Right
		{dx: 1, dy: 0},  // Down
		{dx: 0, dy: -1}, // Left
	}

	grid, guardX, guardY, guardDir := parseLabMap(input)

	simulation := GuardSimulation{
		grid:       grid,
		guardX:     guardX,
		guardY:     guardY,
		guardDir:   guardDir,
		directions: directions,
	}

	// Find all possible positions where an obstruction would create a loop
	possibleObstructions := findObstructionPositions(simulation)

	return possibleObstructions, nil
}

func findObstructionPositions(sim GuardSimulation) int {
	obstructionPositions := make(map[[2]int]bool)
	mapHeight := len(sim.grid)
	mapWidth := len(sim.grid[0])

	for x := 0; x < mapHeight; x++ {
		for y := 0; y < mapWidth; y++ {
			// Skip positions that are walls or the guard's starting position
			if sim.grid[x][y] == '#' || (x == sim.guardX && y == sim.guardY) {
				continue
			}

			// Temporarily place an obstruction
			sim.grid[x][y] = '#'

			// Check if the guard gets stuck in a loop
			if causesLoop(sim) {
				obstructionPositions[[2]int{x, y}] = true
			}

			// Remove the obstruction
			sim.grid[x][y] = '.'
		}
	}

	return len(obstructionPositions)
}

func causesLoop(sim GuardSimulation) bool {
	rows := len(sim.grid)
	cols := len(sim.grid[0])

	row, col := sim.guardX, sim.guardY
	direction := sim.guardDir

	visitedStates := make(map[struct{ x, y, d int }]bool)

	for {
		currentState := struct{ x, y, d int }{row, col, direction}

		// If we've seen this state before, we found a loop
		if visitedStates[currentState] {
			return true
		}

		visitedStates[currentState] = true

		nextRow := row + sim.directions[direction].dx
		nextCol := col + sim.directions[direction].dy

		// Check if guard would exit the grid
		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
			return false
		}

		// If there's an obstacle, turn right
		if sim.grid[nextRow][nextCol] == '#' {
			direction = (direction + 1) % 4
		} else {
			row, col = nextRow, nextCol
		}

		// Prevent infinite execution for paths that don't loop
		if len(visitedStates) > rows*cols*4 {
			return false
		}
	}
}

func (sim *GuardSimulation) simulateMovement() map[[2]int]bool {
	visited := make(map[[2]int]bool)
	visited[[2]int{sim.guardX, sim.guardY}] = true
	mapHeight := len(sim.grid)
	mapWidth := len(sim.grid[0])

	for {
		nextX := sim.guardX + sim.directions[sim.guardDir].dx
		nextY := sim.guardY + sim.directions[sim.guardDir].dy

		if nextX < 0 || nextX >= mapHeight || nextY < 0 || nextY >= mapWidth {
			break
		}

		if sim.grid[nextX][nextY] == '#' {
			sim.guardDir = (sim.guardDir + 1) % 4
		} else {
			visited[[2]int{nextX, nextY}] = true
			sim.guardX, sim.guardY = nextX, nextY
		}
	}

	return visited
}

func parseLabMap(input []string) ([][]rune, int, int, int) {
	grid := make([][]rune, len(input))
	var guardX, guardY, guardDir int

	for i, line := range input {
		grid[i] = []rune(line)
		for j, char := range line {
			switch char {
			case '^':
				guardX, guardY, guardDir = i, j, 0
			case '>':
				guardX, guardY, guardDir = i, j, 1
			case 'v':
				guardX, guardY, guardDir = i, j, 2
			case '<':
				guardX, guardY, guardDir = i, j, 3
			}
		}
	}

	return grid, guardX, guardY, guardDir
}
