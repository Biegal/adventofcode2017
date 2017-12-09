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
	garbageCount := 0
	for i := 0; i < len(content); i++ {
		v := content[i]
		fmt.Println(string(v))

		switch v {
		case '!':
			i++
			continue
		case '<':
			if garbageMode == false {
				garbageMode = true
				continue
			}
		case '>':
			garbageMode = false
			continue
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

		if garbageMode {
			garbageCount++
		}
	}

	fmt.Println(scoreSum)
	fmt.Println(garbageCount)
}
