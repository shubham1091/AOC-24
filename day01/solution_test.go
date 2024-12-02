package day01

import (
    "testing"
)

func TestProcessInput(t *testing.T) {
    // Test case 1: Example from the instructions
    input := []string{
        "3 4",
        "4 3",
        "2 5",
        "1 3",
        "3 9",
        "3 3",
    }

    totalDistance, similarityScore := processInput(input)

    if totalDistance != 11 {
        t.Errorf("Expected total distance to be 11, but got %d", totalDistance)
    }

    if similarityScore != 31 {
        t.Errorf("Expected similarity score to be 31, but got %d", similarityScore)
    }
}
