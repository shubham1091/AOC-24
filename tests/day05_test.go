package tests

import (
	"AOC-24/solutions"
	"reflect"
	"testing"
)

func TestDayFivePartOne(t *testing.T) {
	type testCase struct {
		input       []string
		expected    int
		description string
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
			expected:    143,
			description: "Test basic input with valid page ordering rules and updates.",
		},
		{
			input:       []string{},
			expected:    0,
			description: "Test empty input; should result in no correctly ordered updates.",
		},
		{
			input: []string{
				"47|53",
				"75,47",
			},
			expected:    0,
			description: "Test single update with minimal rules; ensure middle page is calculated.",
		},
	}

	for i, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			day05 := solutions.Day05{}
			result, err := day05.PartOne(tc.input)
			if err != nil {
				t.Errorf("Unexpected error for test case %d: %v", i, err)
			}
			if result != tc.expected {
				t.Errorf("Failed %s: Expected %d, got %d", tc.description, tc.expected, result)
			}
		})
	}
}

func TestDayFivePartTwo(t *testing.T) {
	type testCase struct {
		input       []string
		expected    int
		description string
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
			expected:    123,
			description: "Test reordering incorrectly ordered updates and summing their middle pages.",
		},
	}

	for i, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			day05 := solutions.Day05{}
			result, err := day05.PartTwo(tc.input)
			if err != nil {
				t.Errorf("Unexpected error for test case %d: %v", i, err)
			}
			if result != tc.expected {
				t.Errorf("Failed %s: Expected %d, got %d", tc.description, tc.expected, result)
			}
		})
	}
}

func TestProcessData(t *testing.T) {
	t.Run("Valid rules and updates", func(t *testing.T) {
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

		rules, updates := solutions.ProcessData(input)

		if !reflect.DeepEqual(rules, expectedRules) {
			t.Errorf("Expected rules %v, got %v", expectedRules, rules)
		}
		if !reflect.DeepEqual(updates, expectedUpdates) {
			t.Errorf("Expected updates %v, got %v", expectedUpdates, updates)
		}
	})
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
		{"Non-numeric rule", "abc|def", 0, 0, false},
		{"Empty rule", "", 0, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y, ok := solutions.ParseRule(tt.input)
			if ok != tt.expectedValid {
				t.Errorf("Expected valid=%v, got %v", tt.expectedValid, ok)
			}
			if ok && (x != tt.expectedX || y != tt.expectedY) {
				t.Errorf("Expected (%d, %d), got (%d, %d)", tt.expectedX, tt.expectedY, x, y)
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
		{"Valid update", "75,47,61,53,29", []int{75, 47, 61, 53, 29}},
		{"Empty update", "", []int{}},
		{"Single page", "42", []int{42}},
		{"Mixed spaces", "  75, 47 ,61 ,53  , 29", []int{75, 47, 61, 53, 29}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := solutions.ParseUpdate(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
