package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../files/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += equationExists(scanner.Text())
	}

	fmt.Println("Result:", result)
}

func equationExists(line string) int {
	resultAndOperands := strings.Split(line, ": ")
	requiredResult, err := strconv.Atoi(resultAndOperands[0])
	if err != nil {
		fmt.Println(err)
		return 0
	}

	operands := strings.Split(resultAndOperands[1], " ")

	result, err := strconv.Atoi(operands[0])
	if err != nil {
		fmt.Println(err)
		return 0
	}

	if equationPossible(operands[1:], requiredResult, result) {
		return requiredResult
	}

	return 0
}

func equationPossible(operands []string, requiredResult, result int) bool {
	if result > requiredResult {
		return false
	}

	if len(operands) == 0 {
		return result == requiredResult
	}

	operand, err := strconv.Atoi(operands[0])
	if err != nil {
		fmt.Println(err)
		return false
	}

	if len(operands) == 1 {
		return applyOperation("+", result, operand) == requiredResult ||
			applyOperation("*", result, operand) == requiredResult ||
			applyOperation("||", result, operand) == requiredResult
	}

	return equationPossible(operands[1:], requiredResult, applyOperation("+", result, operand)) ||
		equationPossible(operands[1:], requiredResult, applyOperation("*", result, operand)) ||
		equationPossible(operands[1:], requiredResult, applyOperation("||", result, operand))
}

func applyOperation(operator string, a, b int) int {
	switch operator {
	case "+":
		return a + b
	case "*":
		return a * b
	case "||":
		digitsOfB := len(strconv.Itoa(b))
		return a*int(math.Pow10(digitsOfB)) + b
	default:
		return 0
	}
}
