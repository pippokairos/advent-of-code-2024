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

	antinodes := computeAntinodes(nodes, lines)
	fmt.Println("Antinodes:", antinodes)
}

func computeAntinodes(nodes map[string][][]int, lines []string) int {
	antinodes := 0
	for _, coordinates := range nodes {
		for i, coord := range coordinates {
			for j := i + 1; j < len(coordinates); j++ {
				antinodes += setAndCountAntinodes(coord, coordinates[j], lines)
			}
		}
	}
	return antinodes
}

func setAndCountAntinodes(coord1, coord2 []int, lines []string) int {
	antinodes := 0
	rowDelta := coord2[0] - coord1[0]
	colDelta := coord2[1] - coord1[1]

	antinodeCoords := [][]int{
		{coord1[0] - rowDelta, coord1[1] - colDelta},
		{coord2[0] + rowDelta, coord2[1] + colDelta},
	}

	for _, antinodeCoord := range antinodeCoords {
		if antinodeCoord[0] < 0 || antinodeCoord[0] >= len(lines) || antinodeCoord[1] < 0 || antinodeCoord[1] >= len(lines[0]) {
			continue
		}

		if lines[antinodeCoord[0]][antinodeCoord[1]] != '#' {
			lines[antinodeCoord[0]] = lines[antinodeCoord[0]][:antinodeCoord[1]] + "#" + lines[antinodeCoord[0]][antinodeCoord[1]+1:]
			antinodes++
		}
	}

	return antinodes
}
