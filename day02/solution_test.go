// day02/solution_test.go
package day02

import (
    "testing"
)

func TestProcessSafeData(t *testing.T) {
    // Test case 1: Example from the instructions
    input := []string{
        "7 6 4 2 1",
        "1 2 7 8 9",
        "9 7 6 2 1",
        "1 3 2 4 5",
        "8 6 4 4 1",
        "1 3 6 7 9",
    }

    safeCount, safeCountWithDampener := processSafeData(input)

    if safeCount != 2 {
        t.Errorf("Expected safe count to be 2, but got %d", safeCount)
    }

    if safeCountWithDampener != 4 {
        t.Errorf("Expected safe count with dampener to be 4, but got %d", safeCountWithDampener)
    }

    // Test case 2: Empty input
    input = []string{}

    safeCount, safeCountWithDampener = processSafeData(input)

    if safeCount != 0 {
        t.Errorf("Expected safe count to be 0, but got %d", safeCount)
    }

    if safeCountWithDampener != 0 {
        t.Errorf("Expected safe count with dampener to be 0, but got %d", safeCountWithDampener)
    }

    // Add more test cases as needed
}
