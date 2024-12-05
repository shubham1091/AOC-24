package tests

import (
	"AOC-24/solutions"
	"reflect"
	"strconv"
	"testing"
)

func TestDayOnePartOne(t *testing.T) {
	type testCase struct {
		input       []string
		expected    int
		description string
		expectError bool
	}
	testCases := []testCase{
		{
			input: []string{
				"3 4",
				"4 3",
				"2 5",
				"1 3",
				"3 9",
				"3 3",
			},
			expected:    11,
			description: "Calculates total distance for a standard list with duplicates and mixed numbers.",
			expectError: false,
		},
		{
			input:       []string{},
			expected:    0,
			description: "Handles empty input gracefully by returning zero distance.",
			expectError: false,
		},
	}

	for i, tc := range testCases {
		t.Run("TestDayOnePartOne_"+strconv.Itoa(i), func(t *testing.T) {
			day01 := solutions.Day01{}
			totalDistance, err := day01.PartOne(tc.input)
			if err != nil {
				t.Errorf("Error occurred: %v", err)
				return
			}
			if !reflect.DeepEqual(totalDistance, tc.expected) {
				t.Errorf("Failed %s: For input %v, Expected total distance to be %v, but got %d", tc.description, tc.input, tc.expected, totalDistance)
			}
		})
	}
}

func TestDayOnePartTwo(t *testing.T) {
	type testCase struct {
		input       []string
		expected    int
		description string
		expectError bool
	}
	testCases := []testCase{
		{
			input: []string{
				"3 4",
				"4 3",
				"2 5",
				"1 3",
				"3 9",
				"3 3",
			},
			expected:    31,
			description: "Calculates similarity score correctly for a list with duplicates and mixed numbers.",
			expectError: false,
		},
		{
			input:       []string{},
			expected:    0,
			description: "Returns a similarity score of zero for empty input lists.",
			expectError: false,
		},
	}

	for i, tc := range testCases {
		t.Run("TestDayOnePartTwo_"+strconv.Itoa(i), func(t *testing.T) {
			day01 := solutions.Day01{}
			results, err := day01.PartTwo(tc.input)
			if err != nil {
				t.Errorf("Error occurred: %v", err)
				return
			}
			if !reflect.DeepEqual(results, tc.expected) {
				t.Errorf("Failed %s: For input %v, expected %v but got %v", tc.description, tc.input, tc.expected, results)
			}
		})
	}
}
