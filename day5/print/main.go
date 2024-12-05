package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
			valid := true
			for i := 1; i < len(values); i++ {
				for _, rule := range rules {
					if values[i] == rule[0] && contains(values[:i], rule[1]) {
						valid = false
						break
					}
				}
				if !valid {
					break
				}
			}

			if valid {
				midIndex := len(values) / 2
				value, err := strconv.Atoi(values[midIndex])
				if err != nil {
					return 0, fmt.Errorf("invalid middle value: %s", values[midIndex])
				}
				result += value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return result, nil
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
