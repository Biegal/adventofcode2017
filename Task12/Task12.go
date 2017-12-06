package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Task 11")

	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	input := string(inputBytes)
	stepsAsStrings := strings.Split(input, "\t")
	memAddresses := make([]int, len(stepsAsStrings))

	for i, v := range stepsAsStrings {
		stepAsInt, _ := strconv.Atoi(v)
		memAddresses[i] = stepAsInt
	}

	results := make([][]int, 0)
	redistributions := 0
	hasFound := false
	for {
		highestIndex := 0
		for i, v := range memAddresses {
			if v > memAddresses[highestIndex] {
				highestIndex = i
			}
		}

		memBlocksToAllocate := memAddresses[highestIndex]
		memAddresses[highestIndex] = 0
		for i := 0; i < memBlocksToAllocate; i++ {
			index := math.Mod(float64(highestIndex+i+1), float64(len(memAddresses)))

			memAddresses[int(index)]++
		}
		redistributions++

		if hasSeenThisMemory(memAddresses, results) {
			if hasFound == false {
				hasFound = true
				results = results[:0]
				redistributions = 0
			} else {
				break
			}
		}

		snapshot := make([]int, len(memAddresses))
		copy(snapshot, memAddresses)
		results = append(results, snapshot)
	}

	fmt.Println(redistributions)
}

func hasSeenThisMemory(mem []int, results [][]int) bool {
	for _, v := range results {
		allGood := true
		for j, k := range mem {
			if k != v[j] {
				allGood = false
				break
			}
		}

		if allGood {
			return true
		}
	}
	return false
}
