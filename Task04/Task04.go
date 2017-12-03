package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Task 04")

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
		for i := 0; i < len(nums); i++ {
			num, _ := strconv.Atoi(nums[i])

			isFound := false

			for j := 0; j < len(nums); j++ {
				if i == j {
					continue
				}

				num2, _ := strconv.Atoi(nums[j])

				if num%num2 == 0 {
					sum += num / num2
					isFound = true
					break
				}
				if num2%num == 0 {
					sum += num2 / num
					isFound = true
					break
				}
			}

			if isFound == true {
				break
			}
		}
	}
	fmt.Println(sum)
}
