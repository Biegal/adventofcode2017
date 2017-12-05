package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Task 08")

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

		p := map[string][26]int{}

		isValid := true
		for _, w := range words {
			fmt.Println(w)
			// if p[w] {
			// 	isValid = false
			// 	break
			// }
			// p[w] = true

			lettersInCurrentWord := countLetters(w)

			for _, wordCounter := range p {
				if wordCounter == lettersInCurrentWord {
					isValid = false
					break
				}
			}

			p[w] = lettersInCurrentWord
		}

		if isValid {
			validCounter++
		}
	}
	fmt.Println(validCounter)
}

func countLetters(word string) [26]int {
	ar := [26]int{}
	for _, element := range word {
		ar[element-97]++
	}

	return ar
}
