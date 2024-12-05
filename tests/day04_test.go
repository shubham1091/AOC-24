package tests

import (
	"AOC-24/solutions"
	"reflect"
	"strconv"
	"testing"
)

func TestDayFourPartOne(t *testing.T) {

	type testCase struct {
		input       []string
		expected    int
		description string
		expectError bool
	}

	testCases := []testCase{
		{
			input: []string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			expected:    18,
			description: "Basic grid with multiple horizontal, vertical, and diagonal occurrences of 'XMAS'.",
			expectError: false,
		},
		{
			input:       []string{},
			expected:    0,
			description: "Empty grid input should return 0 occurrences.",
			expectError: false,
		},
		{
			input: []string{
				"MMMMMMMMMM",
				"MMMMMMMMMM",
				"MMMMMMMMMM",
				"MMMMMMMMMM",
				"MMMMMMMMMM",
			},
			expected:    0,
			description: "Grid with no occurrence of 'XMAS' should return 0.",
			expectError: false,
		},
		{
			input: []string{
				"XMAS",
				"SXMA",
				"AMXS",
				"MASX",
				"SSMA",
			},
			expected:    2,
			description: "Small grid with overlapping occurrences of 'XMAS'.",
			expectError: false,
		},
	}

	for i, tc := range testCases {
		t.Run("Day Four PartOne Test Case "+strconv.Itoa(i)+": "+tc.description, func(t *testing.T) {
			day04 := solutions.Day04{}
			result, err := day04.PartOne(tc.input)
			if err != nil {
				t.Errorf("Error occurred: %v", err)
				return
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("For input %v, expected %v but got %d", tc.input, tc.expected, result)
			}
		})
	}
}

func TestDayFourPartTwo(t *testing.T) {

	type testCase struct {
		input       []string
		expected    int
		description string
		expectError bool
	}

	testCases := []testCase{
		{
			input: []string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			expected:    9,
			description: "Basic grid with multiple 'X-MAS' patterns in an 'X' shape.",
			expectError: false,
		},
		{
			input:       []string{},
			expected:    0,
			description: "Empty grid input should return 0 'X-MAS' patterns.",
			expectError: false,
		},
		{
			input: []string{
				"MMMMMMMMMM",
				"MMMMMMMMMM",
				"MMMMMMMMMM",
				"MMMMMMMMMM",
				"MMMMMMMMMM",
			},
			expected:    0,
			description: "Grid with no occurrence of 'X-MAS' patterns should return 0.",
			expectError: false,
		},
		{
			input: []string{
				"MASMASMAS",
				"MASMASMAS",
				"MASMASMAS",
				"MASMASMAS",
				"MASMASMAS",
			},
			expected:    9,
			description: "Grid with only horizontal and vertical 'MAS'.",
			expectError: false,
		},
		{
			input: []string{
				"M.S.M",
				".A.A.",
				"M.S.M",
				".A.A.",
				"M.S.M",
			},
			expected:    4,
			description: "Small grid with overlapping 'X-MAS' patterns in an 'X' shape.",
			expectError: false,
		},
	}

	for i, tc := range testCases {
		t.Run("Day Four PartTwo Test Case "+strconv.Itoa(i)+": "+tc.description, func(t *testing.T) {
			day04 := solutions.Day04{}
			result, err := day04.PartTwo(tc.input)
			if err != nil {
				t.Errorf("Error occurred: %v", err)
				return
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("For input %v, expected %v but got %v", tc.input, tc.expected, result)
			}
		})
	}
}
