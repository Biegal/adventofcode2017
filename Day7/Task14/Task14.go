package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type prg struct {
	name        string
	weight      int
	subNames    string
	subPrograms []*prg
	parent      *prg
}

func (this *prg) GetWeight() int {
	sum := this.weight

	for _, v := range this.subPrograms {
		sum += v.GetWeight()
	}

	return sum
}

func main() {
	fmt.Println("Task 14")

	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)

	programs := make([]*prg, 0)
	programDict := make(map[string]*prg)
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
			for j := 3; j < len(splittedValues); j++ {
				prg.subNames += splittedValues[j]
			}
		}

		programs = append(programs, &prg)
		programDict[prg.name] = &prg
	}

	for _, prg := range programs {
		if len(prg.subNames) > 0 {
			names := strings.Split(prg.subNames, ",")

			for _, v := range names {
				subProgram := programDict[v]
				prg.subPrograms = append(prg.subPrograms, subProgram)

				subProgram.parent = prg
			}
		}
	}

	var root *prg
	for _, prg := range programs {
		if prg.parent == nil {
			root = prg
		}
	}

	for _, v := range root.subPrograms {
		fmt.Printf("%s %d \n", v.name, v.GetWeight())
	}
	//646
	levelTheTree(root)

	//for {

	//}

}

func levelTheTree(root *prg) {
	levelWeight := make([]int, len(root.subPrograms))
	for i, v := range root.subPrograms {
		levelWeight[i] = v.GetWeight()
	}

	isBalanced, weightOff := isBalanced(levelWeight)

	if isBalanced == false {
		fmt.Printf("Unbalanced %d \n", weightOff)

		// var unbalancedProgram *prg
		// for i, v := range levelWeight {
		// 	if v == weightOff {
		// 		unbalancedProgram = root.subPrograms[i]
		// 	}
		// }

		// levelTheTree(unbalancedProgram)
	} else {
		fmt.Println("OKOKOK")
	}

	fmt.Println(levelWeight)

}

func isBalanced(level []int) (bool, int) {
	if len(level) == 1 {
		return true, 0
	}

	counter := map[int]int{}

	for _, v := range level {
		counter[v] += 1
	}

	if len(counter) == 1 {
		return true, 0
	}

	minIndex := 0
	minWeight := 999999
	for i, v := range counter {
		if v < minWeight {
			minIndex = i
		}
	}

	return false, minIndex
}
