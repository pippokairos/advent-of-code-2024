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

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	nodes := make(map[string][][]int)
	for i, line := range lines {
		for j, char := range line {
			if char != '.' {
				nodes[string(char)] = append(nodes[string(char)], []int{i, j})
			}
		}
	}

	antinodes := computeAntinodes(nodes, &lines)
	fmt.Println("Antinodes:", antinodes)
}

func computeAntinodes(nodes map[string][][]int, lines *[]string) int {
	antinodes := 0
	for _, coordinates := range nodes {
		antinodes += len(coordinates)
		for i, coord := range coordinates {
			for j := i + 1; j < len(coordinates); j++ {
				antinodes += setAndCountAntinodes(coord, coordinates[j], lines)
			}
		}
	}
	return antinodes
}

func setAndCountAntinodes(coord1, coord2 []int, lines *[]string) int {
	antinodes := 0

	rowDelta := coord2[0] - coord1[0]
	colDelta := coord2[1] - coord1[1]

	antinodeCoords := [][]int{
		{coord1[0] - rowDelta, coord1[1] - colDelta},
		{coord2[0] + rowDelta, coord2[1] + colDelta},
	}

	setAndCountInDirection(antinodeCoords[0], lines, &antinodes, rowDelta, colDelta, -1)
	setAndCountInDirection(antinodeCoords[1], lines, &antinodes, rowDelta, colDelta, 1)

	return antinodes
}

func setAndCountInDirection(antinodeCoord []int, lines *[]string, antinodes *int, rowDelta, colDelta int, direction int) {
	row := antinodeCoord[0]
	col := antinodeCoord[1]

	if row < 0 || row >= len(*lines) || col < 0 || col >= len((*lines)[0]) {
		return
	}

	if (*lines)[row][col] != '#' {
		if (*lines)[row][col] == '.' {
			*antinodes++
		}
		(*lines)[row] = (*lines)[row][:col] + "#" + (*lines)[row][col+1:]
	}

	newCoords := []int{row + direction*rowDelta, col + direction*colDelta}
	setAndCountInDirection(newCoords, lines, antinodes, rowDelta, colDelta, direction)
}
