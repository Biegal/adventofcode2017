package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day19_1")

	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)

	lines := make([]string, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

	x := strings.Index(lines[0], "|")
	y := 0

	direction := 'D'
	word := make([]rune, 0)
	current_element := '|'

	for {
		if current_element == ' ' {
			break
		}

		if direction == 'D' {
			y += 1

		} else if direction == 'U' {
			y -= 1

		} else if direction == 'L' {
			x -= 1
		} else if direction == 'R' {
			x += 1
		}
		current_element = rune(lines[y][x])
		if current_element == '+' {
			if direction == 'D' || direction == 'U' {
				if lines[y][x-1] != ' ' {
					direction = 'L'
				} else {
					direction = 'R'
				}
			} else {
				if lines[y-1][x] != ' ' {
					direction = 'U'
				} else {
					direction = 'D'
				}
			}
		} else if current_element != '|' && current_element != '-' {
			word = append(word, current_element)
		}
	}

	for _, v := range word {
		fmt.Print(string(v))
	}
}
