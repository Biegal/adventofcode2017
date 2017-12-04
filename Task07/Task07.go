package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Task 07")

	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)

	validCounter := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		words := strings.Split(line, " ")

		p := map[string]bool{}

		isValid := true
		for _, w := range words {
			fmt.Println(w)
			if p[w] {
				isValid = false
				break
			}
			p[w] = true
		}

		if isValid {
			validCounter++
		}
	}
	fmt.Println(validCounter)
}
