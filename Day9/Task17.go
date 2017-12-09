package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Task 17")

	contentBytes, _ := ioutil.ReadFile("input.txt")
	content := string(contentBytes)

	openGroupsCount := 0
	garbageMode := false
	scoreAmount := 0
	scoreSum := 0
	for i := 0; i < len(content); i++ {
		v := content[i]
		fmt.Println(string(v))

		switch v {
		case '!':
			i++
		case '<':
			garbageMode = true
		case '>':
			garbageMode = false
		case '{':
			if garbageMode == false {
				openGroupsCount++
				scoreAmount++
			}
		case '}':
			if garbageMode == false {
				openGroupsCount--
				scoreSum += scoreAmount
				scoreAmount--
			}

		}
	}

	fmt.Println(scoreSum)
	fmt.Println(openGroupsCount)
}
