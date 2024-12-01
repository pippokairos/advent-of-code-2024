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

	var leftValues []int
	rightOccurrences := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var l, r int
		fmt.Sscanf(scanner.Text(), "%d %d", &l, &r)
		leftValues = append(leftValues, l)
		rightOccurrences[r]++
	}

	total := 0
	for i := 0; i < len(leftValues); i++ {
		total += leftValues[i] * rightOccurrences[leftValues[i]]
	}

	fmt.Println(total)
}
