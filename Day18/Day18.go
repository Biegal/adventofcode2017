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
	fmt.Println("Day18_2")

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
	program0.Register["p"] = 0

	program1 := Program{}
	program1.Lines = lines
	program1.Pid = 1
	program1.Register = map[string]int64{}
	program1.Register["p"] = 1

	counter := 0
	signal1 := int64(0)
	signal2 := int64(0)
	wait1 := false
	wait2 := false
	for {
		signal1, wait1 = processOperations(&program0)
		if !wait1 {
			program1.Memory = append(program1.Memory, signal1)
		}

		signal2, wait2 = processOperations(&program1)
		if !wait2 {
			counter++
			program0.Memory = append(program0.Memory, signal2)
		}

		if wait1 && wait2 {
			break
		}
	}

	fmt.Println(counter)
}

func processOperations(prg *Program) (int64, bool) {
	register := prg.Register

	for {
		if prg.Idx < 0 || prg.Idx > int64(len(prg.Lines)) {
			break
		}

		currentOperation := prg.Lines[prg.Idx]

		splitted := strings.Split(currentOperation, " ")

		switch splitted[0] {
		case "set":
			{
				register[splitted[1]] = GetParameter(register, splitted[2])
			}
		case "add":
			{
				CheckRegister(register, splitted[1])
				parameter := GetParameter(register, splitted[2])
				register[splitted[1]] += parameter
			}
		case "mul":
			{
				CheckRegister(register, splitted[1])
				parameter := GetParameter(register, splitted[2])
				register[splitted[1]] *= parameter
			}
		case "mod":
			{
				CheckRegister(register, splitted[1])
				parameter := GetParameter(register, splitted[2])
				value := register[splitted[1]] % parameter
				register[splitted[1]] = value
			}
		case "snd":
			{
				parameter := GetParameter(register, splitted[1])
				prg.Idx++

				return parameter, false
			}
		case "rcv":
			{
				// CheckRegister(register, splitted[1])
				// value := register[splitted[1]]
				// if value > 0 {
				if len(prg.Memory) == 0 {
					return 0, true
				}

				CheckRegister(register, splitted[1])
				register[splitted[1]] = prg.Memory[0]

				prg.Memory = prg.Memory[1:]
				// }
			}
		case "jgz":
			{
				value := int64(0)
				parseVal, err := strconv.ParseInt(splitted[1], 10, 64)

				if err == nil {
					value = parseVal
				} else {
					CheckRegister(register, splitted[1])
					value = register[splitted[1]]
				}

				if value > 0 {
					offset := GetParameter(register, splitted[2])
					prg.Idx += offset
					continue
				}
			}
		}

		prg.Idx++
	}
	panic("no go")
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
