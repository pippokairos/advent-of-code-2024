package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	blinks = 25
)

func main() {
	now := time.Now()
	fileBytes, err := os.ReadFile("../files/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	line := strings.TrimSuffix(string(fileBytes), "\n")
	values := strings.Split(line, " ")

	result := computeStones(values, blinks)

	fmt.Println("Stones:", result)
	fmt.Println("Execution time:", time.Since(now))
}

func computeStones(values []string, blinks int) int {
	var wg sync.WaitGroup
	resultChannel := make(chan int, len(values))
	for _, value := range values {
		wg.Add(1)
		func(v string) {
			defer wg.Done()
			resultChannel <- blink(v, blinks)
		}(value)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	result := 0
	for r := range resultChannel {
		result += r
	}

	return result
}

func blink(value string, blinks int) int {
	if blinks == 0 {
		return 1
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	if intValue == 0 {
		return blink("1", blinks-1)
	}

	valueString := strconv.Itoa(intValue)

	length := len(valueString)
	if length%2 == 0 {
		return blink(valueString[:length/2], blinks-1) + blink(valueString[length/2:], blinks-1)
	}

	newValue := intValue * 2024
	newStringValue := strconv.Itoa(newValue)

	return blink(newStringValue, blinks-1)
}
