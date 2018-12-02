package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	findSameRunes(readFile())
}

func readFile() []string {
	file, _ := ioutil.ReadFile("2nd/input.txt")
	fileSlice := strings.Split(string(file), "\r\n")
	return fileSlice
}

func findSameRunes(fileSlice []string) {
	//Nested loop for checking every line against each other
	for i, current := range fileSlice {
		currentSlice := []byte(current)

		for j := i+1; j < len(fileSlice); j++ {

			alikeRunes := 0
			commonRunes := []byte("")
			//Loop for checking each character in the t'th position of both lines
			for t, b := range currentSlice {
				if b == []byte(fileSlice[j])[t] {
					alikeRunes++
					commonRunes = append(commonRunes, currentSlice[t])
				}
			}
			if alikeRunes == len(current) - 1 {
				fmt.Println("Found the ID!")
				fmt.Println("These are the common characters:")
				for _, s := range commonRunes {
					fmt.Print(string(s))
				}
			}
		}

	}
}
