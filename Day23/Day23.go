package main

import (
	"bufio"
	"fmt"
	"os"
)

type Program struct {
	Lines    []string
	Idx      int64
	Register map[string]int64
	Pid      int
	Memory   []int64
}

func main() {
	fmt.Println("Day23_1")

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

	b := 105700
	c := 105700 + 17000
	counter := 0
	for num := b; num <= c; num += 17 {
		if !IsPrime(num) {
			counter++
		}
	}

	fmt.Println(counter)
}

func IsPrime(number int) bool {
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
