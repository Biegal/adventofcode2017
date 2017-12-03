package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Task 05")

	input := 289326.0
	//input := 1024.0
	//max := 0.0

	level := 1
	for i := 1; i < 999; i += 2 {
		max := math.Pow(float64(i), 2.0)

		fmt.Println(max)

		if max >= input {
			fmt.Println("-------------------------")

			findDistance(max, float64(i), input, float64(level))

			break
		}

		level++
	}
}

func findDistance(maxElement float64, elementsInRow float64, input float64, level float64) {
	fmt.Printf("Calculating distance for %f \n", input)
	fmt.Println(maxElement)
	fmt.Println(elementsInRow)

	numInRow := elementsInRow
	for i := maxElement; i > 0; i-- {

		numInRow--

		if numInRow == 1 {
			numInRow = elementsInRow
		}

		if float64(i) == input {
			fmt.Printf("Level %f \n", level)
			fmt.Printf("Number in row %f \n", numInRow)
			fmt.Printf("Elements in row %f \n", elementsInRow)

			middleElemNo := math.Floor(elementsInRow / 2.0)

			fmt.Printf("Middle elem %f \n", middleElemNo)

			if numInRow > middleElemNo {
				fmt.Printf("Distance %f \n", (numInRow-middleElemNo)+(level-1))
			} else {
				fmt.Printf("Distance %f \n", math.Abs(middleElemNo-numInRow)+(level-1))
			}

		}
	}

}
