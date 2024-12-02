package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// FetchInput fetches the input data for a given day from the Advent of Code website.
// It uses the GetSessionCookie function to retrieve the session cookie.
// The function returns the input data as a byte slice and logs any errors encountered.
func FetchInput(day int) ([]byte, error) {
	// Retrieve session cookie from GetSessionCookie
	sessionCookie, err := GetSessionCookie()
	if err != nil {
		return nil, fmt.Errorf("failed to get session cookie: %v", err)
	}

	// Construct the URL for the input data
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)

	// Create an HTTP client
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Add the session cookie to the request
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch input data: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Log successful fetch
	log.Printf("Successfully fetched input for Day %d\n", day)
	return data, nil
}
