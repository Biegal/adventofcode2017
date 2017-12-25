package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	program0 := Program{}
	program0.Lines = lines
	program0.Pid = 0
	program0.Register = map[string]int64{}

	processOperations(&program0)
}

func processOperations(prg *Program) {
	register := prg.Register
	mulCounter := 0
	for {
		if prg.Idx < 0 || prg.Idx >= int64(len(prg.Lines)) {
			break
		}

		currentOperation := prg.Lines[prg.Idx]
		fmt.Println(currentOperation)

		splitted := strings.Split(currentOperation, " ")

		switch splitted[0] {
		case "set":
			{
				register[splitted[1]] = GetParameter(register, splitted[2])
			}
		case "sub":
			{
				CheckRegister(register, splitted[1])
				parameter := GetParameter(register, splitted[2])
				register[splitted[1]] -= parameter
			}
		case "mul":
			{
				CheckRegister(register, splitted[1])
				parameter := GetParameter(register, splitted[2])
				register[splitted[1]] *= parameter
				mulCounter++
				fmt.Println(mulCounter)
			}
		case "jnz":
			{
				value := int64(0)
				parseVal, err := strconv.ParseInt(splitted[1], 10, 64)

				if err == nil {
					value = parseVal
				} else {
					CheckRegister(register, splitted[1])
					value = register[splitted[1]]
				}

				if value != 0 {
					offset := GetParameter(register, splitted[2])
					prg.Idx += offset
					continue
				}
			}
		}

		prg.Idx++
	}
}

func CheckRegister(register map[string]int64, name string) {
	_, exist := register[name]
	if !exist {
		register[name] = 0
	}
}

func GetParameter(register map[string]int64, paramValue string) int64 {
	val, err := strconv.ParseInt(paramValue, 10, 64)

	if err == nil {
		return val
	}

	return register[paramValue]
}
