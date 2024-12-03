package day03

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
				"xmul(2,4)%&mul[3,7]!@^mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
				"mul(10, 5)*&^mul(15, 3)mul(6,7)mul(9,9)",
				"randomtextmul(7,9)moretextmul(3,8)mul(12,12)",
				"startmul(4,4)endmul(20, 6)mul(1,1)mul(2,2)",
			},
			expected: 2*4 + 5*5 + 11*8 + 8*5 + 6*7 + 9*9 + 7*9 + 3*8 + 12*12 + 4*4 + 1*1 + 2*2, // Adjusted expected value
		},
		{
			input: []string{
				"mul(1,1)mul(2,2)mul(3,3)mul(4,4)mul(5,5)",
			},
			expected: 1*1 + 2*2 + 3*3 + 4*4 + 5*5, // Sum of all valid mul instructions
		},
		{
			input: []string{
				"mul(0,0)mul(0,1)mul(1,0)mul(0,0)",
			},
			expected: 0, // All multiplications involve zero
		},
		{
			input: []string{
				"mul(10,10)mul(20,20)mul(30,30)mul(40,40)",
			},
			expected: 10*10 + 20*20 + 30*30 + 40*40, // Sum of all valid mul instructions
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
				"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			expected: 2*4 + 8*5, // Only mul(2,4) and mul(8,5) are enabled
		},
		{
			input: []string{
				"do()mul(2,3)don't()mul(4,5)do()mul(6,7)",
			},
			expected: 2*3 + 6*7, // Only mul(2,3) and mul(6,7) are enabled
		},
		{
			input: []string{
				"don't()mul(1,1)do()mul(2,2)",
			},
			expected: 2 * 2, // Only mul(2,2) is enabled
		},
		{
			input: []string{
				"do()mul(3,3)don't()mul(4,4)do()mul(5,5)don't()mul(6,6)",
			},
			expected: 3*3 + 5*5, // Only mul(3,3) and mul(5,5) are enabled
		},
	}

	for i, tc := range testCases {
		t.Run("PartTwoTestCase"+strconv.Itoa(i), func(t *testing.T) {
			result := PartTwo(tc.input)
			if result != tc.expected {
				t.Errorf("For input %v, expected %v but got %v", tc.input, tc.expected, result)
			}
		})
	}
}
