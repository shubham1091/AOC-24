package day05

import (
	"reflect"
	"strconv"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := []string{
		"47|53",
		"97|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
	}

	expectedRules := map[int][]int{
		47: {53},
		97: {13},
	}
	expectedUpdates := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
	}

	rules, updates := processData(input)

	if !reflect.DeepEqual(rules, expectedRules) {
		t.Errorf("Expected rules %v, got %v", expectedRules, rules)
	}
	if !reflect.DeepEqual(updates, expectedUpdates) {
		t.Errorf("Expected updates %v, got %v", expectedUpdates, updates)
	}
}

func TestParseRule(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedX     int
		expectedY     int
		expectedValid bool
	}{
		{"Valid rule", "47|53", 47, 53, true},
		{"Invalid format", "47-53", 0, 0, false},
		{"Invalid numbers", "abc|def", 0, 0, false},
		{"Empty string", "", 0, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y, ok := parseRule(tt.input)
			if ok != tt.expectedValid {
				t.Errorf("Expected valid=%v, got %v", tt.expectedValid, ok)
			}
			if ok && (x != tt.expectedX || y != tt.expectedY) {
				t.Errorf("Expected (%d,%d), got (%d,%d)", tt.expectedX, tt.expectedY, x, y)
			}
		})
	}
}

func TestParseUpdate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{
			name:     "Valid update",
			input:    "75,47,61,53,29",
			expected: []int{75, 47, 61, 53, 29},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: []int{},
		},
		{
			name:     "Single number",
			input:    "42",
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseUpdate(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	type testCase struct {
		input    []string
		expected int
	}

	testCases := []testCase{
		{
			input: []string{
				"47|53",
				"97|13",
				"97|61",
				"97|47",
				"75|29",
				"61|13",
				"75|53",
				"29|13",
				"97|29",
				"53|29",
				"61|53",
				"97|53",
				"61|29",
				"47|13",
				"75|47",
				"97|75",
				"47|61",
				"75|61",
				"47|29",
				"75|13",
				"53|13",
				"",
				"75,47,61,53,29",
				"97,61,53,29,13",
				"75,29,13",
				"75,97,47,61,53",
				"61,13,29",
				"97,13,75,29,47",
			},
			expected: 143,
		},
	}

	for i, tc := range testCases {
		t.Run("PartOneTestCase"+strconv.Itoa(i), func(t *testing.T) {
			rules, updates := processData(tc.input)
			results := PartOne(rules, updates)
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
				"47|53",
				"97|13",
				"97|61",
				"97|47",
				"75|29",
				"61|13",
				"75|53",
				"29|13",
				"97|29",
				"53|29",
				"61|53",
				"97|53",
				"61|29",
				"47|13",
				"75|47",
				"97|75",
				"47|61",
				"75|61",
				"47|29",
				"75|13",
				"53|13",
				"",
				"75,47,61,53,29",
				"97,61,53,29,13",
				"75,29,13",
				"75,97,47,61,53",
				"61,13,29",
				"97,13,75,29,47",
			},
			expected: 123,
		},
	}

	for i, tc := range testCases {
		t.Run("PartOneTestCase"+strconv.Itoa(i), func(t *testing.T) {
			rules, updates := processData(tc.input)
			results := PartTwo(rules, updates)
			if !reflect.DeepEqual(results, tc.expected) {
				t.Errorf("For input %v, expected %v but got %v", tc.input, tc.expected, results)
			}
		})
	}
}
