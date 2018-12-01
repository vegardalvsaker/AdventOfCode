package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	findFirstDuplicate()
}

func findFirstDuplicate() {
	data, err := ioutil.ReadFile("1st/input.txt")

	if err != nil {
		fmt.Errorf("error %s", err)
	}

	file := string(data)

	fileSlice := strings.Split(file, "\r\n")

	result := 0
	var frequencyHistory []int

	tall := 1
	for stop := false; stop == false; {
		fmt.Println("iteration #", tall)
		tall++
		for i := 0; i < len(fileSlice); i++ {
			tall, _ := strconv.Atoi(fileSlice[i])
			result += tall
			frequencyHistory = append(frequencyHistory, result)
			for k := 0; k < len(frequencyHistory)-1; k++ {
				if result == frequencyHistory[k] {
					fmt.Println("FIRST DUPLICATE", result)
					stop = true
					k = len(frequencyHistory) + 1
					i = len(fileSlice) +1
				}
			}
		}
	}
}
