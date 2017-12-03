package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Task 06")

	resp, _ := http.Get("https://oeis.org/A141481/b141481.txt")
	defer resp.Body.Close()

	lookupScanner := bufio.NewScanner(resp.Body)

	for lookupScanner.Scan() {
		line := lookupScanner.Text()
		nums := strings.Split(line, " ")
		index, _ := strconv.Atoi(nums[0])
		value, _ := strconv.Atoi(nums[1])

		if value > 289326 {
			fmt.Printf("%d %d \n", index, value)
			break
		}

	}
}
