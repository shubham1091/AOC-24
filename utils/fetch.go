package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LoggedClient struct {
	http.Client
	Logger *log.Logger
}

func (c *LoggedClient) Do(req *http.Request) (*http.Response, error) {
	c.Logger.Printf("Fetching input for day %s", req.URL)
	resp, err := c.Client.Do(req)
	if err != nil {
		c.Logger.Printf("Failed to fetch: %v", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		c.Logger.Printf("Unexpected status code: %d", resp.StatusCode)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return resp, nil
}

func FetchInput(day int) ([]byte, error) {
	sessionCookie, err := GetEnvVar("SESSION_COOKIE")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})

	client := &LoggedClient{Client: *http.DefaultClient, Logger: log.Default()}
	return client.fetchAndRead(req)
}

func (c *LoggedClient) fetchAndRead(req *http.Request) ([]byte, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.Logger.Printf("Successfully fetched input for URL %s\n", req.URL)
	return data, nil
}

func ParseInput(data []byte) []string {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}
