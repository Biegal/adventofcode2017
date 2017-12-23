package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type operation struct {
	register    string
	appendValue int
	leftSide    string
	operator    string
	rightSide   string
}

func main() {
	fmt.Println("Task 15")

	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)

	register := map[string]int{}

	operations := make([]operation, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()

		splittedValues := strings.Split(line, " ")

		operation := operation{}
		operation.register = splittedValues[0]

		operation.appendValue, _ = strconv.Atoi(splittedValues[2])

		if splittedValues[1] == "dec" {
			operation.appendValue = -1 * operation.appendValue
		}

		operation.leftSide = splittedValues[4]
		operation.operator = splittedValues[5]
		operation.rightSide = splittedValues[6]

		operations = append(operations, operation)
		register[operation.leftSide] = 0
	}

	maxEver := 0
	for _, operation := range operations {
		if runCondition(operation, register) {
			register[operation.register] += operation.appendValue
		}

		maxInIteration := findMaxInRegister(register)
		if maxInIteration > maxEver {
			maxEver = maxInIteration
		}
	}

	fmt.Println(maxEver)
}

func findMaxInRegister(register map[string]int) int {
	max := 0
	for _, v := range register {
		if v > max {
			max = v
		}
	}
	return max
}

func runCondition(operation operation, register map[string]int) bool {
	comparer, err := strconv.Atoi(operation.rightSide)
	if err != nil {
		log.Fatal("Error in conversion")
	}

	switch operation.operator {
	case "==":
		return register[operation.leftSide] == comparer
	case ">":
		return register[operation.leftSide] > comparer
	case "<":
		return register[operation.leftSide] < comparer
	case ">=":
		return register[operation.leftSide] >= comparer
	case "<=":
		return register[operation.leftSide] <= comparer
	case "!=":
		return register[operation.leftSide] != comparer
	default:
		log.Fatal("Not handled operator " + operation.operator)
	}
	return false
}
