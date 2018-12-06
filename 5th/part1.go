package main

import (
	"unicode"
	"fmt"
	"io/ioutil"
)

func main () {
	byteSlice, err := ioutil.ReadFile("5th/input.txt")
	fmt.Println("entries in byteSlice", len(byteSlice))
	if err != nil {
		fmt.Println("feil", err)
	}

	runeSlice := []rune{}
	for _, b := range byteSlice {
		runeSlice = append(runeSlice, rune(b))
	}

	loop(runeSlice)
	}


func loop(runeSlice []rune) {
	stop := false
	i := 0
	for stop == false {
		if i == len(runeSlice)-2 {
			stop = true
			fmt.Println("Amount of units", len(runeSlice))
		}
		if unicode.ToLower(runeSlice[i]) == unicode.ToLower(runeSlice[i+1]) {
			if unicode.IsLower(runeSlice[i]) && unicode.IsUpper(runeSlice[i+1]) || !unicode.IsLower(runeSlice[i]) && !unicode.IsUpper(runeSlice[i+1]) {
				runeSlice = append(runeSlice[:i], runeSlice[i+2:]...)
				stop = true
				loop(runeSlice)
			}
		}
		i++
	}
}