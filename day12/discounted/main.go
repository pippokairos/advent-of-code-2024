package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
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
				area := 0
				section := []string{}
				exploreArea(lines, char, i, j, price, visited, &area, &section)
				perimeter := computePerimeter(lines, char, &section)
				*price += area * perimeter
			}
		}
	}
}

func exploreArea(lines [][]string, char string, i, j int, price *int, visited *map[string]bool, area *int, section *[]string) {
	if _, ok := (*visited)[fmt.Sprintf("%d-%d", i, j)]; ok {
		return
	}
	key := fmt.Sprintf("%d-%d", i, j)
	(*visited)[key] = true
	*section = append(*section, key)

	*area++
	for _, direction := range directions {
		x, y := i+direction[0], j+direction[1]
		if x < 0 || x >= len(lines) || y < 0 || y >= len(lines[i]) {
			continue
		} else if lines[x][y] == char {
			exploreArea(lines, char, x, y, price, visited, area, section)
		}
	}
}

func computePerimeter(lines [][]string, char string, section *[]string) int {
	perimeter := 0
	for _, key := range *section {
		var edges [4]bool
		var x, y int
		fmt.Sscanf(key, "%d-%d", &x, &y)
		for i, direction := range directions {
			x1, y1 := x+direction[0], y+direction[1]
			if x1 < 0 || x1 >= len(lines) || y1 < 0 || y1 >= len(lines[x]) || lines[x1][y1] != char {
				edges[i] = true
			}
		}
		for i := 0; i < 4; i++ {
			if !edges[i] {
				continue
			}

			if edges[(i+1)%4] {
				perimeter++
			}
			x1, y1 := x+directions[i][0], y+directions[i][1]
			x2, y2 := x+directions[(i+3)%4][0], y+directions[(i+3)%4][1]
			x3, y3 := x2+directions[i][0], y2+directions[i][1]
			if x1 >= 0 && x1 < len(lines) && y1 >= 0 && y1 < len(lines[x]) && x2 >= 0 && x2 < len(lines) && y2 >= 0 && y2 < len(lines[x]) && x3 >= 0 && x3 < len(lines) && y3 >= 0 && y3 < len(lines[x]) && lines[x2][y2] == char && lines[x3][y3] == char {
				perimeter++
			}
		}
	}

	return perimeter
}
