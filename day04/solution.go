package day04

import (
	"AOC-24/utils"
	"fmt"
)

func Run() {
	// Fetch the input data
	data, err := utils.FetchInput(4)
	if err != nil {
		fmt.Println("Error fetching input:", err)
		return
	}

	// Parse input data into a slice of strings
	grid := utils.ParseInput(data)

	// Count occurrences of "XMAS"
	count := PartOne(grid)
	fmt.Println("Number of times XMAS appears:", count)

	// Solve Part Two
	countTwo := PartTwo(grid)
	fmt.Println("Number of times X-MAS appears:", countTwo)
}

func PartOne(grid []string) int {
	word := "XMAS"
	wordLen := len(word)
	count := 0

	// All eight directions: right, down, diagonal down-right, diagonal up-right,
	// and their reverse counterparts (left, up, diagonal down-left, diagonal up-left)
	directions := [][2]int{
		{0, 1},   // right
		{1, 0},   // down
		{1, 1},   // diagonal down-right
		{-1, 1},  // diagonal up-right
		{0, -1},  // left
		{-1, 0},  // up
		{-1, -1}, // diagonal up-left
		{1, -1},  // diagonal down-left
	}

	// Traverse each cell in the grid
	for x := range grid {
		for y := range grid[x] {
			// Check each direction from the current cell
			for _, dir := range directions {
				if checkWord(grid, word, x, y, dir[0], dir[1], wordLen) {
					count++
				}
			}
		}
	}

	return count
}

func checkWord(grid []string, word string, x, y, dx, dy, wordLen int) bool {
	for i := 0; i < wordLen; i++ {
		nx, ny := x+i*dx, y+i*dy
		// Check bounds and character match
		if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) || grid[nx][ny] != word[i] {
			return false
		}
	}
	return true
}

func PartTwo(grid []string) int {
	count := 0

	// Traverse each cell in the grid
	for x, row := range grid {
		for y := range row {
			// Check for "X-MAS" pattern centered at (x, y)
			if isXMas(grid, x, y) {
				count++
			}
		}
	}

	return count
}

func isXMas(grid []string, x, y int) bool {
	// First, check if center is 'A'
	if !inBounds(grid, x, y) || grid[x][y] != 'A' {
		return false
	}

	// Check both diagonals for MAS pattern
	// Positions relative to center A: up-left, up-right, down-left, down-right
	positions := [][2]int{
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}

	// Check if we can form valid X-MAS pattern

	return checkDiagonal(grid, x, y, positions[1], positions[2]) && checkDiagonal(grid, x, y, positions[0], positions[3])
}

func checkDiagonal(grid []string, centerX, centerY int, pos1, pos2 [2]int) bool {
	// Get positions for the diagonal
	x1, y1 := centerX+pos1[0], centerY+pos1[1]
	x2, y2 := centerX+pos2[0], centerY+pos2[1]

	if !inBounds(grid, x1, y1) || !inBounds(grid, x2, y2) {
		return false
	}

	// Check both forward and backward MAS patterns
	return (grid[x1][y1] == 'M' && grid[x2][y2] == 'S') ||
		(grid[x1][y1] == 'S' && grid[x2][y2] == 'M')
}

func inBounds(grid []string, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
}
