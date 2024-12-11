package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	blinks = 75
)

func main() {
	fileBytes, err := os.ReadFile("../files/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	line := strings.TrimSuffix(string(fileBytes), "\n")
	values := strings.Split(line, " ")

	result := computeStones(values, blinks)

	fmt.Println("Stones:", result)
}

func computeStones(values []string, blinks int) int {
	valToResMap := make(map[string]int)
	result := 0
	for _, value := range values {
		result += blink(value, blinks, valToResMap)
	}

	return result
}

func blink(value string, blinks int, valToResMap map[string]int) int {
	if blinks == 0 {
		return 1
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	valueString := strconv.Itoa(intValue)

	key := fmt.Sprintf("%s-%d", valueString, blinks)
	if res, ok := valToResMap[key]; ok {
		return res
	}

	if intValue == 0 {
		valToResMap[key] = blink("1", blinks-1, valToResMap)
		return valToResMap[key]
	}

	length := len(valueString)
	if length%2 == 0 {
		left := blink(valueString[:length/2], blinks-1, valToResMap)
		right := blink(valueString[length/2:], blinks-1, valToResMap)
		valToResMap[key] = left + right
		return valToResMap[key]
	}

	newValue := intValue * 2024
	newStringValue := strconv.Itoa(newValue)
	valToResMap[key] = blink(newStringValue, blinks-1, valToResMap)
	return valToResMap[key]
}
