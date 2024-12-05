package day04

import (
	"reflect"
	"strconv"
	"testing"
)

func TestPartOne(t *testing.T) {
	type testCase struct {
		input    []string
		expected int
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
			expected: 18,
		},
	}

	for i, tc := range testCases {
		t.Run("PartOneTestCase"+strconv.Itoa(i), func(t *testing.T) {
			results := PartOne(tc.input)
			if !reflect.DeepEqual(results, tc.expected) {
				t.Errorf("For input %v, expected %v but got %v", tc.input, tc.expected, results)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	type testCase struct {
		input    []string
		expected int
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
			expected: 9,
		},
	}

	for i, tc := range testCases {
		t.Run("PartOneTestCase"+strconv.Itoa(i), func(t *testing.T) {
			results := PartTwo(tc.input)
			if !reflect.DeepEqual(results, tc.expected) {
				t.Errorf("For input %v, expected %v but got %v", tc.input, tc.expected, results)
			}
		})
	}
}
