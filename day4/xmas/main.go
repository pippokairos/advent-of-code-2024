package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = []struct {
	rowDelta, colDelta int
}{
	{0, 1},   // Right
	{0, -1},  // Left
	{-1, 0},  // Up
	{1, 0},   // Down
	{-1, 1},  // Diagonal up-right
	{-1, -1}, // Diagonal up-left
	{1, 1},   // Diagonal down-right
	{1, -1},  // Diagonal down-left
}

func main() {
	file, err := os.Open("../files/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	result := 0
	for lineIndex, line := range lines {
		for charIndex, char := range line {
			if char == 'X' {
				result += countOccurrencesFrom(lineIndex, charIndex, &lines)
			}
		}
	}

	fmt.Println(result)
}

func countOccurrencesFrom(lineIndex, charIndex int, lines *[]string) int {
	result := 0

	for _, dir := range directions {
		if checkXmasInDirection(lineIndex, charIndex, lines, dir.rowDelta, dir.colDelta) {
			result++
		}
	}

	return result
}

func checkXmasInDirection(lineIndex, charIndex int, lines *[]string, rowDelta, colDelta int) bool {
	word := "MAS"

	for i := 1; i <= len(word); i++ {
		newRow := lineIndex + rowDelta*i
		newCol := charIndex + colDelta*i

		// Bounds check
		if newRow < 0 || newRow >= len(*lines) || newCol < 0 || newCol >= len((*lines)[newRow]) {
			return false
		}

		// Character check
		if (*lines)[newRow][newCol] != word[i-1] {
			return false
		}
	}

	return true
}
