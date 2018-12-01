package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fileSlice := readFromFile()
	fmt.Println(calculateResult(fileSlice))

}

func readFromFile() []string {
	data, _ := ioutil.ReadFile("1st/input.txt")
	return strings.Split(string(data), "\r\n")
}

func calculateResult(fileSlice []string) int {
	result := 0
	for _, current := range fileSlice {
		tall, _ := strconv.Atoi(current)
		result += tall
	}
	return result
}
