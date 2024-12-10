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
			reachable := map[string]bool{}
			fillReachable(&reachable, lines, i, j, 0)
			scores += len(reachable)
		}
	}
	return scores
}

func fillReachable(reachable *map[string]bool, lines [][]int, x, y, height int) {
	if x < 0 || y < 0 || x >= len(lines) || y >= len(lines[x]) || lines[x][y] != height {
		return
	}

	if height == 9 {
		key := fmt.Sprintf("%d-%d", x, y)
		(*reachable)[key] = true
		return
	}

	coordDiffs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, coordDiff := range coordDiffs {
		fillReachable(reachable, lines, x+coordDiff[0], y+coordDiff[1], height+1)
	}
}
