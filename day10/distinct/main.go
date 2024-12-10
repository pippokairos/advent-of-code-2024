package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../files/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	lines := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []int{}
		for _, char := range scanner.Text() {
			line = append(line, int(char-'0'))
		}
		lines = append(lines, line)
	}

	scores := computeScores(lines)
	fmt.Println(scores)
}

func computeScores(lines [][]int) int {
	scores := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			scores += findTrailheads(lines, i, j, 0)
		}
	}
	return scores
}

func findTrailheads(lines [][]int, i, j, score int) int {
	if i < 0 || i >= len(lines) || j < 0 || j >= len(lines[i]) || lines[i][j] != score {
		return 0
	}

	if lines[i][j] == 9 {
		return 1
	}

	return findTrailheads(lines, i-1, j, score+1) +
		findTrailheads(lines, i+1, j, score+1) +
		findTrailheads(lines, i, j-1, score+1) +
		findTrailheads(lines, i, j+1, score+1)
}
