package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileBytes, err := os.ReadFile("../files/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	inputLine := string(fileBytes)
	line, spacesCount := getExpandedLine(inputLine)
	defragLine(&line, spacesCount)

	checksum := 0
	for i := 0; i < len(line); i++ {
		if line[i] != -1 {
			checksum += i * line[i]
		} else {
			break
		}
	}
	fmt.Println("Checksum:", checksum)
}

func getExpandedLine(inputLine string) ([]int, int) {
	line := []int{}
	spacesCount := 0
	for i := 0; i < len(inputLine); i++ {
		if inputLine[i] == '\n' {
			break
		}

		s := string(inputLine[i])
		count, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			return line, spacesCount
		}

		for j := 0; j < count; j++ {
			if i%2 == 0 {
				line = append(line, i/2)
			} else {
				spacesCount++
				line = append(line, -1)
			}
		}
	}
	return line, spacesCount
}

func defragLine(line *[]int, spacesCount int) {
	if isDefragged(line, spacesCount) {
		return
	}

	for i := len(*line) - 1; i >= 0; i-- {
		if (*line)[i] != -1 {
			moveToFirstSpace(line, i)
		}
	}
}

func isDefragged(line *[]int, spacesCount int) bool {
	for i := spacesCount; i < len(*line); i++ {
		if (*line)[i] != -1 {
			return false
		}
	}
	return true
}

func moveToFirstSpace(line *[]int, index int) {
	for i := 0; i < index; i++ {
		if (*line)[i] == -1 {
			(*line)[i] = (*line)[index]
			(*line)[index] = -1
			break
		}
	}
}
