package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("../files/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	multiplyRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	disableRe := regexp.MustCompile(`don\'t\(\)`)
	enableRe := regexp.MustCompile(`do\(\)`)
	enabled := true

	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		line := scanner.Text()
		validPart := findValidPart(line, &enabled, enableRe, disableRe)
		matches := multiplyRe.FindAllStringSubmatch(validPart, -1)
		for _, match := range matches {
			if len(match) == 3 {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				result += num1 * num2
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total Result:", result)
}

func findValidPart(line string, enabled *bool, enableRe *regexp.Regexp, disableRe *regexp.Regexp) string {
	if len(line) == 0 {
		return ""
	}

	if *enabled {
		dontIndex := disableRe.FindStringIndex(line)
		if dontIndex != nil {
			*enabled = false
			return line[:dontIndex[0]] + findValidPart(line[dontIndex[1]:], enabled, enableRe, disableRe)
		}
		return line
	} else {
		doIndex := enableRe.FindStringIndex(line)
		if doIndex != nil {
			*enabled = true
			return findValidPart(line[doIndex[1]:], enabled, enableRe, disableRe)
		}
		return ""
	}
}
