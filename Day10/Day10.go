package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Day10_2")

	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	input := string(inputBytes)
	lenghtArray := make([]int, 0)

	for _, v := range input {
		lenghtArray = append(lenghtArray, int(v))
	}

	lenghtArray = append(lenghtArray, 17)
	lenghtArray = append(lenghtArray, 31)
	lenghtArray = append(lenghtArray, 73)
	lenghtArray = append(lenghtArray, 47)
	lenghtArray = append(lenghtArray, 23)

	fmt.Println(lenghtArray)

	positions := make([]int, 0)

	for i := 0; i < 256; i++ {
		positions = append(positions, i)
	}

	currentIndex := 0
	skipSize := 0
	for k := 0; k < 64; k++ {
		for _, lenght := range lenghtArray {
			reverse(positions, currentIndex, lenght)

			if currentIndex+lenght+skipSize > len(positions) {
				currentIndex = currentIndex + lenght + skipSize - len(positions)
			} else {
				currentIndex += lenght + skipSize
			}

			if currentIndex > len(positions) {
				currentIndex -= len(positions)
			}

			skipSize++
			if skipSize == len(positions) {
				skipSize = 0
			}
		}
	}

	denseHash := [16]int{}

	result := ""
	for i := 0; i < 16; i++ {
		blockStartIndex := i * 16
		sequenceBlock := positions[blockStartIndex]
		for j := 1; j < 16; j++ {
			sequenceBlock = sequenceBlock ^ positions[blockStartIndex+j]
		}
		denseHash[i] = sequenceBlock

		h := fmt.Sprintf("%x", denseHash[i])
		if len(h) == 1 {
			result += "0"
		}
		result += h
	}

	fmt.Printf("Result %s \n", result)
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

		temp := input[firstIndex]
		input[firstIndex] = input[lastIndex]
		input[lastIndex] = temp
	}
}
