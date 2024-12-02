package day01

import (
	"AOC-24/utils"
	"bufio"
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Run() {
    // Fetch input data
    data, err := utils.FetchInput(1)
    if err != nil {
        fmt.Println("Error fetching input:", err)
        return
    }

    scanner := bufio.NewScanner(bytes.NewReader(data))
    var input []string
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }

    totalDistance, similarityScore := processInput(input)

    fmt.Println("Total distance:", totalDistance)
    fmt.Println("Similarity score:", similarityScore)
}


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


func processInput(input []string) (int, int) {
    var leftList, rightList []int

    // Parse input
    for _, line := range input {
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

    return totalDistance, similarityScore
}
