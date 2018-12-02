package main

import (
	"io/ioutil"
	"strings"
	"bytes"
	"fmt"
)

func main(){
	fileSlice := readFile()
	fmt.Println("Rudimentary checksum: ",countDuplicate(fileSlice) * countTriplets(fileSlice))
}

func readFile() []string {
	file, _ := ioutil.ReadFile("2nd/input.txt")
	fileSlice := strings.Split(string(file), "\r\n")
	return fileSlice
}

func countDuplicate(fileSlice []string) int{
	duplicates := 0
	for _, current := range fileSlice {
		byteSlice := []byte(current)
		InnerLoop:
		for _, b := range byteSlice {
			s := string(b)
			if bytes.Count(byteSlice, []byte(s)) == 2 {
				duplicates++
				break InnerLoop
				}
			}
		}
	fmt.Println("Counts of double characters found", duplicates)
	return duplicates
	}

func countTriplets(fileSlice []string) int{
	triplets := 0
	for _, current := range fileSlice {
		byteSlice := []byte(current)
	InnerLoop:
		for _, b := range byteSlice {
			s := string(b)
			if bytes.Count(byteSlice, []byte(s)) == 3 {
				triplets++
				break InnerLoop
			}
		}
	}
	fmt.Println("Counts of triple characters found", triplets)
	return triplets
}



