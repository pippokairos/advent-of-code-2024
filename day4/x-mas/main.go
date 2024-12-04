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

	result := 0
	for lineIndex, line := range lines {
		for charIndex, char := range line {
			if char == 'A' && isXmas(lineIndex, charIndex, &lines) {
				result += 1
			}
		}
	}

	fmt.Println(result)
}

func isXmas(lineIndex int, charIndex int, lines *[]string) bool {
	if lineIndex == 0 || lineIndex == len(*lines)-1 || charIndex == 0 || charIndex == len((*lines)[lineIndex])-1 {
		return false
	}

	crossValues := []byte{
		(*lines)[lineIndex-1][charIndex-1], // Top left
		(*lines)[lineIndex+1][charIndex-1], // Bottom left
		(*lines)[lineIndex-1][charIndex+1], // Top right
		(*lines)[lineIndex+1][charIndex+1], // Bottom right
	}

	return areMAndS(crossValues[0], crossValues[3]) && areMAndS(crossValues[1], crossValues[2])
}

func areMAndS(a byte, b byte) bool {
	return (a == 'M' && b == 'S' || a == 'S' && b == 'M')
}
