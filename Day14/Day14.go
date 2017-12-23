package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day10_2")

	input := "xlqgujun"

	rows := [128]string{}
	memGrid := [128][128]int{}

	for i := 0; i < 128; i++ {
		rows[i] = calculateHash(input + "-" + strconv.Itoa(i))
	}

	// usedBlocks := 0
	for i, v := range rows {

		binString := ""
		for _, hex := range v {
			hexString := "0" + string(hex)
			val, _ := strconv.ParseInt(hexString, 16, 0)
			asBin := Hex2Bin(byte(val))
			binString += LeftPad2Len(asBin, "0", 4)
		}

		for j, b := range binString {
			if b == '1' {
				memGrid[i][j] = 1
			} else {
				memGrid[i][j] = 0
			}
		}
	}

	for _, row := range memGrid {
		fmt.Println(row)
	}
}

func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func Hex2Bin(in byte) string {
	var out []byte
	for i := 7; i >= 0; i-- {
		b := (in >> uint(i))
		out = append(out, (b%2)+48)
	}
	return string(out)
}

func calculateHash(input string) string {
	lenghtArray := make([]int, 0)

	for _, v := range input {
		lenghtArray = append(lenghtArray, int(v))
	}

	lenghtArray = append(lenghtArray, 17)
	lenghtArray = append(lenghtArray, 31)
	lenghtArray = append(lenghtArray, 73)
	lenghtArray = append(lenghtArray, 47)
	lenghtArray = append(lenghtArray, 23)

	positions := make([]int, 0)

	for i := 0; i < 256; i++ {
		positions = append(positions, i)
	}

	currentIndex := 0
	skipSize := 0
	for k := 0; k < 64; k++ {
		for _, lenght := range lenghtArray {
			reverse(positions, currentIndex, lenght)

			if currentIndex+lenght+skipSize > len(positions) {
				currentIndex = currentIndex + lenght + skipSize - len(positions)
			} else {
				currentIndex += lenght + skipSize
			}

			if currentIndex > len(positions) {
				currentIndex -= len(positions)
			}

			skipSize++
			if skipSize == len(positions) {
				skipSize = 0
			}
		}
	}

	denseHash := [16]int{}

	result := ""
	for i := 0; i < 16; i++ {
		blockStartIndex := i * 16
		sequenceBlock := positions[blockStartIndex]
		for j := 1; j < 16; j++ {
			sequenceBlock = sequenceBlock ^ positions[blockStartIndex+j]
		}
		denseHash[i] = sequenceBlock

		h := fmt.Sprintf("%x", denseHash[i])
		if len(h) == 1 {
			result += "0"
		}
		result += h
	}

	return result
}

func reverse(input []int, startIndex int, lenght int) {
	for i := 0; i < lenght/2; i++ {
		firstIndex := 0
		if startIndex+i < len(input) {
			firstIndex = startIndex + i
		} else {
			firstIndex = (startIndex + i) - len(input)
		}

		lastIndex := 0
		if startIndex+lenght-1-i >= len(input) {
			lastIndex = startIndex + lenght - i - 1 - len(input)
		} else {
			lastIndex = startIndex + lenght - i - 1
		}

		temp := input[firstIndex]
		input[firstIndex] = input[lastIndex]
		input[lastIndex] = temp
	}
}
