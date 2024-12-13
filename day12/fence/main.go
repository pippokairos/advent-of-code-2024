package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func main() {
	file, err := os.Open("../files/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	lines := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []string{}
		for _, char := range scanner.Text() {
			line = append(line, string(char))
		}
		lines = append(lines, line)
	}

	price := 0
	visited := make(map[string]bool)
	for _, line := range lines {
		for _, char := range line {
			computePrice(lines, char, &price, &visited)
		}
	}
	fmt.Println(price)
}

func computePrice(lines [][]string, char string, price *int, visited *map[string]bool) {
	if _, ok := (*visited)[char]; ok {
		return
	}

	for i, line := range lines {
		for j, c := range line {
			if c == char {
				(*visited)[char] = true
				perimeter := 0
				area := 0
				exploreArea(lines, char, i, j, price, visited, &area, &perimeter)
				*price += area * perimeter
			}
		}
	}
}

func exploreArea(lines [][]string, char string, i, j int, price *int, visited *map[string]bool, area, perimeter *int) {
	if _, ok := (*visited)[fmt.Sprintf("%d-%d", i, j)]; ok {
		return
	}
	(*visited)[fmt.Sprintf("%d-%d", i, j)] = true

	*area++
	*perimeter += 4
	for _, direction := range directions {
		x, y := i+direction[0], j+direction[1]
		if x < 0 || x >= len(lines) || y < 0 || y >= len(lines[i]) {
			continue
		} else if lines[x][y] == char {
			*perimeter--
			exploreArea(lines, char, x, y, price, visited, area, perimeter)
		}
	}
}
