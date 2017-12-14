package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	fmt.Println("Day11_1")

	contentBytes, _ := ioutil.ReadFile("input.txt")
	content := string(contentBytes)

	moves := strings.Split(content, ",")

	x := 0.0
	y := 0.0

	for _, move := range moves {

		switch move {
		case "n":
			y += 2
		case "s":
			y -= 2
		case "ne":
			y++
			x++
		case "nw":
			y++
			x--
		case "se":
			y--
			x++
		case "sw":
			y--
			x--
		}
	}

	stps := 0.0
	x = math.Abs(x)
	y = math.Abs(y)
	if x >= y {
		stps = x
	} else {
		stps = x + ((y - x) / 2)
	}
	fmt.Println(stps)
}
