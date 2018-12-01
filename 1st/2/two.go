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
	data, _ := ioutil.ReadFile("1st/input.txt")
	fileSlice := strings.Split(string(data), "\r\n")

	frequency := 0
	var frequencyHistory []int

	tall := 1
	for stop := false; stop == false; {
		fmt.Println("iteration #", tall)
		tall++
		for i := 0; i < len(fileSlice); i++ {
			tall, _ := strconv.Atoi(fileSlice[i])
			frequency += tall
			frequencyHistory = append(frequencyHistory, frequency)
			for k := 0; k < len(frequencyHistory)-1; k++ {
				if frequency == frequencyHistory[k] {
					fmt.Println("FIRST DUPLICATE", frequency)
					stop = true
					k = len(frequencyHistory) + 1
					i = len(fileSlice) +1
				}
			}
		}
	}
}
