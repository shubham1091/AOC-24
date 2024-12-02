package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
	// Fetch input data
	data, err := fetchInput()
	if err != nil {
		fmt.Println("Error fetching input:", err)
		return
	}

	var leftList, rightList []int
	scanner := bufio.NewScanner(bytes.NewReader(data))

	// Parse input
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			continue
		}

		left, err1 := strconv.Atoi(numbers[0])
		right, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing numbers:", err1, err2)
			continue
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	// Sort both lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	// Calculate total distance
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := abs(leftList[i] - rightList[i])
		totalDistance += distance
	}

	freq := make(map[int]int)
	for _, num := range rightList {
		freq[num]++
	}
	var similarityScore int
	for _, num := range leftList {
		similarityScore += num * freq[num]
	}

	fmt.Println("Total distance:", totalDistance)
	fmt.Println("Similarity score:", similarityScore)
}

func fetchInput() ([]byte, error) {
	sessionCookie := os.Getenv("AOC_SESSION")
	if sessionCookie == "" {
		return nil, fmt.Errorf("AOC_SESSION environment variable not set")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://adventofcode.com/2024/day/1/input", nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
