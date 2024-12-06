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

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	row, col, direction := findStart(lines)
	if row == -1 {
		fmt.Println("No start found!")
		return
	}

	result := 0

	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			if lines[r][c] == '#' {
				continue
			}

			visited := make(map[string]bool)

			// Temporarily add obstruction
			originalChar := lines[r][c]
			lines[r] = lines[r][:c] + "#" + lines[r][c+1:]

			if isLoop(&lines, row, col, direction, &visited, 0) {
				result++
			}

			// Restore the original character
			lines[r] = lines[r][:c] + string(originalChar) + lines[r][c+1:]
		}
	}

	fmt.Println(result)
}

func findStart(lines []string) (int, int, string) {
	for i, line := range lines {
		for j, char := range line {
			switch char {
			case '^':
				return i, j, "up"
			case 'v':
				return i, j, "down"
			case '<':
				return i, j, "left"
			case '>':
				return i, j, "right"
			}
		}
	}
	return -1, -1, ""
}

func isLoop(lines *[]string, row int, col int, direction string, visited *map[string]bool, result int) bool {
	for row >= 0 && row < len(*lines) && col >= 0 && col < len((*lines)[row]) {
		if peekChar(row, col, direction, lines) == '#' {
			newDirection := turn(direction)
			return isLoop(lines, row, col, newDirection, visited, result)
		}

		coordAndDir := fmt.Sprintf("%d,%d,%s", row, col, direction)
		if (*visited)[coordAndDir] { // Loop detected
			return true
		}
		(*visited)[coordAndDir] = true

		switch direction {
		case "up":
			row--
		case "right":
			col++
		case "down":
			row++
		default:
			col--
		}
	}

	return false
}

func turn(direction string) string {
	switch direction {
	case "up":
		return "right"
	case "right":
		return "down"
	case "down":
		return "left"
	default:
		return "up"
	}
}

func peekChar(row int, col int, direction string, lines *[]string) byte {
	if direction == "up" && row > 0 {
		return (*lines)[row-1][col]
	} else if direction == "right" && col < len((*lines)[row])-1 {
		return (*lines)[row][col+1]
	} else if direction == "down" && row < len(*lines)-1 {
		return (*lines)[row+1][col]
	} else if direction == "left" && col > 0 {
		return (*lines)[row][col-1]
	}

	return ' '
}
