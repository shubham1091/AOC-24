package tests

import (
	"AOC-24/solutions"
	"reflect"
	"strconv"
	"testing"
)

func TestDayThreePartOne(t *testing.T) {
	type testCase struct {
		input       []string
		expected    int
		description string
		expectError bool
	}

	testCases := []testCase{
		{
			input: []string{
				"xmul(2,4)%&mul[3,7]!@^mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			},
			expected:    2*4 + 5*5 + 11*8 + 8*5,
			description: "Valid mul instructions with mixed invalid characters and nested patterns",
			expectError: false,
		},
		{
			input: []string{
				"mul(10,5)&^mul(15,3)*!mul(6,7)randomtextmul(9,9)end",
			},
			expected:    10*5 + 15*3 + 6*7 + 9*9,
			description: "Valid mul instructions mixed with unrelated text and symbols",
			expectError: false,
		},
		{
			input: []string{
				"invalidtextmul(3,8)moretextmul(12,12)extrastuff",
			},
			expected:    3*8 + 12*12,
			description: "Instructions embedded within arbitrary strings",
			expectError: false,
		},
		{
			input: []string{
				"mul(0,0)mul(0,1)mul(1,0)mul(0,0)",
			},
			expected:    0,
			description: "All valid instructions, but results are zero due to multiplication by zero",
			expectError: false,
		},
		{
			input: []string{
				"mul(123,456)mul(7,8)mul(90,10)invalidmul(4,5)end",
			},
			expected:    123*456 + 7*8 + 90*10 + 4*5,
			description: "Large numbers with mixed invalid instructions",
			expectError: false,
		},
	}

	for i, tc := range testCases {
		t.Run("DayThreePartOneTestCase"+strconv.Itoa(i), func(t *testing.T) {
			day03 := solutions.Day03{}
			result, err := day03.PartOne(tc.input)
			if err != nil && !tc.expectError {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if err == nil && tc.expectError {
				t.Errorf("Expected error but got none")
				return
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("For input %v, expected %v but got %v", tc.input, tc.expected, result)
			}
		})
	}
}

func TestDayThreePartTwo(t *testing.T) {
	type testCase struct {
		input       []string
		expected    int
		description string
		expectError bool
	}

	testCases := []testCase{
		{
			input: []string{
				"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			expected:    2*4 + 8*5,
			description: "Mixed do() and don't() instructions with valid and invalid mul instructions",
			expectError: false,
		},
		{
			input: []string{
				"do()mul(2,3)don't()mul(4,5)do()mul(6,7)",
			},
			expected:    2*3 + 6*7,
			description: "Switching between enabled and disabled states",
			expectError: false,
		},
		{
			input: []string{
				"don't()mul(1,1)do()mul(2,2)",
			},
			expected:    2 * 2,
			description: "Initial disable state, followed by enabling mul instruction",
			expectError: false,
		},
		{
			input: []string{
				"do()mul(3,3)don't()mul(4,4)do()mul(5,5)don't()mul(6,6)",
			},
			expected:    3*3 + 5*5,
			description: "Alternating enable and disable instructions",
			expectError: false,
		},
		{
			input: []string{
				"randomtextdo()mul(4,4)don't()mul(20,6)extrastuff",
			},
			expected:    4 * 4,
			description: "Valid do() and don't() instructions with noise in input",
			expectError: false,
		},
	}

	for i, tc := range testCases {
		t.Run("DayThreePartTwoTestCase"+strconv.Itoa(i), func(t *testing.T) {
			day03 := solutions.Day03{}
			result, err := day03.PartTwo(tc.input)
			if err != nil && !tc.expectError {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if err == nil && tc.expectError {
				t.Errorf("Expected error but got none")
				return
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("For input %v, expected %v but got %v", tc.input, tc.expected, result)
			}
		})
	}
}
