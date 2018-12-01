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
	data, err := ioutil.ReadFile("1st/input.txt")

	if err != nil {
		fmt.Errorf("error %s", err)
	}

	file := string(data)
	return strings.Split(file, "\r\n")
}

func calculateResult(fileSlice []string) int {
	result := 0
	for _, current := range fileSlice {
			tall, _ := strconv.Atoi(current)
			result += tall
	}
	return result
}