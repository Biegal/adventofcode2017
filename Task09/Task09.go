package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Task 09")

	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	input := string(inputBytes)
	stepsAsStrings := strings.Split(input, "\n")
	steps := make([]int, len(stepsAsStrings))

	for i, v := range stepsAsStrings {
		stepAsInt, _ := strconv.Atoi(v)
		steps[i] = stepAsInt
	}

	lastStep := 0
	currentStep := 0
	jumps := 0
	fmt.Printf("Current step: %d \n", currentStep)

	for {
		if currentStep >= len(steps) {
			break
		}

		lastStep = currentStep

		currentStep += steps[currentStep]
		steps[lastStep]++

		fmt.Printf("Current step: %d \n", currentStep)

		jumps++
	}

	fmt.Println(jumps)
}
