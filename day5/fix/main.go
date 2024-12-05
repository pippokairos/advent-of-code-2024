package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../files/input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	rules, err := getRules(file)
	if err != nil {
		fmt.Println("Error reading rules:", err)
		return
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println("Error resetting file pointer:", err)
		return
	}

	result, err := computeValidRules(file, rules)
	if err != nil {
		fmt.Println("Error computing valid rules:", err)
		return
	}

	fmt.Println("Result:", result)
}

func getRules(file *os.File) ([][]string, error) {
	var rules [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		rule := strings.Split(line, "|")
		if len(rule) != 2 {
			return nil, fmt.Errorf("invalid rule format: %s", line)
		}
		rules = append(rules, rule)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading rules: %v", err)
	}

	return rules, nil
}

func computeValidRules(file *os.File, rules [][]string) (int, error) {
	scanner := bufio.NewScanner(file)
	skipLine := true
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			skipLine = false
			continue
		}

		if skipLine {
			continue
		}

		values := strings.Split(line, ",")
		if len(values) == 0 {
			continue
		}

		if len(values) == 1 {
			value, err := strconv.Atoi(values[0])
			if err != nil {
				return 0, fmt.Errorf("invalid value: %s", values[0])
			}
			result += value
		} else {
			copiedValues := make([]string, len(values))
			copy(copiedValues, values)
			validValues := getValidRow(copiedValues, rules)
			if reflect.DeepEqual(validValues, values) {
				continue
			}

			midIndex := len(validValues) / 2
			value, err := strconv.Atoi(validValues[midIndex])
			if err != nil {
				return 0, fmt.Errorf("invalid middle value: %s", validValues[midIndex])
			}
			result += value
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return result, nil
}

func getValidRow(values []string, rules [][]string) []string {
	for i := 1; i < len(values); i++ {
		for _, rule := range rules {
			righSideIndex := contains(values[:i], rule[1])
			if values[i] == rule[0] && righSideIndex != -1 {
				values[i], values[righSideIndex] = values[righSideIndex], values[i]
				return getValidRow(values, rules)
			}
		}
	}
	return values
}

func contains(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}
