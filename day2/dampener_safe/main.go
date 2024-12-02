package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../files/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	safeReports := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values, err := parseLineToIntSlice(line)
		if err != nil {
			fmt.Printf("Skipping invalid line: %s (error: %v)\n", line, err)
			continue
		}
		if isSafe(values, 1) {
			safeReports++
		}
	}
	fmt.Println("Number of safe reports:", safeReports)
}

func parseLineToIntSlice(line string) ([]int, error) {
	parts := strings.Split(line, " ")
	values := make([]int, len(parts))
	for i, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid integer: %s", part)
		}
		values[i] = value
	}
	return values, nil
}

func isSafe(values []int, tolerance int) bool {
	if len(values) < 2 {
		return true
	}

	increasing := values[1] > values[0]
	for i := 1; i < len(values); i++ {
		if !isDifferenceSafe(values[i-1], values[i], increasing) {
			if tolerance > 0 {
				for j := 0; j < len(values); j++ {
					smallerValues := make([]int, len(values)-1)
					copy(smallerValues, values[:j])
					copy(smallerValues[j:], values[j+1:])
					if isSafe(smallerValues, tolerance-1) {
						return true
					}
				}
			}
			return false
		}
	}

	return true
}

func isDifferenceSafe(prev, current int, increasing bool) bool {
	diff := current - prev
	if increasing {
		return diff > 0 && diff < 4
	}
	return diff < 0 && -diff < 4
}
