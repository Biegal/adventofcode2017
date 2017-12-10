package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Task 19")

	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	input := string(inputBytes)
	stepsAsStrings := strings.Split(input, ",")
	lenghtArray := make([]int, len(stepsAsStrings))

	for i, v := range stepsAsStrings {
		stepAsInt, _ := strconv.Atoi(v)
		lenghtArray[i] = stepAsInt
	}

	positions := make([]int, 0)

	for i := 0; i < 256; i++ {
		positions = append(positions, i)
	}
	fmt.Println(positions)

	currentIndex := 0
	skipSize := 0
	for _, lenght := range lenghtArray {
		fmt.Printf("Current position %d \n", currentIndex)
		reverse(positions, currentIndex, lenght)

		if currentIndex+lenght+skipSize > len(positions) {
			currentIndex = currentIndex + lenght + skipSize - len(positions)
		} else {
			currentIndex += lenght + skipSize
		}
		skipSize++

		fmt.Println(positions)
	}

	fmt.Printf("Result %d \n", positions[0]*positions[1])
}

func reverse(input []int, startIndex int, lenght int) {

	for i := 0; i < lenght/2; i++ {
		firstIndex := 0
		if startIndex+i < len(input) {
			firstIndex = startIndex + i
		} else {
			firstIndex = (startIndex + i) - len(input)
		}

		lastIndex := 0
		if startIndex+lenght-1-i >= len(input) {
			lastIndex = startIndex + lenght - i - 1 - len(input)
		} else {
			lastIndex = startIndex + lenght - i - 1
		}

		fmt.Printf("Lenght %d \n", lenght)
		fmt.Printf("First index %d  \n", firstIndex)
		fmt.Printf("Last index %d  \n", lastIndex)

		temp := input[firstIndex]
		input[firstIndex] = input[lastIndex]
		input[lastIndex] = temp
	}
}
