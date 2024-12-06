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
	visited := make(map[string]bool)
	result := walkMap(&lines, row, col, direction, &visited, 0)

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
	fmt.Println("No start found!")
	return -1, -1, ""
}

func walkMap(lines *[]string, row int, col int, direction string, visited *map[string]bool, result int) int {
	for row >= 0 && row < len(*lines) && col >= 0 && col < len((*lines)[row]) {
		if peekChar(row, col, direction, lines) == '#' {
			newDirection := turn(direction)
			return walkMap(lines, row, col, newDirection, visited, result)
		}

		coord := fmt.Sprintf("%d,%d", row, col)
		if !(*visited)[coord] {
			(*visited)[coord] = true
			result++
		}

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

	return result
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
