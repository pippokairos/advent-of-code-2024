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
	line := getExpandedLine(inputLine)
	defragLine(&line, len(line)-1)

	checksum := 0
	for i := 0; i < len(line); i++ {
		if line[i] != -1 {
			checksum += i * line[i]
		} else {
			continue
		}
	}
	fmt.Println("Checksum:", checksum)
}

func getExpandedLine(inputLine string) []int {
	line := []int{}
	for i := 0; i < len(inputLine); i++ {
		if inputLine[i] == '\n' {
			break
		}

		s := string(inputLine[i])
		count, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			return line
		}

		for j := 0; j < count; j++ {
			if i%2 == 0 {
				line = append(line, i/2)
			} else {
				line = append(line, -1)
			}
		}
	}
	return line
}

func defragLine(line *[]int, index int) {
	currentValue := -1
	var count int

	for i := index; i >= 0; i-- {
		if (*line)[i] == currentValue {
			if currentValue == -1 {
				continue
			} else {
				count++
			}
		} else {
			if currentValue != -1 {
				moveToAvailableSpaces(line, currentValue, i+1, count)
				defragLine(line, i)
				return
			} else {
				count = 1
				currentValue = (*line)[i]
			}
		}
	}
}

func moveToAvailableSpaces(line *[]int, value int, index int, count int) {
	spaceIndex := findSpaceIndex(line, count)
	if spaceIndex == -1 || spaceIndex >= index {
		return
	}
	for i := spaceIndex; i < spaceIndex+count; i++ {
		(*line)[i] = value
	}
	for i := index; i < index+count; i++ {
		(*line)[i] = -1
	}
}

func findSpaceIndex(line *[]int, count int) int {
	var spacesCount int
	for i := 0; i < len(*line); i++ {
		if (*line)[i] == -1 {
			spacesCount++
			if spacesCount == count {
				return i - spacesCount + 1
			}
		} else {
			spacesCount = 0
		}
	}
	return -1
}
