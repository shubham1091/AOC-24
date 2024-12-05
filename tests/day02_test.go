package tests

import (
	"AOC-24/solutions"
	"reflect"
	"strconv"
	"testing"
)

func TestDayTwoPartOne(t *testing.T) {
	type testCase struct {
		input       []string
		expected    int
		description string
		expectError bool
	}

	testCases := []testCase{
		{
			input: []string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			expected:    2,
			description: "Standard example with mixed safe and unsafe reports.",
			expectError: false,
		},
		{
			input: []string{
				"1 1 1 1 1",
				"5 4 3 2 1",
				"1 4 7 10 13",
			},
			expected:    2,
			description: "Edge cases with constant, decreasing, and increasing levels.",
			expectError: false,
		},
	}

	for i, tc := range testCases {
		t.Run("DayTwoPartOne_"+strconv.Itoa(i)+"_"+tc.description, func(t *testing.T) {
			day02 := solutions.Day02{}
			safeCount, err := day02.PartOne(tc.input)
			if err != nil {
				t.Errorf("Error occurred: %v", err)
				return
			}
			if !reflect.DeepEqual(safeCount, tc.expected) {
				t.Errorf("Failed %s: Expected %d, got %d", tc.description, tc.expected, safeCount)
			}
		})
	}
}

func TestDayTwoPartTwo(t *testing.T) {
	type testCase struct {
		input       []string
		expected    int
		description string
		expectError bool
	}

	testCases := []testCase{
		{
			input: []string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			expected:    4,
			description: "Standard example with Problem Dampener adjustments.",
			expectError: false,
		},
		{
			input: []string{
				"2 5 8 11 14",
				"10 8 6 5 3",
				"3 3 3 3 3",
			},
			expected:    2,
			description: "Edge cases with large increases, decreases, and constant levels.",
			expectError: false,
		},
	}

	for i, tc := range testCases {
		t.Run("DayTwoPartTwo_"+strconv.Itoa(i)+"_"+tc.description, func(t *testing.T) {
			day02 := solutions.Day02{}
			safeCountWithDampener, err := day02.PartTwo(tc.input)
			if err != nil {
				t.Errorf("Error occurred: %v", err)
				return
			}
			if !reflect.DeepEqual(safeCountWithDampener, tc.expected) {
				t.Errorf("Failed %s: Expected %d, got %d", tc.description, tc.expected, safeCountWithDampener)
			}
		})
	}
}
