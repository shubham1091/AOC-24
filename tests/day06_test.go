package tests

import (
	"AOC-24/solutions"
	"testing"
)

func TestDaySixPartOne(t *testing.T) {
	type testCase struct {
		input       []string
		expected    int
		description string
	}

	testCases := []testCase{
		{
			input: []string{
				"....#.....",
				".........#",
				"..........",
				"..#.......",
				".......#..",
				"..........",
				".#..^.....",
				"........#.",
				"#.........",
				"......#...",
			},
			expected:    41,
			description: "Test basic input with guard facing up.",
		},
	}

	for i, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			day06 := solutions.Day06{}
			result, err := day06.PartOne(tc.input)
			if err != nil {
				t.Errorf("Unexpected error for test case %d: %v", i, err)
			}
			if result != tc.expected {
				t.Errorf("Failed %s: Expected %d, got %d", tc.description, tc.expected, result)
			}
		})
	}
}


func TestDaySixPartTwo(t *testing.T) {
	type testCase struct {
		input       []string
		expected    int
		description string
	}

	testCases := []testCase{
		{
			input: []string{
				"....#.....",
				".........#",
				"..........",
				"..#.......",
				".......#..",
				"..........",
				".#..^.....",
				"........#.",
				"#.........",
				"......#...",
			},
			expected:    6,
			description: "Test basic input with guard facing up.",
		},
	}

	for i, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			day06 := solutions.Day06{}
			result, err := day06.PartTwo(tc.input)
			if err != nil {
				t.Errorf("Unexpected error for test case %d: %v", i, err)
			}
			if result != tc.expected {
				t.Errorf("Failed %s: Expected %d, got %d", tc.description, tc.expected, result)
			}
		})
	}
}
