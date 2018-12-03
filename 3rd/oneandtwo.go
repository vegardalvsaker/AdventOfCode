package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

/**
Not entirely my code. Had to sneek at
 */
func main () {
	fabric := make([][]int, 1000)

	for x := range fabric {
		fabric[x] = make([]int, 1000)
	}

	file, _ := ioutil.ReadFile("3rd/input.txt")
	fileSlice := strings.Split(string(file), "\r\n")

	inches := 0
	for _, line := range fileSlice {

		line = strings.TrimPrefix(line, "#")
		a := strings.FieldsFunc(line, Split)
		for j, t := range a {
			a[j] = strings.TrimSpace(t)
		}
		data := make([]int, 5)
		for i, s := range a {
			data[i], _ = strconv.Atoi(s)
		}

		_, width, height := data[0],data[3],data[4]
		for  offsetX := data[1]; offsetX < data[1]+width; offsetX++ {
			for offsetY := data[2]; offsetY < data[2] + height; offsetY++ {
				if fabric[offsetX][offsetY] == 1 {
					inches++
				}
				fabric[offsetX][offsetY]++

			}
		}
	}

	fmt.Println("this many inches:", inches)

//Part two
	for _, line := range fileSlice {
		line = strings.TrimPrefix(line, "#")
		a := strings.FieldsFunc(line, Split)
		for j, t := range a {
			a[j] = strings.TrimSpace(t)
		}
		data := make([]int, 5)
		for i, s := range a {
			data[i], _ = strconv.Atoi(s)
		}

		id, width, height := data[0],data[3],data[4]
		overlaps := false
		for  offsetX := data[1]; offsetX < data[1]+width; offsetX++ {
			for offsetY := data[2]; offsetY < data[2] + height; offsetY++ {
				if fabric[offsetX][offsetY] > 1 {
					overlaps = true
				}
			}
		}
		if overlaps == false {
			fmt.Println("This ID does not overlap: ", id)
		}
	}


}

func Split(r rune) bool {
	return r == '@' || r == ',' || r == ':' || r == 'x'
}
