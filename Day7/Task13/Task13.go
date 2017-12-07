package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type prg struct {
	name           string
	weight         int
	hasSubPrograms bool
	subNames       string
}

func main() {
	fmt.Println("Task 13")

	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)

	programs := make([]prg, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) == 0 {
			break
		}

		splittedValues := strings.Split(line, " ")

		prg := prg{}
		prg.name = splittedValues[0]

		weightAsString := strings.Replace(splittedValues[1], "(", "", 1)
		weightAsString = strings.Replace(weightAsString, ")", "", 1)
		prg.weight, _ = strconv.Atoi(weightAsString)
		prg.subNames = ""

		if len(splittedValues) > 3 {
			prg.hasSubPrograms = true

			for j := 3; j < len(splittedValues); j++ {
				prg.subNames += splittedValues[j]
			}
		}

		programs = append(programs, prg)
	}

	minWeight := 9999
	name := ""
	for _, v := range programs {
		if v.weight < minWeight {
			if isNotSubProgram(programs, v) {
				minWeight = v.weight
				name = v.name
			}
		}
	}

	fmt.Println(name)
}

func isNotSubProgram(programs []prg, node prg) bool {
	for _, v := range programs {
		if strings.Contains(v.subNames, node.name) {
			return false
		}
	}
	return true
}
