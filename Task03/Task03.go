package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task02 Solution
func main() {
	fmt.Println("Task 03")

	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		nums := strings.Split(line, "\t")
		lowest, _ := strconv.Atoi(nums[0])
		highest := lowest
		for i := 1; i < len(nums); i++ {
			num, _ := strconv.Atoi(nums[i])
			if num < lowest {
				lowest = num
			}
			if num > highest {
				highest = num
			}
		}
		sum += highest - lowest
	}
	fmt.Println(sum)
}
