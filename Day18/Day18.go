package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day18_1")

	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)

	register := map[string]int{}

	lines := make([]string, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

	lastRcv := parseOperations(lines, register)
	fmt.Println(lastRcv)
}

func parseOperations(lines []string, register map[string]int) int {
	idx := 0
	lastSound := 0
	lastRcv := 0
	for {
		if idx < 0 || idx > len(lines) {
			break
		}

		currentOperation := lines[idx]

		fmt.Println(currentOperation)

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
				lastSound = register[splitted[1]]
			}
		case "rcv":
			{
				CheckRegister(register, splitted[1])
				value := register[splitted[1]]
				if value > 0 {
					lastRcv = lastSound
					return lastRcv
				}
			}
		case "jgz":
			{
				CheckRegister(register, splitted[1])
				value := register[splitted[1]]

				if value > 0 {
					offset := GetParameter(register, splitted[2])
					idx += offset
					continue
				}
			}
		}

		idx++
		fmt.Println(register)
	}

	return -1
}

func CheckRegister(register map[string]int, name string) {
	_, exist := register[name]
	if !exist {
		register[name] = 0
	}
}

func GetParameter(register map[string]int, paramValue string) int {
	val, err := strconv.Atoi(paramValue)

	if err == nil {
		return val
	}

	return register[paramValue]
}
